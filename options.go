package druid

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

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

func WithBasicAuth(username, password string) ClientOption {
	return func(opts *clientOptions) {
		if username != "" && password != "" {
			opts.username = username
			opts.password = password
		}
	}
}

func WithSkipTLSVerify() ClientOption {
	return func(opts *clientOptions) {
		if opts.httpClient.Transport == nil {
			opts.httpClient.Transport = &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			return
		}

		transport, ok := opts.httpClient.Transport.(*http.Transport)
		if !ok {
			// If it's not an *http.Transport, create a new one
			opts.httpClient.Transport = &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			return
		}

		if transport.TLSClientConfig == nil {
			transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		} else {
			transport.TLSClientConfig.InsecureSkipVerify = true
		}
	}
}

func WithCustomBackoff(backoff retryablehttp.Backoff) ClientOption {
	return func(opts *clientOptions) {
		if backoff != nil {
			opts.backoff = backoff
		}
	}
}

func WithCustomRetry(retry retryablehttp.CheckRetry) ClientOption {
	return func(opts *clientOptions) {
		if retry != nil {
			opts.retry = retry
		}
	}
}

func WithCustomErrorHandler(h retryablehttp.ErrorHandler) ClientOption {
	return func(opts *clientOptions) {
		if h != nil {
			opts.errorHandler = h
		}
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(opts *clientOptions) {
		if httpClient != nil {
			opts.httpClient = httpClient
		}
	}
}

func WithRetryWaitMin(retryWaitMin time.Duration) ClientOption {
	return func(opts *clientOptions) {
		if retryWaitMin > 0 {
			opts.retryWaitMin = retryWaitMin
		}
	}
}

func WithRetryWaitMax(retryWaitMax time.Duration) ClientOption {
	return func(opts *clientOptions) {
		if retryWaitMax > 0 {
			opts.retryWaitMax = retryWaitMax
		}
	}
}

func WithRetryMax(retryMax int) ClientOption {
	return func(opts *clientOptions) {
		if retryMax >= 0 {
			opts.retryMax = retryMax
		}
	}
}

func defaultHTTPClient() *http.Client {
	return &http.Client{}
}
