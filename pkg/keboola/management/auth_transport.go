// Hand-written companion to the generated client — not produced by
// openapi-generator, keep when regenerating.
package management

import (
	"errors"
	"net/http"
	"strings"
)

// authTransport injects Auth.AuthHeaders() into every outgoing request.
//
// Note: the generated callAPI dumps the request for debug logging before the
// transport runs, so headers set here never appear in the debug dump and the
// token is never logged.
type authTransport struct {
	base http.RoundTripper // nil falls back to http.DefaultTransport
	auth Auth
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	headers, err := t.auth.AuthHeaders()
	if err != nil {
		return nil, err
	}

	// The RoundTripper contract forbids mutating the supplied request.
	req = req.Clone(req.Context())
	for k, v := range headers {
		// The generated prepareRequest assigns header map keys directly,
		// bypassing canonicalization — drop any case-insensitive duplicates
		// (e.g. from ContextAPIKeys) so the Auth value is the only one sent.
		for existing := range req.Header {
			if strings.EqualFold(existing, k) {
				delete(req.Header, existing)
			}
		}
		req.Header.Set(k, v)
	}

	base := t.base
	if base == nil {
		base = http.DefaultTransport
	}

	return base.RoundTrip(req)
}

// NewAPIClientWithAuth creates an APIClient that resolves authentication
// headers via the given Auth on every request, so rotating credentials
// (e.g. a projected Kubernetes ServiceAccount token) are picked up
// automatically. When an Auth is installed it takes precedence over headers
// set via ContextAPIKeys — use one mechanism or the other, not both.
//
// A nil auth is rejected with an error so a misconfigured client fails at
// construction instead of panicking on the first request.
//
// cfg.HTTPClient is shallow-copied before its transport is wrapped, so a
// shared client (e.g. http.DefaultClient) is never mutated.
func NewAPIClientWithAuth(cfg *Configuration, auth Auth) (*APIClient, error) {
	if auth == nil {
		return nil, errors.New("auth must not be nil")
	}

	httpClient := &http.Client{}
	if cfg.HTTPClient != nil {
		clientCopy := *cfg.HTTPClient
		httpClient = &clientCopy
	}
	httpClient.Transport = &authTransport{base: httpClient.Transport, auth: auth}
	cfg.HTTPClient = httpClient

	return NewAPIClient(cfg), nil
}
