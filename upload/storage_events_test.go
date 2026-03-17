package upload_test

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
	"github.com/keboola/keboola-sdk-go/v2/upload"
)

// createTestTableWithEvents creates a bucket, table, and generates events for testing.
// Returns the TableID of the created table.
// The bucket is automatically cleaned up after the test completes.
func createTestTableWithEvents(t *testing.T, ctx context.Context, api *keboola.AuthorizedAPI, nameSuffix string) keboola.TableID {
	t.Helper()

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	// Create bucket with unique name
	randomSuffix := rand.Int()
	bucket := &keboola.Bucket{
		BucketKey: keboola.BucketKey{
			BranchID: defBranch.ID,
			BucketID: keboola.BucketID{
				Stage:      keboola.BucketStageIn,
				BucketName: fmt.Sprintf("c-test-events-%s-%d", nameSuffix, randomSuffix),
			},
		},
	}
	_, err = api.CreateBucketRequest(bucket).Send(ctx)
	require.NoError(t, err)

	// Register cleanup to delete bucket after test completes
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cleanupCancel()
		_, _ = api.DeleteBucketRequest(bucket.BucketKey, keboola.WithForce()).Send(cleanupCtx)
	})

	tableKey := keboola.TableKey{
		BranchID: defBranch.ID,
		TableID: keboola.TableID{
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
	_, err = upload.Upload(ctx, file, bytes.NewReader(content))
	require.NoError(t, err)

	// Create table from file (this generates events)
	_, err = api.CreateTableFromFileRequest(tableKey, file.FileKey).Send(ctx)
	require.NoError(t, err)

	return tableKey.TableID
}

func TestListTableEvents(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

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
	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Create table with events using helper
	tableID := createTestTableWithEvents(t, ctx, api, "limit")

	// List events with limit
	events, err := api.ListTableEventsRequest(tableID, keboola.WithTableEventsLimit(1)).Send(ctx)
	require.NoError(t, err)

	// Should have at most 1 event due to limit
	assert.LessOrEqual(t, len(*events), 1)
}
