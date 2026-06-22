package keboola

import "strings"

// Prefixes of Connection programmatic bearer tokens. Unlike legacy Storage
// tokens, these are sent in the "Authorization: Bearer <token>" header and
// must be exchanged for a legacy Storage token before calling Storage-token
// based APIs — see StorageTokenExchanger.
const (
	// TokenPrefixAccessToken marks a short-lived session access token
	// issued by Connection's programmatic login.
	TokenPrefixAccessToken = "kbc_at_"
	// TokenPrefixPersonalAccessToken marks a long-lived personal access token.
	TokenPrefixPersonalAccessToken = "kbc_pat_" //nolint:gosec // token prefix, not a credential
)

// IsProgrammaticToken reports whether the token is a Connection programmatic
// bearer token (kbc_at_* or kbc_pat_*). An optional "Bearer " scheme prefix
// (case-insensitive) and surrounding whitespace are ignored, so both raw
// tokens and Authorization header values match.
func IsProgrammaticToken(token string) bool {
	token = stripBearerScheme(token)

	return strings.HasPrefix(token, TokenPrefixAccessToken) ||
		strings.HasPrefix(token, TokenPrefixPersonalAccessToken)
}

// stripBearerScheme removes an optional, case-insensitive "Bearer " scheme
// prefix and surrounding whitespace. HTTP auth schemes are case-insensitive
// (RFC 7235), so "Bearer ", "bearer " and "  BEARER  <token>" are all handled.
func stripBearerScheme(token string) string {
	token = strings.TrimSpace(token)

	const scheme = "bearer "
	if len(token) >= len(scheme) && strings.EqualFold(token[:len(scheme)], scheme) {
		return strings.TrimSpace(token[len(scheme):])
	}

	return token
}
