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

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	// Create bucket
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	bucket := &Bucket{
		BucketKey: BucketKey{
			BranchID: defBranch.ID,
			BucketID: BucketID{
				Stage:      BucketStageIn,
				BucketName: fmt.Sprintf("c-test-events-%d", rng.Int()),
			},
		},
	}
	_, err = api.CreateBucketRequest(bucket).Send(ctx)
	require.NoError(t, err)

	tableKey := TableKey{
		BranchID: defBranch.ID,
		TableID: TableID{
			BucketID:  bucket.BucketID,
			TableName: fmt.Sprintf("test_events_%d", rng.Int()),
		},
	}

	// Create file for table data
	fileName := fmt.Sprintf("file_events_%d", rng.Int())
	file, err := api.CreateFileResourceRequest(defBranch.ID, fileName).Send(ctx)
	require.NoError(t, err)

	// Upload file
	content := []byte("col1,col2\nval1,val2\n")
	_, err = Upload(ctx, file, bytes.NewReader(content))
	require.NoError(t, err)

	// Create table from file (this generates events)
	_, err = api.CreateTableFromFileRequest(tableKey, file.FileKey).Send(ctx)
	require.NoError(t, err)

	// List events for the table
	events, err := api.ListTableEventsRequest(tableKey.TableID).Send(ctx)
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

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	// Create bucket
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	bucket := &Bucket{
		BucketKey: BucketKey{
			BranchID: defBranch.ID,
			BucketID: BucketID{
				Stage:      BucketStageIn,
				BucketName: fmt.Sprintf("c-test-events-limit-%d", rng.Int()),
			},
		},
	}
	_, err = api.CreateBucketRequest(bucket).Send(ctx)
	require.NoError(t, err)

	tableKey := TableKey{
		BranchID: defBranch.ID,
		TableID: TableID{
			BucketID:  bucket.BucketID,
			TableName: fmt.Sprintf("test_events_limit_%d", rng.Int()),
		},
	}

	// Create file for table data
	fileName := fmt.Sprintf("file_events_limit_%d", rng.Int())
	file, err := api.CreateFileResourceRequest(defBranch.ID, fileName).Send(ctx)
	require.NoError(t, err)

	// Upload file
	content := []byte("col1,col2\nval1,val2\n")
	_, err = Upload(ctx, file, bytes.NewReader(content))
	require.NoError(t, err)

	// Create table from file
	_, err = api.CreateTableFromFileRequest(tableKey, file.FileKey).Send(ctx)
	require.NoError(t, err)

	// List events with limit
	events, err := api.ListTableEventsRequest(tableKey.TableID, WithTableEventsLimit(1)).Send(ctx)
	require.NoError(t, err)

	// Should have at most 1 event due to limit
	assert.LessOrEqual(t, len(*events), 1)
}
