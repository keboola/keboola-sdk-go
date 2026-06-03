package management

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	management "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

// newAuthTestServer returns a test server that records the headers of the
// last request and responds with a valid token verification payload.
func newAuthTestServer(headers *http.Header) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*headers = r.Header.Clone()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"scopes": ["git-repo:manage"]}`))
	}))
}

// newAuthTestClient returns an API client with the given Auth pointing at the test server.
func newAuthTestClient(serverURL string, auth management.Auth) *management.APIClient {
	cfg := management.NewConfiguration()
	cfg.Servers = management.ServerConfigurations{{URL: serverURL}}
	return management.NewAPIClientWithAuth(cfg, auth)
}

func Test_management_Auth_ManageAPITokenHeader(t *testing.T) {
	t.Parallel()

	var gotHeader http.Header
	server := newAuthTestServer(&gotHeader)
	defer server.Close()

	auth, err := management.NewManageAPITokenAuth("manage-token")
	require.NoError(t, err)

	resp, httpRes, err := newAuthTestClient(server.URL, auth).TokenVerificationAPI.TokenVerification(t.Context()).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, httpRes.StatusCode)
	assert.Equal(t, "manage-token", gotHeader.Get("X-KBC-ManageApiToken"))
	assert.Empty(t, gotHeader.Get("X-Kubernetes-Authorization"))
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

	var gotHeader http.Header
	server := newAuthTestServer(&gotHeader)
	defer server.Close()

	auth := management.NewKeboolaServiceAccountAuth(tokenPath)

	resp, httpRes, err := newAuthTestClient(server.URL, auth).TokenVerificationAPI.TokenVerification(t.Context()).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, httpRes.StatusCode)
	assert.Equal(t, "Bearer k8s-token", gotHeader.Get("X-Kubernetes-Authorization"))
	assert.Empty(t, gotHeader.Get("X-KBC-ManageApiToken"))
}

func Test_management_Auth_ServiceAccountTokenRotation(t *testing.T) {
	t.Parallel()

	tokenPath := filepath.Join(t.TempDir(), "token")
	require.NoError(t, os.WriteFile(tokenPath, []byte("token-a"), 0o600))

	var gotHeader http.Header
	server := newAuthTestServer(&gotHeader)
	defer server.Close()

	client := newAuthTestClient(server.URL, management.NewKeboolaServiceAccountAuth(tokenPath))

	_, _, err := client.TokenVerificationAPI.TokenVerification(t.Context()).Execute()
	require.NoError(t, err)
	assert.Equal(t, "Bearer token-a", gotHeader.Get("X-Kubernetes-Authorization"))

	// Simulate kubelet rotating the projected token; the next request must
	// pick up the new value without rebuilding the client.
	require.NoError(t, os.WriteFile(tokenPath, []byte("token-b"), 0o600))

	_, _, err = client.TokenVerificationAPI.TokenVerification(t.Context()).Execute()
	require.NoError(t, err)
	assert.Equal(t, "Bearer token-b", gotHeader.Get("X-Kubernetes-Authorization"))
}

func Test_management_Auth_ServiceAccountTokenFileMissing(t *testing.T) {
	t.Parallel()

	var gotHeader http.Header
	server := newAuthTestServer(&gotHeader)
	defer server.Close()

	auth := management.NewKeboolaServiceAccountAuth(filepath.Join(t.TempDir(), "missing"))

	_, _, err := newAuthTestClient(server.URL, auth).TokenVerificationAPI.TokenVerification(t.Context()).Execute()
	require.Error(t, err)
	assert.ErrorContains(t, err, "failed to read service account token file")
}

func Test_management_Auth_ServiceAccountTokenFileEmpty(t *testing.T) {
	t.Parallel()

	tokenPath := filepath.Join(t.TempDir(), "token")
	require.NoError(t, os.WriteFile(tokenPath, []byte("  \n"), 0o600))

	var gotHeader http.Header
	server := newAuthTestServer(&gotHeader)
	defer server.Close()

	auth := management.NewKeboolaServiceAccountAuth(tokenPath)

	_, _, err := newAuthTestClient(server.URL, auth).TokenVerificationAPI.TokenVerification(t.Context()).Execute()
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
	management.NewAPIClientWithAuth(management.NewConfiguration(), auth)
	assert.Equal(t, before, http.DefaultClient.Transport)
}
