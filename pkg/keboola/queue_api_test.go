package keboola

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/keboola/go-utils/pkg/orderedmap"
	"github.com/keboola/go-utils/pkg/wildcards"
	"github.com/stretchr/testify/assert"

	"github.com/keboola/keboola-sdk-go/v2/pkg/client"
	"github.com/keboola/keboola-sdk-go/v2/pkg/client/trace"
)

func TestQueueApiCalls(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := APIClientForAnEmptyProject(t, ctx)

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, branch)

	// List - no component/config
	components, err := api.ListConfigsAndRowsFrom(branch.BranchKey).Send(ctx)
	assert.NoError(t, err)
	assert.Empty(t, components)

	// Create config
	config := &ConfigWithRows{
		Config: &Config{
			ConfigKey: ConfigKey{
				BranchID:    branch.ID,
				ComponentID: "ex-generic-v2",
			},
			Name:              "Test",
			Description:       "Test description",
			ChangeDescription: "My test",
			Content: orderedmap.FromPairs([]orderedmap.Pair{
				{Key: "foo", Value: "bar"},
			}),
		},
		Rows: []*ConfigRow{},
	}
	resConfig, err := api.CreateConfigRequest(config, true).Send(ctx)
	assert.NoError(t, err)
	assert.Same(t, config, resConfig)
	assert.NotEmpty(t, config.ID)

	// Run a job on the config
	resJob, err := api.NewCreateJobRequest("ex-generic-v2").WithConfig(config.ID).Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resJob.ID)

	// Wait for the job
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancelFn()
	err = api.WaitForQueueJob(timeoutCtx, resJob.ID)
	// The job payload is malformed, so it fails. We are checking just that it finished.
	assert.ErrorContains(t, err, "Unrecognized option \"foo\" under \"container\"")
}

func TestCreateQueueJobRequestBuilder(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := APIClientForAnEmptyProject(t, ctx)

	builder := api.NewCreateJobRequest("ex-generic-v2").
		WithTag("latest").
		WithBranch(1234).
		WithConfig("987654321").
		WithConfigRowIDs([]string{"config-row-a", "config-row-b"}).
		WithConfigData(map[string]any{"a": "b"}).
		WithBackendSize("xsmall").
		WithVariableValuesID("variable-values-id").
		WithVariableValuesData([]VariableData{{Name: "a", Value: "b"}})

	data, err := json.MarshalIndent(builder.config, "", "  ")
	assert.NoError(t, err)
	assert.Equal(t,
		`{
  "tag": "latest",
  "branchId": 1234,
  "component": "ex-generic-v2",
  "config": "987654321",
  "configRowIds": [
    "config-row-a",
    "config-row-b"
  ],
  "configData": {
    "a": "b"
  },
  "variableValuesId": "variable-values-id",
  "variableValuesData": {
    "values": [
      {
        "name": "a",
        "value": "b"
      }
    ]
  },
  "backend": "xsmall"
}`,
		string(data),
	)
}

func TestQueueWaitForQueueJobTimeout(t *testing.T) {
	t.Parallel()

	job := QueueJob{
		JobKey:     JobKey{ID: "1234"},
		Status:     "waiting",
		IsFinished: false,
	}

	// Trace client activity
	var traceOut bytes.Buffer

	// Create mocked timeout
	c, transport := client.NewMockedClient()
	c = c.WithBaseURL("https://connection.test").AndTrace(trace.LogTracer(&traceOut))
	transport.RegisterResponder("GET", `https://connection.test/v2/storage/?exclude=components`, newJSONResponder(`{
		"services": [
			{
				"id": "queue",
				"url": "https://queue.connection.test"
			}
		],
		"features": []
	}`))
	transport.RegisterResponder("GET", `=~^https://queue.connection.test/jobs/1234`, httpmock.NewJsonResponderOrPanic(200, job))
	api, err := NewAuthorizedAPI(context.Background(), "https://connection.test", "my-token", WithClient(&c))
	assert.NoError(t, err)

	// Create context with deadline
	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFn()

	// Error - deadline exceeded
	err = api.WaitForQueueJob(ctx, job.ID)
	assert.Error(t, err)
	assert.Equal(t, `error while waiting for the job "1234" to complete: context deadline exceeded`, err.Error())

	// Check calls count
	assert.Equal(t, 3, transport.GetCallCountInfo()["GET https://queue.connection.test/jobs/1234"])

	// Check client activity
	wildcards.Assert(t, strings.TrimSpace(`
HTTP_REQUEST[0001] START GET "https://connection.test/v2/storage/?exclude=components"
HTTP_REQUEST[0001] DONE  GET "https://connection.test/v2/storage/?exclude=components" | 200 | %s
HTTP_REQUEST[0001] BODY  GET "https://connection.test/v2/storage/?exclude=components" | %s
HTTP_REQUEST[0002] START GET "https://queue.connection.test/jobs/1234"
HTTP_REQUEST[0002] DONE  GET "https://queue.connection.test/jobs/1234" | 200 | %s
HTTP_REQUEST[0002] BODY  GET "https://queue.connection.test/jobs/1234" | %s
HTTP_REQUEST[0003] START GET "https://queue.connection.test/jobs/1234"
HTTP_REQUEST[0003] DONE  GET "https://queue.connection.test/jobs/1234" | 200 | %s
HTTP_REQUEST[0003] BODY  GET "https://queue.connection.test/jobs/1234" | %s
HTTP_REQUEST[0004] START GET "https://queue.connection.test/jobs/1234"
HTTP_REQUEST[0004] DONE  GET "https://queue.connection.test/jobs/1234" | 200 | %s
HTTP_REQUEST[0004] BODY  GET "https://queue.connection.test/jobs/1234" | %s
`), traceOut.String())
}

func TestDeprecatedQueueApiCalls(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := APIClientForAnEmptyProject(t, ctx)

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, branch)

	// List - no component/config
	components, err := api.ListConfigsAndRowsFrom(branch.BranchKey).Send(ctx)
	assert.NoError(t, err)
	assert.Empty(t, components)

	// Create config
	config := &ConfigWithRows{
		Config: &Config{
			ConfigKey: ConfigKey{
				BranchID:    branch.ID,
				ComponentID: "ex-generic-v2",
			},
			Name:              "Test",
			Description:       "Test description",
			ChangeDescription: "My test",
			Content: orderedmap.FromPairs([]orderedmap.Pair{
				{Key: "foo", Value: "bar"},
			}),
		},
		Rows: []*ConfigRow{},
	}
	resConfig, err := api.CreateConfigRequest(config, true).Send(ctx)
	assert.NoError(t, err)
	assert.Same(t, config, resConfig)
	assert.NotEmpty(t, config.ID)

	// Run a job on the config
	job, err := api.NewCreateJobRequest("ex-generic-v2").WithConfig(config.ID).Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, job.ID)

	// Wait for the job
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancelFn()
	err = api.WaitForQueueJob(timeoutCtx, job.ID)
	// The job payload is malformed, so it fails. We are checking just that it finished.
	assert.ErrorContains(t, err, "Unrecognized option \"foo\" under \"container\"")
}

func newJSONResponder(response string) httpmock.Responder {
	r := httpmock.NewStringResponse(200, response)
	r.Header.Set("Content-Type", "application/json")
	return httpmock.ResponderFromResponse(r)
}

func TestJobResultExtended_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	t.Run("empty array returns empty struct", func(t *testing.T) {
		t.Parallel()
		var result JobResultExtended
		err := json.Unmarshal([]byte("[]"), &result)
		assert.NoError(t, err)
		assert.Empty(t, result.Message)
		assert.Nil(t, result.Input)
		assert.Nil(t, result.Output)
	})

	t.Run("object with message", func(t *testing.T) {
		t.Parallel()
		var result JobResultExtended
		err := json.Unmarshal([]byte(`{"message": "Job completed"}`), &result)
		assert.NoError(t, err)
		assert.Equal(t, "Job completed", result.Message)
	})

	t.Run("object with input/output tables", func(t *testing.T) {
		t.Parallel()
		var result JobResultExtended
		err := json.Unmarshal([]byte(`{
			"message": "Success",
			"input": {
				"tables": [{"id": "in.c-bucket.table1", "name": "table1"}]
			},
			"output": {
				"tables": [{"id": "out.c-bucket.table2", "name": "table2"}]
			}
		}`), &result)
		assert.NoError(t, err)
		assert.Equal(t, "Success", result.Message)
		assert.NotNil(t, result.Input)
		assert.Len(t, result.Input.Tables, 1)
		assert.Equal(t, "in.c-bucket.table1", result.Input.Tables[0].ID)
		assert.NotNil(t, result.Output)
		assert.Len(t, result.Output.Tables, 1)
		assert.Equal(t, "out.c-bucket.table2", result.Output.Tables[0].ID)
	})

	t.Run("object with error", func(t *testing.T) {
		t.Parallel()
		var result JobResultExtended
		err := json.Unmarshal([]byte(`{
			"message": "Job failed",
			"error": {"code": 500, "message": "Internal error"}
		}`), &result)
		assert.NoError(t, err)
		assert.Equal(t, "Job failed", result.Message)
		assert.NotNil(t, result.Error)
		assert.Equal(t, float64(500), result.Error["code"])
	})
}

func TestJobMetrics_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	t.Run("empty array returns empty struct", func(t *testing.T) {
		t.Parallel()
		var metrics JobMetrics
		err := json.Unmarshal([]byte("[]"), &metrics)
		assert.NoError(t, err)
		assert.Nil(t, metrics.Storage)
		assert.Nil(t, metrics.Backend)
	})

	t.Run("null returns empty struct", func(t *testing.T) {
		t.Parallel()
		var metrics JobMetrics
		err := json.Unmarshal([]byte("null"), &metrics)
		assert.NoError(t, err)
		assert.Nil(t, metrics.Storage)
		assert.Nil(t, metrics.Backend)
	})

	t.Run("object with storage metrics", func(t *testing.T) {
		t.Parallel()
		var metrics JobMetrics
		err := json.Unmarshal([]byte(`{
			"storage": {
				"inputTablesBytesSum": 1024,
				"outputTablesBytesSum": 2048
			}
		}`), &metrics)
		assert.NoError(t, err)
		assert.NotNil(t, metrics.Storage)
		assert.Equal(t, int64(1024), metrics.Storage.InputTablesBytesSum)
		assert.Equal(t, int64(2048), metrics.Storage.OutputTablesBytesSum)
		assert.Nil(t, metrics.Backend)
	})

	t.Run("object with backend metrics", func(t *testing.T) {
		t.Parallel()
		var metrics JobMetrics
		err := json.Unmarshal([]byte(`{
			"backend": {
				"size": "small",
				"containerSize": "medium",
				"context": "wlm"
			}
		}`), &metrics)
		assert.NoError(t, err)
		assert.Nil(t, metrics.Storage)
		assert.NotNil(t, metrics.Backend)
		assert.Equal(t, "small", metrics.Backend.Size)
		assert.Equal(t, "medium", metrics.Backend.ContainerSize)
		assert.Equal(t, "wlm", metrics.Backend.Context)
	})

	t.Run("object with both storage and backend metrics", func(t *testing.T) {
		t.Parallel()
		var metrics JobMetrics
		err := json.Unmarshal([]byte(`{
			"storage": {
				"inputTablesBytesSum": 5000,
				"outputTablesBytesSum": 10000
			},
			"backend": {
				"size": "large"
			}
		}`), &metrics)
		assert.NoError(t, err)
		assert.NotNil(t, metrics.Storage)
		assert.Equal(t, int64(5000), metrics.Storage.InputTablesBytesSum)
		assert.NotNil(t, metrics.Backend)
		assert.Equal(t, "large", metrics.Backend.Size)
	})
}

func TestGetQueueJobDetailRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := APIClientForAnEmptyProject(t, ctx)

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	assert.NoError(t, err)

	// Create config
	config := &ConfigWithRows{
		Config: &Config{
			ConfigKey: ConfigKey{
				BranchID:    branch.ID,
				ComponentID: "ex-generic-v2",
			},
			Name:              "Test Detail",
			Description:       "Test for GetQueueJobDetailRequest",
			ChangeDescription: "Test",
			Content:           orderedmap.New(),
		},
		Rows: []*ConfigRow{},
	}
	_, err = api.CreateConfigRequest(config, true).Send(ctx)
	assert.NoError(t, err)

	// Run a job
	job, err := api.NewCreateJobRequest("ex-generic-v2").WithConfig(config.ID).Send(ctx)
	assert.NoError(t, err)

	// Wait for job to finish
	timeoutCtx, cancelFn := context.WithTimeout(ctx, time.Minute*5)
	defer cancelFn()
	_ = api.WaitForQueueJob(timeoutCtx, job.ID)

	// Get job detail using JobKey
	jobDetail, err := api.GetQueueJobDetailRequest(job.JobKey).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, jobDetail)
	assert.Equal(t, job.ID, jobDetail.ID)
	assert.Equal(t, ComponentID("ex-generic-v2"), jobDetail.ComponentID)
	assert.True(t, jobDetail.IsFinished)

	// Verify extended fields are accessible (may be nil/empty depending on job type)
	// The Result struct should be populated with at least a message for failed jobs
	assert.NotEmpty(t, jobDetail.Result.Message, "Result message should be populated for finished jobs")
	// Type field indicates job type (standard, orchestrationContainer, etc.)
	assert.NotEmpty(t, jobDetail.Type, "Type field should be populated")
	// Metrics may be nil for some job types, but the field should be accessible
	// Input/Output tables depend on job configuration, so we just verify the struct is correctly deserialized
}

func TestSearchJobsDetailRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := APIClientForAnEmptyProject(t, ctx)

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	assert.NoError(t, err)

	// Create config
	config := &ConfigWithRows{
		Config: &Config{
			ConfigKey: ConfigKey{
				BranchID:    branch.ID,
				ComponentID: "ex-generic-v2",
			},
			Name:              "Test Search Detail",
			Description:       "Test for SearchJobsDetailRequest",
			ChangeDescription: "Test",
			Content:           orderedmap.New(),
		},
		Rows: []*ConfigRow{},
	}
	_, err = api.CreateConfigRequest(config, true).Send(ctx)
	assert.NoError(t, err)

	// Run a job
	job, err := api.NewCreateJobRequest("ex-generic-v2").WithConfig(config.ID).Send(ctx)
	assert.NoError(t, err)

	// Wait for job to finish
	timeoutCtx, cancelFn := context.WithTimeout(ctx, time.Minute*5)
	defer cancelFn()
	_ = api.WaitForQueueJob(timeoutCtx, job.ID)

	// Search for jobs with detail
	jobs, err := api.SearchJobsDetailRequest(
		WithSearchJobsComponent(ComponentID("ex-generic-v2")),
		WithSearchJobsLimit(10),
	).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, jobs)
	assert.NotEmpty(t, *jobs)

	// Find our job in results and verify extended fields
	found := false
	for _, j := range *jobs {
		if j.ID == job.ID {
			found = true
			assert.Equal(t, ComponentID("ex-generic-v2"), j.ComponentID)
			assert.True(t, j.IsFinished)

			// Verify extended fields are accessible (may be nil/empty depending on job type)
			// The Result struct should be populated with at least a message for failed jobs
			assert.NotEmpty(t, j.Result.Message, "Result message should be populated for finished jobs")
			// Type field indicates job type (standard, orchestrationContainer, etc.)
			assert.NotEmpty(t, j.Type, "Type field should be populated")
			// Metrics may be nil for some job types, but the field should be accessible
			// Input/Output tables depend on job configuration, so we just verify the struct is correctly deserialized
			break
		}
	}
	assert.True(t, found, "Created job should be in search results")
}
