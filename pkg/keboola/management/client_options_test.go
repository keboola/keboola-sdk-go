package management

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

// rotatingAuth hands out a new token on every call, like a projected ServiceAccount token that is rotated.
type rotatingAuth struct {
	calls atomic.Int32
}

func (a *rotatingAuth) AuthHeaders() (map[string]string, error) {
	return map[string]string{HeaderManageAPIToken: fmt.Sprintf("token-%d", a.calls.Add(1))}, nil
}

func resolveStorageToken(t *testing.T, api *APIClient) error {
	t.Helper()

	_, resp, err := api.AuthBridgeAPI().
		ResolveStorageToken(t.Context()).
		SubjectToken("kbc_at_secret").
		AuthBridgeStorageTokenResolveRequest(AuthBridgeStorageTokenResolveRequest{ProjectID: 123}).
		Execute()
	if resp != nil {
		_ = resp.Body.Close()
	}

	return err
}

// With WithRetry the client sends a failed request again, with the same body and with the Auth headers resolved again;
// without it the Manage API client sends a request once, as it has operations that change data.
func TestNewAPIClientWithAuth_Retry(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		opts           []ClientOption
		expectedBodies int
		expectedTokens []string
	}{
		"with retry":       {opts: []ClientOption{WithRetry(request.TestingRetry())}, expectedBodies: 2, expectedTokens: []string{"token-1", "token-2"}},
		"no retry default": {expectedBodies: 1, expectedTokens: []string{"token-1"}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var lock sync.Mutex
			var bodies, tokens []string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				body, err := io.ReadAll(r.Body)
				assert.NoError(t, err)
				lock.Lock()
				defer lock.Unlock()
				bodies = append(bodies, string(body))
				tokens = append(tokens, r.Header.Get(HeaderManageAPIToken))
				if len(bodies) == 1 {
					w.WriteHeader(http.StatusServiceUnavailable)

					return
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"storageToken": "legacy-token", "tokenDetail": {"id": "456"}}`))
			}))
			defer server.Close()

			api := newAuthBridgeTestClient(t, server, &rotatingAuth{}, tc.opts...)
			err := resolveStorageToken(t, api)

			assert.Equal(t, tc.expectedBodies == 2, err == nil)
			require.Len(t, bodies, tc.expectedBodies)
			assert.Equal(t, tc.expectedTokens, tokens, "the Auth headers are resolved for every attempt")
			for _, body := range bodies {
				assert.JSONEq(t, `{"projectId": 123}`, body, "the body is sent again")
			}
		})
	}
}

// The TotalRequestTimeout of the retry config bounds a request whose attempt hangs.
func TestNewAPIClientWithAuth_RetryTimeout(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read the body so the server notices that the client closed the connection.
		_, _ = io.Copy(io.Discard, r.Body)
		select {
		case <-r.Context().Done():
		case <-time.After(3 * time.Second):
		}
	}))
	defer server.Close()

	retry := request.TestingRetry()
	retry.TotalRequestTimeout = 100 * time.Millisecond
	api := newAuthBridgeTestClient(t, server, &rotatingAuth{}, WithRetry(retry))

	start := time.Now()
	err := resolveStorageToken(t, api)

	require.Error(t, err)
	assert.Less(t, time.Since(start), 2*time.Second)
}
