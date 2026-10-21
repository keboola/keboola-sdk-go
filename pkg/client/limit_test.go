package client_test

import (
	"bytes"
	"compress/gzip"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/client"
	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

func TestWithMaxResponseBytes(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name     string
		limit    int64
		body     string
		tooLarge bool
	}{
		{name: "below the limit", limit: 11, body: "0123456789", tooLarge: false},
		{name: "exactly the limit", limit: 10, body: "0123456789", tooLarge: false},
		{name: "one byte over", limit: 9, body: "0123456789", tooLarge: true},
		{name: "no limit", limit: 0, body: strings.Repeat("x", 1<<16), tooLarge: false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			transport := httpmock.NewMockTransport()
			transport.RegisterResponder(http.MethodGet, "https://example.com/body", httpmock.NewStringResponder(200, tc.body))
			c := client.New().WithTransport(transport).WithRetry(request.TestingRetry()).WithMaxResponseBytes(tc.limit)

			var got string
			_, _, err := request.NewHTTPRequest(c).WithGet("https://example.com/body").WithResult(&got).Send(t.Context())
			if tc.tooLarge {
				var tooLarge *http.MaxBytesError
				require.ErrorAs(t, err, &tooLarge)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.body, got)
		})
	}
}

// The limit is on the decoded body, so a small compressed answer cannot expand past it.
func TestWithMaxResponseBytes_Decoded(t *testing.T) {
	t.Parallel()

	var compressed bytes.Buffer
	zw := gzip.NewWriter(&compressed)
	_, err := zw.Write([]byte(strings.Repeat("a", 100_000)))
	require.NoError(t, err)
	require.NoError(t, zw.Close())
	require.Less(t, compressed.Len(), 1000, "the test needs a small wire body")

	transport := httpmock.NewMockTransport()
	transport.RegisterResponder(http.MethodGet, "https://example.com/zip", func(*http.Request) (*http.Response, error) {
		resp := httpmock.NewBytesResponse(200, compressed.Bytes())
		resp.Header.Set("Content-Encoding", "gzip")
		resp.Header.Set("Content-Type", "text/plain")
		return resp, nil
	})
	c := client.New().WithTransport(transport).WithRetry(request.TestingRetry()).WithMaxResponseBytes(4096)

	var got string
	_, _, err = request.NewHTTPRequest(c).WithGet("https://example.com/zip").WithResult(&got).Send(t.Context())
	var tooLarge *http.MaxBytesError
	require.ErrorAs(t, err, &tooLarge)
}

// The limit holds for an error answer too, not only for a result.
func TestWithMaxResponseBytes_ErrorAnswers(t *testing.T) {
	t.Parallel()

	transport := httpmock.NewMockTransport()
	transport.RegisterResponder(http.MethodGet, "https://example.com/big", httpmock.NewStringResponder(200, strings.Repeat("x", 100)))
	transport.RegisterResponder(http.MethodGet, "https://example.com/err", httpmock.NewStringResponder(500, strings.Repeat("x", 100)))
	c := client.New().WithTransport(transport).WithRetry(request.TestingRetry()).WithMaxResponseBytes(10)

	var got string
	for _, path := range []string{"big", "err"} {
		_, _, err := request.NewHTTPRequest(c).WithGet("https://example.com/" + path).WithResult(&got).WithError(&testError{}).Send(t.Context())
		var tooLarge *http.MaxBytesError
		require.ErrorAs(t, err, &tooLarge, path)
	}
}

// With redirects off the request goes to the first URL only, and a 3xx is an error rather than an empty result. Without
// the option the same request is followed to the second server: that is what the option prevents.
func TestWithoutRedirects(t *testing.T) {
	t.Parallel()

	var hits atomic.Int32
	target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		_, _ = w.Write([]byte("target"))
	}))
	t.Cleanup(target.Close)
	redirector := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, target.URL, http.StatusFound)
	}))
	t.Cleanup(redirector.Close)

	send := func(c client.Client) (string, error) {
		var got string
		_, _, err := request.NewHTTPRequest(c).WithGet(redirector.URL).AndHeader("X-Secret", "token").WithResult(&got).Send(t.Context())
		return got, err
	}

	_, err := send(client.New().WithRetry(request.TestingRetry()).WithoutRedirects())
	require.ErrorIs(t, err, client.ErrRedirectNotFollowed)
	var redirect *client.RedirectError
	require.ErrorAs(t, err, &redirect)
	assert.Equal(t, http.StatusFound, redirect.StatusCode)
	assert.Zero(t, hits.Load(), "the redirect target was not called")

	got, err := send(client.New().WithRetry(request.TestingRetry()))
	require.NoError(t, err)
	assert.Equal(t, "target", got)
	assert.Equal(t, int32(1), hits.Load(), "the default client follows the redirect")
}

// The JSON decoder reports a failed read as text, the cap error must still be the typed one.
func TestWithMaxResponseBytes_JSONResult(t *testing.T) {
	t.Parallel()

	transport := httpmock.NewMockTransport()
	body := `{"foo":"` + strings.Repeat("x", 200) + `"}`
	transport.RegisterResponder(http.MethodGet, "https://example.com/json", func(*http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, body)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	c := client.New().WithTransport(transport).WithRetry(request.TestingRetry()).WithMaxResponseBytes(50)

	var got testStruct
	_, _, err := request.NewHTTPRequest(c).WithGet("https://example.com/json").WithResult(&got).Send(t.Context())
	var tooLarge *http.MaxBytesError
	require.ErrorAs(t, err, &tooLarge)
}

// A body that never ends and never completes a JSON value is cut off at the cap, whatever the decoder does with it.
func TestWithMaxResponseBytes_EndlessJSON(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"foo":[`))
		w.(http.Flusher).Flush()
		chunk := []byte(strings.Repeat(" ", 1024))
		for range 10_000 {
			if _, err := w.Write(chunk); err != nil {
				return
			}
			w.(http.Flusher).Flush()
		}
	}))
	t.Cleanup(srv.Close)
	c := client.New().WithRetry(request.TestingRetry()).WithMaxResponseBytes(16 << 10)

	var got map[string]any
	_, _, err := request.NewHTTPRequest(c).WithGet(srv.URL).WithResult(&got).Send(t.Context())
	var tooLarge *http.MaxBytesError
	require.ErrorAs(t, err, &tooLarge)
}
