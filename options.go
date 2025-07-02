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

func defaultHTTPClient() *http.Client {
	return &http.Client{}
}
