package trace_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"

	. "github.com/keboola/keboola-sdk-go/v2/pkg/client"
	. "github.com/keboola/keboola-sdk-go/v2/pkg/client/trace"
	. "github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

func TestTrace(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("POST", `https://example.com/redirect1`, func(request *http.Request) (*http.Response, error) {
		header := make(http.Header)
		header.Set("Location", "https://example.com/redirect2")
		_, err := io.ReadAll(request.Body)
		assert.NoError(t, err)
		return &http.Response{
			StatusCode: http.StatusMovedPermanently,
			Header:     header,
		}, nil
	})
	transport.RegisterResponder("GET", `https://example.com/redirect2`, func(request *http.Request) (*http.Response, error) {
		header := make(http.Header)
		header.Set("Location", "https://example.com/index")
		return &http.Response{
			StatusCode: http.StatusMovedPermanently,
			Header:     header,
		}, nil
	})
	transport.RegisterResponder("GET", `https://example.com/index`, httpmock.ResponderFromMultipleResponses([]*http.Response{
		{StatusCode: http.StatusLocked},
		{StatusCode: http.StatusTooManyRequests},
		{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("OK"))},
	}))

	// Logs for trace testing
	var logs strings.Builder

	// Create client
	ctx := context.Background()
	c := New().
		WithTransport(transport).
		WithRetry(RetryConfig{
			Condition:     DefaultRetryCondition(),
			Count:         3,
			WaitTimeStart: 1 * time.Microsecond,
			WaitTimeMax:   20 * time.Microsecond,
		}).
		AndTrace(func(ctx context.Context, reqDef HTTPRequest) (context.Context, *ClientTrace) {
			s := spew.NewDefaultConfig()
			s.DisablePointerAddresses = true
			s.DisableCapacities = true
			logs.WriteString(fmt.Sprintf("GotRequest        %s %s\n", reqDef.Method(), reqDef.URL()))
			return ctx, &ClientTrace{
				RequestProcessed: func(result any, err error) {
					logs.WriteString(fmt.Sprintf("RequestProcessed  result=%s err=%v\n", strings.TrimSpace(s.Sdump(result)), err))
				},
				HTTPRequestStart: func(request *http.Request) {
					logs.WriteString(fmt.Sprintf("HTTPRequestStart  %s %s\n", request.Method, request.URL))
				},
				HTTPRequestDone: func(response *http.Response, send, received int64, err error) {
					logs.WriteString(fmt.Sprintf(
						"HttpRequestDone   %d %s send=%vB received=%vB err=%v\n",
						response.StatusCode, http.StatusText(response.StatusCode), send, received, err),
					)
				},
				RetryDelay: func(attempt int, delay time.Duration) {
					logs.WriteString(fmt.Sprintf("HttpRequestRetry  attempt=%d delay=%s\n", attempt, delay))
				},
				BodyParseStart: func(response *http.Response) {
					logs.WriteString("BodyParseStart\n")
				},
				BodyParseDone: func(response *http.Response, result any, err error, parseError error) {
					logs.WriteString(fmt.Sprintf("BodyParseDone     result=%v err=%v parseError=%v\n", strings.TrimSpace(s.Sdump(result)), err, parseError))
				},
			}
		})

	// Expected events
	expected := `
GotRequest        POST https://example.com/redirect1
HTTPRequestStart  POST https://example.com/redirect1
HttpRequestDone   301 Moved Permanently send=7B received=0B err=<nil>
HTTPRequestStart  GET https://example.com/redirect2
HttpRequestDone   301 Moved Permanently send=0B received=0B err=<nil>
HTTPRequestStart  GET https://example.com/index
HttpRequestDone   423 Locked send=0B received=0B err=<nil>
HttpRequestRetry  attempt=1 delay=1µs
HTTPRequestStart  GET https://example.com/index
HttpRequestDone   429 Too Many Requests send=0B received=0B err=<nil>
HttpRequestRetry  attempt=2 delay=2µs
HTTPRequestStart  GET https://example.com/index
BodyParseStart
HttpRequestDone   200 OK send=0B received=2B err=<nil>
BodyParseDone     result=(*string)((len=2) "OK") err=<nil> parseError=<nil>
RequestProcessed  result=(*string)((len=2) "OK") err=<nil>
`

	// Test
	str := ""
	_, result, err := NewHTTPRequest(c).WithPost("https://example.com/redirect1").WithBody("my-body").WithResult(&str).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, "OK", *result.(*string))
	assert.Equal(t, strings.TrimLeft(expected, "\n"), logs.String())
}

func TestTrace_Multiple(t *testing.T) {
	t.Parallel()

	// Mocked response
	transport := httpmock.NewMockTransport()
	transport.RegisterResponder("GET", `https://example.com`, httpmock.NewStringResponder(200, "OK"))

	// Logs for trace testing
	var logs strings.Builder

	// Create client
	ctx := context.Background()
	c := New().
		WithTransport(transport).
		WithRetry(TestingRetry()).
		AndTrace(func(ctx context.Context, reqDef HTTPRequest) (context.Context, *ClientTrace) {
			logs.WriteString(fmt.Sprintf("1: GotRequest        %s %s\n", reqDef.Method(), reqDef.URL()))
			return ctx, &ClientTrace{
				RequestProcessed: func(result any, err error) {
					s := spew.NewDefaultConfig()
					s.DisablePointerAddresses = true
					s.DisableCapacities = true
					logs.WriteString(fmt.Sprintf("1: RequestProcessed  result=%s err=%v\n", strings.TrimSpace(s.Sdump(result)), err))
				},
				HTTPRequestStart: func(request *http.Request) {
					logs.WriteString(fmt.Sprintf("1: HTTPRequestStart  %s %s\n", request.Method, request.URL))
				},
				HTTPRequestDone: func(response *http.Response, send, received int64, err error) {
					logs.WriteString(fmt.Sprintf(
						"1: HttpRequestDone   %d %s send=%vB received=%vB err=%v\n",
						response.StatusCode, http.StatusText(response.StatusCode), send, received, err),
					)
				},
			}
		}).
		AndTrace(func(ctx context.Context, reqDef HTTPRequest) (context.Context, *ClientTrace) {
			logs.WriteString(fmt.Sprintf("2: GotRequest        %s %s\n", reqDef.Method(), reqDef.URL()))
			return ctx, &ClientTrace{
				HTTPRequestStart: func(request *http.Request) {
					logs.WriteString(fmt.Sprintf("2: HTTPRequestStart  %s %s\n", request.Method, request.URL))
				},
				HTTPRequestDone: func(response *http.Response, send, received int64, err error) {
					logs.WriteString(fmt.Sprintf(
						"2: HttpRequestDone   %d %s send=%vB received=%vB err=%v\n",
						response.StatusCode, http.StatusText(response.StatusCode), send, received, err),
					)
				},
			}
		}).
		AndTrace(func(ctx context.Context, _ HTTPRequest) (context.Context, *ClientTrace) {
			return ctx, &ClientTrace{
				RequestProcessed: func(result any, err error) {
					s := spew.NewDefaultConfig()
					s.DisablePointerAddresses = true
					s.DisableCapacities = true
					logs.WriteString(fmt.Sprintf("3: RequestProcessed  result=%s err=%v\n", strings.TrimSpace(s.Sdump(result)), err))
				},
				HTTPRequestStart: func(request *http.Request) {
					logs.WriteString(fmt.Sprintf("3: HTTPRequestStart  %s %s\n", request.Method, request.URL))
				},
				HTTPRequestDone: func(response *http.Response, send, received int64, err error) {
					logs.WriteString(fmt.Sprintf(
						"3: HttpRequestDone   %d %s send=%vB received=%vB err=%v\n",
						response.StatusCode, http.StatusText(response.StatusCode), send, received, err),
					)
				},
			}
		})

	// Expected events
	expected := `
1: GotRequest        GET https://example.com
2: GotRequest        GET https://example.com
1: HTTPRequestStart  GET https://example.com
2: HTTPRequestStart  GET https://example.com
3: HTTPRequestStart  GET https://example.com
1: HttpRequestDone   200 OK send=0B received=2B err=<nil>
2: HttpRequestDone   200 OK send=0B received=2B err=<nil>
3: HttpRequestDone   200 OK send=0B received=2B err=<nil>
1: RequestProcessed  result=(*string)((len=2) "OK") err=<nil>
3: RequestProcessed  result=(*string)((len=2) "OK") err=<nil>
`

	// Test
	str := ""
	_, result, err := NewHTTPRequest(c).WithGet("https://example.com").WithResult(&str).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, "OK", *result.(*string))
	assert.Equal(t, strings.TrimLeft(expected, "\n"), logs.String())
}
