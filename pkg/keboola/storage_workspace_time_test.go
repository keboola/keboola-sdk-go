package keboola

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStorageWorkspaceTime_UnmarshalJSON(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		jsonData string
		expected time.Time
		wantErr  bool
	}{
		{
			name:     "parse +0200 timezone format",
			jsonData: `"2025-08-06T14:12:37+0200"`,
			expected: time.Date(2025, 8, 6, 14, 12, 37, 0, time.FixedZone("", 2*60*60)),
			wantErr:  false,
		},
		{
			name:     "parse standard RFC3339 format",
			jsonData: `"2025-08-06T14:12:37+02:00"`,
			expected: time.Date(2025, 8, 6, 14, 12, 37, 0, time.FixedZone("", 2*60*60)),
			wantErr:  false,
		},
		{
			name:     "parse negative timezone offset",
			jsonData: `"2025-08-06T14:12:37-0500"`,
			expected: time.Date(2025, 8, 6, 14, 12, 37, 0, time.FixedZone("", -5*60*60)),
			wantErr:  false,
		},
		{
			name:     "parse UTC timezone",
			jsonData: `"2025-08-06T14:12:37Z"`,
			expected: time.Date(2025, 8, 6, 14, 12, 37, 0, time.UTC),
			wantErr:  false,
		},
		{
			name:     "invalid time format",
			jsonData: `"invalid-time"`,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		// capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var result StorageWorkspaceTime
			err := json.Unmarshal([]byte(tt.jsonData), &result)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result.Time())
		})
	}
}

func TestStorageWorkspaceTime_MarshalJSON(t *testing.T) {
	t.Parallel()
	// Create a time with +0200 timezone
	testTime := time.Date(2025, 8, 6, 14, 12, 37, 0, time.FixedZone("", 2*60*60))
	storageTime := StorageWorkspaceTime(testTime)

	// Marshal to JSON
	jsonData, err := json.Marshal(storageTime)
	assert.NoError(t, err)

	// Unmarshal back
	var result StorageWorkspaceTime
	err = json.Unmarshal(jsonData, &result)
	assert.NoError(t, err)

	// Verify the time is preserved
	assert.Equal(t, testTime, result.Time())
}

func TestStorageWorkspaceTime_String(t *testing.T) {
	t.Parallel()
	testTime := time.Date(2025, 8, 6, 14, 12, 37, 0, time.FixedZone("", 2*60*60))
	storageTime := StorageWorkspaceTime(testTime)

	// String should return RFC3339 format
	expected := "2025-08-06T14:12:37+02:00"
	assert.Equal(t, expected, storageTime.String())
}
