package apptoken_test

import (
	"crypto/ecdsa"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/apptoken"
)

// The signing side writes a token and publishes its key, the verifying side reads both. This is the test that fails when the two
// drift apart: it signs the way apps-proxy does and parses the keys the way apps-proxy publishes them.
func TestRoundTrip_SignedByTheProxyVerifiedByTheReader(t *testing.T) {
	t.Parallel()

	const issuer = "https://hub.example.com"
	key := newKey(t)
	issued := time.Now().Truncate(time.Second)
	claims := apptoken.Claims{
		Issuer: issuer, Audience: apptoken.Audience, Subject: "user-1", IDP: "https://login.example.com",
		Email: "user@example.com", Groups: []string{"sales", "finance"},
		Kbc: apptoken.KbcClaims{ProjectID: "512", AppID: "9", WorkspaceID: "77"},
		ID:  "jti-1", IssuedAt: jwt.NewNumericDate(issued), ExpiresAt: jwt.NewNumericDate(issued.Add(5 * time.Minute)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["typ"] = apptoken.TokenType
	token.Header["kid"] = "2026-10"
	raw, err := token.SignedString(key)
	require.NoError(t, err)

	keys, err := apptoken.ParseJWKS(publishedJWKS(t, map[string]*ecdsa.PrivateKey{"2026-10": key}))
	require.NoError(t, err)

	var got apptoken.Claims
	parsed, err := jwt.NewParser(
		jwt.WithValidMethods([]string{apptoken.Algorithm}), jwt.WithIssuer(issuer), jwt.WithExpirationRequired(),
	).ParseWithClaims(raw, &got, func(token *jwt.Token) (any, error) {
		assert.Equal(t, apptoken.TokenType, token.Header["typ"])
		kid, _ := token.Header["kid"].(string)
		require.True(t, apptoken.ValidKID(kid))

		return keys[kid], nil
	})

	require.NoError(t, err)
	assert.True(t, parsed.Valid)
	assert.Equal(t, claims, got)
}
