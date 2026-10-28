// Package apptoken is the contract of the app user token: the JWT that apps-proxy signs for the end user of a Data App, and that
// the app forwards to the services it calls (Query Service, ...) in the X-Kbc-App-User-Token header. The services that sign it and
// the services that verify it share these constants and types, and the parsing of the JWKS in which apps-proxy publishes the keys.
//
// The package holds the contract only. Signing (the private keys) belongs to apps-proxy, and fetching and caching the keys, and
// deciding which identity provider, email or group to trust, belong to the verifying service.
//
// A verifier takes the keys from ParseJWKS (of an address from its configuration, never from the token), pins Algorithm, checks
// the TokenType header, the Issuer, the Audience, requires `exp` and `iat`, and checks that kbc.project_id is the project of the
// request. The golang-jwt parser options that do it are the verifier's: a token without `exp` never expires and a token without
// `aud` passes unless the parser is told to check them.
package apptoken

import (
	"regexp"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// HeaderName is the request header in which a Data App forwards the token.
	HeaderName = "X-Kbc-App-User-Token"
	// TokenType is the `typ` of the JOSE header. A verifier must check it, so that it never takes this token for another JWT.
	TokenType = "kbc-app-user+jwt" //nolint:gosec // G101: the JOSE typ value, not a credential
	// Audience is the `aud` of the token: the platform, not one service. It is a plain string, never an array.
	Audience = "keboola"
	// Algorithm is the only signing algorithm, ECDSA on P-256 with SHA-256.
	Algorithm = "ES256"
	// JWKSPath is the path, on the public URL of apps-proxy, where the public keys are published.
	JWKSPath = "/.well-known/jwks.json"

	// KIDPattern is the shape of a key id that apps-proxy accepts at startup, as a regular expression.
	KIDPattern = `^[A-Za-z0-9._-]{1,64}$`

	// KeyType, Curve and KeyUse are the `kty`, `crv` and `use` of every key in the JWKS.
	KeyType = "EC"
	Curve   = "P-256"
	KeyUse  = "sig"
)

var kidPattern = regexp.MustCompile(KIDPattern)

// Claims are the claims of the token as they are on the wire. It is written by hand, not built on jwt.RegisteredClaims: that type
// accepts `aud` as an array, and verifiers expect the plain string "keboola". A token whose `aud` is an array does not decode.
//
// A group list that is empty is not written (omitempty), so "the user is in no group" and "the groups are unknown" both read as a
// missing claim, which is unknown.
type Claims struct {
	Issuer   string `json:"iss"`
	Audience string `json:"aud"`
	Subject  string `json:"sub"`
	// IDP is the issuer of the identity provider that authenticated the user. `sub` is unique only per IDP.
	IDP string `json:"idp"`
	// Email is present only when the identity provider verified it.
	Email     string           `json:"email,omitempty"`
	Groups    []string         `json:"groups,omitempty"`
	Kbc       KbcClaims        `json:"kbc"`
	ID        string           `json:"jti"`
	SID       string           `json:"sid,omitempty"`
	IssuedAt  *jwt.NumericDate `json:"iat"`
	ExpiresAt *jwt.NumericDate `json:"exp"`
}

// KbcClaims say where the user is acting. The ids are decimal strings, not numbers.
//
//nolint:tagliatelle // the claim names are part of the token contract.
type KbcClaims struct {
	ProjectID   string `json:"project_id"`
	AppID       string `json:"app_id"`
	WorkspaceID string `json:"workspace_id,omitempty"`
}

// ValidKID reports whether kid has the shape of a key id of apps-proxy.
func ValidKID(kid string) bool {
	return kidPattern.MatchString(kid)
}

// The methods below make Claims a jwt.Claims. There is no `nbf` claim.

func (c Claims) GetExpirationTime() (*jwt.NumericDate, error) { return c.ExpiresAt, nil }

func (c Claims) GetIssuedAt() (*jwt.NumericDate, error) { return c.IssuedAt, nil }

func (c Claims) GetNotBefore() (*jwt.NumericDate, error) { return nil, nil } //nolint:nilnil // there is no nbf claim

func (c Claims) GetIssuer() (string, error) { return c.Issuer, nil }

func (c Claims) GetSubject() (string, error) { return c.Subject, nil }

func (c Claims) GetAudience() (jwt.ClaimStrings, error) { return jwt.ClaimStrings{c.Audience}, nil }
