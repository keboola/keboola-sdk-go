package request

import (
	"net/http"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v5"
)

// RetriesCount - default retries count.
const RetriesCount = 5

// RequestTimeout - default request timeout.
const RequestTimeout = 30 * time.Second

// RetryWaitTimeStart - default retry interval.
const RetryWaitTimeStart = 100 * time.Millisecond

// RetryWaitTimeMax - default maximum retry interval.
const RetryWaitTimeMax = 3 * time.Second

// RetryConfig configures request retries.
// It can be set at the client level (for all requests) or per-request (to override client defaults).
type RetryConfig struct {
	// Condition defines which responses should be retried.
	// If nil, no retries will be performed.
	Condition RetryCondition
	// Count is the maximum number of retry attempts.
	Count int
	// TotalRequestTimeout is the maximum time allowed for the entire request including retries.
	TotalRequestTimeout time.Duration
	// WaitTimeStart is the initial delay before the first retry.
	WaitTimeStart time.Duration
	// WaitTimeMax is the maximum delay between retries.
	WaitTimeMax time.Duration
}

// RetryCondition defines which responses should be retried.
// It receives the HTTP response and error, and returns true if the request should be retried.
type RetryCondition func(*http.Response, error) bool

// TestingRetry returns a fast retry configuration for use in tests.
// It uses minimal wait times to speed up test execution.
func TestingRetry() RetryConfig {
	v := DefaultRetry()
	v.WaitTimeStart = 1 * time.Millisecond
	v.WaitTimeMax = 1 * time.Millisecond
	return v
}

// DefaultRetry returns a default RetryConfig with sensible defaults.
// It includes retry condition that handles common network and HTTP errors.
func DefaultRetry() RetryConfig {
	return RetryConfig{
		TotalRequestTimeout: RequestTimeout,
		Count:               RetriesCount,
		WaitTimeStart:       RetryWaitTimeStart,
		WaitTimeMax:         RetryWaitTimeMax,
		Condition:           DefaultRetryCondition(),
	}
}

// DefaultRetryCondition returns a retry condition that retries on common network and HTTP errors.
// It retries on network errors (except hostname resolution failures) and specific HTTP status codes
// that indicate temporary failures (timeouts, rate limits, server errors).
func DefaultRetryCondition() RetryCondition {
	return func(response *http.Response, err error) bool {
		// On network errors - except hostname not found
		if response == nil || response.StatusCode == 0 {
			switch {
			case strings.Contains(err.Error(), "No address associated with hostname"):
				return false
			case strings.Contains(err.Error(), "no such host"):
				return false
			default:
				return true
			}
		}

		// On HTTP status codes
		switch response.StatusCode {
		case
			http.StatusRequestTimeout,
			http.StatusConflict,
			http.StatusLocked,
			http.StatusTooManyRequests,
			http.StatusInternalServerError,
			http.StatusBadGateway,
			http.StatusServiceUnavailable,
			http.StatusGatewayTimeout:
			return true
		default:
			return false
		}
	}
}

// NewBackoff returns an exponential backoff for HTTP retries.
// The backoff uses exponential backoff strategy with the configured wait times.
func (c RetryConfig) NewBackoff() backoff.BackOff {
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = c.WaitTimeStart
	b.MaxInterval = c.WaitTimeMax
	b.Multiplier = 2
	b.RandomizationFactor = 0
	b.Reset()
	return b
}
