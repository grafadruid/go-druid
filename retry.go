package druid

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

var (
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

type druidErrorResponse struct {
	Error        string
	ErrorMessage string
	ErrorClass   string
	Host         string
}

// Non-retryable Druid error types based on official documentation
// https://druid.apache.org/docs/latest/querying/#query-execution-failures
var nonRetryableDruidErrors = map[string]bool{
	"SQL parse failed":       true,
	"Plan validation failed": true,
	"Unsupported operation":  true,
	"Query cancelled":        true,
	"Unknown exception":      true,
}

// shouldRetryConnectionError determines if a connection-level error should be retried
func shouldRetryConnectionError(err error) (bool, error) {
	urlErr, ok := err.(*url.Error)
	if !ok {
		return true, nil // Retry unknown connection errors
	}

	// Don't retry if the error was due to too many redirects
	if redirectsErrorRe.MatchString(urlErr.Error()) {
		return false, urlErr
	}

	// Don't retry if the error was due to an invalid protocol scheme
	if schemeErrorRe.MatchString(urlErr.Error()) {
		return false, urlErr
	}

	// Don't retry if the error was due to TLS cert verification failure
	if _, ok := urlErr.Err.(x509.UnknownAuthorityError); ok {
		return false, urlErr
	}

	return true, nil
}

// processDruidErrorResponse reads and analyzes Druid error responses
func processDruidErrorResponse(resp *http.Response) (bool, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return true, fmt.Errorf("failed to read the response from Druid: %w", err)
	}

	var errResp druidErrorResponse
	if err := json.Unmarshal(body, &errResp); err != nil {
		return true, fmt.Errorf("failed to read the response from Druid: %w", err)
	}

	// Check if this is a non-retryable Druid error
	if nonRetryableDruidErrors[errResp.Error] {
		// When aborting the retry, the response body should be closed:
		// https://pkg.go.dev/github.com/hashicorp/go-retryablehttp#CheckRetry
		resp.Body.Close()
		return false, fmt.Errorf("failed to query Druid: %+v", errResp)
	}

	return true, fmt.Errorf("error response from Druid: %+v", errResp)
}

func defaultRetry(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	// Handle connection-level errors
	// As explained here https://golang.org/pkg/net/http/#Client.Do,
	// An error is returned if caused by client policy (such as CheckRedirect), or failure to speak HTTP (such as a network connectivity problem). A non-2xx status code doesn't cause an error.
	if err != nil {
		return shouldRetryConnectionError(err)
	}

	// Success case - no retry needed
	if resp.StatusCode == http.StatusOK {
		return false, nil
	}

	// Handle Druid error responses
	return processDruidErrorResponse(resp)
}

func defaultErrorHandler(resp *http.Response, err error, numTries int) (*http.Response, error) {
	// Drain and close the response body so the connection can be reused:
	// https://pkg.go.dev/github.com/hashicorp/go-retryablehttp#ErrorHandler
	defer resp.Body.Close()
	io.Copy(io.Discard, io.LimitReader(resp.Body, respReadLimit))

	return resp, fmt.Errorf("failed after %d attempt(s). Last error: %w", numTries, err)
}
