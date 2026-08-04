package management

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// retryRecorder records the backoff delay of each retry, and runs an optional
// action to simulate the kubelet finishing a rotation between two reads.
type retryRecorder struct {
	delays []time.Duration
	onCall func(call int)
}

func (v *retryRecorder) notify(_ error, delay time.Duration) {
	v.delays = append(v.delays, delay)
	if v.onCall != nil {
		v.onCall(len(v.delays))
	}
}

// newTestAuth returns an auth with the retry delays shortened, so a test that
// exercises the retry does not wait for the production backoff.
func newTestAuth(tokenPath string, recorder *retryRecorder) *KeboolaServiceAccountAuth {
	return &KeboolaServiceAccountAuth{
		tokenPath: tokenPath,
		baseDelay: time.Millisecond,
		onRetry:   recorder.notify,
	}
}

func TestKeboolaServiceAccountAuth_RetriesMissingFile(t *testing.T) {
	t.Parallel()

	// The token path resolution can hit ENOENT while the kubelet swaps the
	// "..data" symlink, even though the token is there before and after.
	tokenPath := filepath.Join(t.TempDir(), "token")
	recorder := &retryRecorder{onCall: func(call int) {
		if call == 2 {
			require.NoError(t, os.WriteFile(tokenPath, []byte("rotated-token\n"), 0o600))
		}
	}}

	headers, err := newTestAuth(tokenPath, recorder).AuthHeaders()
	require.NoError(t, err)
	assert.Equal(t, map[string]string{HeaderKubernetesAuthorization: "Bearer rotated-token"}, headers)
	assert.Len(t, recorder.delays, 2)
}

func TestKeboolaServiceAccountAuth_RetriesEmptyFile(t *testing.T) {
	t.Parallel()

	tokenPath := filepath.Join(t.TempDir(), "token")
	require.NoError(t, os.WriteFile(tokenPath, []byte("  \n"), 0o600))
	recorder := &retryRecorder{onCall: func(int) {
		require.NoError(t, os.WriteFile(tokenPath, []byte("late-token"), 0o600))
	}}

	headers, err := newTestAuth(tokenPath, recorder).AuthHeaders()
	require.NoError(t, err)
	assert.Equal(t, "Bearer late-token", headers[HeaderKubernetesAuthorization])
	assert.Len(t, recorder.delays, 1)
}

func TestKeboolaServiceAccountAuth_GivesUpAfterMaxTries(t *testing.T) {
	t.Parallel()

	tokenPath := filepath.Join(t.TempDir(), "missing")
	recorder := &retryRecorder{}

	_, err := newTestAuth(tokenPath, recorder).AuthHeaders()
	require.Error(t, err)
	assert.ErrorContains(t, err, "failed to read service account token file")
	assert.ErrorContains(t, err, tokenPath)

	// The default budget is 4 reads, so 3 retries.
	assert.Len(t, recorder.delays, defaultTokenReadTries-1)
}

// TestKeboolaServiceAccountAuth_NoRetryOnPermanentError checks that a real
// misconfiguration fails immediately - a wrong path or wrong permissions is not
// going to fix itself, and a proxy must not spend the backoff on every request.
func TestKeboolaServiceAccountAuth_NoRetryOnPermanentError(t *testing.T) {
	t.Parallel()

	// A directory instead of a file fails with EISDIR on read, regardless of the
	// user the tests run as.
	recorder := &retryRecorder{}

	_, err := newTestAuth(t.TempDir(), recorder).AuthHeaders()
	require.Error(t, err)
	assert.ErrorContains(t, err, "failed to read service account token file")
	assert.Empty(t, recorder.delays)
}

// TestKeboolaServiceAccountAuth_BackOffDelays checks the production delays, which
// the tests above shorten. The sequence doubles and is capped, so the default
// budget of 4 reads spends at most 20+40+80 = 140 ms.
func TestKeboolaServiceAccountAuth_BackOffDelays(t *testing.T) {
	t.Parallel()

	b := NewKeboolaServiceAccountAuth("/token").newBackOff()

	var delays []time.Duration
	for range 6 {
		delays = append(delays, b.NextBackOff())
	}

	assert.Equal(t, []time.Duration{
		20 * time.Millisecond,
		40 * time.Millisecond,
		80 * time.Millisecond,
		160 * time.Millisecond,
		maxTokenReadDelay,
		maxTokenReadDelay,
	}, delays)
}

// TestKeboolaServiceAccountAuth_ReadsTokenOnFirstAttempt checks the common case:
// a readable token costs a single read and no backoff.
func TestKeboolaServiceAccountAuth_ReadsTokenOnFirstAttempt(t *testing.T) {
	t.Parallel()

	tokenPath := filepath.Join(t.TempDir(), "token")
	require.NoError(t, os.WriteFile(tokenPath, []byte("k8s-token\n"), 0o600))
	recorder := &retryRecorder{}

	headers, err := newTestAuth(tokenPath, recorder).AuthHeaders()
	require.NoError(t, err)
	assert.Equal(t, "Bearer k8s-token", headers[HeaderKubernetesAuthorization])
	assert.Empty(t, recorder.delays)
}
