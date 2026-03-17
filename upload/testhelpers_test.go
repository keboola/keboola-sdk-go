package upload_test

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/relvacode/iso8601"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func createBucketAndTableKey(branch *keboola.Branch) (*keboola.Bucket, keboola.TableKey) {
	bucket := &keboola.Bucket{
		BucketKey: keboola.BucketKey{
			BranchID: branch.ID,
			BucketID: keboola.BucketID{
				Stage:      keboola.BucketStageIn,
				BucketName: fmt.Sprintf("c-test_%d", rnd.Int()),
			},
		},
	}
	tableKey := keboola.TableKey{
		BranchID: branch.ID,
		TableID: keboola.TableID{
			BucketID:  bucket.BucketID,
			TableName: fmt.Sprintf("test_%d", rnd.Int()),
		},
	}
	return bucket, tableKey
}

func removeDynamicValueFromTable(table *keboola.Table) {
	table.Created = iso8601.Time{}
	table.LastImportDate = iso8601.Time{}
	table.LastChangeDate = nil
	table.Bucket.Created = iso8601.Time{}
	table.Bucket.LastChangeDate = nil
}

func ptr[T any](v T) *T {
	return &v
}
