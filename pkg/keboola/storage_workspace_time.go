package keboola

import (
	"time"
)

// StorageWorkspaceTimeFormat is the time format used by the Storage Workspace API.
// The API returns times in format like "2025-08-06T14:12:37+0200".
const StorageWorkspaceTimeFormat = "2006-01-02T15:04:05-0700"

// StorageWorkspaceTime is a custom time type that handles the timezone format used by the Storage Workspace API.
// The API returns times in format like "2025-08-06T14:12:37+0200" which Go's standard time parsing can handle.
type StorageWorkspaceTime time.Time

// UnmarshalJSON implements JSON decoding for StorageWorkspaceTime.
func (t *StorageWorkspaceTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.Parse(`"`+StorageWorkspaceTimeFormat+`"`, string(data))
	*t = StorageWorkspaceTime(now)
	return
}

// MarshalJSON implements JSON encoding for StorageWorkspaceTime.
func (t StorageWorkspaceTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(StorageWorkspaceTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, StorageWorkspaceTimeFormat)
	b = append(b, '"')
	return b, nil
}

// String returns the time formatted as StorageWorkspaceTimeFormat.
func (t StorageWorkspaceTime) String() string {
	return time.Time(t).Format(StorageWorkspaceTimeFormat)
}

// Time returns the underlying time.Time value.
func (t StorageWorkspaceTime) Time() time.Time {
	return time.Time(t)
}
