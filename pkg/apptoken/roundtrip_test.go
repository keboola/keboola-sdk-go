package apptoken_test

import (
	"cmp"
	"crypto/ecdsa"
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/apptoken"
)

const (
	issuer  = "https://hub.example.com"
	testKID = "2026-10"
)

// sign writes a token the way apps-proxy does.
func sign(t *testing.T, key *ecdsa.PrivateKey, kid string, claims apptoken.Claims) string {
	t.Helper()

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["typ"] = apptoken.TokenType
	token.Header["kid"] = kid
	raw, err := token.SignedString(key)
	require.NoError(t, err)

	return raw
}

// verify reads a token the way a verifier does: the keys are the published ones and a key id that is not among them is an
// error, never a nil key.
func verify(t *testing.T, key *ecdsa.PrivateKey, raw string, options ...jwt.ParserOption) (apptoken.Claims, error) {
	t.Helper()

	keys, err := apptoken.ParseJWKS(publishedJWKS(t, map[string]*ecdsa.PrivateKey{testKID: key}))
	require.NoError(t, err)

	var got apptoken.Claims
	options = append([]jwt.ParserOption{
		jwt.WithValidMethods([]string{apptoken.Algorithm}), jwt.WithIssuer(issuer), jwt.WithExpirationRequired(),
		jwt.WithAudience(apptoken.Audience), jwt.WithSubject("user-1"), jwt.WithIssuedAt(),
	}, options...)
	_, err = jwt.NewParser(options...).ParseWithClaims(raw, &got, func(token *jwt.Token) (any, error) {
		if token.Header["typ"] != apptoken.TokenType {
			return nil, errors.New("not an app user token")
		}
		id, _ := token.Header["kid"].(string)
		public, ok := keys[id]
		if !ok {
			return nil, errors.New("unknown key id")
		}

		return public, nil
	})

	return got, err
}

func validClaims(now time.Time) apptoken.Claims {
	return apptoken.Claims{
		Issuer: issuer, Audience: apptoken.Audience, Subject: "user-1", IDP: "https://login.example.com",
		Email: "user@example.com", Groups: []string{"sales", "finance"},
		Kbc: apptoken.KbcClaims{ProjectID: "512", AppID: "9", WorkspaceID: "77"},
		ID:  "jti-1", IssuedAt: jwt.NewNumericDate(now), ExpiresAt: jwt.NewNumericDate(now.Add(5 * time.Minute)),
	}
}

// The signing side writes a token and publishes its key, the verifying side reads both. This is the test that fails when the two
// drift apart: it signs the way apps-proxy does and parses the keys the way apps-proxy publishes them.
func TestRoundTrip_SignedByTheProxyVerifiedByTheReader(t *testing.T) {
	t.Parallel()

	key := newKey(t)
	claims := validClaims(time.Now().Truncate(time.Second))

	got, err := verify(t, key, sign(t, key, testKID, claims))

	require.NoError(t, err)
	assert.Equal(t, claims, got)
}

// The claims are what a verifier checks, through the jwt.Claims methods of the type: a token that is not for the platform,
// that is not yet issued, that has expired or that is signed by another key is refused.
func TestRoundTrip_WhatTheVerifierRefuses(t *testing.T) {
	t.Parallel()

	now := time.Now().Truncate(time.Second)
	key := newKey(t)

	for name, tc := range map[string]struct {
		claims func(*apptoken.Claims)
		signer *ecdsa.PrivateKey
		kid    string
		want   error
	}{
		"another audience": {claims: func(c *apptoken.Claims) { c.Audience = "other" }, want: jwt.ErrTokenInvalidAudience},
		"another subject":  {claims: func(c *apptoken.Claims) { c.Subject = "user-2" }, want: jwt.ErrTokenInvalidSubject},
		"another issuer":   {claims: func(c *apptoken.Claims) { c.Issuer = "https://hub.other.example.com" }, want: jwt.ErrTokenInvalidIssuer},
		"issued in the future": {
			claims: func(c *apptoken.Claims) { c.IssuedAt = jwt.NewNumericDate(now.Add(time.Hour)) },
			want:   jwt.ErrTokenUsedBeforeIssued,
		},
		"expired": {
			claims: func(c *apptoken.Claims) { c.ExpiresAt = jwt.NewNumericDate(now.Add(-time.Minute)) },
			want:   jwt.ErrTokenExpired,
		},
		"without an expiry": {claims: func(c *apptoken.Claims) { c.ExpiresAt = nil }, want: jwt.ErrTokenRequiredClaimMissing},
		"another key":       {signer: newKey(t), want: jwt.ErrTokenSignatureInvalid},
		"an unknown key id": {kid: "2026-09", want: jwt.ErrTokenUnverifiable},
	} {
		claims := validClaims(now)
		if tc.claims != nil {
			tc.claims(&claims)
		}
		signer := key
		if tc.signer != nil {
			signer = tc.signer
		}

		signedKID := cmp.Or(tc.kid, testKID)

		_, err := verify(t, key, sign(t, signer, signedKID, claims))

		require.ErrorIs(t, err, tc.want, name)
	}
}
