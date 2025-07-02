package druid

import (
	"encoding/json"
	"net/http"
	"net/url"

	querystring "github.com/google/go-querystring/query"
	"github.com/hashicorp/go-retryablehttp"
)

// buildURL constructs the full URL for a request
func (c *Client) buildURL(path string) (*url.URL, error) {
	u := *c.baseURL
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped
	return &u, nil
}

// prepareBody prepares the request body based on method and options
func (c *Client) prepareBody(method string, opt any, u *url.URL, reqHeaders http.Header) (any, error) {
	var body any
	if opt != nil {
		switch method {
		case "POST", "PUT":
			reqHeaders.Set("Content-Type", "application/json")
			var err error
			body, err = json.Marshal(opt)
			if err != nil {
				return nil, err
			}
		default:
			q, err := querystring.Values(opt)
			if err != nil {
				return nil, err
			}
			u.RawQuery = q.Encode()
		}
	}
	return body, nil
}

// setupHeaders applies headers and authentication to the request
func (c *Client) setupHeaders(r *retryablehttp.Request, reqHeaders http.Header) {
	r.Header = reqHeaders
	if c.basicAuth {
		r.SetBasicAuth(c.username, c.password)
	}
}

func (c *Client) NewRequest(method, path string, opt any) (*retryablehttp.Request, error) {
	u, err := c.buildURL(path)
	if err != nil {
		return nil, err
	}

	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	body, err := c.prepareBody(method, opt, u, reqHeaders)
	if err != nil {
		return nil, err
	}

	r, err := retryablehttp.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	c.setupHeaders(r, reqHeaders)
	return r, nil
}
