package keboola

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
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

func TestSearchJobsRequest(t *testing.T) {
	t.Parallel()

	setupMockedAPI := func(t *testing.T) (*AuthorizedAPI, *httpmock.MockTransport) {
		t.Helper()
		c, transport := client.NewMockedClient()
		c = c.WithBaseURL("https://connection.test")
		transport.RegisterResponder("GET", `https://connection.test/v2/storage/?exclude=components`, newJSONResponder(`{
			"services": [
				{
					"id": "queue",
					"url": "https://queue.connection.test"
				}
			],
			"features": []
		}`))
		api, err := NewAuthorizedAPI(context.Background(), "https://connection.test", "my-token", WithClient(&c))
		assert.NoError(t, err)
		return api, transport
	}

	t.Run("default request with limit 100", func(t *testing.T) {
		t.Parallel()
		api, transport := setupMockedAPI(t)

		transport.RegisterResponder("GET", `=~^https://queue\.connection\.test/search/jobs`,
			func(req *http.Request) (*http.Response, error) {
				// Verify default limit is 100
				assert.Equal(t, "100", req.URL.Query().Get("limit"))
				// Verify offset is 0 (always added)
				assert.Equal(t, "0", req.URL.Query().Get("offset"))
				return httpmock.NewJsonResponse(200, []*QueueJob{})
			})

		jobs, err := api.SearchJobsRequest().Send(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, jobs)
		assert.Equal(t, 1, transport.GetCallCountInfo()["GET =~^https://queue\\.connection\\.test/search/jobs"])
	})

	t.Run("with branch filter", func(t *testing.T) {
		t.Parallel()
		api, transport := setupMockedAPI(t)

		transport.RegisterResponder("GET", `=~^https://queue\.connection\.test/search/jobs`,
			func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "123", req.URL.Query().Get("branchId"))
				return httpmock.NewJsonResponse(200, []*QueueJob{})
			})

		_, err := api.SearchJobsRequest(WithSearchJobsBranch(BranchID(123))).Send(context.Background())
		assert.NoError(t, err)
	})

	t.Run("with component filter", func(t *testing.T) {
		t.Parallel()
		api, transport := setupMockedAPI(t)

		transport.RegisterResponder("GET", `=~^https://queue\.connection\.test/search/jobs`,
			func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "ex-generic-v2", req.URL.Query().Get("component"))
				return httpmock.NewJsonResponse(200, []*QueueJob{})
			})

		_, err := api.SearchJobsRequest(WithSearchJobsComponent(ComponentID("ex-generic-v2"))).Send(context.Background())
		assert.NoError(t, err)
	})

	t.Run("with config filter", func(t *testing.T) {
		t.Parallel()
		api, transport := setupMockedAPI(t)

		transport.RegisterResponder("GET", `=~^https://queue\.connection\.test/search/jobs`,
			func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "config-123", req.URL.Query().Get("config"))
				return httpmock.NewJsonResponse(200, []*QueueJob{})
			})

		_, err := api.SearchJobsRequest(WithSearchJobsConfig(ConfigID("config-123"))).Send(context.Background())
		assert.NoError(t, err)
	})

	t.Run("with status filter", func(t *testing.T) {
		t.Parallel()
		api, transport := setupMockedAPI(t)

		transport.RegisterResponder("GET", `=~^https://queue\.connection\.test/search/jobs`,
			func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "success", req.URL.Query().Get("status"))
				return httpmock.NewJsonResponse(200, []*QueueJob{})
			})

		_, err := api.SearchJobsRequest(WithSearchJobsStatus("success")).Send(context.Background())
		assert.NoError(t, err)
	})

	t.Run("with custom limit", func(t *testing.T) {
		t.Parallel()
		api, transport := setupMockedAPI(t)

		transport.RegisterResponder("GET", `=~^https://queue\.connection\.test/search/jobs`,
			func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "50", req.URL.Query().Get("limit"))
				return httpmock.NewJsonResponse(200, []*QueueJob{})
			})

		_, err := api.SearchJobsRequest(WithSearchJobsLimit(50)).Send(context.Background())
		assert.NoError(t, err)
	})

	t.Run("with limit 0 omits limit parameter", func(t *testing.T) {
		t.Parallel()
		api, transport := setupMockedAPI(t)

		transport.RegisterResponder("GET", `=~^https://queue\.connection\.test/search/jobs`,
			func(req *http.Request) (*http.Response, error) {
				// limit should NOT be present when set to 0
				assert.Empty(t, req.URL.Query().Get("limit"))
				return httpmock.NewJsonResponse(200, []*QueueJob{})
			})

		_, err := api.SearchJobsRequest(WithSearchJobsLimit(0)).Send(context.Background())
		assert.NoError(t, err)
	})

	t.Run("with negative limit omits limit parameter", func(t *testing.T) {
		t.Parallel()
		api, transport := setupMockedAPI(t)

		transport.RegisterResponder("GET", `=~^https://queue\.connection\.test/search/jobs`,
			func(req *http.Request) (*http.Response, error) {
				// limit should NOT be present when negative
				assert.Empty(t, req.URL.Query().Get("limit"))
				return httpmock.NewJsonResponse(200, []*QueueJob{})
			})

		_, err := api.SearchJobsRequest(WithSearchJobsLimit(-1)).Send(context.Background())
		assert.NoError(t, err)
	})

	t.Run("with offset", func(t *testing.T) {
		t.Parallel()
		api, transport := setupMockedAPI(t)

		transport.RegisterResponder("GET", `=~^https://queue\.connection\.test/search/jobs`,
			func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "10", req.URL.Query().Get("offset"))
				return httpmock.NewJsonResponse(200, []*QueueJob{})
			})

		_, err := api.SearchJobsRequest(WithSearchJobsOffset(10)).Send(context.Background())
		assert.NoError(t, err)
	})

	t.Run("with all options combined", func(t *testing.T) {
		t.Parallel()
		api, transport := setupMockedAPI(t)

		transport.RegisterResponder("GET", `=~^https://queue\.connection\.test/search/jobs`,
			func(req *http.Request) (*http.Response, error) {
				query := req.URL.Query()
				assert.Equal(t, "789", query.Get("branchId"))
				assert.Equal(t, "keboola.ex-db-mysql", query.Get("component"))
				assert.Equal(t, "my-config-id", query.Get("config"))
				assert.Equal(t, "waiting", query.Get("status"))
				assert.Equal(t, "25", query.Get("limit"))
				assert.Equal(t, "5", query.Get("offset"))
				return httpmock.NewJsonResponse(200, []*QueueJob{})
			})

		_, err := api.SearchJobsRequest(
			WithSearchJobsBranch(BranchID(789)),
			WithSearchJobsComponent(ComponentID("keboola.ex-db-mysql")),
			WithSearchJobsConfig(ConfigID("my-config-id")),
			WithSearchJobsStatus("waiting"),
			WithSearchJobsLimit(25),
			WithSearchJobsOffset(5),
		).Send(context.Background())
		assert.NoError(t, err)
	})
}

func TestSearchJobsResponseParsing(t *testing.T) {
	t.Parallel()

	c, transport := client.NewMockedClient()
	c = c.WithBaseURL("https://connection.test")
	transport.RegisterResponder("GET", `https://connection.test/v2/storage/?exclude=components`, newJSONResponder(`{
		"services": [
			{
				"id": "queue",
				"url": "https://queue.connection.test"
			}
		],
		"features": []
	}`))

	// Response with all QueueJob fields including BranchID as string (Queue API format)
	transport.RegisterResponder("GET", `=~^https://queue\.connection\.test/search/jobs`, newJSONResponder(`[
		{
			"id": "12345",
			"status": "success",
			"isFinished": true,
			"url": "https://queue.connection.test/jobs/12345",
			"result": {
				"message": "Job completed successfully"
			},
			"createdTime": "2024-01-15T10:30:00Z",
			"startTime": "2024-01-15T10:30:05Z",
			"endTime": "2024-01-15T10:35:00Z",
			"component": "ex-generic-v2",
			"config": "config-123",
			"branchId": "789",
			"operationName": "run",
			"durationSeconds": 295,
			"runId": "run-abc-123"
		},
		{
			"id": 67890,
			"status": "waiting",
			"isFinished": false,
			"url": "https://queue.connection.test/jobs/67890",
			"result": {},
			"createdTime": "2024-01-15T11:00:00Z",
			"component": "keboola.ex-db-mysql",
			"config": "config-456",
			"branchId": "999"
		}
	]`))

	api, err := NewAuthorizedAPI(context.Background(), "https://connection.test", "my-token", WithClient(&c))
	assert.NoError(t, err)

	jobs, err := api.SearchJobsRequest().Send(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, jobs)
	assert.Len(t, *jobs, 2)

	// Verify first job with all fields
	job1 := (*jobs)[0]
	assert.Equal(t, JobID("12345"), job1.ID)
	assert.Equal(t, "success", job1.Status)
	assert.True(t, job1.IsFinished)
	assert.Equal(t, "https://queue.connection.test/jobs/12345", job1.URL)
	assert.Equal(t, "Job completed successfully", job1.Result.Message)
	assert.Equal(t, ComponentID("ex-generic-v2"), job1.ComponentID)
	assert.Equal(t, ConfigID("config-123"), job1.ConfigID)
	assert.Equal(t, BranchID(789), job1.BranchID) // BranchID parsed from string
	assert.Equal(t, "run", job1.OperationName)
	assert.Equal(t, 295, job1.DurationSeconds)
	assert.Equal(t, "run-abc-123", job1.RunID)
	assert.NotNil(t, job1.StartTime)
	assert.NotNil(t, job1.EndTime)

	// Verify second job with numeric ID
	job2 := (*jobs)[1]
	assert.Equal(t, JobID("67890"), job2.ID) // JobID handles both string and int
	assert.Equal(t, "waiting", job2.Status)
	assert.False(t, job2.IsFinished)
	assert.Equal(t, ComponentID("keboola.ex-db-mysql"), job2.ComponentID)
	assert.Equal(t, BranchID(999), job2.BranchID) // BranchID parsed from string
	assert.Nil(t, job2.StartTime)
	assert.Nil(t, job2.EndTime)
}
