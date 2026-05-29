package client_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"

	. "github.com/keboola/keboola-sdk-go/v2/pkg/client"
	. "github.com/keboola/keboola-sdk-go/v2/pkg/client/trace"
	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

type testStruct struct {
	Foo string `json:"foo"`
}

type testError struct {
	ErrorMsg string `json:"error"`
}

type testWriteCloser struct {
	io.Writer
}

func (v testWriteCloser) Close() error {
	_, err := v.Write([]byte("<CLOSE>"))
	return err
}

func (e testError) Error() string {
	return e.ErrorMsg
}

// responseError is a minimal test error type that implements the client's
// errorWithResponse interface (via SetResponse). Tests for empty-body 4xx
// handling use it to assert that the SDK still populates the typed errDef
// with the HTTP status code.
type responseError struct {
	Message  string `json:"error"`
	response *http.Response
}

func (e *responseError) Error() string {
	if e.Message == "" {
		return http.StatusText(e.StatusCode())
	}
	return e.Message
}

func (e *responseError) SetResponse(r *http.Response) { e.response = r }

func (e *responseError) StatusCode() int {
	if e.response == nil {
		return 0
	}
	return e.response.StatusCode
}

func TestNew(t *testing.T) {
	t.Parallel()
	c := New()
	assert.NotNil(t, c)
}

func TestRequest(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", `https://example.com`, httpmock.NewStringResponder(200, "test"))

	ctx := context.Background()
	c := New().WithTransport(transport).WithRetry(request.TestingRetry())
	_, _, err := request.NewHTTPRequest(c).WithGet("https://example.com").Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET https://example.com"])
}

func TestBytesResult(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", "https://example.com", httpmock.NewJsonResponderOrPanic(200, map[string]any{"foo": "bar"}))

	ctx := context.Background()
	c := New().WithTransport(transport).WithRetry(request.TestingRetry())
	var resultDef []byte
	_, result, err := request.NewHTTPRequest(c).WithGet("https://example.com").WithResult(&resultDef).Send(ctx)
	assert.NoError(t, err)
	assert.Same(t, &resultDef, result)
	assert.Equal(t, []byte(`{"foo":"bar"}`), resultDef)
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET https://example.com"])
}

func TestWriterResult(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", "https://example.com", httpmock.NewJsonResponderOrPanic(200, map[string]any{"foo": "bar"}))

	ctx := context.Background()
	c := New().WithTransport(transport).WithRetry(request.TestingRetry())
	var out strings.Builder
	_, _, err := request.NewHTTPRequest(c).WithGet("https://example.com").WithResult(io.Writer(&out)).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, `{"foo":"bar"}`, out.String())
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET https://example.com"])
}

func TestWriteCloserResult(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", "https://example.com", httpmock.NewJsonResponderOrPanic(200, map[string]any{"foo": "bar"}))

	ctx := context.Background()
	c := New().WithTransport(transport).WithRetry(request.TestingRetry())
	var out strings.Builder
	_, _, err := request.NewHTTPRequest(c).WithGet("https://example.com").WithResult(testWriteCloser{Writer: &out}).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, `{"foo":"bar"}<CLOSE>`, out.String())
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET https://example.com"])
}

func TestJsonMapResult(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", `https://example.com`, httpmock.NewJsonResponderOrPanic(200, map[string]any{"foo": "bar"}))

	ctx := context.Background()
	c := New().WithTransport(transport).WithRetry(request.TestingRetry())
	resultDef := make(map[string]any)
	_, result, err := request.NewHTTPRequest(c).WithGet("https://example.com").WithResult(&resultDef).Send(ctx)
	assert.NoError(t, err)
	assert.Same(t, &resultDef, result)
	assert.Equal(t, &map[string]any{"foo": "bar"}, result)
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET https://example.com"])
}

func TestJsonMapResult_ContentTypeWithCharset(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", `https://example.com`, func(request *http.Request) (*http.Response, error) {
		response := httpmock.NewBytesResponse(200, []byte(`{"foo":"bar"}`))
		response.Header.Set("Content-Type", "application/json; charset=utf-8")
		return response, nil
	})

	ctx := context.Background()
	c := New().WithTransport(transport).WithRetry(request.TestingRetry())
	resultDef := make(map[string]any)
	_, result, err := request.NewHTTPRequest(c).WithGet("https://example.com").WithResult(&resultDef).Send(ctx)
	assert.NoError(t, err)
	assert.Same(t, &resultDef, result)
	assert.Equal(t, &map[string]any{"foo": "bar"}, result)
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET https://example.com"])
}

func TestJsonStructResult(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", `https://example.com`, httpmock.NewJsonResponderOrPanic(200, map[string]any{"foo": "bar"}))

	ctx := context.Background()
	c := New().WithTransport(transport).WithRetry(request.TestingRetry())
	resultDef := &testStruct{}
	_, result, err := request.NewHTTPRequest(c).WithGet("https://example.com").WithResult(resultDef).Send(ctx)
	assert.NoError(t, err)
	assert.Same(t, resultDef, result)
	assert.Equal(t, &testStruct{Foo: "bar"}, result)
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET https://example.com"])
}

func TestJsonErrorResult(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", `https://example.com`, httpmock.NewJsonResponderOrPanic(400, map[string]any{"error": "error message"}))

	ctx := context.Background()
	c := New().WithTransport(transport).WithRetry(request.TestingRetry())
	errDef := &testError{}
	_, _, err := request.NewHTTPRequest(c).WithGet("https://example.com").WithError(errDef).Send(ctx)
	assert.Error(t, err)
	assert.Same(t, errDef, err)
	assert.Equal(t, &testError{ErrorMsg: "error message"}, err)
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET https://example.com"])
}

// TestJsonErrorResult_EmptyBody covers the case where the server returns a 4xx
// status with Content-Type: application/json but an empty body — as Storage API
// does under workspace credentials lock contention (409). The caller must still
// receive the typed errDef with the HTTP status code accessible, instead of a
// generic "cannot decode JSON error: EOF" string that hides the real status.
func TestJsonErrorResult_EmptyBody(t *testing.T) {
	t.Parallel()

	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("POST", "https://example.com/workspaces/credentials", func(*http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(http.StatusConflict, "")
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	ctx := context.Background()
	c := New().WithTransport(transport).WithRetry(request.RetryConfig{
		Count:               0, // Skip retries — we want to observe the empty-body handling, not loop.
		WaitTimeStart:       time.Millisecond,
		WaitTimeMax:         time.Millisecond,
		TotalRequestTimeout: time.Second,
	})
	errDef := &responseError{}
	_, _, err := request.NewHTTPRequest(c).
		WithPost("https://example.com/workspaces/credentials").
		WithError(errDef).
		Send(ctx)

	assert.Error(t, err)
	assert.NotContains(t, err.Error(), "cannot decode JSON error: EOF",
		"empty body must not surface a JSON parse error to the caller")
	assert.Same(t, errDef, err, "typed errDef must be returned even when the body is empty")
	assert.Equal(t, http.StatusConflict, errDef.StatusCode(),
		"errDef must have the HTTP status code populated so callers can branch on it")
}

// TestJsonErrorResult_MalformedBody confirms that a non-empty but invalid JSON
// payload still produces a parse error — only empty/whitespace bodies are
// handled specially.
func TestJsonErrorResult_MalformedBody(t *testing.T) {
	t.Parallel()

	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", "https://example.com", func(*http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(http.StatusBadRequest, "{not valid json")
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	ctx := context.Background()
	c := New().WithTransport(transport).WithRetry(request.RetryConfig{Count: 0, TotalRequestTimeout: time.Second})
	errDef := &responseError{}
	_, _, err := request.NewHTTPRequest(c).WithGet("https://example.com").WithError(errDef).Send(ctx)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot decode JSON error",
		"malformed (non-empty) JSON body must still surface the decode error")
}

func TestWithBaseUrl(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", "https://example.com/baz", httpmock.NewStringResponder(200, "test"))

	ctx := context.Background()
	c := New().WithTransport(transport).WithRetry(request.TestingRetry()).WithBaseURL("https://example.com")
	_, _, err := request.NewHTTPRequest(c).WithGet("baz").Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET https://example.com/baz"])
}

func TestRequestContext(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", `https://example.com`, func(request *http.Request) (*http.Response, error) {
		// Request context should be used by HTTP request
		assert.Equal(t, "testValue", request.Context().Value("testKey"))
		return httpmock.NewStringResponse(200, "test"), nil
	})
	ctx := context.WithValue(context.Background(), "testKey", "testValue") //nolint:staticcheck
	c := New().WithTransport(transport).WithRetry(request.TestingRetry())
	_, _, err := request.NewHTTPRequest(c).WithGet("https://example.com").Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET https://example.com"])
}

func TestDefaultHeaders(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", "https://example.com", func(request *http.Request) (*http.Response, error) {
		assert.Equal(t, http.Header{
			"User-Agent":      []string{"keboola-sdk-go"},
			"Accept-Encoding": []string{"gzip, br"},
		}, request.Header)
		return httpmock.NewStringResponse(200, "test"), nil
	})

	ctx := context.Background()
	c := New().WithTransport(transport).WithRetry(request.TestingRetry())
	_, _, err := request.NewHTTPRequest(c).WithGet("https://example.com").Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET https://example.com"])
}

func TestWithUserAgent(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", `https://example.com`, func(request *http.Request) (*http.Response, error) {
		assert.Equal(t, http.Header{
			"User-Agent":      []string{"my-user-agent"},
			"Accept-Encoding": []string{"gzip, br"},
		}, request.Header)
		return httpmock.NewStringResponse(200, "test"), nil
	})

	ctx := context.Background()
	c := New().WithTransport(transport).WithRetry(request.TestingRetry()).WithUserAgent("my-user-agent")
	_, _, err := request.NewHTTPRequest(c).WithGet("https://example.com").Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET https://example.com"])
}

func TestWithHeader(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", `https://example.com`, func(request *http.Request) (*http.Response, error) {
		assert.Equal(t, http.Header{
			"User-Agent":      []string{"keboola-sdk-go"},
			"Accept-Encoding": []string{"gzip, br"},
			"My-Header":       []string{"my-value"},
		}, request.Header)
		return httpmock.NewStringResponse(200, "test"), nil
	})

	ctx := context.Background()
	c := New().WithTransport(transport).WithRetry(request.TestingRetry()).WithHeader("my-header", "my-value")
	_, _, err := request.NewHTTPRequest(c).WithGet("https://example.com").Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET https://example.com"])
}

func TestWithHeaders(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", `https://example.com`, func(request *http.Request) (*http.Response, error) {
		assert.Equal(t, http.Header{
			"User-Agent":      []string{"keboola-sdk-go"},
			"Accept-Encoding": []string{"gzip, br"},
			"Key1":            []string{"value1"},
			"Key2":            []string{"value2"},
		}, request.Header)
		return httpmock.NewStringResponse(200, "test"), nil
	})

	ctx := context.Background()
	c := New().WithTransport(transport).WithRetry(request.TestingRetry()).WithHeaders(map[string]string{
		"key1": "value1",
		"key2": "value2",
	})
	_, _, err := request.NewHTTPRequest(c).WithGet("https://example.com").Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET https://example.com"])
}

func TestRequestTimeout(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", "https://example.com", func(request *http.Request) (*http.Response, error) {
		time.Sleep(100 * time.Millisecond) // <<<<<<<
		return httpmock.NewStringResponse(504, "test"), nil
	})

	// Create client
	ctx := context.Background()
	c := New().
		WithTransport(transport).
		WithRetry(request.RetryConfig{
			Condition:           request.DefaultRetryCondition(),
			Count:               10,
			TotalRequestTimeout: 5 * time.Millisecond, // <<<<<<<
		})

	// Get
	_, _, err := request.NewHTTPRequest(c).WithGet("https://example.com").Send(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), `request GET "https://example.com" failed: timeout after`)
}

func TestContext_DeadlineExceeded(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", "https://example.com", func(request *http.Request) (*http.Response, error) {
		time.Sleep(100 * time.Millisecond) // <<<<<<<
		return httpmock.NewStringResponse(504, "test"), nil
	})

	// Create client
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(50*time.Millisecond))
	defer cancel()
	c := New().WithTransport(transport)

	wg := request.NewWaitGroup(ctx)
	wg.Send(request.NewHTTPRequest(c).WithGet("https://example.com"))
	err := wg.Wait()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), `request GET "https://example.com" failed: timeout after`)
}

func TestContext_Canceled(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", "https://example.com", func(request *http.Request) (*http.Response, error) {
		time.Sleep(100 * time.Millisecond) // <<<<<<<
		return httpmock.NewStringResponse(504, "test"), nil
	})

	// Create client
	ctx, cancel := context.WithCancel(context.Background())
	c := New().WithTransport(transport)

	wg := request.NewWaitGroup(ctx)
	wg.Send(request.NewHTTPRequest(c).WithGet("https://example.com"))

	time.Sleep(50 * time.Millisecond)
	cancel()

	err := wg.Wait()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), `request GET "https://example.com" failed: canceled after`)
}

func TestStopRetryOnRequestTimeout(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", "https://example.com", httpmock.NewStringResponder(504, "test"))

	// Setup
	var delays []time.Duration

	// Create client
	ctx := context.Background()
	c := New().
		WithTransport(transport).
		WithRetry(request.RetryConfig{
			Condition:           request.DefaultRetryCondition(),
			Count:               10,
			TotalRequestTimeout: 30 * time.Millisecond, // <<<<<<<
			WaitTimeStart:       40 * time.Millisecond, // <<<<<<<
			WaitTimeMax:         40 * time.Millisecond,
		}).
		AndTrace(func(ctx context.Context, _ request.HTTPRequest) (context.Context, *ClientTrace) {
			return ctx, &ClientTrace{
				RetryDelay: func(_ int, delay time.Duration) {
					delays = append(delays, delay)
				},
			}
		})

	// Get
	_, _, err := request.NewHTTPRequest(c).WithGet("https://example.com").Send(ctx)
	assert.Error(t, err)
	assert.Equal(t, `request GET "https://example.com" failed: 504 Gateway Timeout`, err.Error())

	// Check number of requests
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET https://example.com"])

	// Check delays
	assert.Empty(t, delays)
}
