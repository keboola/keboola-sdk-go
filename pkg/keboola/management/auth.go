// Hand-written companion to the generated client — not produced by
// openapi-generator, keep when regenerating.
package management

import (
	"fmt"
	"os"
	"strings"
)

// DefaultServiceAccountTokenPath is the projected Kubernetes ServiceAccount
// token mounted by the kbc-stacks chart.
const DefaultServiceAccountTokenPath = "/var/run/secrets/connection.keboola.com/serviceaccount/token" //nolint:gosec // file path, not a credential

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
	data, err := os.ReadFile(a.tokenPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read service account token file %q: %w", a.tokenPath, err)
	}

	token := strings.TrimSpace(string(data))
	if token == "" {
		return nil, fmt.Errorf("service account token file is empty: %q", a.tokenPath)
	}

	return map[string]string{HeaderKubernetesAuthorization: "Bearer " + token}, nil
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
