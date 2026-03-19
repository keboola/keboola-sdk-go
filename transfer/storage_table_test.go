package transfer_test

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/keboola/go-utils/pkg/testproject"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
	"github.com/keboola/keboola-sdk-go/v2/transfer"
)

func TestTableApiCalls(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	project, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	bucket, tableKey := createBucketAndTableKey(defBranch)

	// Create bucket
	resBucket, err := api.CreateBucketRequest(bucket).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, bucket, resBucket)

	table := &keboola.Table{
		TableKey:   tableKey,
		Bucket:     bucket,
		Name:       tableKey.TableID.TableName,
		Columns:    []string{"first", "second", "third", "fourth"},
		PrimaryKey: []string{"first", "fourth"},
	}

	// Create table
	res, err := transfer.CreateTable(ctx, api, tableKey, table.Columns)
	assert.NoError(t, err)
	assert.Equal(t, table.Name, res.Name)

	// List tables
	resList, err := api.ListTablesRequest(defBranch.ID).Send(ctx)
	assert.NoError(t, err)
	tableFound := false
	for _, resTable := range *resList {
		assert.Equal(t, resTable.BranchID, defBranch.ID)
		if resTable.TableID == table.TableID {
			tableFound = true
		}
	}
	assert.True(t, tableFound)

	// Get table (without table and columns metadata)
	respGet1, err := api.GetTableRequest(tableKey).Send(ctx)
	assert.NoError(t, err)
	removeDynamicValueFromTable(respGet1)
	assert.Equal(t, &keboola.Table{
		TableKey:      tableKey,
		URI:           "https://" + project.StorageAPIHost() + "/v2/storage/tables/" + tableKey.TableID.String(),
		Name:          tableKey.TableID.TableName,
		DisplayName:   tableKey.TableID.TableName,
		PrimaryKey:    []string{},
		RowsCount:     0,
		DataSizeBytes: 0,
		Columns:       []string{"first", "second", "third", "fourth"},
		Bucket: &keboola.Bucket{
			BucketKey:   bucket.BucketKey,
			DisplayName: bucket.DisplayName,
			URI:         "https://" + project.StorageAPIHost() + "/v2/storage/buckets/" + tableKey.BucketKey().BucketID.String(),
		},
		Metadata:       keboola.TableMetadata{},
		ColumnMetadata: keboola.ColumnsMetadata{},
	}, respGet1)

	// Set metadata
	resMetadata, err := api.
		CreateOrUpdateTableMetadata(
			tableKey,
			"keboola-sdk-go-test",
			[]keboola.TableMetadataRequest{
				{Key: "tableMetadata1", Value: "value1"},
				{Key: "tableMetadata2", Value: "value2"},
			},
			[]keboola.ColumnMetadataRequest{
				{ColumnName: "first", Key: "columnMetadata1", Value: "value3"},
				{ColumnName: "first", Key: "columnMetadata2", Value: "value4"},
				{ColumnName: "second", Key: "columnMetadata3", Value: "value5"},
				{ColumnName: "second", Key: "columnMetadata4", Value: "value6"},
			},
		).
		Send(ctx)
	assert.NoError(t, err)

	// Check metadata response
	removeDynamicValuesFromTableMetadata(resMetadata.Metadata)
	removeDynamicValuesFromColumnsMetadata(resMetadata.ColumnMetadata)
	assert.Equal(t, &keboola.TableMetadataResponse{
		Metadata: keboola.TableMetadata{
			{Key: "tableMetadata1", Value: "value1", Provider: "keboola-sdk-go-test"},
			{Key: "tableMetadata2", Value: "value2", Provider: "keboola-sdk-go-test"},
		},
		ColumnMetadata: keboola.ColumnsMetadata{
			"first": {
				{Key: "columnMetadata1", Value: "value3", Provider: "keboola-sdk-go-test"},
				{Key: "columnMetadata2", Value: "value4", Provider: "keboola-sdk-go-test"},
			},
			"second": {
				{Key: "columnMetadata3", Value: "value5", Provider: "keboola-sdk-go-test"},
				{Key: "columnMetadata4", Value: "value6", Provider: "keboola-sdk-go-test"},
			},
		},
	}, resMetadata)

	// Get table (with table and columns metadata)
	respGet2, err := api.GetTableRequest(tableKey).Send(ctx)
	assert.NoError(t, err)
	removeDynamicValuesFromTableMetadata(respGet2.Metadata)
	removeDynamicValuesFromColumnsMetadata(respGet2.ColumnMetadata)
	removeDynamicValueFromTable(respGet2)
	assert.Equal(t, &keboola.Table{
		TableKey:      tableKey,
		URI:           "https://" + project.StorageAPIHost() + "/v2/storage/tables/" + tableKey.TableID.String(),
		Name:          tableKey.TableID.TableName,
		DisplayName:   tableKey.TableID.TableName,
		PrimaryKey:    []string{},
		RowsCount:     0,
		DataSizeBytes: 0,
		Columns:       []string{"first", "second", "third", "fourth"},
		Bucket: &keboola.Bucket{
			BucketKey:   bucket.BucketKey,
			DisplayName: bucket.DisplayName,
			URI:         "https://" + project.StorageAPIHost() + "/v2/storage/buckets/" + tableKey.BucketKey().BucketID.String(),
		},
		Metadata: keboola.TableMetadata{
			{Key: "tableMetadata1", Value: "value1", Provider: "keboola-sdk-go-test"},
			{Key: "tableMetadata2", Value: "value2", Provider: "keboola-sdk-go-test"},
		},
		ColumnMetadata: keboola.ColumnsMetadata{
			"first": {
				{Key: "columnMetadata1", Value: "value3", Provider: "keboola-sdk-go-test"},
				{Key: "columnMetadata2", Value: "value4", Provider: "keboola-sdk-go-test"},
			},
			"second": {
				{Key: "columnMetadata3", Value: "value5", Provider: "keboola-sdk-go-test"},
				{Key: "columnMetadata4", Value: "value6", Provider: "keboola-sdk-go-test"},
			},
		},
	}, respGet2)

	// add new keys and update existing key for table metadata
	err = api.
		CreateOrUpdateTableMetadata(
			respGet2.TableKey,
			"keboola-sdk-go-test",
			[]keboola.TableMetadataRequest{
				{Key: "tableMetadata1", Value: "value1-updated"},
				{Key: "tableMetadata3", Value: "value3"},
				{Key: "tableMetadata4", Value: "value4"},
			},
			[]keboola.ColumnMetadataRequest{},
		).
		SendOrErr(ctx)
	assert.NoError(t, err)

	// Get table (with table and columns metadata)
	respGet3, err := api.GetTableRequest(tableKey).Send(ctx)
	assert.NoError(t, err)
	removeDynamicValuesFromTableMetadata(respGet3.Metadata)
	removeDynamicValuesFromColumnsMetadata(respGet3.ColumnMetadata)
	removeDynamicValueFromTable(respGet3)
	// check table metadata
	assert.Equal(t, keboola.TableMetadata{
		{Key: "tableMetadata1", Value: "value1-updated", Provider: "keboola-sdk-go-test"},
		{Key: "tableMetadata2", Value: "value2", Provider: "keboola-sdk-go-test"},
		{Key: "tableMetadata3", Value: "value3", Provider: "keboola-sdk-go-test"},
		{Key: "tableMetadata4", Value: "value4", Provider: "keboola-sdk-go-test"},
	}, respGet3.Metadata)

	// Delete table
	_, err = api.DeleteTableRequest(tableKey, keboola.WithForce()).Send(ctx)
	assert.NoError(t, err)

	// List tables again - without the deleted table
	resList, err = api.ListTablesRequest(defBranch.ID).Send(ctx)
	assert.NoError(t, err)
	tableFound = false
	for _, tbl := range *resList {
		if tbl.TableID == table.TableID {
			tableFound = true
		}
	}
	assert.False(t, tableFound)
}

func TestTableCreateLoadDataFromFile(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	bucket, tableKey := createBucketAndTableKey(defBranch)

	// Create bucket
	resBucket, err := api.CreateBucketRequest(bucket).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, bucket, resBucket)

	// Create file
	fileName1 := fmt.Sprintf("file_%d", rand.Int()) //nolint:gosec
	file1, err := api.CreateFileResourceRequest(defBranch.ID, fileName1).Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, file1.FileID)

	// Upload file
	content := []byte("col1,col2\nval1,val2\n")
	written, err := transfer.Upload(ctx, file1, bytes.NewReader(content))
	assert.NoError(t, err)
	assert.Equal(t, int64(len(content)), written)

	// Create table
	_, err = api.CreateTableFromFileRequest(tableKey, file1.FileKey, keboola.WithPrimaryKey([]string{"col1", "col2"})).Send(ctx)
	assert.NoError(t, err)

	// Create file
	fileName2 := fmt.Sprintf("file_%d", rand.Int()) //nolint:gosec
	file2, err := api.CreateFileResourceRequest(defBranch.ID, fileName2).Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, file2.FileID)

	// Check rows count
	table, err := api.GetTableRequest(tableKey).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), table.RowsCount)

	// Upload file
	content = []byte("val2,val3\nval3,val4\nval4,val5\n")
	written, err = transfer.Upload(ctx, file2, bytes.NewReader(content))
	assert.NoError(t, err)
	assert.Equal(t, int64(len(content)), written)

	// Load data to table - added three rows
	waitCtx2, waitCancelFn2 := context.WithTimeout(ctx, time.Minute*5)
	defer waitCancelFn2()
	job, err := api.LoadDataFromFileRequest(tableKey, file2.FileKey, keboola.WithColumnsHeaders([]string{"col2", "col1"}), keboola.WithIncrementalLoad(true)).Send(ctx)
	assert.NoError(t, err)
	assert.NoError(t, api.WaitForStorageJob(waitCtx2, job))

	// Check rows count
	table, err = api.GetTableRequest(tableKey).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, uint64(4), table.RowsCount)
}

func TestTableCreateFromSlicedFile(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithStagingStorageS3())

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	bucket, tableKey := createBucketAndTableKey(defBranch)

	// Create bucket
	_, err = api.CreateBucketRequest(bucket).Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, bucket.BucketID)

	// Create whole file
	wholeFile, err := api.CreateFileResourceRequest(defBranch.ID, "file name").Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, wholeFile.FileID)

	// Upload file
	content := []byte("col1,col2\nval1,val2\n")
	written, err := transfer.Upload(ctx, wholeFile, bytes.NewReader(content))
	assert.NoError(t, err)
	assert.Equal(t, int64(len(content)), written)

	// Create non-sliced table.
	// Table cannot be created from a sliced file (https://keboola.atlassian.net/browse/KBC-1861).
	_, err = api.CreateTableFromFileRequest(tableKey, wholeFile.FileKey, keboola.WithPrimaryKey([]string{"col1", "col2"})).Send(ctx)
	assert.NoError(t, err)

	// Check rows count
	table, err := api.GetTableRequest(tableKey).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), table.RowsCount)

	// Create sliced file
	slicedFile, err := api.CreateFileResourceRequest(defBranch.ID, "file name", keboola.WithIsSliced(true)).Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, slicedFile.FileID)

	// Upload slice 1 file
	content = []byte("val3,val4\nval5,val6\n")
	_, err = transfer.UploadSlice(ctx, slicedFile, "slice1", bytes.NewReader(content))
	assert.NoError(t, err)

	// Upload slice 2 file
	content = []byte("val7,val8\nval9,val10\n")
	_, err = transfer.UploadSlice(ctx, slicedFile, "slice2", bytes.NewReader(content))
	assert.NoError(t, err)

	// Upload manifest
	_, err = transfer.UploadSlicedFileManifest(ctx, slicedFile, []string{"slice1", "slice2"})
	assert.NoError(t, err)

	// Load data to table
	waitCtx, waitCancelFn := context.WithTimeout(ctx, time.Minute*5)
	defer waitCancelFn()
	job, err := api.LoadDataFromFileRequest(tableKey, slicedFile.FileKey, keboola.WithIncrementalLoad(true), keboola.WithColumnsHeaders([]string{"col1", "col2"})).Send(ctx)
	assert.NoError(t, err)
	assert.NoError(t, api.WaitForStorageJob(waitCtx, job))

	// Check rows count
	table, err = api.GetTableRequest(tableKey).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, uint64(5), table.RowsCount)
}

func TestTableCreateFromFileOtherOptions(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	bucket, tableKey := createBucketAndTableKey(defBranch)

	// Create bucket
	resBucket, err := api.CreateBucketRequest(bucket).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, bucket, resBucket)

	// Create file
	fileName1 := fmt.Sprintf("file_%d", rand.Int()) //nolint:gosec
	file1, err := api.CreateFileResourceRequest(defBranch.ID, fileName1).Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, file1.FileID)

	// Upload file
	content := []byte("'col1'&'col2'\n'val1'&'val2'\n")
	written, err := transfer.Upload(ctx, file1, bytes.NewReader(content))
	assert.NoError(t, err)
	assert.Equal(t, int64(len(content)), written)

	// Create table
	_, err = api.CreateTableFromFileRequest(tableKey, file1.FileKey, keboola.WithDelimiter("&"), keboola.WithEnclosure("'")).Send(ctx)
	assert.NoError(t, err)

	// Check rows count
	table, err := api.GetTableRequest(tableKey).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), table.RowsCount)
}

func TestTableUnloadRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	bucket, tableKey := createBucketAndTableKey(defBranch)

	// Create bucket
	resBucket, err := api.CreateBucketRequest(bucket).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, bucket, resBucket)

	// Create file
	fileName1 := fmt.Sprintf("file_%d", rand.Int()) //nolint:gosec
	inputFile, err := api.CreateFileResourceRequest(defBranch.ID, fileName1).Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, inputFile.FileID)

	// Upload file
	content := []byte("col1,col2\nval1,val2\n")
	written, err := transfer.Upload(ctx, inputFile, bytes.NewReader(content))
	assert.NoError(t, err)
	assert.Equal(t, int64(len(content)), written)

	// Create table
	_, err = api.CreateTableFromFileRequest(tableKey, inputFile.FileKey, keboola.WithPrimaryKey([]string{"col1", "col2"})).Send(ctx)
	assert.NoError(t, err)

	// Unload table as CSV
	outputFileInfo, err := api.NewTableUnloadRequest(tableKey).
		WithColumns("col1").
		WithChangedSince("-2 days").
		WithFormat(keboola.UnloadFormatCSV).
		WithLimitRows(100).
		WithOrderBy("col1", keboola.OrderAsc).
		WithWhere("col1", keboola.CompareEq, []string{"val1"}).
		SendAndWait(ctx, time.Minute*5)
	assert.NoError(t, err)

	// Download file
	credentials, err := api.GetFileWithCredentialsRequest(outputFileInfo.File.FileKey).Send(ctx)
	assert.NoError(t, err)

	data, err := downloadAllSlices(ctx, credentials)
	assert.NoError(t, err)

	row, err := csv.NewReader(bytes.NewReader(data)).ReadAll()
	assert.NoError(t, err)
	assert.Equal(t, [][]string{{"val1"}}, row)
}

func downloadAllSlices(ctx context.Context, file *keboola.FileDownloadCredentials) ([]byte, error) {
	if !file.IsSliced {
		return nil, fmt.Errorf("cannot download a whole file as a sliced file")
	}

	out := []byte{}

	slices, err := transfer.DownloadManifest(ctx, file)
	if err != nil {
		return nil, err
	}

	for _, slice := range slices {
		data, err := transfer.DownloadSlice(ctx, file, slice)
		if err != nil {
			return nil, err
		}
		out = append(out, data...)
	}

	return out, nil
}

func removeDynamicValuesFromTableMetadata(in keboola.TableMetadata) {
	for i := range in {
		meta := &in[i]
		meta.ID = ""
		meta.Timestamp = ""
	}
}

func removeDynamicValuesFromColumnsMetadata(in keboola.ColumnsMetadata) {
	for _, colMetadata := range in {
		for i := range colMetadata {
			item := &colMetadata[i]
			item.ID = ""
			item.Timestamp = ""
		}
	}
}
