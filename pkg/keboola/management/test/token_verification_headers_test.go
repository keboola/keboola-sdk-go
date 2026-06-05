package management

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	management "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

// newTestClient returns an API client pointing at the given test server.
func newTestClient(serverURL string) *management.APIClient {
	cfg := management.NewConfiguration()
	cfg.Servers = management.ServerConfigurations{{URL: serverURL}}
	return management.NewAPIClient(cfg)
}

func Test_management_TokenVerification_ApiKeyAuthHeader(t *testing.T) {
	t.Parallel()

	server, gotHeader := newAuthTestServer()
	defer server.Close()

	ctx := context.WithValue(context.Background(), management.ContextAPIKeys, map[string]management.APIKey{
		management.AuthSchemeAPIKey: {Key: "manage-token"},
	})

	resp, httpRes, err := newTestClient(server.URL).TokenVerificationAPI.TokenVerification(ctx).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, httpRes.StatusCode)
	assert.Equal(t, "manage-token", gotHeader().Get(management.HeaderManageAPIToken))
	assert.Empty(t, gotHeader().Get(management.HeaderKubernetesAuthorization))
	assert.Equal(t, []interface{}{"git-repo:manage"}, resp.GetScopes())
}

func Test_management_TokenVerification_K8sAuthHeader(t *testing.T) {
	t.Parallel()

	server, gotHeader := newAuthTestServer()
	defer server.Close()

	ctx := context.WithValue(context.Background(), management.ContextAPIKeys, map[string]management.APIKey{
		management.AuthSchemeK8s: {Key: "k8s-token"},
	})

	resp, httpRes, err := newTestClient(server.URL).TokenVerificationAPI.TokenVerification(ctx).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, httpRes.StatusCode)
	assert.Equal(t, "k8s-token", gotHeader().Get(management.HeaderKubernetesAuthorization))
	assert.Empty(t, gotHeader().Get(management.HeaderManageAPIToken))
	assert.Equal(t, []interface{}{"git-repo:manage"}, resp.GetScopes())
}
