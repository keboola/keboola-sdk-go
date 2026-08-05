// Hand-written companion to the generated client — not produced by
// openapi-generator, keep when regenerating.
package management

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v5"
)

// DefaultServiceAccountTokenPath is the projected Kubernetes ServiceAccount
// token mounted by the kbc-stacks chart.
const DefaultServiceAccountTokenPath = "/var/run/secrets/connection.keboola.com/serviceaccount/token" //nolint:gosec // file path, not a credential

// Bounds of the retry that covers a token read racing with a kubelet rotation,
// see KeboolaServiceAccountAuth.readToken. With the defaults the delays are
// 20, 40 and 80 ms, so a read failure is reported after 140 ms at the latest.
const (
	defaultTokenReadTries     = 4
	defaultTokenReadBaseDelay = 20 * time.Millisecond
	maxTokenReadDelay         = 250 * time.Millisecond
)

// Auth resolves the authentication headers to attach to a Manage API request.
// AuthHeaders is called per request so file-backed strategies (e.g. a projected
// Kubernetes ServiceAccount token) pick up rotated tokens automatically.
//
// Use NewAPIClientWithAuth to build an APIClient that applies an Auth to
// every outgoing request.
type Auth interface {
	AuthHeaders() (map[string]string, error)
}

// ManageAPITokenAuth authenticates with a static Keboola Manage API token
// sent in the HeaderManageAPIToken header.
type ManageAPITokenAuth struct {
	token string
}

// NewManageAPITokenAuth creates a ManageAPITokenAuth. It errors if token is empty.
func NewManageAPITokenAuth(token string) (*ManageAPITokenAuth, error) {
	if token == "" {
		return nil, fmt.Errorf("manage API token must not be empty")
	}

	return &ManageAPITokenAuth{token: token}, nil
}

func (a *ManageAPITokenAuth) AuthHeaders() (map[string]string, error) {
	return map[string]string{HeaderManageAPIToken: a.token}, nil
}

// KeboolaServiceAccountAuth authenticates with a projected Kubernetes
// ServiceAccount token sent as "Bearer <token>" in the
// HeaderKubernetesAuthorization header. The file is read on every request so
// kubelet-rotated tokens are picked up automatically.
type KeboolaServiceAccountAuth struct {
	tokenPath string
	// maxTries and baseDelay bound the retry of a transient read, zero values
	// mean the defaults above. onRetry observes each retry. All three are
	// unexported: no caller has asked to tune the retry, and keeping them
	// internal leaves the retry an implementation detail. Tests set them.
	maxTries  uint
	baseDelay time.Duration
	onRetry   backoff.Notify
}

// NewKeboolaServiceAccountAuth creates a KeboolaServiceAccountAuth reading from tokenPath.
// An empty tokenPath falls back to DefaultServiceAccountTokenPath.
func NewKeboolaServiceAccountAuth(tokenPath string) *KeboolaServiceAccountAuth {
	if tokenPath == "" {
		tokenPath = DefaultServiceAccountTokenPath
	}

	return &KeboolaServiceAccountAuth{tokenPath: tokenPath}
}

func (a *KeboolaServiceAccountAuth) AuthHeaders() (map[string]string, error) {
	token, err := a.readToken()
	if err != nil {
		return nil, err
	}

	return map[string]string{HeaderKubernetesAuthorization: "Bearer " + token}, nil
}

// readToken reads the token file, retrying a read that is missing or empty.
//
// The kubelet rotates a projected token by writing a new timestamped directory,
// atomically swapping the "..data" symlink the token path resolves through, and
// then removing the old directory. A read whose path resolution interleaves with
// that removal fails with ENOENT even though the token is there before and after,
// so a short bounded backoff turns that race into a slightly slower request
// instead of a failed one.
//
// Anything else - no permission, path is a directory - is a real
// misconfiguration that is not going to fix itself, so readTokenOnce marks it
// backoff.Permanent and it fails without spending the backoff.
//
// This mirrors KeboolaServiceAccountAuthenticator in php-api-client-base, with
// one deliberate difference: PHP fails fast on an unreadable file and retries an
// empty read, because there the race surfaces through the process-global stat
// cache. Go has no stat cache and os.ReadFile is a single open(2), so here the
// race surfaces as ENOENT and that is what has to be retried. The PHP budget
// (~1.2 s) covers a cached stat that can persist for a whole request; a window
// of one syscall needs far less.
func (a *KeboolaServiceAccountAuth) readToken() (string, error) {
	tries := a.maxTries
	if tries < 1 {
		tries = defaultTokenReadTries
	}

	// Auth.AuthHeaders takes no context, and the retry is bounded by the tries
	// and the delays above, so a background context is enough here.
	return backoff.Retry(
		context.Background(),
		a.readTokenOnce,
		backoff.WithBackOff(a.newBackOff()),
		backoff.WithMaxTries(tries),
		backoff.WithNotify(a.onRetry),
	)
}

// readTokenOnce is a single attempt of readToken.
func (a *KeboolaServiceAccountAuth) readTokenOnce() (string, error) {
	data, err := os.ReadFile(a.tokenPath)
	if err != nil {
		err = fmt.Errorf("failed to read service account token file %q: %w", a.tokenPath, err)
		if errors.Is(err, fs.ErrNotExist) {
			return "", err
		}

		return "", backoff.Permanent(err)
	}

	token := strings.TrimSpace(string(data))
	if token == "" {
		return "", fmt.Errorf("service account token file is empty: %q", a.tokenPath)
	}

	return token, nil
}

// newBackOff returns the delays between the token reads: doubling from baseDelay,
// capped at maxTokenReadDelay, without jitter - the same shape as
// RetryConfig.NewBackoff in pkg/request.
func (a *KeboolaServiceAccountAuth) newBackOff() backoff.BackOff {
	base := a.baseDelay
	if base <= 0 {
		base = defaultTokenReadBaseDelay
	}

	b := backoff.NewExponentialBackOff()
	b.InitialInterval = base
	b.MaxInterval = maxTokenReadDelay
	b.Multiplier = 2
	b.RandomizationFactor = 0
	b.Reset()

	return b
}

// NewAutoAuth selects an Auth strategy automatically: a non-empty token yields
// a ManageAPITokenAuth, an empty token falls back to a KeboolaServiceAccountAuth
// reading the projected ServiceAccount token from DefaultServiceAccountTokenPath.
func NewAutoAuth(token string) (Auth, error) {
	if token != "" {
		return NewManageAPITokenAuth(token)
	}

	return NewKeboolaServiceAccountAuth(""), nil
}

// Ensure both strategies implement Auth.
var (
	_ Auth = (*ManageAPITokenAuth)(nil)
	_ Auth = (*KeboolaServiceAccountAuth)(nil)
)
