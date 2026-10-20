package keboola

import (
	"fmt"
	"net/http"
)

// MetastoreError represents an error from the Metastore API.
type MetastoreError struct {
	// HTTPStatus is the HTTP status code the Metastore reports in its body ("error").
	HTTPStatus  int            `json:"error"`
	Code        string         `json:"code"`
	Message     string         `json:"exception"`
	ExceptionID string         `json:"exceptionId"`
	Status      string         `json:"status"`
	Context     map[string]any `json:"context,omitempty"`
	// Errors is the detail of a validation failure, its shape is defined by the Metastore.
	Errors   any `json:"errors,omitempty"`
	request  *http.Request
	response *http.Response
}

// Error returns the error message with HTTP context.
func (e *MetastoreError) Error() string {
	msg := e.Message
	if e.request != nil {
		msg += fmt.Sprintf(`, method: "%s", url: "%s"`, e.request.Method, e.request.URL)
	}
	if e.response != nil {
		msg += fmt.Sprintf(`, httpCode: "%d"`, e.response.StatusCode)
	}
	if len(e.ExceptionID) > 0 {
		msg += fmt.Sprintf(`, exceptionId: "%s"`, e.ExceptionID)
	}
	return msg
}

// ErrorName returns a human-readable name for the error.
func (e MetastoreError) ErrorName() string {
	if e.response != nil {
		return e.response.Status
	}
	return "metastore api error"
}

// ErrorUserMessage returns the error message suitable for end users.
func (e MetastoreError) ErrorUserMessage() string {
	return e.Message
}

// ErrorExceptionID returns the exception ID if available.
func (e MetastoreError) ErrorExceptionID() string {
	return e.ExceptionID
}

// StatusCode returns the HTTP status code of the error response.
func (e MetastoreError) StatusCode() int {
	if e.response != nil {
		return e.response.StatusCode
	}
	return e.HTTPStatus
}

// ResponseHeader returns the headers of the error response, nil when there was none. The Metastore marks some refusals
// with a header (for example the refusal of the ServiceAccount step-up).
func (e MetastoreError) ResponseHeader() http.Header {
	if e.response == nil {
		return nil
	}
	return e.response.Header
}

// SetRequest sets the HTTP request that caused the error.
func (e *MetastoreError) SetRequest(request *http.Request) {
	e.request = request
}

// SetResponse sets the HTTP response that contains the error.
func (e *MetastoreError) SetResponse(response *http.Response) {
	e.response = response
}
