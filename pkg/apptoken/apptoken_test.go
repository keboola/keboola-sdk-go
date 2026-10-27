package apptoken_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/apptoken"
)

var _ jwt.Claims = apptoken.Claims{}

func TestValidKID(t *testing.T) {
	t.Parallel()

	for kid, want := range map[string]bool{
		"2026-10":  true,
		"a.b_c-D9": true,
		"k":        true,
		"0123456789012345678901234567890123456789012345678901234567890123":  true,  // 64 characters
		"01234567890123456789012345678901234567890123456789012345678901234": false, // 65 characters
		"":           false,
		"with space": false,
		"slash/kid":  false,
		"žluťoučký":  false,
		"new\nline":  false,
	} {
		assert.Equal(t, want, apptoken.ValidKID(kid), kid)
	}
}

func TestClaims_WireShape(t *testing.T) {
	t.Parallel()

	issued := jwt.NewNumericDate(time.Unix(1_700_000_000, 0))
	expires := jwt.NewNumericDate(time.Unix(1_700_000_300, 0))

	t.Run("a full token", func(t *testing.T) {
		t.Parallel()

		raw, err := json.Marshal(apptoken.Claims{
			Issuer: "https://hub.example.com", Audience: apptoken.Audience, Subject: "user-1", IDP: "https://login.example.com",
			Email: "user@example.com", Groups: []string{"sales"},
			Kbc: apptoken.KbcClaims{ProjectID: "512", AppID: "9", WorkspaceID: "77"},
			ID:  "jti-1", SID: "sid-1", IssuedAt: issued, ExpiresAt: expires,
		})

		require.NoError(t, err)
		assert.JSONEq(t, `{"iss":"https://hub.example.com","aud":"keboola","sub":"user-1","idp":"https://login.example.com",
			"email":"user@example.com","groups":["sales"],"kbc":{"project_id":"512","app_id":"9","workspace_id":"77"},
			"jti":"jti-1","sid":"sid-1","iat":1700000000,"exp":1700000300}`, string(raw), "aud is a string, ids are decimal strings")
	})

	t.Run("the optional claims are left out", func(t *testing.T) {
		t.Parallel()

		raw, err := json.Marshal(apptoken.Claims{
			Audience: apptoken.Audience, Groups: []string{}, Kbc: apptoken.KbcClaims{ProjectID: "512"},
		})

		require.NoError(t, err)
		var got map[string]any
		require.NoError(t, json.Unmarshal(raw, &got))
		for _, name := range []string{"email", "groups", "sid"} {
			assert.NotContains(t, got, name, "a known-empty group list is written as a missing claim, which a reader takes as unknown")
		}
		assert.NotContains(t, got["kbc"], "workspace_id")
	})

	t.Run("aud as an array does not decode", func(t *testing.T) {
		t.Parallel()

		var claims apptoken.Claims
		err := json.Unmarshal([]byte(`{"aud":["keboola"]}`), &claims)

		require.Error(t, err)
	})

	t.Run("a reader tells a missing group list from an empty one", func(t *testing.T) {
		t.Parallel()

		var missing, empty apptoken.Claims
		require.NoError(t, json.Unmarshal([]byte(`{"aud":"keboola"}`), &missing))
		require.NoError(t, json.Unmarshal([]byte(`{"aud":"keboola","groups":[]}`), &empty))

		assert.Nil(t, missing.Groups)
		assert.NotNil(t, empty.Groups)
		assert.Empty(t, empty.Groups)
	})
}
