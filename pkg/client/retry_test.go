package client_test

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"

	. "github.com/keboola/keboola-sdk-go/v2/pkg/client"
	. "github.com/keboola/keboola-sdk-go/v2/pkg/client/trace"
	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

func TestRetryCount(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", `https://example.com`, httpmock.NewStringResponder(504, "test"))

	// Setup
	retryCount := 10
	var delays []time.Duration

	// Create client
	ctx := context.Background()
	c := New().
		WithTransport(transport).
		WithRetry(request.RetryConfig{
			Condition:     request.DefaultRetryCondition(),
			Count:         retryCount,
			WaitTimeStart: 1 * time.Microsecond,
			WaitTimeMax:   20 * time.Microsecond,
		}).
		AndTrace(func(ctx context.Context, reqDef request.HTTPRequest) (context.Context, *ClientTrace) {
			return ctx, &ClientTrace{
				RetryDelay: func(_ int, delay time.Duration) {
					delays = append(delays, delay)
				},
			}
		})

	// Get
	_, _, err := request.NewHTTPRequest(c).
		WithGet("https://example.com").
		WithOnComplete(func(ctx context.Context, response request.HTTPResponse, err error) error {
			// Check context
			attempt, found := ContextRetryAttempt(response.RawRequest().Context())
			assert.True(t, found)
			assert.Equal(t, retryCount, attempt)
			return err
		}).
		Send(ctx)
	assert.Error(t, err)
	assert.Equal(t, `request GET "https://example.com" failed: 504 Gateway Timeout`, err.Error())

	// Check number of requests
	assert.Equal(t, 1+retryCount, transport.GetCallCountInfo()["GET https://example.com"])

	// Check delays
	assert.Equal(t, []time.Duration{
		1 * time.Microsecond,
		2 * time.Microsecond,
		4 * time.Microsecond,
		8 * time.Microsecond,
		16 * time.Microsecond,
		20 * time.Microsecond,
		20 * time.Microsecond,
		20 * time.Microsecond,
		20 * time.Microsecond,
		20 * time.Microsecond,
	}, delays)
}

func TestRetryBodyRewind(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("POST", `https://example.com`, func(req *http.Request) (*http.Response, error) {
		requestBody, err := io.ReadAll(req.Body)
		assert.NoError(t, err)
		// Each retry attempt must send same body
		assert.Equal(t, `{"foo":"bar"}`, string(requestBody))
		return httpmock.NewStringResponse(502, "retry!"), nil
	})

	// Create client
	ctx := context.Background()
	c := New().
		WithTransport(transport).
		WithRetry(request.TestingRetry())

	// Post
	jsonBody := map[string]any{"foo": "bar"}
	_, _, err := request.NewHTTPRequest(c).WithPost("https://example.com").WithJSONBody(jsonBody).Send(ctx)
	assert.Error(t, err)
	assert.Equal(t, `request POST "https://example.com" failed: 502 Bad Gateway`, err.Error())

	// Check number of requests
	assert.Equal(t, 1+5, transport.GetCallCountInfo()["POST https://example.com"])
}

func TestDoNotRetry(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", "https://example.com", httpmock.NewStringResponder(403, "test"))

	// Setup
	var delays []time.Duration

	// Create client
	ctx := context.Background()
	c := New().
		WithTransport(transport).
		WithRetry(request.RetryConfig{
			Condition:     request.DefaultRetryCondition(),
			Count:         10,
			WaitTimeStart: 1 * time.Microsecond,
			WaitTimeMax:   20 * time.Microsecond,
		}).
		AndTrace(func(ctx context.Context, reqDef request.HTTPRequest) (context.Context, *ClientTrace) {
			return ctx, &ClientTrace{
				RetryDelay: func(_ int, delay time.Duration) {
					delays = append(delays, delay)
				},
			}
		})

	// Get
	_, _, err := request.NewHTTPRequest(c).WithGet("https://example.com").Send(ctx)
	assert.Error(t, err)
	assert.Equal(t, `request GET "https://example.com" failed: 403 Forbidden`, err.Error())

	// Check number of requests
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET https://example.com"])

	// Check delays
	assert.Empty(t, delays)
}

// TestWithRetryOverridesDefault tests that WithRetry overrides the default client retry behavior.
// It verifies that a custom retry config (10 retries) is used instead of the default (5 retries).
func TestWithRetryOverridesDefault(t *testing.T) {
	t.Parallel()

	// Mocked response that fails for first 8 attempts, then succeeds on 9th attempt.
	// This tests that WithRetry(10) allows more retries than the default (5).
	attemptCount := 0
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", `https://example.com`, func(req *http.Request) (*http.Response, error) {
		attemptCount++
		// Fail for first 8 attempts with 504 Gateway Timeout (retryable error)
		if attemptCount < 9 {
			return httpmock.NewStringResponse(504, "Gateway Timeout"), nil
		}
		// Succeed on 9th attempt
		return httpmock.NewStringResponse(200, "Success"), nil
	})

	// Setup: Track retry attempts and delays
	var retryAttempts []int
	var delays []time.Duration

	// Create client with default retry config (5 retries), then override with WithRetry to 10 retries.
	// This demonstrates that WithRetry overrides the default behavior.
	ctx := context.Background()
	c := New().
		WithTransport(transport).
		WithRetry(request.RetryConfig{
			Condition:     request.DefaultRetryCondition(),
			Count:         10, // Override default 5 retries to 10
			WaitTimeStart: 1 * time.Microsecond,
			WaitTimeMax:   20 * time.Microsecond,
		}).
		AndTrace(func(ctx context.Context, reqDef request.HTTPRequest) (context.Context, *ClientTrace) {
			return ctx, &ClientTrace{
				RetryDelay: func(attempt int, delay time.Duration) {
					retryAttempts = append(retryAttempts, attempt)
					delays = append(delays, delay)
				},
			}
		})

	// Send request - should succeed after 8 retries (9th attempt total)
	_, _, err := request.NewHTTPRequest(c).
		WithGet("https://example.com").
		Send(ctx)
	assert.NoError(t, err, "request should succeed after retries")

	// Verify total number of requests: 1 initial + 8 retries = 9 total
	// This proves we used 10 retries (not the default 5), because with default 5 retries,
	// we would have given up after 6 total attempts (1 initial + 5 retries).
	assert.Equal(t, 9, transport.GetCallCountInfo()["GET https://example.com"],
		"should have made 9 requests (1 initial + 8 retries), proving 10 retry limit was used")

	// Verify we retried 8 times
	assert.Equal(t, 8, len(retryAttempts), "should have retried 8 times")
	assert.Equal(t, 8, len(delays), "should have 8 retry delays")

	// Verify retry attempts are sequential (1, 2, 3, ..., 8)
	expectedAttempts := []int{1, 2, 3, 4, 5, 6, 7, 8}
	assert.Equal(t, expectedAttempts, retryAttempts, "retry attempts should be sequential")

	// Verify exponential backoff delays
	expectedDelays := []time.Duration{
		1 * time.Microsecond,  // 1st retry
		2 * time.Microsecond,  // 2nd retry
		4 * time.Microsecond,  // 3rd retry
		8 * time.Microsecond,  // 4th retry
		16 * time.Microsecond, // 5th retry
		20 * time.Microsecond, // 6th retry (capped at max)
		20 * time.Microsecond, // 7th retry (capped at max)
		20 * time.Microsecond, // 8th retry (capped at max)
	}
	assert.Equal(t, expectedDelays, delays, "delays should follow exponential backoff pattern")
}

// TestPerRequestRetryOverride tests that request-level WithRetry overrides client default.
// This is a regression test for the bug where per-request retry config was not being applied
// due to a type mismatch (storing *RetryConfig in context but asserting to RetryConfig).
func TestPerRequestRetryOverride(t *testing.T) {
	t.Parallel()

	// Mocked response that fails for first 8 attempts, then succeeds on 9th attempt.
	attemptCount := 0
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", `https://example.com`, func(req *http.Request) (*http.Response, error) {
		attemptCount++
		// Fail for first 8 attempts with 504 Gateway Timeout (retryable error)
		if attemptCount < 9 {
			return httpmock.NewStringResponse(504, "Gateway Timeout"), nil
		}
		// Succeed on 9th attempt
		return httpmock.NewStringResponse(200, "Success"), nil
	})

	// Setup: Track retry attempts and delays
	var retryAttempts []int
	var delays []time.Duration

	// Create client with default retry config (5 retries)
	ctx := context.Background()
	c := New().
		WithTransport(transport).
		WithRetry(request.RetryConfig{
			Condition:     request.DefaultRetryCondition(),
			Count:         5, // Default client-level retry count
			WaitTimeStart: 1 * time.Microsecond,
			WaitTimeMax:   20 * time.Microsecond,
		}).
		AndTrace(func(ctx context.Context, reqDef request.HTTPRequest) (context.Context, *ClientTrace) {
			return ctx, &ClientTrace{
				RetryDelay: func(attempt int, delay time.Duration) {
					retryAttempts = append(retryAttempts, attempt)
					delays = append(delays, delay)
				},
			}
		})

	// Send request with per-request retry override (10 retries) - should succeed after 8 retries
	httpReq := request.NewHTTPRequest(c).WithGet("https://example.com")
	result := ""
	apiReq := request.NewAPIRequest(&result, httpReq).
		WithRetry(request.RetryConfig{
			Condition:     request.DefaultRetryCondition(),
			Count:         10, // Override to 10 retries for this request
			WaitTimeStart: 1 * time.Microsecond,
			WaitTimeMax:   20 * time.Microsecond,
		})
	_, err := apiReq.Send(ctx)
	assert.NoError(t, err, "request should succeed after retries")

	// Verify total number of requests: 1 initial + 8 retries = 9 total
	// This proves the per-request override worked. If it didn't, we would have only made
	// 6 requests (1 initial + 5 retries from client default) and failed.
	assert.Equal(t, 9, transport.GetCallCountInfo()["GET https://example.com"],
		"should have made 9 requests (1 initial + 8 retries), proving per-request retry override worked")

	// Verify we retried 8 times
	assert.Equal(t, 8, len(retryAttempts), "should have retried 8 times")
}
