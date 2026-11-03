package keboola_test

import (
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/client"
	. "github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

// createTestTableWithEvents, TestListTableEvents, TestListTableEventsWithLimit are in
// transfer/storage_events_test.go as they require transfer.Upload.

func TestSendEvent(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	_, api := APIClientForRandomProject(t, ctx)
	event, err := api.CreateEventRequest(&Event{
		ComponentID: "keboola.keboola-as-code",
		Type:        "info",
		Message:     "Test event",
		Params:      map[string]any{"command": "bar1"},
		Results:     map[string]any{"projectId": 123, "error": "err"},
		Duration:    client.DurationSeconds(123456 * time.Millisecond),
		RunID:       "123",
	}).Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, event.ID)
}

func TestSendEventInBranch(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	_, api := APIClientForRandomProject(t, ctx)
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)
	event, err := api.CreateEventInBranchRequest(defBranch.ID, &Event{
		ComponentID: "keboola.keboola-as-code",
		Type:        "info",
		Message:     "Test event",
		Params:      map[string]any{"command": "bar1"},
		Results:     map[string]any{"projectId": 123, "error": "err"},
		Duration:    client.DurationSeconds(123456 * time.Millisecond),
		RunID:       "123",
	}).Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, event.ID)
}

func TestCreateEventInBranchRequest(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	c, transport := client.NewMockedClient()
	idx := &Index{Services: Services{{ID: "storage", URL: "https://connection.keboola.mock"}}}
	api := NewPublicAPIFromIndex("https://connection.keboola.mock", idx, WithClient(&c)).NewAuthorizedAPI("my-token", time.Minute)

	bodies := make(map[string]string)
	responder := func(req *http.Request) (*http.Response, error) {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		bodies[req.URL.Path] = string(body)
		return httpmock.NewJsonResponse(http.StatusCreated, map[string]any{"uuid": "event-id"})
	}
	transport.RegisterResponder(http.MethodPost, "https://connection.keboola.mock/v2/storage/events", responder)
	transport.RegisterResponder(http.MethodPost, "https://connection.keboola.mock/v2/storage/branch/456/events", responder)

	newEvent := func() *Event {
		return &Event{
			ComponentID: "keboola.keboola-as-code",
			Type:        "info",
			Message:     "Test event",
			Params:      map[string]any{"command": "bar1"},
			Results:     map[string]any{"projectId": 123},
			RunID:       "123",
		}
	}
	_, err := api.CreateEventRequest(newEvent()).Send(ctx)
	require.NoError(t, err)
	event, err := api.CreateEventInBranchRequest(456, newEvent()).Send(ctx)
	require.NoError(t, err)

	assert.Equal(t, EventID("event-id"), event.ID)
	require.Contains(t, bodies, "/v2/storage/branch/456/events")
	assert.JSONEq(t, bodies["/v2/storage/events"], bodies["/v2/storage/branch/456/events"])
	assert.JSONEq(t, `{
		"component": "keboola.keboola-as-code",
		"type": "info",
		"message": "Test event",
		"duration": 0,
		"params": "{\"command\":\"bar1\"}",
		"results": "{\"projectId\":123}",
		"runId": "123"
	}`, bodies["/v2/storage/branch/456/events"])
}
