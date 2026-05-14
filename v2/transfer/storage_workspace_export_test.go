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

// TestWorkspaceTableExportSuccess tests the successful export of a table from a Snowflake workspace.
func TestWorkspaceTableExportSuccess(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	t.Cleanup(cancelFn)

	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	bucket, tableKey := createBucketAndTableKey(defBranch)
	_, err = api.CreateBucketRequest(bucket).Send(ctx)
	require.NoError(t, err)

	file, err := api.CreateFileResourceRequest(defBranch.ID, "test_data.csv").Send(ctx)
	require.NoError(t, err)

	_, err = transfer.Upload(ctx, file, bytes.NewReader([]byte("col1,col2\nval1,val2\nval3,val4\n")))
	require.NoError(t, err)

	_, err = api.CreateTableFromFileRequest(tableKey, file.FileKey).Send(ctx)
	require.NoError(t, err)

	networkPolicy := "user"
	createdWorkspace, err := api.StorageWorkspaceCreateRequest(defBranch.ID, &keboola.StorageWorkspacePayload{
		Backend:       keboola.StorageWorkspaceBackendSnowflake,
		BackendSize:   new(keboola.StorageWorkspaceBackendSizeMedium),
		NetworkPolicy: &networkPolicy,
		LoginType:     keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair,
		PublicKey:     new(os.Getenv("TEST_SNOWFLAKE_PUBLIC_KEY")), //nolint: forbidigo
	}).Send(ctx)
	require.NoError(t, err)
	t.Cleanup(func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		_, _ = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).Send(cleanupCtx)
	})

	assert.Equal(t, keboola.StorageWorkspaceBackendSnowflake, createdWorkspace.StorageWorkspaceDetails.Backend)
	assert.Equal(t, keboola.StorageWorkspaceBackendSizeMedium, *createdWorkspace.BackendSize)
	assert.Equal(t, string(keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair), *createdWorkspace.StorageWorkspaceDetails.LoginType)

	loadTableIntoWorkspace(t, ctx, api, defBranch.ID, createdWorkspace.ID, tableKey)

	result, err := api.NewWorkspaceTableExportRequest(defBranch.ID, createdWorkspace.ID, "test_table").
		WithFileName("test_export.csv").
		WithFileType("csv").
		SendAndWait(ctx, time.Second*30)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotZero(t, result.File.FileID)

	file2, err := api.GetFileRequest(keboola.FileKey{BranchID: defBranch.ID, FileID: result.File.FileID}).Send(ctx)
	require.NoError(t, err)
	require.Equal(t, result.File.FileID, file2.FileID)
	require.Equal(t, defBranch.ID, file2.BranchID)
}

// TestWorkspaceTableExportSuccessBigQuery tests the successful export of a table from a BigQuery workspace.
func TestWorkspaceTableExportSuccessBigQuery(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithBigQueryBackend())

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	t.Cleanup(cancelFn)

	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	bucket, tableKey := createBucketAndTableKey(defBranch)
	_, err = api.CreateBucketRequest(bucket).Send(ctx)
	require.NoError(t, err)

	file, err := api.CreateFileResourceRequest(defBranch.ID, "test_data.csv").Send(ctx)
	require.NoError(t, err)

	_, err = transfer.Upload(ctx, file, bytes.NewReader([]byte("col1,col2\nval1,val2\nval3,val4\n")))
	require.NoError(t, err)

	_, err = api.CreateTableFromFileRequest(tableKey, file.FileKey).Send(ctx)
	require.NoError(t, err)

	createdWorkspace, err := api.StorageWorkspaceCreateRequest(defBranch.ID, &keboola.StorageWorkspacePayload{
		Backend:   keboola.StorageWorkspaceBackendBigQuery,
		LoginType: keboola.StorageWorkspaceLoginTypeDefault,
	}).Send(ctx)
	require.NoError(t, err)
	t.Cleanup(func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		_, _ = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).Send(cleanupCtx)
	})

	assert.Equal(t, keboola.StorageWorkspaceBackendBigQuery, createdWorkspace.StorageWorkspaceDetails.Backend)

	loadTableIntoWorkspace(t, ctx, api, defBranch.ID, createdWorkspace.ID, tableKey)

	result, err := api.NewWorkspaceTableExportRequest(defBranch.ID, createdWorkspace.ID, "test_table").
		WithFileName("test_export.csv").
		WithFileType("csv").
		SendAndWait(ctx, time.Second*30)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotZero(t, result.File.FileID)

	file2, err := api.GetFileRequest(keboola.FileKey{BranchID: defBranch.ID, FileID: result.File.FileID}).Send(ctx)
	require.NoError(t, err)
	require.Equal(t, result.File.FileID, file2.FileID)
	require.Equal(t, defBranch.ID, file2.BranchID)
}

// TestWorkspaceTableExportGzip tests that WithGzip controls compression for a Snowflake workspace export.
func TestWorkspaceTableExportGzip(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	t.Cleanup(cancelFn)

	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	bucket, tableKey := createBucketAndTableKey(defBranch)
	_, err = api.CreateBucketRequest(bucket).Send(ctx)
	require.NoError(t, err)

	file, err := api.CreateFileResourceRequest(defBranch.ID, "test_data_gzip.csv").Send(ctx)
	require.NoError(t, err)

	_, err = transfer.Upload(ctx, file, bytes.NewReader([]byte("col1,col2\nval1,val2\nval3,val4\n")))
	require.NoError(t, err)

	_, err = api.CreateTableFromFileRequest(tableKey, file.FileKey).Send(ctx)
	require.NoError(t, err)

	networkPolicy := "user"
	createdWorkspace, err := api.StorageWorkspaceCreateRequest(defBranch.ID, &keboola.StorageWorkspacePayload{
		Backend:       keboola.StorageWorkspaceBackendSnowflake,
		BackendSize:   new(keboola.StorageWorkspaceBackendSizeMedium),
		NetworkPolicy: &networkPolicy,
		LoginType:     keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair,
		PublicKey:     new(os.Getenv("TEST_SNOWFLAKE_PUBLIC_KEY")), //nolint: forbidigo
	}).Send(ctx)
	require.NoError(t, err)
	t.Cleanup(func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		_, _ = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).Send(cleanupCtx)
	})

	loadTableIntoWorkspace(t, ctx, api, defBranch.ID, createdWorkspace.ID, tableKey)

	runGzipSubtests(t, ctx, api, defBranch, createdWorkspace.ID)
}

// TestWorkspaceTableExportGzipBigQuery tests that WithGzip controls compression for a BigQuery workspace export.
func TestWorkspaceTableExportGzipBigQuery(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithBigQueryBackend())

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	t.Cleanup(cancelFn)

	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	bucket, tableKey := createBucketAndTableKey(defBranch)
	_, err = api.CreateBucketRequest(bucket).Send(ctx)
	require.NoError(t, err)

	file, err := api.CreateFileResourceRequest(defBranch.ID, "test_data_gzip.csv").Send(ctx)
	require.NoError(t, err)

	_, err = transfer.Upload(ctx, file, bytes.NewReader([]byte("col1,col2\nval1,val2\nval3,val4\n")))
	require.NoError(t, err)

	_, err = api.CreateTableFromFileRequest(tableKey, file.FileKey).Send(ctx)
	require.NoError(t, err)

	createdWorkspace, err := api.StorageWorkspaceCreateRequest(defBranch.ID, &keboola.StorageWorkspacePayload{
		Backend:   keboola.StorageWorkspaceBackendBigQuery,
		LoginType: keboola.StorageWorkspaceLoginTypeDefault,
	}).Send(ctx)
	require.NoError(t, err)
	t.Cleanup(func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		_, _ = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).Send(cleanupCtx)
	})

	loadTableIntoWorkspace(t, ctx, api, defBranch.ID, createdWorkspace.ID, tableKey)

	runGzipSubtests(t, ctx, api, defBranch, createdWorkspace.ID)
}

func loadTableIntoWorkspace(t *testing.T, ctx context.Context, api *keboola.AuthorizedAPI, branchID keboola.BranchID, workspaceID uint64, tableKey keboola.TableKey) {
	t.Helper()
	job, err := api.StorageWorkspaceLoadDataRequest(branchID, workspaceID, &keboola.WorkspaceLoadDataPayload{
		Input: []keboola.WorkspaceLoadDataInput{
			{Source: tableKey.TableID.String(), Destination: "test_table"},
		},
	}).Send(ctx)
	require.NoError(t, err)

	waitCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	t.Cleanup(cancel)
	require.NoError(t, api.WaitForStorageJob(waitCtx, job))
	assert.Equal(t, "success", job.Status)
}

func runGzipSubtests(t *testing.T, ctx context.Context, api *keboola.AuthorizedAPI, defBranch *keboola.Branch, workspaceID uint64) {
	t.Helper()

	downloadExportedFile := func(t *testing.T, fileID keboola.FileID) []byte {
		t.Helper()
		fileCreds, err := api.GetFileWithCredentialsRequest(keboola.FileKey{BranchID: defBranch.ID, FileID: fileID}).Send(ctx)
		require.NoError(t, err)

		var rawData []byte
		if fileCreds.IsSliced {
			slices, sliceErr := transfer.DownloadManifest(ctx, fileCreds)
			require.NoError(t, sliceErr)
			for _, s := range slices {
				sliceData, dlErr := transfer.DownloadSlice(ctx, fileCreds, s)
				require.NoError(t, dlErr)
				rawData = append(rawData, sliceData...)
			}
		} else {
			rawData, err = transfer.Download(ctx, fileCreds)
			require.NoError(t, err)
		}
		require.NotEmpty(t, rawData)
		return rawData
	}

	t.Run("without_gzip", func(t *testing.T) {
		t.Parallel()
		result, err := api.NewWorkspaceTableExportRequest(defBranch.ID, workspaceID, "test_table").
			WithFileName("test_export_plain").
			WithFileType("csv").
			SendAndWait(ctx, 2*time.Minute)
		require.NoError(t, err)
		require.NotZero(t, result.File.FileID)

		rawData := downloadExportedFile(t, result.File.FileID)
		_, gzErr := gzip.NewReader(bytes.NewReader(rawData))
		require.Error(t, gzErr, "exported file without gzip should not be valid gzip")
	})

	t.Run("with_gzip", func(t *testing.T) {
		t.Parallel()
		result, err := api.NewWorkspaceTableExportRequest(defBranch.ID, workspaceID, "test_table").
			WithFileName("test_export_gzip.csv").
			WithFileType("csv").
			WithGzip(true).
			SendAndWait(ctx, 2*time.Minute)
		require.NoError(t, err)
		require.NotZero(t, result.File.FileID)

		rawData := downloadExportedFile(t, result.File.FileID)
		gzReader, gzErr := gzip.NewReader(bytes.NewReader(rawData))
		require.NoError(t, gzErr, "exported file with gzip should be valid gzip")
		require.NoError(t, gzReader.Close())
	})
}
