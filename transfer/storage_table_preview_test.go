package transfer_test

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
	"github.com/keboola/keboola-sdk-go/v2/transfer"
)

func TestPreviewTableRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	bucketID := keboola.BucketID{
		Stage:      keboola.BucketStageIn,
		BucketName: fmt.Sprintf("c-bucket_%d", rand.Int()),
	}
	bucket := &keboola.Bucket{
		BucketKey: keboola.BucketKey{
			BranchID: defBranch.ID,
			BucketID: bucketID,
		},
	}

	tableKey := keboola.TableKey{
		BranchID: defBranch.ID,
		TableID: keboola.TableID{
			BucketID:  bucketID,
			TableName: fmt.Sprintf("table_%d", rand.Int()),
		},
	}

	// Create bucket
	resBucket, err := api.CreateBucketRequest(bucket).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, bucket, resBucket)

	// Create file
	fileName1 := fmt.Sprintf("file_%d", rand.Int())
	file1, err := api.CreateFileResourceRequest(defBranch.ID, fileName1).Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, file1.FileID)

	// Upload file
	content := []byte("id,value\n")
	for i := range 100 {
		content = append(content, fmt.Sprintf("%d,%d\n", i, i)...)
	}
	written, err := transfer.Upload(ctx, file1, bytes.NewReader(content))
	assert.NoError(t, err)
	assert.Equal(t, int64(len(content)), written)

	// Create table
	_, err = api.CreateTableFromFileRequest(tableKey, file1.FileKey, keboola.WithPrimaryKey([]string{"id"})).Send(ctx)
	assert.NoError(t, err)

	// Preview table
	preview, err := api.PreviewTableRequest(tableKey,
		keboola.WithWhere("value", "ge", []int{10}, keboola.TypeInteger),
		keboola.WithWhere("value", keboola.CompareLe, []int{15}, keboola.TypeInteger),
		keboola.WithOrderBy("value", keboola.OrderDesc, keboola.TypeInteger),
	).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t,
		&keboola.TablePreview{
			Columns: []string{"id", "value"},
			Rows: [][]string{
				{"15", "15"},
				{"14", "14"},
				{"13", "13"},
				{"12", "12"},
				{"11", "11"},
				{"10", "10"},
			},
		},
		preview,
	)
}
