package management

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newAuthBridgeTestClient(t *testing.T, server *httptest.Server, auth Auth) *APIClient {
	t.Helper()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: server.URL}}
	cfg.HTTPClient = server.Client()

	if auth == nil {
		return NewAPIClient(cfg)
	}

	return NewAPIClientWithAuth(cfg, auth)
}

func TestResolveStorageToken_Success(t *testing.T) {
	t.Parallel()

	saTokenPath := filepath.Join(t.TempDir(), "token")
	require.NoError(t, os.WriteFile(saTokenPath, []byte("sa-jwt"), 0o600))

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/manage/internal/auth-bridge/resolve-storage-token", r.URL.Path)
		assert.Equal(t, "Bearer kbc_at_secret", r.Header.Get(HeaderSubjectToken))
		assert.Equal(t, "Bearer sa-jwt", r.Header.Get(HeaderKubernetesAuthorization))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, float64(123), body["projectId"])

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"storageToken": "legacy-token",
			"projectId": 123,
			"tokenId": "456",
			"userId": "789",
			"expiresAt": "2026-06-04T12:30:00+02:00"
		}`))
	}))
	defer server.Close()

	apiClient := newAuthBridgeTestClient(t, server, NewKeboolaServiceAccountAuth(saTokenPath))

	result, httpResponse, err := apiClient.AuthBridgeAPI().
		ResolveStorageToken(context.Background()).
		SubjectToken("kbc_at_secret").
		AuthBridgeStorageTokenResolveRequest(AuthBridgeStorageTokenResolveRequest{ProjectID: 123}).
		Execute()

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
	assert.Equal(t, "legacy-token", result.StorageToken)
	assert.Equal(t, 123, result.ProjectID)
	assert.Equal(t, "456", result.TokenID)
	assert.Equal(t, "789", result.UserID)
	require.NotNil(t, result.ExpiresAt)
	assert.Equal(t, "2026-06-04T12:30:00+02:00", *result.ExpiresAt)
}

func TestResolveStorageToken_BearerPrefixNotDuplicated(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Bearer kbc_pat_secret", r.Header.Get(HeaderSubjectToken))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"storageToken": "t", "projectId": 1, "tokenId": "1", "userId": "1", "expiresAt": null}`))
	}))
	defer server.Close()

	apiClient := newAuthBridgeTestClient(t, server, nil)

	result, _, err := apiClient.AuthBridgeAPI().
		ResolveStorageToken(context.Background()).
		SubjectToken("Bearer kbc_pat_secret").
		AuthBridgeStorageTokenResolveRequest(AuthBridgeStorageTokenResolveRequest{ProjectID: 1}).
		Execute()

	require.NoError(t, err)
	assert.Nil(t, result.ExpiresAt)
}

func TestResolveStorageToken_ErrorStatus(t *testing.T) {
	t.Parallel()

	for _, status := range []int{
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusInternalServerError,
	} {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(status)
		}))

		apiClient := newAuthBridgeTestClient(t, server, nil)

		result, httpResponse, err := apiClient.AuthBridgeAPI().
			ResolveStorageToken(context.Background()).
			SubjectToken("kbc_at_secret").
			AuthBridgeStorageTokenResolveRequest(AuthBridgeStorageTokenResolveRequest{ProjectID: 123}).
			Execute()

		require.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, status, httpResponse.StatusCode)

		server.Close()
	}
}

func TestResolveStorageToken_MissingInputs(t *testing.T) {
	t.Parallel()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: "http://localhost"}}
	apiClient := NewAPIClient(cfg)

	_, _, err := apiClient.AuthBridgeAPI().
		ResolveStorageToken(context.Background()).
		AuthBridgeStorageTokenResolveRequest(AuthBridgeStorageTokenResolveRequest{ProjectID: 123}).
		Execute()
	require.ErrorContains(t, err, "subjectToken is required")

	_, _, err = apiClient.AuthBridgeAPI().
		ResolveStorageToken(context.Background()).
		SubjectToken("kbc_at_secret").
		Execute()
	require.ErrorContains(t, err, "authBridgeStorageTokenResolveRequest is required")
}

func TestRedactSensitiveData_SubjectTokenAndStorageToken(t *testing.T) {
	t.Parallel()

	dump := []byte(
		"POST /x HTTP/1.1\r\n" +
			"X-Subject-Token: Bearer kbc_at_secret\r\n" +
			"\r\n" +
			`{"storageToken": "legacy-secret", "subjectToken": "kbc_pat_secret"}`,
	)

	redacted := string(redactSensitiveData(dump))
	assert.NotContains(t, redacted, "kbc_at_secret")
	assert.NotContains(t, redacted, "legacy-secret")
	assert.NotContains(t, redacted, "kbc_pat_secret")
}
