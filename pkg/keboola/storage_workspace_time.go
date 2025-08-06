package keboola

import "time"

// StorageWorkspaceTime is a custom time type that handles the +0200 timezone format used by the Storage Workspace API.
// The API returns times in format like "2025-08-06T14:12:37+0200" which Go's standard time parsing cannot handle.
type StorageWorkspaceTime time.Time

// UnmarshalJSON implements JSON decoding for StorageWorkspaceTime.
// It handles the +0200 timezone format by converting it to +02:00 format before parsing.
func (t *StorageWorkspaceTime) UnmarshalJSON(data []byte) error {
	// Remove quotes from the JSON string
	timeStr := string(data)
	if len(timeStr) >= 2 && timeStr[0] == '"' && timeStr[len(timeStr)-1] == '"' {
		timeStr = timeStr[1 : len(timeStr)-1]
	}

	// Try to parse with standard format first
	if parsed, err := time.Parse(time.RFC3339, timeStr); err == nil {
		*t = StorageWorkspaceTime(parsed)
		return nil
	}

	// If standard format fails, try to handle +0200 format by converting to +02:00
	// Look for timezone offset pattern like +0200 or -0500
	if len(timeStr) >= 5 {
		lastChar := timeStr[len(timeStr)-1]
		secondLastChar := timeStr[len(timeStr)-2]
		thirdLastChar := timeStr[len(timeStr)-3]
		fourthLastChar := timeStr[len(timeStr)-4]
		fifthLastChar := timeStr[len(timeStr)-5]

		// Check if we have a timezone offset pattern
		if (fifthLastChar == '+' || fifthLastChar == '-') &&
			(fourthLastChar >= '0' && fourthLastChar <= '9') &&
			(thirdLastChar >= '0' && thirdLastChar <= '9') &&
			(secondLastChar >= '0' && secondLastChar <= '9') &&
			(lastChar >= '0' && lastChar <= '9') {
			// Convert +0200 to +02:00 format
			convertedTimeStr := timeStr[:len(timeStr)-2] + ":" + timeStr[len(timeStr)-2:]
			if parsed, err := time.Parse(time.RFC3339, convertedTimeStr); err == nil {
				*t = StorageWorkspaceTime(parsed)
				return nil
			}
		}
	}

	// If all parsing attempts fail, return the original error
	_, err := time.Parse(time.RFC3339, timeStr)
	return err
}

// MarshalJSON implements JSON encoding for StorageWorkspaceTime.
func (t StorageWorkspaceTime) MarshalJSON() ([]byte, error) {
	return time.Time(t).MarshalJSON()
}

// String returns the time formatted as RFC3339.
func (t StorageWorkspaceTime) String() string {
	return time.Time(t).Format(time.RFC3339)
}

// Time returns the underlying time.Time value.
func (t StorageWorkspaceTime) Time() time.Time {
	return time.Time(t)
}
