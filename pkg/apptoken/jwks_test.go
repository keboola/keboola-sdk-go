package apptoken_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/go-jose/go-jose/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/apptoken"
)

func newKey(t *testing.T) *ecdsa.PrivateKey {
	t.Helper()

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	return key
}

// publishedJWKS is the document as apps-proxy writes it: go-jose encodes the public keys.
func publishedJWKS(t *testing.T, keys map[string]*ecdsa.PrivateKey) []byte {
	t.Helper()

	set := jose.JSONWebKeySet{}
	for kid, key := range keys {
		set.Keys = append(set.Keys, jose.JSONWebKey{Key: &key.PublicKey, KeyID: kid, Algorithm: apptoken.Algorithm, Use: apptoken.KeyUse})
	}
	raw, err := json.Marshal(set)
	require.NoError(t, err)

	return raw
}

// ecJWK is a key by hand, to break one field of it.
func ecJWK(kid string, key *ecdsa.PublicKey) map[string]string {
	point, err := key.Bytes()
	if err != nil {
		panic(err)
	}

	return map[string]string{
		"kty": apptoken.KeyType, "crv": apptoken.Curve, "use": apptoken.KeyUse, "alg": apptoken.Algorithm, "kid": kid,
		"x": base64.RawURLEncoding.EncodeToString(point[1:33]), "y": base64.RawURLEncoding.EncodeToString(point[33:]),
	}
}

func document(t *testing.T, keys ...map[string]string) []byte {
	t.Helper()

	raw, err := json.Marshal(map[string]any{"keys": keys})
	require.NoError(t, err)

	return raw
}

func TestParseJWKS_WhatApsProxyPublishes(t *testing.T) {
	t.Parallel()

	active, previous := newKey(t), newKey(t)

	keys, err := apptoken.ParseJWKS(publishedJWKS(t, map[string]*ecdsa.PrivateKey{"2026-10": active, "2026-09": previous}))

	require.NoError(t, err)
	require.Len(t, keys, 2)
	assert.True(t, active.PublicKey.Equal(keys["2026-10"]))
	assert.True(t, previous.PublicKey.Equal(keys["2026-09"]))
}

func TestParseJWKS_EmptySet(t *testing.T) {
	t.Parallel()

	keys, err := apptoken.ParseJWKS([]byte(`{"keys":[]}`))

	require.NoError(t, err)
	assert.Empty(t, keys)
}

func TestParseJWKS_NotAJWKS(t *testing.T) {
	t.Parallel()

	for name, body := range map[string]string{
		"html":         `<html>`,
		"keys is text": `{"keys":"none"}`,
		"empty":        ``,
	} {
		_, err := apptoken.ParseJWKS([]byte(body))

		require.Error(t, err, name)
	}
}

func TestParseJWKS_UnusableKeysAreSkippedAndTheOthersStay(t *testing.T) {
	t.Parallel()

	good, other := newKey(t), newKey(t)
	p384, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	require.NoError(t, err)

	with := func(name, value string) map[string]string {
		k := ecJWK("bad", &other.PublicKey)
		k[name] = value
		return k
	}
	offCurve := ecJWK("bad", &other.PublicKey)
	offCurve["y"] = offCurve["x"]
	p384Point, err := p384.PublicKey.Bytes()
	require.NoError(t, err)
	p384JWK := map[string]string{
		"kty": "EC", "crv": "P-384", "use": "sig", "alg": "ES256", "kid": "bad",
		"x": base64.RawURLEncoding.EncodeToString(p384Point[1:49]), "y": base64.RawURLEncoding.EncodeToString(p384Point[49:]),
	}

	for name, bad := range map[string]map[string]string{
		"another key type":       with("kty", "RSA"),
		"another curve":          p384JWK,
		"an encryption key":      with("use", "enc"),
		"another algorithm":      with("alg", "ES384"),
		"no key id":              with("kid", ""),
		"a key id with a slash":  with("kid", "a/b"),
		"a short coordinate":     with("x", "AAAA"),
		"coordinates not base64": with("y", "!!!"),
		"a point off the curve":  offCurve,
	} {
		keys, err := apptoken.ParseJWKS(document(t, bad, ecJWK("good", &good.PublicKey)))

		require.NoError(t, err, name)
		require.Len(t, keys, 1, name)
		assert.True(t, good.PublicKey.Equal(keys["good"]), name)
	}
}

func TestParseJWKS_RepeatedKidKeepsTheFirst(t *testing.T) {
	t.Parallel()

	first, second := newKey(t), newKey(t)

	keys, err := apptoken.ParseJWKS(document(t, ecJWK("k", &first.PublicKey), ecJWK("k", &second.PublicKey)))

	require.NoError(t, err)
	require.Len(t, keys, 1)
	assert.True(t, first.PublicKey.Equal(keys["k"]), "a key id is never taken over by a later key")
}
