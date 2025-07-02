package druid

import (
	"encoding/json"
	"net/url"
	"strings"
	"time"

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
)

var defaultBackoff = retryablehttp.DefaultBackoff

type Client struct {
	http      *retryablehttp.Client
	baseURL   *url.URL
	username  string
	password  string
	basicAuth bool
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

func (c *Client) Do(r *retryablehttp.Request, result any) (*Response, error) {
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

func (c *Client) ExecuteRequest(method, path string, opt, result any) (*Response, error) {
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
