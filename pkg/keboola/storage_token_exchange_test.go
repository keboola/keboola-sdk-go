package keboola

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

// resolverReset is a scripted resolver reply that closes the connection without
// an HTTP response, which the client sees as a transport error.
const resolverReset = -1

// resolverOK is the body of a successful resolver reply.
const resolverOK = `{"storageToken": "legacy-token", "projectId": 123, "tokenDetail": {"id": "456"}}`

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
	exchanger.retryBaseDelay = time.Millisecond

	return exchanger
}

// newScriptedResolver answers the n-th call with replies[n] (the last reply
// repeats), either a status or resolverReset, and counts the calls. A status of
// 200 carries resolverOK.
func newScriptedResolver(t *testing.T, replies []int, calls *atomic.Int32) *httptest.Server {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := int(calls.Add(1)) - 1
		reply := replies[min(n, len(replies)-1)]

		if reply == resolverReset {
			conn, _, err := w.(http.Hijacker).Hijack()
			require.NoError(t, err)
			require.NoError(t, conn.Close())

			return
		}
		if reply == http.StatusOK {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(resolverOK))

			return
		}
		w.WriteHeader(reply)
	}))
	t.Cleanup(server.Close)

	return server
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
// 502, never echoes the response body, and is not retried.
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

			calls := &atomic.Int32{}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls.Add(1)
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(responseBody))
			}))
			defer server.Close()

			result, err := newTestExchanger(t, server).Exchange(context.Background(), "kbc_at_secret", 123)

			assert.Equal(t, int32(1), calls.Load())
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

// A connection URL with a trailing slash must not produce a double slash in
// the resolver request path.
func TestStorageTokenExchanger_TrailingSlashURL(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/manage/internal/auth-bridge/resolve-storage-token", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"storageToken": "legacy-token", "tokenDetail": {"id": "456"}}`))
	}))
	defer server.Close()

	saTokenPath := t.TempDir() + "/token"
	require.NoError(t, os.WriteFile(saTokenPath, []byte("sa-jwt"), 0o600))

	exchanger, err := NewStorageTokenExchanger(
		server.URL+"/",
		management.NewKeboolaServiceAccountAuth(saTokenPath),
		WithExchangerHTTPClient(server.Client()),
	)
	require.NoError(t, err)

	result, err := exchanger.Exchange(context.Background(), "kbc_at_secret", 123)
	require.NoError(t, err)
	assert.Equal(t, "legacy-token", result.StorageToken)
}

func TestNewStorageTokenExchanger_Validation(t *testing.T) {
	t.Parallel()

	_, err := NewStorageTokenExchanger("", management.NewKeboolaServiceAccountAuth(""))
	require.ErrorContains(t, err, "connection URL must not be empty")

	_, err = NewStorageTokenExchanger("https://connection.keboola.com", nil)
	require.ErrorContains(t, err, "auth must not be nil")
}

// A transient resolver failure (502, 503, 504, a dropped connection) is retried
// up to exchangeMaxTries calls; any other failure is answered by the first call.
func TestStorageTokenExchanger_Retry(t *testing.T) {
	t.Parallel()

	ok := http.StatusOK
	cases := map[string]struct {
		replies        []int
		expectedCalls  int32
		expectedStatus int // resolver status of the error, 0 for a success or a transport error
		expectedOK     bool
	}{
		"503 then ok":              {replies: []int{503, ok}, expectedCalls: 2, expectedOK: true},
		"502 then ok":              {replies: []int{502, ok}, expectedCalls: 2, expectedOK: true},
		"504 then ok":              {replies: []int{504, ok}, expectedCalls: 2, expectedOK: true},
		"connection reset then ok": {replies: []int{resolverReset, ok}, expectedCalls: 2, expectedOK: true},
		"two failures then ok":     {replies: []int{503, 502, ok}, expectedCalls: 3, expectedOK: true},
		"503 forever":              {replies: []int{503}, expectedCalls: exchangeMaxTries, expectedStatus: 503},
		"connection reset forever": {replies: []int{resolverReset}, expectedCalls: exchangeMaxTries},
		"400 is not retried":       {replies: []int{400, ok}, expectedCalls: 1, expectedStatus: 400},
		"401 is not retried":       {replies: []int{401, ok}, expectedCalls: 1, expectedStatus: 401},
		"403 is not retried":       {replies: []int{403, ok}, expectedCalls: 1, expectedStatus: 403},
		"404 is not retried":       {replies: []int{404, ok}, expectedCalls: 1, expectedStatus: 404},
		"500 is not retried":       {replies: []int{500, ok}, expectedCalls: 1, expectedStatus: 500},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			calls := &atomic.Int32{}
			server := newScriptedResolver(t, tc.replies, calls)

			result, err := newTestExchanger(t, server).Exchange(context.Background(), "kbc_at_secret", 123)

			assert.Equal(t, tc.expectedCalls, calls.Load())
			if tc.expectedOK {
				require.NoError(t, err)
				assert.Equal(t, "legacy-token", result.StorageToken)

				return
			}
			assert.Nil(t, result)
			exchangeErr := &StorageTokenExchangeError{}
			require.ErrorAs(t, err, &exchangeErr)
			assert.Equal(t, tc.expectedStatus, exchangeErr.StatusCode)
		})
	}
}

// A context that ends before or while the exchange waits for the next attempt
// stops the retry at once and still comes back as a StorageTokenExchangeError.
func TestStorageTokenExchanger_RetryStopsOnCancelledContext(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		cancelBeforeCall bool
		expectedCalls    int32
	}{
		"cancelled before the call":        {cancelBeforeCall: true, expectedCalls: 0},
		"cancelled while waiting to retry": {expectedCalls: 1},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			calls := &atomic.Int32{}
			server := newScriptedResolver(t, []int{503}, calls)
			exchanger := newTestExchanger(t, server)
			// A long wait between attempts: the exchange must not sit it out.
			exchanger.retryBaseDelay = 10 * time.Second
			if tc.cancelBeforeCall {
				cancel()
			} else {
				time.AfterFunc(50*time.Millisecond, cancel)
			}

			start := time.Now()
			result, err := exchanger.Exchange(ctx, "kbc_at_secret", 123)

			assert.Nil(t, result)
			assert.Equal(t, tc.expectedCalls, calls.Load())
			assert.Less(t, time.Since(start), 5*time.Second)
			exchangeErr := &StorageTokenExchangeError{}
			require.ErrorAs(t, err, &exchangeErr)
			require.ErrorIs(t, err, context.Canceled)
		})
	}
}

// The default delays between attempts are 100 ms, then 200 ms, without jitter.
func TestStorageTokenExchanger_DefaultBackOff(t *testing.T) {
	t.Parallel()

	b := (&StorageTokenExchanger{}).newBackOff()

	assert.Equal(t, 100*time.Millisecond, b.NextBackOff())
	assert.Equal(t, 200*time.Millisecond, b.NextBackOff())
}
