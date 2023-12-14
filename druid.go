package druid

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	querystring "github.com/google/go-querystring/query"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	processInformationPathPrefix = "status/"
	coordinatorPathPrefix        = "druid/coordinator/v1/"
	overlordPathPrefix           = "druid/indexer/v1/"
	middleManagerPathPrefix      = "druid/worker/v1/"
	peonPathPrefix               = "druid/worker/v1/chat/"
	historicalPathPrefix         = "druid/historical/v1/"
	supervisorPathPrefix         = "druid/indexer/v1/supervisor"
	defaultRetryWaitMin          = 100 * time.Millisecond
	defaultRetryWaitMax          = 3 * time.Second
	defaultRetryMax              = 5
)

var (
	defaultBackoff = retryablehttp.DefaultBackoff
	// A regular expression to match the error returned by net/http when the
	// configured number of redirects is exhausted. This error isn't typed
	// specifically so we resort to matching on the error string.
	redirectsErrorRe = regexp.MustCompile(`stopped after \d+ redirects\z`)

	// A regular expression to match the error returned by net/http when the
	// scheme specified in the URL is invalid. This error isn't typed
	// specifically so we resort to matching on the error string.
	schemeErrorRe = regexp.MustCompile(`unsupported protocol scheme`)

	// We need to consume response bodies to maintain http connections, but
	// limit the size we consume to respReadLimit.
	respReadLimit = int64(4096)
)

type Client struct {
	http      *retryablehttp.Client
	baseURL   *url.URL
	username  string
	password  string
	basicAuth bool
}

type clientOptions struct {
	httpClient   *http.Client
	username     string
	password     string
	backoff      retryablehttp.Backoff
	errorHandler retryablehttp.ErrorHandler
	retry        retryablehttp.CheckRetry
	retryWaitMin time.Duration
	retryWaitMax time.Duration
	retryMax     int
}

type ClientOption func(*clientOptions)

type druidErrorReponse struct {
	Error        string
	ErrorMessage string
	ErrorClass   string
	Host         string
}

func NewClient(baseURL string, options ...ClientOption) (*Client, error) {
	opts := &clientOptions{
		httpClient:   defaultHTTPClient(),
		backoff:      defaultBackoff,
		errorHandler: defaultErrorHandler,
		retry:        defaultRetry,
		retryWaitMin: defaultRetryWaitMin,
		retryWaitMax: defaultRetryWaitMax,
		retryMax:     defaultRetryMax,
	}
	for _, opt := range options {
		opt(opts)
	}
	c := &Client{
		http: &retryablehttp.Client{
			Backoff:      opts.backoff,
			CheckRetry:   opts.retry,
			HTTPClient:   opts.httpClient,
			RetryWaitMin: opts.retryWaitMin,
			RetryWaitMax: opts.retryWaitMax,
			RetryMax:     opts.retryMax,
		},
		username:  opts.username,
		password:  opts.password,
		basicAuth: opts.username != "" && opts.password != "",
	}
	if err := c.setBaseURL(baseURL); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Close() error {
	return nil
}

func (c *Client) NewRequest(method, path string, opt interface{}) (*retryablehttp.Request, error) {
	u := *c.baseURL
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	var body interface{}
	if opt != nil {
		switch {
		case method == "POST" || method == "PUT":
			reqHeaders.Set("Content-Type", "application/json")
			if opt != nil {
				body, err = json.Marshal(opt)
				if err != nil {
					return nil, err
				}
			}
		default:
			q, err := querystring.Values(opt)
			if err != nil {
				return nil, err
			}
			u.RawQuery = q.Encode()
		}
	}

	r, err := retryablehttp.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	r.Header = reqHeaders
	if c.basicAuth {
		r.SetBasicAuth(c.username, c.password)
	}

	return r, nil
}

func (c *Client) Do(r *retryablehttp.Request, result interface{}) (*Response, error) {
	resp, err := c.http.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response := &Response{resp}
	if err = response.ExtractError(); err != nil {
		return nil, err
	}
	if result != nil {
		if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
			return nil, err
		}
	}
	return response, nil
}

func (c *Client) ExecuteRequest(method, path string, opt, result interface{}) (*Response, error) {
	req, err := c.NewRequest(method, path, opt)
	if err != nil {
		return nil, err
	}
	return c.Do(req, result)
}

func defaultRetry(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	// As explained here https://golang.org/pkg/net/http/#Client.Do,
	// An error is returned if caused by client policy (such as CheckRedirect), or failure to speak HTTP (such as a network connectivity problem). A non-2xx status code doesn't cause an error.
	if err != nil {
		if v, ok := err.(*url.Error); ok {
			// Don't retry if the error was due to too many redirects.
			if redirectsErrorRe.MatchString(v.Error()) {
				return false, v
			}

			// Don't retry if the error was due to an invalid protocol scheme.
			if schemeErrorRe.MatchString(v.Error()) {
				return false, v
			}

			// Don't retry if the error was due to TLS cert verification failure.
			if _, ok := v.Err.(x509.UnknownAuthorityError); ok {
				return false, v
			}
		}

		return true, nil
	}

	if resp.StatusCode == http.StatusOK {
		return false, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return true, fmt.Errorf("failed to read the response from Druid: %w", err)
	}
	var errResp druidErrorReponse
	err = json.Unmarshal(body, &errResp)
	if err != nil {
		return true, fmt.Errorf("failed to read the response from Druid: %w", err)
	}

	// https://druid.apache.org/docs/latest/querying/querying.html#query-execution-failures
	switch errResp.Error {
	case "SQL parse failed":
		goto ABORT
	case "Plan validation failed":
		goto ABORT
	case "Unsupported operation":
		goto ABORT
	case "Query cancelled":
		goto ABORT
	case "Unknown exception":
		goto ABORT
	case "Request body content type is not in JSON format":
		goto ABORT
	case "Invalid supervisor ID":
		goto ABORT
	case "Invalid supervisor ID or supervisor not running":
		goto ABORT
	default:
		return true, fmt.Errorf("error response from Druid: %+v", errResp)
	}

ABORT:
	// When aborting the retry, the response body should be closed:
	// https://pkg.go.dev/github.com/hashicorp/go-retryablehttp#CheckRetry
	resp.Body.Close()
	return false, fmt.Errorf("failed to query Druid: %+v", errResp)
}

func defaultErrorHandler(resp *http.Response, err error, numTries int) (*http.Response, error) {
	// Drain and close the response body so the connection can be reused:
	// https://pkg.go.dev/github.com/hashicorp/go-retryablehttp#ErrorHandler
	defer resp.Body.Close()
	io.Copy(io.Discard, io.LimitReader(resp.Body, respReadLimit))

	return resp, fmt.Errorf("Failed after %d attempt(s). Last error: %w", numTries, err)
}

func (c *Client) setBaseURL(urlStr string) error {
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}
	baseURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return err
	}
	c.baseURL = baseURL
	return nil
}

func (c *Client) Common() *CommonService {
	return &CommonService{client: c}
}

func (c *Client) Query() *QueryService {
	return &QueryService{client: c}
}

func (c *Client) Supervisor() *SupervisorService {
	return &SupervisorService{client: c}
}

func (c *Client) Tasks() *TasksService {
	return &TasksService{client: c}
}

func WithBasicAuth(username, password string) ClientOption {
	return func(opts *clientOptions) {
		opts.username = username
		opts.password = password
	}
}

func WithSkipTLSVerify() ClientOption {
	return func(opts *clientOptions) {
		if nil == opts.httpClient.Transport {
			opts.httpClient.Transport = &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
		}
		opts.httpClient.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify = true
	}
}

func WithCustomBackoff(backoff retryablehttp.Backoff) ClientOption {
	return func(opts *clientOptions) {
		opts.backoff = backoff
	}
}

func WithCustomRetry(retry retryablehttp.CheckRetry) ClientOption {
	return func(opts *clientOptions) {
		opts.retry = retry
	}
}

func WithCustomErrorHandler(h retryablehttp.ErrorHandler) ClientOption {
	return func(opts *clientOptions) {
		opts.errorHandler = h
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(opts *clientOptions) {
		opts.httpClient = httpClient
	}
}

func WithRetryWaitMin(retryWaitMin time.Duration) ClientOption {
	return func(opts *clientOptions) {
		opts.retryWaitMin = retryWaitMin
	}
}

func WithRetryWaitMax(retryWaitMax time.Duration) ClientOption {
	return func(opts *clientOptions) {
		opts.retryWaitMax = retryWaitMax
	}
}

func WithRetryMax(retryMax int) ClientOption {
	return func(opts *clientOptions) {
		opts.retryMax = retryMax
	}
}

type Response struct {
	*http.Response
}

func (r *Response) ExtractError() error {
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	}
	errorResponse := &errResponse{Response: r.Response}
	data, err := io.ReadAll(r.Body)
	if err == nil && data != nil {
		errorResponse.Body = data
		var raw interface{}
		if err := json.Unmarshal(data, &raw); err != nil {
			errorResponse.Message = r.Status
		} else {
			errorResponse.Message = parseError(raw)
		}
	}
	return errorResponse
}

type errResponse struct {
	Body     []byte
	Response *http.Response
	Message  string
}

func (e *errResponse) Error() string {
	path, _ := url.QueryUnescape(e.Response.Request.URL.Path)
	return fmt.Sprintf(
		"error with code %d %s %s message: %s",
		e.Response.StatusCode,
		e.Response.Request.Method,
		fmt.Sprintf("%s://%s%s", e.Response.Request.URL.Scheme, e.Response.Request.URL.Host, path),
		e.Message,
	)
}

func parseError(raw interface{}) string {
	if raw, isMapSI := raw.(map[string]interface{}); isMapSI {
		if errStr, hasErrorStr := raw["error"]; hasErrorStr {
			return errStr.(string)
		}
	}
	return fmt.Sprintf("failed to parse unexpected error type: %T", raw)
}

func defaultHTTPClient() *http.Client {
	return &http.Client{}
}
