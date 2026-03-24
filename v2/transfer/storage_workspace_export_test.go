package transfer_test

import (
	"bytes"
	"compress/gzip"
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

// TestWorkspaceTableExportSuccess tests the successful export of a table from workspace.
func TestWorkspaceTableExportSuccess(t *testing.T) {
	t.Skip()
	t.Parallel()
	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	t.Cleanup(cancelFn)

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
		_, _ = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).Send(cleanupCtx)
	})

	assert.Equal(t, keboola.StorageWorkspaceBackendSnowflake, createdWorkspace.StorageWorkspaceDetails.Backend)
	assert.Equal(t, keboola.StorageWorkspaceBackendSizeMedium, *createdWorkspace.BackendSize)
	assert.Equal(t, string(keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair), *createdWorkspace.StorageWorkspaceDetails.LoginType)

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

	// Wait for load job to complete
	waitCtx1, waitCancelFn1 := context.WithTimeout(ctx, time.Minute*5)
	t.Cleanup(waitCancelFn1)
	err = api.WaitForStorageJob(waitCtx1, job)
	require.NoError(t, err)
	assert.Equal(t, "success", job.Status)

	// Export the table from workspace
	result, err := api.NewWorkspaceTableExportRequest(defBranch.ID, createdWorkspace.ID, "test_table").
		WithFileName("test_export.csv").
		WithFileType("csv").
		SendAndWait(ctx, time.Second*30)

	// Verify successful export
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotZero(t, result.File.FileID)

	// Verify the exported file exists in Storage
	fileKey := keboola.FileKey{
		BranchID: defBranch.ID,
		FileID:   result.File.FileID,
	}
	file2, err := api.GetFileRequest(fileKey).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, file2)
	require.Equal(t, result.File.FileID, file2.FileID)
	require.Equal(t, defBranch.ID, file2.BranchID)
}

// TestWorkspaceTableExportGzip tests that WithGzip option controls whether the exported file is gzip-compressed.
func TestWorkspaceTableExportGzip(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	t.Cleanup(cancelFn)

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	// Create bucket and table in Storage
	bucket, tableKey := createBucketAndTableKey(defBranch)
	_, err = api.CreateBucketRequest(bucket).Send(ctx)
	require.NoError(t, err)

	// Create file with test data
	fileName := "test_data_gzip.csv"
	file, err := api.CreateFileResourceRequest(defBranch.ID, fileName).Send(ctx)
	require.NoError(t, err)

	// Upload file with test data
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
		_, _ = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).Send(cleanupCtx)
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

	waitCtx1, waitCancelFn1 := context.WithTimeout(ctx, time.Minute*5)
	t.Cleanup(waitCancelFn1)
	err = api.WaitForStorageJob(waitCtx1, job)
	require.NoError(t, err)
	assert.Equal(t, "success", job.Status)

	// downloadExportedFile downloads the exported file raw bytes.
	downloadExportedFile := func(t *testing.T, fileID keboola.FileID) []byte {
		t.Helper()
		fileKey := keboola.FileKey{
			BranchID: defBranch.ID,
			FileID:   fileID,
		}
		fileCreds, err := api.GetFileWithCredentialsRequest(fileKey).Send(ctx)
		require.NoError(t, err)

		var rawData []byte
		if fileCreds.IsSliced {
			slices, sliceErr := transfer.DownloadManifest(ctx, fileCreds)
			require.NoError(t, sliceErr)
			for _, slice := range slices {
				sliceData, dlErr := transfer.DownloadSlice(ctx, fileCreds, slice)
				require.NoError(t, dlErr)
				rawData = append(rawData, sliceData...)
			}
		} else {
			rawData, err = transfer.Download(ctx, fileCreds)
			require.NoError(t, err)
		}
		require.NotEmpty(t, rawData, "exported file should not be empty")
		return rawData
	}

	// Export WITHOUT gzip — data should not be valid gzip
	t.Run("without_gzip", func(t *testing.T) {
		t.Parallel()
		result, err := api.NewWorkspaceTableExportRequest(defBranch.ID, createdWorkspace.ID, "test_table").
			WithFileName("test_export_plain").
			WithFileType("csv").
			SendAndWait(ctx, time.Minute*2)
		require.NoError(t, err)
		require.NotZero(t, result.File.FileID)

		rawData := downloadExportedFile(t, result.File.FileID)
		_, gzErr := gzip.NewReader(bytes.NewReader(rawData))
		require.Error(t, gzErr, "exported file without gzip should not be valid gzip")
	})

	// Export WITH gzip — data should be valid gzip
	t.Run("with_gzip", func(t *testing.T) {
		t.Parallel()
		result, err := api.NewWorkspaceTableExportRequest(defBranch.ID, createdWorkspace.ID, "test_table").
			WithFileName("test_export_gzip.csv").
			WithFileType("csv").
			WithGzip(true).
			SendAndWait(ctx, time.Minute*2)
		require.NoError(t, err)
		require.NotZero(t, result.File.FileID)

		rawData := downloadExportedFile(t, result.File.FileID)
		gzReader, gzErr := gzip.NewReader(bytes.NewReader(rawData))
		require.NoError(t, gzErr, "exported file with gzip should be valid gzip")
		require.NoError(t, gzReader.Close())
	})
}
