package keboola_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/keboola/keboola-sdk-go/v2/pkg/client"
	. "github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

func TestSendEvent(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
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
