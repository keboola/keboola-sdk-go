package keboola

import (
	"fmt"
	"net/http"
)

// NotificationError represents an error from the Notification API.
type NotificationError struct {
	Message     string `json:"error"`
	ErrCode     int    `json:"code"`
	ExceptionID string `json:"exceptionId"`
	request     *http.Request
	response    *http.Response
}

// Error returns the error message with HTTP context.
func (e *NotificationError) Error() string {
	msg := e.Message
	if e.request != nil {
		msg += fmt.Sprintf(`, method: "%s", url: "%s"`, e.request.Method, e.request.URL)
	}
	if e.response != nil {
		msg += fmt.Sprintf(`, httpCode: "%d"`, e.response.StatusCode)
	}
	if e.ErrCode != 0 {
		msg += fmt.Sprintf(`, code: "%d"`, e.ErrCode)
	}
	if len(e.ExceptionID) > 0 {
		msg += fmt.Sprintf(`, exceptionId: "%s"`, e.ExceptionID)
	}
	return msg
}

// ErrorName returns a human-readable name for the error.
func (e NotificationError) ErrorName() string {
	if e.response != nil {
		return fmt.Sprintf("%s", e.response.Status)
	}
	return "notification api error"
}

// ErrorUserMessage returns the error message suitable for end users.
func (e NotificationError) ErrorUserMessage() string {
	return e.Message
}

// ErrorExceptionID returns the exception ID if available.
func (e NotificationError) ErrorExceptionID() string {
	return e.ExceptionID
}

// StatusCode returns the HTTP status code of the error response.
func (e NotificationError) StatusCode() int {
	if e.response != nil {
		return e.response.StatusCode
	}
	return 0
}

// SetRequest sets the HTTP request that caused the error.
func (e *NotificationError) SetRequest(request *http.Request) {
	e.request = request
}

// SetResponse sets the HTTP response that contains the error.
func (e *NotificationError) SetResponse(response *http.Response) {
	e.response = response
}
