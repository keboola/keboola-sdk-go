package keboola

import (
	"strconv"
	"time"
)

// TimeFormat used in SandboxWorkspaces API.
const TimeFormat = "2006-01-02T15:04:05Z"

// SandboxWorkspacesTime is encoded/decoded in TimeFormat used in Workspaces API.
type SandboxWorkspacesTime time.Time

// UnmarshalJSON implements JSON decoding.
func (t *SandboxWorkspacesTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = SandboxWorkspacesTime(now)
	return
}

// MarshalJSON implements JSON encoding.
func (t SandboxWorkspacesTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t SandboxWorkspacesTime) String() string {
	return time.Time(t).Format(TimeFormat)
}

// DurationSeconds is time.Duration encoded/decoded as number of seconds.
type DurationSeconds time.Duration

// UnmarshalJSON implements JSON decoding.
func (d *DurationSeconds) UnmarshalJSON(data []byte) (err error) {
	v, err := time.ParseDuration(string(data) + "s")
	if err != nil {
		return err
	}
	*d = DurationSeconds(v)
	return
}

// MarshalJSON implements JSON encoding.
func (d DurationSeconds) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d DurationSeconds) String() string {
	return strconv.Itoa(int(time.Duration(d).Seconds()))
}
