package management

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

// rotatingAuth hands out a new token on every call, like a projected ServiceAccount token that is rotated.
type rotatingAuth struct {
	calls atomic.Int32
}

func (a *rotatingAuth) AuthHeaders() (map[string]string, error) {
	return map[string]string{HeaderManageAPIToken: "token-" + string(rune('0'+a.calls.Add(1)))}, nil
}

// With WithRetry the client sends a failed request again with the same body and with the Auth headers resolved again.
func TestNewAPIClientWithAuth_WithRetry(t *testing.T) {
	t.Parallel()

	var bodies, tokens []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		bodies = append(bodies, string(body))
		tokens = append(tokens, r.Header.Get(HeaderManageAPIToken))
		if len(bodies) == 1 {
			w.WriteHeader(http.StatusServiceUnavailable)

			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.HTTPClient = server.Client()
	api, err := NewAPIClientWithAuth(cfg, &rotatingAuth{}, WithRetry(request.TestingRetry()))
	require.NoError(t, err)

	req, err := http.NewRequestWithContext(t.Context(), http.MethodPost, server.URL, strings.NewReader(`{"foo":"bar"}`))
	require.NoError(t, err)
	resp, err := api.GetConfig().HTTPClient.Do(req)
	require.NoError(t, err)
	_ = resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, []string{`{"foo":"bar"}`, `{"foo":"bar"}`}, bodies, "the body is sent again")
	assert.Equal(t, []string{"token-1", "token-2"}, tokens, "the Auth headers are resolved for every attempt")
}

// A client built without WithRetry sends a request once: the Manage API has operations that change data.
func TestNewAPIClientWithAuth_NoRetryByDefault(t *testing.T) {
	t.Parallel()

	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.HTTPClient = server.Client()
	api, err := NewAPIClientWithAuth(cfg, &rotatingAuth{})
	require.NoError(t, err)

	req, err := http.NewRequestWithContext(t.Context(), http.MethodPost, server.URL, nil)
	require.NoError(t, err)
	resp, err := api.GetConfig().HTTPClient.Do(req)
	require.NoError(t, err)
	_ = resp.Body.Close()

	assert.Equal(t, http.StatusServiceUnavailable, resp.StatusCode)
	assert.Equal(t, int32(1), calls.Load())
}
