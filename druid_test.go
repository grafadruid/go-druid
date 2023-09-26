package druid

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBaseURLWithSuffix(t *testing.T) {
	assert := assert.New(t)

	d, err := NewClient("localhost:8082")
	assert.Nil(err, "error should be nil")
	assert.NotNil(d, "client should not be nil")

	wantBaseURL, _ := url.ParseRequestURI("/")
	err = d.setBaseURL("")
	assert.Nil(err, "error should be nil")
	assert.Equal(d.baseURL, wantBaseURL, "they should not be equal")
}

func TestNewClientWithSkipVerify(t *testing.T) {
	assert := assert.New(t)

	var druidOpts []ClientOption
	druidOpts = append(druidOpts, WithSkipTLSVerify())

	d, err := NewClient("localhost:8082", druidOpts...)
	assert.Nil(err, "error should be nil")
	assert.NotNil(d, "client should not be nil")
	assert.True(d.http.HTTPClient.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify, "insecure skip verify should be true")
}

// TODO: at some point use https://golang.org/src/crypto/tls/example_test.go this to create server with bad cert and test

func TestDefaultRetry(t *testing.T) {
	type testCase struct {
		statusCode int
		response   string
		wantErr    string
		wantRetry  bool
	}

	run := func(t *testing.T, tc testCase) {
		// Given
		ctx := context.Background()
		resp := buildMockResp(tc.statusCode, tc.response)

		// When
		retry, err := defaultRetry(ctx, &resp, nil)

		// Then
		assert := assert.New(t)
		if tc.wantErr != "" {
			assert.Equal(tc.wantErr, err.Error())
		} else {
			assert.NoError(err)
		}
		assert.Equal(tc.wantRetry, retry)
	}

	testCases := map[string]testCase{
		"OK": {
			statusCode: 200,
			response:   `{ "id"": "12345"}`,
			wantErr:    "",
			wantRetry:  false,
		},
		"SQL parse error": {
			statusCode: 400,
			response: `{
				"error": "SQL parse failed", "errorMessage" : "incorrect input."
			}`,
			wantErr:   "failed to query Druid: {Error:SQL parse failed ErrorMessage:incorrect input. ErrorClass: Host:}",
			wantRetry: false,
		},
		"SQL plan validatio error": {
			statusCode: 400,
			response: `{
				"error": "Plan validation failed", "errorMessage" : "validation error."
			}`,
			wantErr:   "failed to query Druid: {Error:Plan validation failed ErrorMessage:validation error. ErrorClass: Host:}",
			wantRetry: false,
		},
		"Resource limit error": {
			statusCode: 400,
			response: `{
				"error": "Resource limit exceeded", "errorMessage" : "Something bad happened."
			}`,
			wantErr:   "error response from Druid: {Error:Resource limit exceeded ErrorMessage:Something bad happened. ErrorClass: Host:}",
			wantRetry: true,
		},
		"Query capacity exceeded": {
			statusCode: 429,
			response: `{
				"error": "Query capacity exceeded", "errorMessage" : "capacity exceeded."
			}`,
			wantErr:   "error response from Druid: {Error:Query capacity exceeded ErrorMessage:capacity exceeded. ErrorClass: Host:}",
			wantRetry: true,
		},
		"Unsupported operation": {
			statusCode: 501,
			response: `{
				"error": "Unsupported operation", "errorMessage" : "wrong operation."
			}`,
			wantErr:   "failed to query Druid: {Error:Unsupported operation ErrorMessage:wrong operation. ErrorClass: Host:}",
			wantRetry: false,
		},
		"Query timeout": {
			statusCode: 504,
			response: `{
				"error": "Query timeout", "errorMessage" : "timeout."
			}`,
			wantErr:   "error response from Druid: {Error:Query timeout ErrorMessage:timeout. ErrorClass: Host:}",
			wantRetry: true,
		},
		"Query cancelled": {
			statusCode: 500,
			response: `{
				"error": "Query cancelled", "errorMessage" : "cancelled."
			}`,
			wantErr:   "failed to query Druid: {Error:Query cancelled ErrorMessage:cancelled. ErrorClass: Host:}",
			wantRetry: false,
		},
		"Unknown exception": {
			statusCode: 500,
			response: `{
				"error": "Unknown exception", "errorMessage" : "failure."
			}`,
			wantErr:   "failed to query Druid: {Error:Unknown exception ErrorMessage:failure. ErrorClass: Host:}",
			wantRetry: false,
		},
		"Invalid json": {
			statusCode: 500,
			response:   `invalid json`,
			wantErr:    "failed to read the response from Druid: invalid character 'i' looking for beginning of value",
			wantRetry:  true,
		},
		"Request body content type is not in JSON format": {
			statusCode: 415,
			response: `{
				"error": "Request body content type is not in JSON format."
			}`,
			wantErr:   "error response from Druid: {Error:Request body content type is not in JSON format. ErrorMessage: ErrorClass: Host:}",
			wantRetry: true,
		},
		"Query Supervisor Status: Invalid supervisor ID": {

			statusCode: 404,
			response: `{
				"error": "Invalid supervisor ID."
			}`,
			wantErr:   "error response from Druid: {Error:Invalid supervisor ID. ErrorMessage: ErrorClass: Host:}",
			wantRetry: true,
		},
		"Terminate Query Supervisor: Invalid supervisor ID": {

			statusCode: 404,
			response: `{
				"error": "Invalid supervisor ID or supervisor not running."
			}`,
			wantErr:   "error response from Druid: {Error:Invalid supervisor ID or supervisor not running. ErrorMessage: ErrorClass: Host:}",
			wantRetry: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { run(t, tc) })

	}
}

func buildMockResp(statusCode int, body string) http.Response {
	var st string
	switch statusCode {
	case 200:
		st = "200 OK"
	case 400:
		st = "400 Bad Request"
	case 404:
		st = "404 Not Found"
	case 415:
		st = "415 Unsupported Media Type"
	case 429:
		st = "429 Too Many Requests"
	case 500:
		st = "500 Internal Server Error"
	case 501:
		st = "Not Implemented"
	case 504:
		st = "Gateway Timeout"
	default:
		panic(fmt.Errorf("Unsupported mock status code: %d", statusCode))
	}
	return http.Response{
		Status: st, StatusCode: statusCode,
		Body: io.NopCloser(strings.NewReader(body)),
	}
}
