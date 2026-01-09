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

func TestSearchJobsOptions(t *testing.T) {
	t.Parallel()

	t.Run("WithSearchJobsBranch", func(t *testing.T) {
		t.Parallel()
		config := SearchJobsConfig{}
		WithSearchJobsBranch(BranchID(123))(&config)
		assert.Equal(t, BranchID(123), config.BranchID)
	})

	t.Run("WithSearchJobsComponent", func(t *testing.T) {
		t.Parallel()
		config := SearchJobsConfig{}
		WithSearchJobsComponent(ComponentID("my-component"))(&config)
		assert.Equal(t, ComponentID("my-component"), config.ComponentID)
	})

	t.Run("WithSearchJobsConfig", func(t *testing.T) {
		t.Parallel()
		config := SearchJobsConfig{}
		WithSearchJobsConfig(ConfigID("my-config"))(&config)
		assert.Equal(t, ConfigID("my-config"), config.ConfigID)
	})

	t.Run("WithSearchJobsStatus", func(t *testing.T) {
		t.Parallel()
		config := SearchJobsConfig{}
		WithSearchJobsStatus("success")(&config)
		assert.Equal(t, "success", config.Status)
	})

	t.Run("WithSearchJobsLimit", func(t *testing.T) {
		t.Parallel()
		config := SearchJobsConfig{}
		WithSearchJobsLimit(50)(&config)
		assert.Equal(t, 50, config.Limit)
	})

	t.Run("WithSearchJobsOffset", func(t *testing.T) {
		t.Parallel()
		config := SearchJobsConfig{}
		WithSearchJobsOffset(10)(&config)
		assert.Equal(t, 10, config.Offset)
	})

	t.Run("multiple options", func(t *testing.T) {
		t.Parallel()
		config := SearchJobsConfig{}
		WithSearchJobsBranch(BranchID(456))(&config)
		WithSearchJobsComponent(ComponentID("ex-generic-v2"))(&config)
		WithSearchJobsStatus("waiting")(&config)
		WithSearchJobsLimit(25)(&config)
		WithSearchJobsOffset(5)(&config)

		assert.Equal(t, BranchID(456), config.BranchID)
		assert.Equal(t, ComponentID("ex-generic-v2"), config.ComponentID)
		assert.Equal(t, "waiting", config.Status)
		assert.Equal(t, 25, config.Limit)
		assert.Equal(t, 5, config.Offset)
	})
}

func TestSearchJobsAPICalls(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	_, api := APIClientForAnEmptyProject(t, ctx)

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, branch)

	// Create config for the job
	config := &ConfigWithRows{
		Config: &Config{
			ConfigKey: ConfigKey{
				BranchID:    branch.ID,
				ComponentID: "ex-generic-v2",
			},
			Name:              "Test SearchJobs",
			Description:       "Test config for SearchJobs API",
			ChangeDescription: "Test",
			Content:           orderedmap.New(),
		},
		Rows: []*ConfigRow{},
	}
	_, err = api.CreateConfigRequest(config, true).Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, config.ID)

	// Run a job on the config
	job, err := api.NewCreateJobRequest("ex-generic-v2").WithConfig(config.ID).Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, job.ID)

	// Wait for the job to finish (it will fail due to empty config, but that's OK)
	timeoutCtx, cancelFn := context.WithTimeout(ctx, time.Minute*5)
	defer cancelFn()
	_ = api.WaitForQueueJob(timeoutCtx, job.ID) // Ignore error, job will fail

	// Test SearchJobsRequest - search for jobs by component
	jobs, err := api.SearchJobsRequest(
		WithSearchJobsComponent(ComponentID("ex-generic-v2")),
		WithSearchJobsLimit(10),
	).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, jobs)
	assert.NotEmpty(t, *jobs, "Expected at least one job")

	// Verify the job we created is in the results
	found := false
	for _, j := range *jobs {
		if j.ID == job.ID {
			found = true
			assert.Equal(t, ComponentID("ex-generic-v2"), j.ComponentID)
			assert.Equal(t, config.ID, j.ConfigID)
			assert.Equal(t, branch.ID, j.BranchID)
			break
		}
	}
	assert.True(t, found, "Created job should be in search results")

	// Test SearchJobsRequest - search by status
	// The job we created should have failed, so search for "error" status
	jobsByStatus, err := api.SearchJobsRequest(
		WithSearchJobsStatus("error"),
		WithSearchJobsComponent(ComponentID("ex-generic-v2")),
		WithSearchJobsLimit(10),
	).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, jobsByStatus)
	// Verify our failed job is in the error status results
	foundInStatus := false
	for _, j := range *jobsByStatus {
		if j.ID == job.ID {
			foundInStatus = true
			assert.Equal(t, "error", j.Status)
			break
		}
	}
	assert.True(t, foundInStatus, "Failed job should be found when searching by error status")

	// Test SearchJobsRequest - search by branch
	jobsByBranch, err := api.SearchJobsRequest(
		WithSearchJobsBranch(branch.ID),
		WithSearchJobsLimit(10),
	).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, jobsByBranch)
	assert.NotEmpty(t, *jobsByBranch)

	// Test SearchJobsRequest - search by config
	jobsByConfig, err := api.SearchJobsRequest(
		WithSearchJobsConfig(config.ID),
		WithSearchJobsLimit(10),
	).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, jobsByConfig)
	assert.NotEmpty(t, *jobsByConfig)

	// Test SearchJobsRequest - combined filters
	jobsCombined, err := api.SearchJobsRequest(
		WithSearchJobsBranch(branch.ID),
		WithSearchJobsComponent(ComponentID("ex-generic-v2")),
		WithSearchJobsConfig(config.ID),
		WithSearchJobsLimit(5),
	).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, jobsCombined)
	assert.NotEmpty(t, *jobsCombined)

	// Test SearchJobsRequest - pagination with offset
	// First, get jobs with limit 1 and offset 0
	jobsPage1, err := api.SearchJobsRequest(
		WithSearchJobsComponent(ComponentID("ex-generic-v2")),
		WithSearchJobsLimit(1),
		WithSearchJobsOffset(0),
	).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, jobsPage1)
	assert.Len(t, *jobsPage1, 1, "Expected exactly 1 job with limit=1")

	// Get jobs with limit 1 and offset 1 (should return different job if more exist)
	jobsPage2, err := api.SearchJobsRequest(
		WithSearchJobsComponent(ComponentID("ex-generic-v2")),
		WithSearchJobsLimit(1),
		WithSearchJobsOffset(1),
	).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, jobsPage2)

	// If there are multiple jobs, the offset should return different results
	if len(*jobsPage2) > 0 {
		assert.NotEqual(t, (*jobsPage1)[0].ID, (*jobsPage2)[0].ID,
			"Pagination should return different jobs at different offsets")
	}

	// Test SearchJobsRequest - default (no options)
	jobsDefault, err := api.SearchJobsRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, jobsDefault)
}
