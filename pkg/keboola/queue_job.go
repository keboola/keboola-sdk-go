package keboola

import (
	jsonLib "encoding/json"
	"fmt"

	"github.com/relvacode/iso8601"
)

// JobID is an ID of a component job.
type JobID string

func (j JobID) String() string {
	return string(j)
}

func (j *JobID) UnmarshalJSON(b []byte) error {
	var asString string
	err := jsonLib.Unmarshal(b, &asString)
	if err != nil {
		var asInt int
		err = jsonLib.Unmarshal(b, &asInt)
		if err != nil {
			return fmt.Errorf("failed to unmarshal int or string")
		}
		asString = fmt.Sprint(asInt)
	}
	*j = JobID(asString)
	return nil
}

// JobKey is a unique identifier of QueueJob.
type JobKey struct {
	ID JobID `json:"id"`
}

type JobResult struct {
	Error   map[string]any `json:"error,omitempty"`
	Message string         `json:"message,omitempty"`
}

// UnmarshalJSON implements JSON decoding.
func (r *JobResult) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "[]" {
		*r = JobResult{
			Error:   nil,
			Message: "",
		}
		return nil
	}
	// see https://stackoverflow.com/questions/43176625/call-json-unmarshal-inside-unmarshaljson-function-without-causing-stack-overflow
	type _r JobResult
	return jsonLib.Unmarshal(data, (*_r)(r))
}

// QueueJob is a component job.
type QueueJob struct {
	JobKey
	Status          string        `json:"status"`
	IsFinished      bool          `json:"isFinished"`
	URL             string        `json:"url"`
	Result          JobResult     `json:"result"`
	CreateTime      iso8601.Time  `json:"createdTime"`
	StartTime       *iso8601.Time `json:"startTime"`
	EndTime         *iso8601.Time `json:"endTime"`
	ComponentID     ComponentID   `json:"component,omitempty"`
	ConfigID        ConfigID      `json:"config,omitempty"`
	BranchID        BranchID      `json:"branchId,omitempty"`
	OperationName   string        `json:"operationName,omitempty"`
	DurationSeconds int           `json:"durationSeconds,omitempty"`
	RunID           string        `json:"runId,omitempty"`
}

// QueueJobDetail is a component job with extended result data including input/output tables.
type QueueJobDetail struct {
	JobKey
	Status          string              `json:"status"`
	IsFinished      bool                `json:"isFinished"`
	URL             string              `json:"url"`
	Result          JobResultExtended   `json:"result"`
	CreateTime      iso8601.Time        `json:"createdTime"`
	StartTime       *iso8601.Time       `json:"startTime"`
	EndTime         *iso8601.Time       `json:"endTime"`
	ComponentID     ComponentID         `json:"component,omitempty"`
	ConfigID        ConfigID            `json:"config,omitempty"`
	BranchID        BranchID            `json:"branchId,omitempty"`
	OperationName   string              `json:"operationName,omitempty"`
	DurationSeconds int                 `json:"durationSeconds,omitempty"`
	RunID           string              `json:"runId,omitempty"`
	ParentRunID     string              `json:"parentRunId,omitempty"`
	Type            string              `json:"type,omitempty"` // "standard", "orchestrationContainer", etc.
	Metrics         *JobMetrics         `json:"metrics,omitempty"`
}

// JobResultExtended contains the full result data from a job including input/output tables.
type JobResultExtended struct {
	Error     map[string]any  `json:"error,omitempty"`
	Message   string          `json:"message,omitempty"`
	Input     *JobIO          `json:"input,omitempty"`
	Output    *JobIO          `json:"output,omitempty"`
	Artifacts *JobArtifacts   `json:"artifacts,omitempty"`
	Tasks     []JobTask       `json:"tasks,omitempty"`  // For orchestrations/flows
	Phases    []JobPhase      `json:"phases,omitempty"` // For flows
}

// UnmarshalJSON implements JSON decoding for JobResultExtended.
func (r *JobResultExtended) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "[]" {
		*r = JobResultExtended{}
		return nil
	}
	type _r JobResultExtended
	return jsonLib.Unmarshal(data, (*_r)(r))
}

// JobIO represents input or output tables for a job.
type JobIO struct {
	Tables []JobTable `json:"tables,omitempty"`
}

// JobTable represents a table in job input/output.
type JobTable struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	DisplayName string          `json:"displayName,omitempty"`
	Columns     []JobTableColumn `json:"columns,omitempty"`
}

// JobTableColumn represents a column in a job table.
type JobTableColumn struct {
	Name string `json:"name"`
}

// JobArtifacts represents artifacts uploaded/downloaded by a job.
type JobArtifacts struct {
	Uploaded   []any `json:"uploaded,omitempty"`
	Downloaded []any `json:"downloaded,omitempty"`
}

// JobTask represents a task within an orchestration or flow job.
type JobTask struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Type      string          `json:"type"`
	JobID     string          `json:"jobId,omitempty"`
	Phase     string          `json:"phase,omitempty"`
	Status    string          `json:"status"`
	Component string          `json:"component,omitempty"`
	Duration  int             `json:"duration,omitempty"`
	StartTime *iso8601.Time   `json:"startTime,omitempty"`
	Results   []JobTaskResult `json:"results,omitempty"`
}

// JobTaskResult represents the result of a task within an orchestration or flow.
type JobTaskResult struct {
	JobID    string            `json:"jobId"`
	Status   string            `json:"status"`
	Duration int               `json:"duration,omitempty"`
	Result   JobResultExtended `json:"result"`
}

// JobPhase represents a phase in a flow job.
type JobPhase struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Status    string        `json:"status"`
	Duration  int           `json:"duration,omitempty"`
	StartTime *iso8601.Time `json:"startTime,omitempty"`
}

// JobMetrics represents metrics for a job.
type JobMetrics struct {
	Storage *JobStorageMetrics `json:"storage,omitempty"`
	Backend *JobBackendMetrics `json:"backend,omitempty"`
}

// UnmarshalJSON implements JSON decoding for JobMetrics.
// Handles the case where the API returns an empty array [] instead of an object or null.
func (m *JobMetrics) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "[]" {
		*m = JobMetrics{}
		return nil
	}
	type _m JobMetrics
	return jsonLib.Unmarshal(data, (*_m)(m))
}

// JobStorageMetrics represents storage metrics for a job.
type JobStorageMetrics struct {
	InputTablesBytesSum  int64 `json:"inputTablesBytesSum"`
	OutputTablesBytesSum int64 `json:"outputTablesBytesSum"`
}

// JobBackendMetrics represents backend metrics for a job.
type JobBackendMetrics struct {
	Size          string `json:"size,omitempty"`
	ContainerSize string `json:"containerSize,omitempty"`
	Context       string `json:"context,omitempty"`
}
