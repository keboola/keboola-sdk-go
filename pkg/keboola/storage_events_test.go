package keboola_test

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/client"
	. "github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

// createTestTableWithEvents creates a bucket, table, and generates events for testing.
// Returns the TableID of the created table.
func createTestTableWithEvents(t *testing.T, ctx context.Context, api *AuthorizedAPI, nameSuffix string) TableID {
	t.Helper()

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	// Create bucket with unique name
	// Use single random number for all resource names for consistency
	randomSuffix := rand.Int()
	bucket := &Bucket{
		BucketKey: BucketKey{
			BranchID: defBranch.ID,
			BucketID: BucketID{
				Stage:      BucketStageIn,
				BucketName: fmt.Sprintf("c-test-events-%s-%d", nameSuffix, randomSuffix),
			},
		},
	}
	_, err = api.CreateBucketRequest(bucket).Send(ctx)
	require.NoError(t, err)

	tableKey := TableKey{
		BranchID: defBranch.ID,
		TableID: TableID{
			BucketID:  bucket.BucketID,
			TableName: fmt.Sprintf("test_events_%s_%d", nameSuffix, randomSuffix),
		},
	}

	// Create file for table data
	fileName := fmt.Sprintf("file_events_%s_%d", nameSuffix, randomSuffix)
	file, err := api.CreateFileResourceRequest(defBranch.ID, fileName).Send(ctx)
	require.NoError(t, err)

	// Upload file
	content := []byte("col1,col2\nval1,val2\n")
	_, err = Upload(ctx, file, bytes.NewReader(content))
	require.NoError(t, err)

	// Create table from file (this generates events)
	_, err = api.CreateTableFromFileRequest(tableKey, file.FileKey).Send(ctx)
	require.NoError(t, err)

	return tableKey.TableID
}

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

func TestListTableEvents(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := APIClientForAnEmptyProject(t, ctx)

	// Create table with events using helper
	tableID := createTestTableWithEvents(t, ctx, api, "basic")

	// List events for the table
	events, err := api.ListTableEventsRequest(tableID).Send(ctx)
	require.NoError(t, err)

	// Should have at least one event (tableCreated or tableImportDone)
	assert.NotEmpty(t, *events)

	// Verify event structure
	for _, event := range *events {
		assert.NotEmpty(t, event.ID)
		assert.NotEmpty(t, event.Event)
		assert.NotEmpty(t, event.Created)
	}
}

func TestListTableEventsWithLimit(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := APIClientForAnEmptyProject(t, ctx)

	// Create table with events using helper
	tableID := createTestTableWithEvents(t, ctx, api, "limit")

	// List events with limit
	events, err := api.ListTableEventsRequest(tableID, WithTableEventsLimit(1)).Send(ctx)
	require.NoError(t, err)

	// Should have at most 1 event due to limit
	assert.LessOrEqual(t, len(*events), 1)
}
