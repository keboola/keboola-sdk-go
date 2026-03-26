// nolint: dupl
package keboola

import (
	"fmt"
	"net/http"
)

// EditorError represents the structure of editor service API error.
type EditorError struct {
	Message   string `json:"message"`
	ErrorInfo string `json:"error"`
	request   *http.Request
	response  *http.Response
}

func (e *EditorError) Error() string {
	msg := e.Message
	if e.request != nil {
		msg += fmt.Sprintf(`, method: "%s", url: "%s"`, e.request.Method, e.request.URL)
	}
	if e.response != nil {
		msg += fmt.Sprintf(`, httpCode: "%d"`, e.StatusCode())
	}
	return msg
}

// ErrorName returns a human-readable name of the error.
func (e *EditorError) ErrorName() string {
	return e.ErrorInfo
}

// ErrorUserMessage returns error message for end user.
func (e *EditorError) ErrorUserMessage() string {
	return e.Message
}

// StatusCode returns HTTP status code.
func (e *EditorError) StatusCode() int {
	if e.response == nil {
		return 0
	}
	return e.response.StatusCode
}

// SetRequest method allows injection of HTTP request to the error, it implements client.errorWithRequest.
func (e *EditorError) SetRequest(request *http.Request) {
	e.request = request
}

// SetResponse method allows injection of HTTP response to the error, it implements client.errorWithResponse.
func (e *EditorError) SetResponse(response *http.Response) {
	e.response = response
}
