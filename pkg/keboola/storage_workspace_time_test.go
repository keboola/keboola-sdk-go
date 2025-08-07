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
			name:     "parse negative timezone offset",
			jsonData: `"2025-08-06T14:12:37-0500"`,
			expected: time.Date(2025, 8, 6, 14, 12, 37, 0, time.FixedZone("", -5*60*60)),
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
			// Compare the actual time values and timezone offsets
			assert.Equal(t, tt.expected.Year(), result.Time().Year())
			assert.Equal(t, tt.expected.Month(), result.Time().Month())
			assert.Equal(t, tt.expected.Day(), result.Time().Day())
			assert.Equal(t, tt.expected.Hour(), result.Time().Hour())
			assert.Equal(t, tt.expected.Minute(), result.Time().Minute())
			assert.Equal(t, tt.expected.Second(), result.Time().Second())
			// Compare timezone offsets in seconds
			_, expectedOffset := tt.expected.Zone()
			_, actualOffset := result.Time().Zone()
			assert.Equal(t, expectedOffset, actualOffset)
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

	// Verify the time is preserved by comparing individual components
	assert.Equal(t, testTime.Year(), result.Time().Year())
	assert.Equal(t, testTime.Month(), result.Time().Month())
	assert.Equal(t, testTime.Day(), result.Time().Day())
	assert.Equal(t, testTime.Hour(), result.Time().Hour())
	assert.Equal(t, testTime.Minute(), result.Time().Minute())
	assert.Equal(t, testTime.Second(), result.Time().Second())
	// Compare timezone offsets in seconds
	_, expectedOffset := testTime.Zone()
	_, actualOffset := result.Time().Zone()
	assert.Equal(t, expectedOffset, actualOffset)
}

func TestStorageWorkspaceTime_String(t *testing.T) {
	t.Parallel()
	testTime := time.Date(2025, 8, 6, 14, 12, 37, 0, time.FixedZone("", 2*60*60))
	storageTime := StorageWorkspaceTime(testTime)

	// String should return StorageWorkspaceTimeFormat format (+0200)
	expected := "2025-08-06T14:12:37+0200"
	assert.Equal(t, expected, storageTime.String())
}
