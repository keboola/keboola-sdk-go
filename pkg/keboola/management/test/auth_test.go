package management

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	management "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

// newAuthTestServer returns a test server that responds with a valid token
// verification payload, plus an accessor for the headers of the last request.
// The accessor is mutex-synchronized with the handler goroutine, so reading
// it after a completed request is race-free.
func newAuthTestServer() (*httptest.Server, func() http.Header) {
	var mu sync.Mutex
	var gotHeader http.Header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		gotHeader = r.Header.Clone()
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"scopes": ["git-repo:manage"]}`))
	}))
	return server, func() http.Header {
		mu.Lock()
		defer mu.Unlock()
		return gotHeader
	}
}

// newAuthTestClient returns an API client with the given Auth pointing at the test server.
func newAuthTestClient(t *testing.T, serverURL string, auth management.Auth) *management.APIClient {
	t.Helper()
	cfg := management.NewConfiguration()
	cfg.Servers = management.ServerConfigurations{{URL: serverURL}}
	client, err := management.NewAPIClientWithAuth(cfg, auth)
	require.NoError(t, err)
	return client
}

func Test_management_Auth_ManageAPITokenHeader(t *testing.T) {
	t.Parallel()

	server, gotHeader := newAuthTestServer()
	defer server.Close()

	auth, err := management.NewManageAPITokenAuth("manage-token")
	require.NoError(t, err)

	resp, httpRes, err := newAuthTestClient(t, server.URL, auth).TokenVerificationAPI.TokenVerification(t.Context()).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, httpRes.StatusCode)
	assert.Equal(t, "manage-token", gotHeader().Get("X-KBC-ManageApiToken"))
	assert.Empty(t, gotHeader().Get("X-Kubernetes-Authorization"))
}

func Test_management_Auth_ManageAPIToken_Empty(t *testing.T) {
	t.Parallel()

	_, err := management.NewManageAPITokenAuth("")
	require.Error(t, err)
	assert.ErrorContains(t, err, "must not be empty")
}

func Test_management_Auth_ServiceAccountHeader(t *testing.T) {
	t.Parallel()

	tokenPath := filepath.Join(t.TempDir(), "token")
	require.NoError(t, os.WriteFile(tokenPath, []byte("k8s-token\n"), 0o600))

	server, gotHeader := newAuthTestServer()
	defer server.Close()

	auth := management.NewKeboolaServiceAccountAuth(tokenPath)

	resp, httpRes, err := newAuthTestClient(t, server.URL, auth).TokenVerificationAPI.TokenVerification(t.Context()).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, httpRes.StatusCode)
	assert.Equal(t, "Bearer k8s-token", gotHeader().Get("X-Kubernetes-Authorization"))
	assert.Empty(t, gotHeader().Get("X-KBC-ManageApiToken"))
}

func Test_management_Auth_ServiceAccountTokenRotation(t *testing.T) {
	t.Parallel()

	tokenPath := filepath.Join(t.TempDir(), "token")
	require.NoError(t, os.WriteFile(tokenPath, []byte("token-a"), 0o600))

	server, gotHeader := newAuthTestServer()
	defer server.Close()

	client := newAuthTestClient(t, server.URL, management.NewKeboolaServiceAccountAuth(tokenPath))

	_, _, err := client.TokenVerificationAPI.TokenVerification(t.Context()).Execute()
	require.NoError(t, err)
	assert.Equal(t, "Bearer token-a", gotHeader().Get("X-Kubernetes-Authorization"))

	// Simulate kubelet rotating the projected token; the next request must
	// pick up the new value without rebuilding the client.
	require.NoError(t, os.WriteFile(tokenPath, []byte("token-b"), 0o600))

	_, _, err = client.TokenVerificationAPI.TokenVerification(t.Context()).Execute()
	require.NoError(t, err)
	assert.Equal(t, "Bearer token-b", gotHeader().Get("X-Kubernetes-Authorization"))
}

func Test_management_Auth_ServiceAccountTokenFileMissing(t *testing.T) {
	t.Parallel()

	server, _ := newAuthTestServer()
	defer server.Close()

	auth := management.NewKeboolaServiceAccountAuth(filepath.Join(t.TempDir(), "missing"))

	_, _, err := newAuthTestClient(t, server.URL, auth).TokenVerificationAPI.TokenVerification(t.Context()).Execute()
	require.Error(t, err)
	assert.ErrorContains(t, err, "failed to read service account token file")
}

func Test_management_Auth_ServiceAccountTokenFileEmpty(t *testing.T) {
	t.Parallel()

	tokenPath := filepath.Join(t.TempDir(), "token")
	require.NoError(t, os.WriteFile(tokenPath, []byte("  \n"), 0o600))

	server, _ := newAuthTestServer()
	defer server.Close()

	auth := management.NewKeboolaServiceAccountAuth(tokenPath)

	_, _, err := newAuthTestClient(t, server.URL, auth).TokenVerificationAPI.TokenVerification(t.Context()).Execute()
	require.Error(t, err)
	assert.ErrorContains(t, err, "service account token file is empty")
}

func Test_management_Auth_AutoAuthSelection(t *testing.T) {
	t.Parallel()

	auth, err := management.NewAutoAuth("manage-token")
	require.NoError(t, err)
	assert.IsType(t, &management.ManageAPITokenAuth{}, auth)

	auth, err = management.NewAutoAuth("")
	require.NoError(t, err)
	assert.IsType(t, &management.KeboolaServiceAccountAuth{}, auth)
}

func Test_management_Auth_DoesNotMutateDefaultClient(t *testing.T) {
	t.Parallel()

	auth, err := management.NewManageAPITokenAuth("manage-token")
	require.NoError(t, err)

	before := http.DefaultClient.Transport
	_, err = management.NewAPIClientWithAuth(management.NewConfiguration(), auth)
	require.NoError(t, err)
	assert.Equal(t, before, http.DefaultClient.Transport)
}

func Test_management_Auth_NilAuth(t *testing.T) {
	t.Parallel()

	client, err := management.NewAPIClientWithAuth(management.NewConfiguration(), nil)
	require.Error(t, err)
	assert.ErrorContains(t, err, "auth must not be nil")
	assert.Nil(t, client)
}

func Test_management_Auth_TakesPrecedenceOverContextAPIKeys(t *testing.T) {
	t.Parallel()

	server, gotHeader := newAuthTestServer()
	defer server.Close()

	auth, err := management.NewManageAPITokenAuth("auth-strategy-token")
	require.NoError(t, err)

	// Set auth headers via ContextAPIKeys too — for the same scheme the Auth
	// uses AND for the other scheme. The Auth transport runs last and must be
	// the only auth mechanism on the wire, as documented on NewAPIClientWithAuth.
	ctx := context.WithValue(t.Context(), management.ContextAPIKeys, map[string]management.APIKey{
		management.AuthSchemeAPIKey: {Key: "context-api-keys-token"},
		management.AuthSchemeK8s:    {Key: "context-api-keys-k8s-token"},
	})

	_, _, err = newAuthTestClient(t, server.URL, auth).TokenVerificationAPI.TokenVerification(ctx).Execute()
	require.NoError(t, err)
	assert.Equal(t, []string{"auth-strategy-token"}, gotHeader().Values(management.HeaderManageAPIToken))
	assert.Empty(t, gotHeader().Values(management.HeaderKubernetesAuthorization))
}
