package keboola_test

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

// json mirrors the configuration the client uses to decode responses, so the test
// exercises the same decoder path that produced the original parsing failure.
var clientJSON = jsoniter.ConfigCompatibleWithStandardLibrary //nolint:gochecknoglobals

func TestColumn_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    string
		wantName string
		wantDef  *ColumnDefinition
	}{
		{
			name:     "untyped column with empty array definition (server regression)",
			input:    `{"name":"inward_issue_id","definition":[]}`,
			wantName: "inward_issue_id",
			wantDef:  nil,
		},
		{
			name:     "untyped column with null definition",
			input:    `{"name":"col","definition":null}`,
			wantName: "col",
			wantDef:  nil,
		},
		{
			name:     "untyped column with empty object definition (server regression)",
			input:    `{"name":"col","definition":{}}`,
			wantName: "col",
			wantDef:  nil,
		},
		{
			name:     "untyped column with all-empty object definition",
			input:    `{"name":"col","definition":{"type":"","length":"","nullable":false,"default":""}}`,
			wantName: "col",
			wantDef:  nil,
		},
		{
			name:     "untyped column without definition field",
			input:    `{"name":"col"}`,
			wantName: "col",
			wantDef:  nil,
		},
		{
			name:     "typed column with object definition",
			input:    `{"name":"id","definition":{"type":"INTEGER","length":"","nullable":false,"default":""}}`,
			wantName: "id",
			wantDef:  &ColumnDefinition{Type: "INTEGER"},
		},
		{
			name:     "typed column with nullable definition",
			input:    `{"name":"name","definition":{"type":"STRING","length":"255","nullable":true,"default":"x"}}`,
			wantName: "name",
			wantDef:  &ColumnDefinition{Type: "STRING", Length: "255", Nullable: true, Default: "x"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var col Column
			require.NoError(t, clientJSON.Unmarshal([]byte(tc.input), &col))
			assert.Equal(t, tc.wantName, col.Name)
			assert.Equal(t, tc.wantDef, col.Definition)
		})
	}
}

// TestColumns_UnmarshalJSON reproduces the exact payload from the bug report: a list of
// untyped columns whose "definition" fields are empty arrays.
func TestColumns_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	input := `[{"name":"inward_issue_id","definition":[]},{"name":"inward_issue_key","definition":[]}]`

	var cols Columns
	require.NoError(t, clientJSON.Unmarshal([]byte(input), &cols))
	require.Len(t, cols, 2)
	assert.Equal(t, []string{"inward_issue_id", "inward_issue_key"}, cols.Names())
	for _, c := range cols {
		assert.Nil(t, c.Definition)
	}
}
