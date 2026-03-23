package transfer_test

import (
	"bytes"
	"context"
	"os"
	"testing"
	"time"

	"github.com/keboola/go-utils/pkg/testproject"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
	"github.com/keboola/keboola-sdk-go/v2/transfer"
)

func TestStorageWorkspaceLoadData(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	// Create bucket and table
	bucket, tableKey := createBucketAndTableKey(defBranch)
	_, err = api.CreateBucketRequest(bucket).Send(ctx)
	require.NoError(t, err)

	// Create file with test data
	fileName := "test_data.csv"
	file, err := api.CreateFileResourceRequest(defBranch.ID, fileName).Send(ctx)
	require.NoError(t, err)

	// Upload file
	content := []byte("col1,col2\nval1,val2\nval3,val4\n")
	_, err = transfer.Upload(ctx, file, bytes.NewReader(content))
	require.NoError(t, err)

	// Create table from file
	_, err = api.CreateTableFromFileRequest(tableKey, file.FileKey).Send(ctx)
	require.NoError(t, err)

	// Create workspace
	networkPolicy := "user"
	workspace := &keboola.StorageWorkspacePayload{
		Backend:       keboola.StorageWorkspaceBackendSnowflake,
		BackendSize:   ptr(keboola.StorageWorkspaceBackendSizeMedium),
		NetworkPolicy: &networkPolicy,
		LoginType:     keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair,
		PublicKey:     ptr(os.Getenv("TEST_SNOWFLAKE_PUBLIC_KEY")), //nolint: forbidigo
	}

	createdWorkspace, err := api.StorageWorkspaceCreateRequest(defBranch.ID, workspace).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, createdWorkspace)

	// Ensure workspace is cleaned up even if test fails
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cleanupCancel()
		_ = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).SendOrErr(cleanupCtx)
	})

	// Load data into workspace
	loadPayload := &keboola.WorkspaceLoadDataPayload{
		Input: []keboola.WorkspaceLoadDataInput{
			{
				Source:      tableKey.TableID.String(),
				Destination: "test_table",
			},
		},
	}

	job, err := api.StorageWorkspaceLoadDataRequest(defBranch.ID, createdWorkspace.ID, loadPayload).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, job)

	// Wait for job to complete
	waitCtx, waitCancelFn := context.WithTimeout(ctx, time.Minute*5)
	defer waitCancelFn()
	err = api.WaitForStorageJob(waitCtx, job)
	require.NoError(t, err)
	assert.Equal(t, "success", job.Status)
}
