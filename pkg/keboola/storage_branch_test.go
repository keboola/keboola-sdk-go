package keboola_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/keboola/go-utils/pkg/wildcards"
	"github.com/stretchr/testify/assert"

	. "github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

func TestBranchApiCalls(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := APIClientForAnEmptyProject(t, ctx)

	// Get default branch
	defaultBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, defaultBranch)
	assert.Equal(t, "Main", defaultBranch.Name)
	assert.True(t, defaultBranch.IsDefault)

	// Default branch cannot be created
	assert.PanicsWithError(t, "default branch cannot be created", func() {
		api.CreateBranchRequest(&Branch{
			Name:        "Foo",
			Description: "Foo branch",
			IsDefault:   true,
		}).Send(ctx)
	})

	// Create branch, wait for successful job status
	branchFoo := &Branch{
		Name:        "Foo",
		Description: "Foo branch",
		IsDefault:   false,
	}
	_, err = api.CreateBranchRequest(branchFoo).Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, branchFoo.ID)

	// Get branch
	resultBranch, err := api.GetBranchRequest(branchFoo.BranchKey).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, branchFoo, resultBranch)

	// Create branch, already exists
	branchFooDuplicate := &Branch{
		Name:        "Foo",
		Description: "Foo branch 2",
		IsDefault:   false,
	}
	_, err = api.CreateBranchRequest(branchFooDuplicate).Send(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "There already is a branch with name \"Foo\"")

	// Update branch
	branchFoo.Name = "Foo modified"
	branchFoo.Description = "Foo description modified"
	_, err = api.UpdateBranchRequest(branchFoo, []string{"name", "description"}).Send(ctx)
	assert.NoError(t, err)

	// Update main branch description
	defaultBranch.Description = "Default branch"
	_, err = api.UpdateBranchRequest(defaultBranch, []string{"description"}).Send(ctx)
	assert.NoError(t, err)

	// Can update default branch name
	defaultBranch.Name = "Not Allowed"
	assert.PanicsWithError(t, `the name of the main branch cannot be changed`, func() {
		api.UpdateBranchRequest(defaultBranch, []string{"name", "description"}).Send(ctx)
	})

	// List branches
	branches, err := api.ListBranchesRequest().Send(ctx)
	assert.NotNil(t, branches)
	assert.NoError(t, err)
	branchesJSON, err := json.MarshalIndent(branches, "", "  ")
	assert.NoError(t, err)
	wildcards.Assert(t, expectedBranchesAll(), string(branchesJSON), "Unexpected branches state")

	// Append branch metadata
	_, err = api.AppendBranchMetadataRequest(branchFoo.BranchKey, map[string]string{"KBC.KaC.meta1": "value", "KBC.KaC.meta2": "value"}).Send(ctx)
	assert.NoError(t, err)

	// List metadata
	metadata, err := api.ListBranchMetadataRequest(branchFoo.BranchKey).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, Metadata{"KBC.KaC.meta1": "value", "KBC.KaC.meta2": "value"}, metadata.ToMap())

	// Append metadata with empty value
	_, err = api.AppendBranchMetadataRequest(branchFoo.BranchKey, map[string]string{"KBC.KaC.meta2": ""}).Send(ctx)
	assert.NoError(t, err)

	// Check that metadata is deleted
	metadata, err = api.ListBranchMetadataRequest(branchFoo.BranchKey).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, Metadata{"KBC.KaC.meta1": "value"}, metadata.ToMap())

	// Delete metadata
	_, err = api.DeleteBranchMetadataRequest(branchFoo.BranchKey, (*metadata)[0].ID).Send(ctx)
	assert.NoError(t, err)

	// Check that metadata is deleted
	metadata, err = api.ListBranchMetadataRequest(branchFoo.BranchKey).Send(ctx)
	assert.NoError(t, err)
	assert.Empty(t, metadata)

	// Delete branch
	_, err = api.DeleteBranchRequest(branchFoo.BranchKey).Send(ctx)
	assert.NoError(t, err)

	// Check that branch has been deleted
	branches, err = api.ListBranchesRequest().Send(ctx)
	assert.NotNil(t, branches)
	assert.NoError(t, err)
	branchesJSON, err = json.MarshalIndent(branches, "", "  ")
	assert.NoError(t, err)
	wildcards.Assert(t, expectedBranchesMain(), string(branchesJSON), "Unexpected branches state")
}

func expectedBranchesAll() string {
	return `[
  {
    "id": %s,
    "name": "Foo modified",
    "description": "Foo description modified",
    "created": "%s",
    "isDefault": false
  },
  {
    "id": %s,
    "name": "Main",
    "description": "Default branch",
    "created": "%s",
    "isDefault": true
  }
]`
}

func expectedBranchesMain() string {
	return `[
  {
    "id": %s,
    "name": "Main",
    "description": "Default branch",
    "created": "%s",
    "isDefault": true
  }
]`
}

func TestBranchID_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    string
		expected BranchID
		wantErr  bool
	}{
		{name: "numeric", input: "123", expected: BranchID(123), wantErr: false},
		{name: "string number", input: `"456"`, expected: BranchID(456), wantErr: false},
		{name: "empty string", input: `""`, expected: BranchID(0), wantErr: false},
		{name: "zero numeric", input: "0", expected: BranchID(0), wantErr: false},
		{name: "zero string", input: `"0"`, expected: BranchID(0), wantErr: false},
		{name: "invalid string", input: `"abc"`, expected: BranchID(0), wantErr: true},
		{name: "negative numeric", input: "-1", expected: BranchID(-1), wantErr: false},
		{name: "negative string", input: `"-1"`, expected: BranchID(-1), wantErr: false},
		{name: "large number", input: "999999", expected: BranchID(999999), wantErr: false},
		{name: "large string number", input: `"999999"`, expected: BranchID(999999), wantErr: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var id BranchID
			err := json.Unmarshal([]byte(tc.input), &id)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, id)
			}
		})
	}
}

func TestBranchID_UnmarshalJSON_InStruct(t *testing.T) {
	t.Parallel()

	// Test unmarshaling BranchID as part of a struct (like QueueJob response)
	type testStruct struct {
		BranchID BranchID `json:"branchId"`
	}

	testCases := []struct {
		name     string
		input    string
		expected BranchID
	}{
		{
			name:     "branchId as string (Queue API format)",
			input:    `{"branchId": "12345"}`,
			expected: BranchID(12345),
		},
		{
			name:     "branchId as number (Storage API format)",
			input:    `{"branchId": 67890}`,
			expected: BranchID(67890),
		},
		{
			name:     "branchId as empty string",
			input:    `{"branchId": ""}`,
			expected: BranchID(0),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var s testStruct
			err := json.Unmarshal([]byte(tc.input), &s)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, s.BranchID)
		})
	}
}
