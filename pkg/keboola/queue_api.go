package keboola

// The file contains request definitions for the Jobs Queue API.
// Requests can be sent by any HTTP client that implements the client.Sender interface.

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v5"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

type JobBackendSize string

const (
	JobBackendXSmall JobBackendSize = "xsmall"
	JobBackendSmall  JobBackendSize = "small"
	JobBackendMedium JobBackendSize = "medium"
	JobBackendLarge  JobBackendSize = "large"
)

type JobMode string

const (
	JobModeRun      JobMode = "run"      // JobModeRun is the default mode, runs the job as usual.
	JobModeDebug    JobMode = "debug"    // JobModeDebug outputs a snapshot of configuration and a snapshot of output as storage files, but does not perform output mapping to storage.
	JobModeForceRun JobMode = "forceRun" // JobModeForceRun forces a configuration to run even if it is disabled.
)

type jobConfig struct {
	Tag                string              `json:"tag,omitempty"`
	BranchID           BranchID            `json:"branchId,omitempty"`
	ComponentID        ComponentID         `json:"component"`
	ConfigID           ConfigID            `json:"config,omitempty"`
	ConfigRowIDs       []string            `json:"configRowIds,omitempty"`
	ConfigData         map[string]any      `json:"configData,omitempty"`
	Mode               string              `json:"mode,omitempty"`
	VariableValuesID   string              `json:"variableValuesId,omitempty"`
	VariableValuesData *VariableValuesData `json:"variableValuesData,omitempty"`
	BackendSize        string              `json:"backend,omitempty"`
}

type VariableValuesData struct {
	Values []VariableData `json:"values"`
}

type VariableData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CreateQueueJobRequestBuilder struct {
	config jobConfig
	api    *AuthorizedAPI
}

// NewCreateJobRequest - https://app.swaggerhub.com/apis-docs/keboola/job-queue-api/1.3.2#/Jobs/createJob
func (a *AuthorizedAPI) NewCreateJobRequest(componentID ComponentID) *CreateQueueJobRequestBuilder {
	return &CreateQueueJobRequestBuilder{
		config: jobConfig{
			ComponentID: componentID,
		},
		api: a,
	}
}

// WithTag causes the component job to be run with a specific version of the component's docker image.
//
// If not provided, defaults to the latest tag.
func (b *CreateQueueJobRequestBuilder) WithTag(tag string) *CreateQueueJobRequestBuilder {
	b.config.Tag = tag
	return b
}

// WithBranch starts the component job in a dev branch.
//
// If not provided, defaults to the main branch.
func (b *CreateQueueJobRequestBuilder) WithBranch(id BranchID) *CreateQueueJobRequestBuilder {
	b.config.BranchID = id
	return b
}

// WithConfig starts the component job using a configuration.
//
// At least one of `WithConfig`, `WithConfigData` is required.
// If both are provided, `configData` overrides the `config`.
func (b *CreateQueueJobRequestBuilder) WithConfig(id ConfigID) *CreateQueueJobRequestBuilder {
	b.config.ConfigID = id
	return b
}

// WithConfigRowIDs starts the component job using only the specified row IDs.
func (b *CreateQueueJobRequestBuilder) WithConfigRowIDs(ids []string) *CreateQueueJobRequestBuilder {
	b.config.ConfigRowIDs = ids
	return b
}

// WithConfigData starts the component job using configuration data.
//
// At least one of `WithConfig`, `WithConfigData` is required.
// If both are provided, `configData` overrides the `config`.
func (b *CreateQueueJobRequestBuilder) WithConfigData(data map[string]any) *CreateQueueJobRequestBuilder {
	b.config.ConfigData = data
	return b
}

// WithMode starts the component job in a different mode.
//
// The available modes are:
//
// - Debug, which outputs a snapshot of configuration and a snapshot of output as storage files, but does not perform output mapping to storage.
//
// - ForceRun, which forces a configuration to run even if it is disabled.
func (b *CreateQueueJobRequestBuilder) WithMode(mode JobMode) *CreateQueueJobRequestBuilder {
	b.config.Mode = string(mode)
	return b
}

func (b *CreateQueueJobRequestBuilder) WithVariableValuesID(id string) *CreateQueueJobRequestBuilder {
	b.config.VariableValuesID = id
	return b
}

func (b *CreateQueueJobRequestBuilder) WithVariableValuesData(values []VariableData) *CreateQueueJobRequestBuilder {
	b.config.VariableValuesData = &VariableValuesData{Values: values}
	return b
}

// WithMode starts the component job with a specific backend size.
//
// The available sizes are: xsmall, small, medium, large.
func (b *CreateQueueJobRequestBuilder) WithBackendSize(size JobBackendSize) *CreateQueueJobRequestBuilder {
	b.config.BackendSize = string(size)
	return b
}

// Build finalizes and builds the request.
//
// This is useful if you want to send many of these requests in parallel.
func (b *CreateQueueJobRequestBuilder) Build() request.APIRequest[*QueueJob] {
	result := &QueueJob{}
	req := b.api.newRequest(QueueAPI).
		WithResult(&result).
		WithMethod(http.MethodPost).
		WithURL(QueueAPIJobs).
		WithJSONBody(b.config)
	return request.NewAPIRequest(result, req)
}

// Send builds the request and immediately sends it.
//
// This is a convenience method that simply calls Build() followed by Send(ctx).
func (b *CreateQueueJobRequestBuilder) Send(ctx context.Context) (*QueueJob, error) {
	return b.Build().Send(ctx)
}

// Deprecated: Use `NewCreateJobRequest` instead.
//
// CreateQueueJobRequest - https://app.swaggerhub.com/apis-docs/keboola/job-queue-api/1.3.2#/Jobs/createJob
func (a *AuthorizedAPI) CreateQueueJobRequest(componentID ComponentID, configID ConfigID) request.APIRequest[*QueueJob] {
	data := map[string]string{
		"component": componentID.String(),
		"mode":      "run",
		"config":    configID.String(),
	}
	result := QueueJob{}
	req := a.newRequest(QueueAPI).
		WithResult(&result).
		WithMethod(http.MethodPost).
		WithURL(QueueAPIJobs).
		WithJSONBody(data)
	return request.NewAPIRequest(&result, req)
}

// Deprecated: Use `NewCreateJobRequest` instead.
//
// CreateQueueJobConfigDataRequest - https://app.swaggerhub.com/apis-docs/keboola/job-queue-api/1.3.2#/Jobs/createJob
//
// Allows setting configData.
//
// `configId` can be set to an empty string, and it will be omitted.
func (a *AuthorizedAPI) CreateQueueJobConfigDataRequest(componentID ComponentID, configID ConfigID, configData map[string]any) request.APIRequest[*QueueJob] {
	body := map[string]any{
		"component":  componentID.String(),
		"mode":       "run",
		"configData": configData,
	}
	if len(configID.String()) > 0 {
		body["config"] = configID.String()
	}

	result := &QueueJob{}
	req := a.newRequest(QueueAPI).
		WithResult(result).
		WithMethod(http.MethodPost).
		WithURL(QueueAPIJobs).
		WithJSONBody(body)
	return request.NewAPIRequest(result, req)
}

// GetJobRequest https://app.swaggerhub.com/apis-docs/keboola/job-queue-api/1.3.2#/Jobs/getJob
func (a *AuthorizedAPI) GetQueueJobRequest(key JobKey) request.APIRequest[*QueueJob] {
	return a.getQueueJobRequest(key.ID)
}

func (a *AuthorizedAPI) getQueueJobRequest(id JobID) request.APIRequest[*QueueJob] {
	job := &QueueJob{}
	req := a.newRequest(QueueAPI).
		WithResult(job).
		WithGet(QueueAPIJob).
		AndPathParam("jobId", id.String())
	return request.NewAPIRequest(job, req)
}

// WaitForQueueJob pulls job status until it is completed.
func (a *AuthorizedAPI) WaitForQueueJob(ctx context.Context, id JobID) (err error) {
	_, ok := ctx.Deadline()
	if !ok {
		return fmt.Errorf("timeout for the job was not set")
	}

	// Telemetry
	parentSpan := trace.SpanFromContext(ctx)
	var span trace.Span
	ctx, span = parentSpan.TracerProvider().Tracer(appName).Start(ctx, "keboola.go.api.client.waitFor.queueJob")
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	retry := newQueueJobBackoff()
	for {
		// Get job status
		job, err := a.getQueueJobRequest(id).Send(ctx)
		if err != nil {
			return err
		}

		// Check status
		if job.IsFinished {
			if job.Status == "success" {
				return nil
			}
			return fmt.Errorf(`job "%s" failed: %v`, job.ID, job.Result.Message)
		}

		// Wait and check again
		select {
		case <-ctx.Done():
			return fmt.Errorf(`error while waiting for the job "%s" to complete: %w`, job.ID, ctx.Err())
		case <-time.After(retry.NextBackOff()):
			// try again
		}
	}
}

// newQueueJobBackoff creates retry for WaitForQueueJob.
func newQueueJobBackoff() *backoff.ExponentialBackOff {
	b := backoff.NewExponentialBackOff()
	b.RandomizationFactor = 0
	b.InitialInterval = 3 * time.Second
	b.Multiplier = 2
	b.MaxInterval = 5 * time.Second
	b.Reset()
	return b
}

// SearchJobsConfig holds configuration for SearchJobsRequest.
type SearchJobsConfig struct {
	BranchID    BranchID
	ComponentID ComponentID
	ConfigID    ConfigID
	Status      string
	Limit       int
	Offset      int
}

// SearchJobsOption is a functional option for SearchJobsRequest.
type SearchJobsOption func(c *SearchJobsConfig)

// WithSearchJobsBranch filters jobs by branch ID.
func WithSearchJobsBranch(branchID BranchID) SearchJobsOption {
	return func(c *SearchJobsConfig) {
		c.BranchID = branchID
	}
}

// WithSearchJobsComponent filters jobs by component ID.
func WithSearchJobsComponent(componentID ComponentID) SearchJobsOption {
	return func(c *SearchJobsConfig) {
		c.ComponentID = componentID
	}
}

// WithSearchJobsConfig filters jobs by config ID.
func WithSearchJobsConfig(configID ConfigID) SearchJobsOption {
	return func(c *SearchJobsConfig) {
		c.ConfigID = configID
	}
}

// WithSearchJobsStatus filters jobs by status.
func WithSearchJobsStatus(status string) SearchJobsOption {
	return func(c *SearchJobsConfig) {
		c.Status = status
	}
}

// WithSearchJobsLimit sets the maximum number of jobs to return.
// If not called, defaults to 100. Setting limit to 0 or negative
// will omit the limit parameter from the request, using server default.
func WithSearchJobsLimit(limit int) SearchJobsOption {
	return func(c *SearchJobsConfig) {
		c.Limit = limit
	}
}

// WithSearchJobsOffset sets the offset for pagination.
func WithSearchJobsOffset(offset int) SearchJobsOption {
	return func(c *SearchJobsConfig) {
		c.Offset = offset
	}
}

// SearchJobsRequest searches for jobs in the Jobs Queue API.
// https://app.swaggerhub.com/apis-docs/keboola/job-queue-api/1.3.8#/Jobs/get_search_jobs
func (a *AuthorizedAPI) SearchJobsRequest(opts ...SearchJobsOption) request.APIRequest[*[]*QueueJob] {
	config := SearchJobsConfig{
		Limit: 100, // Default limit
	}
	for _, opt := range opts {
		opt(&config)
	}

	result := make([]*QueueJob, 0)
	req := a.newRequest(QueueAPI).
		WithResult(&result).
		WithGet(QueueAPISearchJobs)

	if config.BranchID != 0 {
		req = req.AndQueryParam("branchId", config.BranchID.String())
	}
	if config.ComponentID != "" {
		req = req.AndQueryParam("component", config.ComponentID.String())
	}
	if config.ConfigID != "" {
		req = req.AndQueryParam("config", config.ConfigID.String())
	}
	if config.Status != "" {
		req = req.AndQueryParam("status", config.Status)
	}
	if config.Limit > 0 {
		req = req.AndQueryParam("limit", fmt.Sprintf("%d", config.Limit))
	}
	if config.Offset > 0 {
		req = req.AndQueryParam("offset", fmt.Sprintf("%d", config.Offset))
	}

	return request.NewAPIRequest(&result, req)
}
