package keboola_test

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
)

func TestWorkspaceTableExport(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	ctx, cancelFn := context.WithTimeout(ctx, time.Second*30)
	defer cancelFn()

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
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
	t.Cleanup(func() {
		// Clean up workspace after test
		_, err := api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).Send(ctx)
		require.NoError(t, err)
	})

	// Test the workspace table export functionality
	// Note: A full end-to-end test would require:
	// 1. Creating a workspace (done above)
	// 2. Loading a Storage table into the workspace
	// 3. Exporting the table from the workspace
	// 4. Verifying the exported file
	//
	// Since workspace data loading is complex and may not be implemented in all backends,
	// these tests focus on verifying the API structure and error handling.

	t.Run("BuildRequest", func(t *testing.T) {
		t.Parallel()
		// Test building the request without sending
		req := api.NewWorkspaceTableExportRequest(defBranch.ID, createdWorkspace.ID, "test_table").
			WithFileName("exported_table.csv").
			WithFormat("csv").
			Build()

		require.NotNil(t, req)
	})

	t.Run("SendRequestWithAllOptions", func(t *testing.T) {
		t.Parallel()
		// Test sending a request with all options
		// This will fail because the table doesn't exist in the workspace
		_, err := api.NewWorkspaceTableExportRequest(defBranch.ID, createdWorkspace.ID, "test_table").
			WithFileName("exported_table.csv").
			WithFormat("csv").
			Send(ctx)
		require.Error(t, err)
	})

	t.Run("SendRequestMinimal", func(t *testing.T) {
		t.Parallel()
		// Test sending a minimal request with only required parameters
		// This will fail because the table doesn't exist in the workspace
		_, err := api.NewWorkspaceTableExportRequest(defBranch.ID, createdWorkspace.ID, "test_table").
			Send(ctx)
		require.Error(t, err)
	})

	t.Run("SendAndWait", func(t *testing.T) {
		t.Parallel()
		// Test SendAndWait method
		// This will fail because the table doesn't exist in the workspace
		_, err := api.NewWorkspaceTableExportRequest(defBranch.ID, createdWorkspace.ID, "test_table").
			WithFileName("exported_table.csv").
			WithFormat("csv").
			SendAndWait(ctx, time.Minute*2)
		require.Error(t, err)
	})
}

// TestWorkspaceTableExportSuccess tests the successful export of a table from workspace.
func TestWorkspaceTableExportSuccess(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	// Create bucket and table in Storage
	bucket, tableKey := createBucketAndTableKey(defBranch)
	_, err = api.CreateBucketRequest(bucket).Send(ctx)
	require.NoError(t, err)

	// Create file with test data
	fileName := "test_data.csv"
	file, err := api.CreateFileResourceRequest(defBranch.ID, fileName).Send(ctx)
	require.NoError(t, err)

	// Upload file with test data
	content := []byte("col1,col2\nval1,val2\nval3,val4\n")
	_, err = keboola.Upload(ctx, file, bytes.NewReader(content))
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
	assert.Equal(t, keboola.StorageWorkspaceBackendSnowflake, createdWorkspace.StorageWorkspaceDetails.Backend)
	assert.Equal(t, keboola.StorageWorkspaceBackendSizeMedium, *createdWorkspace.BackendSize)
	assert.Equal(t, string(keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair), *createdWorkspace.StorageWorkspaceDetails.LoginType)
	t.Cleanup(func() {
		_, err := api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).Send(ctx)
		require.NoError(t, err)
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

	// Wait for load job to complete
	waitCtx1, waitCancelFn1 := context.WithTimeout(ctx, time.Minute*5)
	defer waitCancelFn1()
	err = api.WaitForStorageJob(waitCtx1, job)
	require.NoError(t, err)
	assert.Equal(t, "success", job.Status)

	// Export the table from workspace
	result, err := api.NewWorkspaceTableExportRequest(defBranch.ID, createdWorkspace.ID, "test_table").
		WithFileName("test_export.csv").
		WithFormat("csv").
		SendAndWait(ctx, time.Second*30)

	// Verify successful export
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotZero(t, result.FileID)
	require.Equal(t, "test_export.csv", result.FileName)
	require.NotZero(t, result.SizeBytes)

	// Verify the exported file exists in Storage
	fileKey := keboola.FileKey{
		BranchID: defBranch.ID,
		FileID:   result.FileID,
	}
	file2, err := api.GetFileRequest(fileKey).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, file2)
	require.Equal(t, result.FileID, file2.FileID)
	require.Equal(t, defBranch.ID, file2.BranchID)
}
