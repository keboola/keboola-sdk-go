package keboola

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsProgrammaticToken(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		token    string
		expected bool
	}{
		{name: "access token", token: "kbc_at_abc123", expected: true},
		{name: "personal access token", token: "kbc_pat_abc123", expected: true},
		{name: "access token with bearer scheme", token: "Bearer kbc_at_abc123", expected: true},
		{name: "personal access token with bearer scheme", token: "Bearer kbc_pat_abc123", expected: true},
		{name: "lowercase bearer scheme", token: "bearer kbc_at_abc123", expected: true},
		{name: "uppercase bearer scheme", token: "BEARER kbc_pat_abc123", expected: true},
		{name: "surrounding whitespace", token: "  kbc_at_abc123  ", expected: true},
		{name: "bearer scheme with surrounding whitespace", token: "  Bearer   kbc_at_abc123 ", expected: true},
		{name: "legacy storage token", token: "1234-123456-abcdef", expected: false},
		{name: "legacy token with bearer scheme", token: "Bearer eyJhbGciOi", expected: false},
		{name: "empty", token: "", expected: false},
		{name: "prefix only mid-string", token: "xkbc_at_abc", expected: false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, IsProgrammaticToken(tc.token))
		})
	}
}
