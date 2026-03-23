package transfer_test

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"os"
	"strings"
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

// TestWorkspaceTableExportGzip tests that exporting a table with gzip=true produces a valid gzip-compressed file.
func TestWorkspaceTableExportGzip(t *testing.T) {
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

	// Export the table from workspace with gzip enabled
	result, err := api.NewWorkspaceTableExportRequest(defBranch.ID, createdWorkspace.ID, "test_table").
		WithFileName("test_export_gzip.csv").
		WithFileType("csv").
		WithGzip(true).
		SendAndWait(ctx, time.Minute*2)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotZero(t, result.File.FileID)

	// Download the exported file and verify it is gzip-compressed
	fileKey := keboola.FileKey{
		BranchID: defBranch.ID,
		FileID:   result.File.FileID,
	}
	fileCreds, err := api.GetFileWithCredentialsRequest(fileKey).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, fileCreds)

	// Download and concatenate all slices
	var rawData []byte
	if fileCreds.IsSliced {
		slices, sliceErr := keboola.DownloadManifest(ctx, fileCreds)
		require.NoError(t, sliceErr)
		for _, slice := range slices {
			sliceData, dlErr := keboola.DownloadSlice(ctx, fileCreds, slice)
			require.NoError(t, dlErr)
			rawData = append(rawData, sliceData...)
		}
	} else {
		rawData, err = keboola.Download(ctx, fileCreds)
		require.NoError(t, err)
	}
	require.NotEmpty(t, rawData, "exported file should not be empty")

	// Verify the data is valid gzip by decompressing it
	gzReader, err := gzip.NewReader(bytes.NewReader(rawData))
	require.NoError(t, err, "exported file should be valid gzip data")

	decompressed, err := io.ReadAll(gzReader)
	require.NoError(t, err, "should be able to decompress gzip data")
	require.NoError(t, gzReader.Close())

	// Verify decompressed content contains the original CSV data
	decompressedStr := string(decompressed)
	assert.True(t, strings.Contains(decompressedStr, "col1"), "decompressed data should contain column headers")
	assert.True(t, strings.Contains(decompressedStr, "val1"), "decompressed data should contain row values")
}
