package druid

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

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
		var raw any
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

func parseError(raw any) string {
	if raw, isMapSI := raw.(map[string]any); isMapSI {
		if errStr, hasErrorStr := raw["error"]; hasErrorStr {
			return errStr.(string)
		}
	}
	return fmt.Sprintf("failed to parse unexpected error type: %T", raw)
}
