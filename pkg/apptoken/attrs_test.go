package apptoken_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/apptoken"
)

// The limits and the reserved names are a contract between apps-proxy (which builds `attrs`) and the verifying services, so a
// change of any of them has to be a decision.
func TestAttrs_ContractLiterals(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 100, apptoken.MaxAttrNames)
	assert.Equal(t, 32, apptoken.MaxAttrValues)
	assert.Equal(t, 1024, apptoken.MaxAttrValueChars)
	assert.Equal(t, 4096, apptoken.MaxAttrBytes)
	assert.Equal(t, `^[A-Za-z0-9_.:/-]{1,128}$`, apptoken.AttrNamePattern)
	assert.ElementsMatch(t, []string{
		"sub", "email", "email_verified", "groups", "roles", "iss", "aud", "exp", "iat", "nbf", "nonce", "jti", "sid", "idp",
		"kbc", "azp", "acr", "amr", "auth_time", "at_hash", "c_hash", "attrs",
	}, apptoken.ReservedAttrNames())
}

func TestValidAttrName(t *testing.T) {
	t.Parallel()

	for name, want := range map[string]bool{
		"department":                     true,
		"https://example.com/department": true,
		"a.b_c-d:e/f":                    true,
		"X":                              true,
		strings.Repeat("a", 128):         true,
		strings.Repeat("a", 129):         false,
		"":                               false,
		"with space":                     false,
		"new\nline":                      false,
		"žluťoučký":                      false,
		"semi;colon":                     false,
		// Reserved names are refused whatever their case: they would bypass the verified-email and the groups rules.
		"email":          false,
		"Email":          false,
		"GROUPS":         false,
		"sub":            false,
		"email_verified": false,
		"attrs":          false,
		"kbc":            false,
		"at_hash":        false,
		// A name that merely contains a reserved one is fine.
		"email_domain": true,
		"groupsize":    true,
	} {
		assert.Equal(t, want, apptoken.ValidAttrName(name), name)
	}
}

func TestValidateAttrs(t *testing.T) {
	t.Parallel()

	manyNames := make(map[string][]string, apptoken.MaxAttrNames+1)
	hundredNames := make(map[string][]string, apptoken.MaxAttrNames)
	for i := 0; i <= apptoken.MaxAttrNames; i++ {
		name := fmt.Sprintf("n%03d", i)
		manyNames[name] = []string{"v"}
		if i < apptoken.MaxAttrNames {
			hundredNames[name] = []string{"v"}
		}
	}

	tests := []struct {
		name    string
		attrs   map[string][]string
		wantErr string
	}{
		{name: "nil is unknown, not invalid", attrs: nil},
		{name: "empty is unknown, not invalid", attrs: map[string][]string{}},
		{name: "one name one value", attrs: map[string][]string{"department": {"finance"}}},
		{name: "several values", attrs: map[string][]string{"department": {"finance", "hr"}}},
		{name: "duplicate values are tolerated", attrs: map[string][]string{"department": {"hr", "hr"}}},
		{name: "a value of exactly 1024 characters", attrs: map[string][]string{"d": {strings.Repeat("v", 1024)}}},
		{name: "a value of 1024 multi-byte characters", attrs: map[string][]string{"d": {strings.Repeat("\u00e9", 1024)}}},
		{name: "a value of 1024 three-byte characters", attrs: map[string][]string{"d": {strings.Repeat("\u20ac", 1024)}}},
		{name: "32 values", attrs: map[string][]string{"d": manyValues(32)}},
		{name: "100 names", attrs: hundredNames},
		{name: "101 names", attrs: manyNames, wantErr: "too many names"},
		{name: "a bad name", attrs: map[string][]string{"bad name": {"v"}}, wantErr: "bad name"},
		{name: "a reserved name", attrs: map[string][]string{"email": {"a@b.c"}}, wantErr: "reserved name"},
		{name: "a name with no values", attrs: map[string][]string{"department": {}}, wantErr: "no values"},
		{name: "a nil value list", attrs: map[string][]string{"department": nil}, wantErr: "no values"},
		{name: "an empty value", attrs: map[string][]string{"department": {"finance", ""}}, wantErr: "empty value"},
		{name: "33 values", attrs: map[string][]string{"d": manyValues(33)}, wantErr: "too many values"},
		{name: "a value of 1025 characters", attrs: map[string][]string{"d": {strings.Repeat("v", 1025)}}, wantErr: "value too long"},
		{name: "a value of 1025 multi-byte characters", attrs: map[string][]string{"d": {strings.Repeat("\u00e9", 1025)}}, wantErr: "value too long"},
		{name: "a value that is not valid UTF-8", attrs: map[string][]string{"d": {"a\xffb"}}, wantErr: "bad value"},
		{name: "a value with a control character", attrs: map[string][]string{"d": {"a\nb"}}, wantErr: "bad value"},
		{name: "exactly 4096 bytes", attrs: attrsOfSize(t, apptoken.MaxAttrBytes)},
		{name: "4097 bytes", attrs: attrsOfSize(t, apptoken.MaxAttrBytes+1), wantErr: "too large"},
		{name: "4 values of 1024 three-byte characters", attrs: map[string][]string{"d": {
			strings.Repeat("\u20ac", 1024), strings.Repeat("\u20ac", 1024), strings.Repeat("\u20ac", 1024),
			strings.Repeat("\u20ac", 1024)}}, wantErr: "too large"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := apptoken.ValidateAttrs(tt.attrs)

			if tt.wantErr == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorIs(t, err, apptoken.ErrInvalidAttrs)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

// The reason of a refusal is a fixed phrase: it is logged, so it must never carry a claim name or a value.
func TestValidateAttrs_ErrorHasNoNameOrValue(t *testing.T) {
	t.Parallel()

	err := apptoken.ValidateAttrs(map[string][]string{"secret-name": {"secret-value", ""}})

	require.Error(t, err)
	assert.NotContains(t, err.Error(), "secret")
}

func TestAttrsBytes(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0, apptoken.AttrsBytes(nil))
	assert.Equal(t, len(`{"department":["finance"]}`), apptoken.AttrsBytes(map[string][]string{"department": {"finance"}}))
}

func TestClaims_Attrs_WireShape(t *testing.T) {
	t.Parallel()

	t.Run("written when present, omitted when unknown", func(t *testing.T) {
		t.Parallel()

		with, err := json.Marshal(apptoken.Claims{Audience: apptoken.Audience, Attrs: map[string][]string{"department": {"finance"}}})
		require.NoError(t, err)
		assert.Contains(t, string(with), `"attrs":{"department":["finance"]}`)

		without, err := json.Marshal(apptoken.Claims{Audience: apptoken.Audience, Attrs: map[string][]string{}})
		require.NoError(t, err)
		assert.NotContains(t, string(without), "attrs", "an empty map is written as a missing claim, which a reader takes as unknown")
	})

	t.Run("a token without the claim has nil attrs", func(t *testing.T) {
		t.Parallel()

		var claims apptoken.Claims
		require.NoError(t, json.Unmarshal([]byte(`{"aud":"keboola"}`), &claims))

		assert.Nil(t, claims.Attrs)
	})

	t.Run("a wrong shape does not decode", func(t *testing.T) {
		t.Parallel()

		var claims apptoken.Claims
		require.Error(t, json.Unmarshal([]byte(`{"aud":"keboola","attrs":{"department":"finance"}}`), &claims))
	})
}

func manyValues(n int) []string {
	values := make([]string, n)
	for i := range values {
		values[i] = "v"
	}

	return values
}

// attrsOfSize returns attrs whose raw JSON is exactly size bytes: names `n0`, `n1`, ... with one value each, the values as long as the
// limit of a value allows, the last one as long as the remaining size needs.
func attrsOfSize(t *testing.T, size int) map[string][]string {
	t.Helper()

	attrs := map[string][]string{}
	for i := 0; ; i++ {
		name := fmt.Sprintf("n%d", i)
		attrs[name] = []string{"v"}
		room := size - apptoken.AttrsBytes(attrs)
		require.GreaterOrEqual(t, room, 0)
		if room < apptoken.MaxAttrValueChars {
			attrs[name] = []string{strings.Repeat("v", 1+room)}
			require.Equal(t, size, apptoken.AttrsBytes(attrs))

			return attrs
		}
		attrs[name] = []string{strings.Repeat("v", apptoken.MaxAttrValueChars)}
	}
}
