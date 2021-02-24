package druid

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
	defaultRetryWaitMin          = 100 * time.Millisecond
	defaultRetryWaitMax          = 3 * time.Second
	defaultRetryMax              = 5
	defaultSkipTlsOption         = false
)

var (
	defaultBackoff = retryablehttp.DefaultBackoff
	defaultRetry   = retryablehttp.DefaultRetryPolicy
)

type Client struct {
	http          *retryablehttp.Client
	baseURL       *url.URL
	username      string
	password      string
	basicAuth     bool
	skipTLSVerify bool
}

type clientOptions struct {
	httpClient    *http.Client
	username      string
	password      string
	backoff       retryablehttp.Backoff
	retry         retryablehttp.CheckRetry
	retryWaitMin  time.Duration
	retryWaitMax  time.Duration
	retryMax      int
	skipTLSVerify bool
}

type ClientOption func(*clientOptions)

func NewClient(baseURL string, options ...ClientOption) (*Client, error) {
	opts := &clientOptions{
		httpClient:    defaultHTTPClient(),
		backoff:       defaultBackoff,
		retry:         defaultRetry,
		retryWaitMin:  defaultRetryWaitMin,
		retryWaitMax:  defaultRetryWaitMax,
		retryMax:      defaultRetryMax,
		skipTLSVerify: defaultSkipTlsOption,
	}
	for _, opt := range options {
		opt(opts)
	}
	if opts.skipTLSVerify {
		InsecureSkipVerify(opts)
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
		username:      opts.username,
		password:      opts.password,
		basicAuth:     opts.username != "" && opts.password != "",
		skipTLSVerify: opts.skipTLSVerify,
	}
	if err := c.setBaseURL(baseURL); err != nil {
		return nil, err
	}

	return c, nil
}

func InsecureSkipVerify(opts *clientOptions) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	opts.httpClient.Transport = &http.Transport{TLSClientConfig: tlsConfig}
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

func WithBasicAuth(username, password string) ClientOption {
	return func(opts *clientOptions) {
		opts.username = username
		opts.password = password
	}
}

func WithSkipTLSVerify(skip bool) ClientOption {
	return func(opts *clientOptions) {
		opts.skipTLSVerify = skip
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
	data, err := ioutil.ReadAll(r.Body)
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
