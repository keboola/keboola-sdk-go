package keboola

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

func newTestExchanger(t *testing.T, server *httptest.Server) *StorageTokenExchanger {
	t.Helper()

	saTokenPath := t.TempDir() + "/token"
	require.NoError(t, os.WriteFile(saTokenPath, []byte("sa-jwt"), 0o600))

	exchanger, err := NewStorageTokenExchanger(
		server.URL,
		management.NewKeboolaServiceAccountAuth(saTokenPath),
		WithExchangerHTTPClient(server.Client()),
	)
	require.NoError(t, err)

	return exchanger
}

func TestStorageTokenExchanger_Exchange(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/manage/internal/auth-bridge/resolve-storage-token", r.URL.Path)
		assert.Equal(t, "Bearer kbc_pat_secret", r.Header.Get(management.HeaderSubjectToken))
		assert.Equal(t, "Bearer sa-jwt", r.Header.Get(management.HeaderKubernetesAuthorization))

		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, float64(123), body["projectId"])

		w.Header().Set("Content-Type", "application/json")
		// The resolver returns the full token detail (the same payload as
		// tokens/verify) alongside the legacy token, so a single call yields a
		// ready-to-use Token with no follow-up verify.
		_, _ = w.Write([]byte(`{
			"storageToken": "legacy-token",
			"projectId": 123,
			"tokenId": "456",
			"userId": "789",
			"expiresAt": null,
			"tokenDetail": {
				"id": "456",
				"description": "programmatic token",
				"owner": {"id": 123, "name": "my-project", "features": ["feat-x"]}
			}
		}`))
	}))
	defer server.Close()

	result, err := newTestExchanger(t, server).Exchange(context.Background(), "kbc_pat_secret", 123)

	require.NoError(t, err)
	assert.Equal(t, "legacy-token", result.StorageToken)
	assert.Equal(t, 123, result.ProjectID)
	assert.Equal(t, "456", result.TokenID)
	assert.Equal(t, "789", result.UserID)
	assert.Nil(t, result.ExpiresAt)

	// Token is built directly from the resolver's tokenDetail, with the legacy
	// token set as the token value — no follow-up VerifyTokenRequest needed.
	require.NotNil(t, result.Token)
	assert.Equal(t, "legacy-token", result.Token.Token)
	assert.Equal(t, "456", result.Token.ID)
	assert.Equal(t, "programmatic token", result.Token.Description)
	assert.Equal(t, 123, result.Token.Owner.ID)
	assert.Equal(t, "my-project", result.Token.Owner.Name)
	assert.Contains(t, result.Token.Owner.Features, "feat-x")
}

func TestStorageTokenExchanger_ErrorMapping(t *testing.T) {
	t.Parallel()

	cases := []struct {
		resolverStatus       int
		expectedClientStatus int
	}{
		{resolverStatus: http.StatusBadRequest, expectedClientStatus: http.StatusBadRequest},
		{resolverStatus: http.StatusUnauthorized, expectedClientStatus: http.StatusUnauthorized},
		{resolverStatus: http.StatusForbidden, expectedClientStatus: http.StatusForbidden},
		{resolverStatus: http.StatusNotFound, expectedClientStatus: http.StatusBadGateway},
		{resolverStatus: http.StatusInternalServerError, expectedClientStatus: http.StatusBadGateway},
		{resolverStatus: http.StatusServiceUnavailable, expectedClientStatus: http.StatusBadGateway},
	}

	for _, tc := range cases {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(tc.resolverStatus)
		}))

		result, err := newTestExchanger(t, server).Exchange(context.Background(), "kbc_at_secret", 123)

		assert.Nil(t, result)
		exchangeErr := &StorageTokenExchangeError{}
		require.ErrorAs(t, err, &exchangeErr)
		assert.Equal(t, tc.resolverStatus, exchangeErr.StatusCode)
		assert.Equal(t, tc.expectedClientStatus, exchangeErr.ClientStatusCode())
		assert.NotContains(t, exchangeErr.Error(), "kbc_at_secret")

		server.Close()
	}
}

func TestStorageTokenExchanger_NetworkErrorMapsToBadGateway(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	exchanger := newTestExchanger(t, server)
	server.Close() // connection refused

	result, err := exchanger.Exchange(context.Background(), "kbc_at_secret", 123)

	assert.Nil(t, result)
	exchangeErr := &StorageTokenExchangeError{}
	require.ErrorAs(t, err, &exchangeErr)
	assert.Equal(t, 0, exchangeErr.StatusCode)
	assert.Equal(t, http.StatusBadGateway, exchangeErr.ClientStatusCode())
}

// A 200 response missing the storage token or its detail (e.g. a Connection
// deploy predating keboola/connection#7604) is an our-side incident mapped to
// 502, and never echoes the response body.
func TestStorageTokenExchanger_MissingFieldsMapToBadGateway(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"missing token detail":  `{"storageToken": "legacy-token", "projectId": 123}`,
		"missing storage token": `{"projectId": 123, "tokenDetail": {"id": "456"}}`,
		"empty token detail":    `{"storageToken": "legacy-token", "tokenDetail": {}}`,
	}

	for name, responseBody := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(responseBody))
			}))
			defer server.Close()

			result, err := newTestExchanger(t, server).Exchange(context.Background(), "kbc_at_secret", 123)

			assert.Nil(t, result)
			exchangeErr := &StorageTokenExchangeError{}
			require.ErrorAs(t, err, &exchangeErr)
			assert.Equal(t, http.StatusBadGateway, exchangeErr.ClientStatusCode())
			assert.NotContains(t, exchangeErr.Error(), "legacy-token")
		})
	}
}

func TestStorageTokenExchanger_RejectsNonProgrammaticToken(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("resolver must not be called for a non-programmatic token")
	}))
	defer server.Close()

	result, err := newTestExchanger(t, server).Exchange(context.Background(), "1234-123456-abcdef", 123)

	assert.Nil(t, result)
	exchangeErr := &StorageTokenExchangeError{}
	require.ErrorAs(t, err, &exchangeErr)
	assert.Equal(t, http.StatusUnauthorized, exchangeErr.ClientStatusCode())
	assert.NotContains(t, exchangeErr.Error(), "1234-123456-abcdef")
}

func TestNewStorageTokenExchanger_Validation(t *testing.T) {
	t.Parallel()

	_, err := NewStorageTokenExchanger("", management.NewKeboolaServiceAccountAuth(""))
	require.ErrorContains(t, err, "connection URL must not be empty")

	_, err = NewStorageTokenExchanger("https://connection.keboola.com", nil)
	require.ErrorContains(t, err, "auth must not be nil")
}
