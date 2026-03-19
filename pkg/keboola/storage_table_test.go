package keboola_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/keboola/go-utils/pkg/testproject"
	"github.com/relvacode/iso8601"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	DefaultNumber = "38,0"
	DefaultString = "16777216"
)

func TestTableKey_BucketKey(t *testing.T) {
	t.Parallel()

	tableKey := TableKey{
		BranchID: 123,
		TableID: TableID{
			BucketID: BucketID{
				Stage:      BucketStageIn,
				BucketName: "my-bucket",
			},
			TableName: fmt.Sprintf("test_%d", rnd.Int()),
		},
	}

	assert.Equal(t, BucketKey{
		BranchID: 123,
		BucketID: BucketID{
			Stage:      BucketStageIn,
			BucketName: "my-bucket",
		},
	}, tableKey.BucketKey())
}

func TestListTablesRequest(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	_, api := APIClientForAnEmptyProject(t, ctx)

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	tables, err := api.ListTablesRequest(defBranch.ID).Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *tables, 0)
}

// TestTableApiCalls, TestTableCreateLoadDataFromFile, TestTableCreateFromSlicedFile,
// TestTableCreateFromFileOtherOptions, TestTableUnloadRequest are in transfer/storage_table_test.go
// as they require transfer.Upload or transfer.CreateTable.

func TestCreateTableDefinition(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	project, api := APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	bucket, tableKey := createBucketAndTableKey(defBranch)

	// Create bucket
	resBucket, err := api.CreateBucketRequest(bucket).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, bucket, resBucket)

	// min use-case Create Table
	tableDef := TableDefinition{
		PrimaryKeyNames: []string{"name"},
		Columns: Columns{
			{
				Name:       "name",
				BaseType:   ptr(TypeString),
				Definition: &ColumnDefinition{Type: "STRING"},
			},
			{
				Name:       "age",
				BaseType:   ptr(TypeNumeric),
				Definition: &ColumnDefinition{Type: "INT"},
			},
			{
				Name:       "time",
				BaseType:   ptr(TypeDate),
				Definition: &ColumnDefinition{Type: "DATE"},
			},
		},
	}
	assert.Equal(t, []string{"name", "age", "time"}, tableDef.Columns.Names())

	// Create a new table
	newTable, err := api.CreateTableDefinitionRequest(tableKey, tableDef).Send(ctx)
	require.NoError(t, err)
	assert.Equal(t, tableKey.TableID.TableName, newTable.Name)

	for _, column := range newTable.Definition.Columns {
		for _, primaryKey := range newTable.PrimaryKey {
			if column.Name == primaryKey {
				assert.False(t, column.Definition.Nullable)
			}
		}
	}

	// Get a list of the tables
	resTables, err := api.ListTablesRequest(defBranch.ID).Send(ctx)
	require.NoError(t, err)

	tableFound := false
	for _, table := range *resTables {
		if table.TableID == tableKey.TableID {
			tableFound = true
		}
	}
	assert.True(t, tableFound)

	// Get a specific table by tableID
	resTab, err := api.GetTableRequest(newTable.TableKey).Send(ctx)
	removeDynamicValueFromTable(resTab)
	resTab.Metadata = TableMetadata{}
	resTab.ColumnMetadata = ColumnsMetadata{}
	require.NoError(t, err)

	assert.Equal(t, &Table{
		TableKey:    newTable.TableKey,
		URI:         newTable.URI,
		Name:        newTable.Name,
		DisplayName: newTable.DisplayName,
		SourceTable: nil,
		PrimaryKey:  newTable.PrimaryKey,
		Definition: &TableDefinition{
			PrimaryKeyNames: tableDef.PrimaryKeyNames,
			Columns: Columns{
				{
					Name:       "age",
					BaseType:   ptr(TypeInt),
					Definition: &ColumnDefinition{Type: "NUMBER", Length: DefaultNumber, Nullable: false},
				},
				{
					Name:       "name",
					BaseType:   ptr(TypeString),
					Definition: &ColumnDefinition{Type: "VARCHAR", Length: DefaultString, Nullable: false},
				},
				{
					Name:       "time",
					BaseType:   ptr(TypeDate),
					Definition: &ColumnDefinition{Type: "DATE", Nullable: false},
				},
			},
		},
		RowsCount:      0,
		DataSizeBytes:  0,
		Columns:        newTable.Columns,
		Metadata:       TableMetadata{},
		ColumnMetadata: ColumnsMetadata{},
		Bucket: &Bucket{
			BucketKey:   bucket.BucketKey,
			DisplayName: bucket.DisplayName,
			URI:         "https://" + project.StorageAPIHost() + "/v2/storage/buckets/" + tableKey.TableID.BucketID.String(),
		},
	}, resTab)
	assert.Equal(t, tableKey.TableID.TableName, resTab.Name)
	assert.Equal(t, len(newTable.Columns), len(resTab.Columns))

	// Delete the table that was created in the CreateTableDefinitionRequest func
	_, err = api.DeleteTableRequest(tableKey).Send(ctx)
	require.NoError(t, err)

	// Get a list of the tables
	res, err := api.ListTablesRequest(defBranch.ID).Send(ctx)
	require.NoError(t, err)
	assert.Empty(t, res)

	// CreateTable: maximum use-case
	{
		maxUseCaseTableKey := TableKey{BranchID: defBranch.ID, TableID: TableID{BucketID: bucket.BucketID, TableName: "maxUseCase"}}
		maxUseCaseTableDef := TableDefinition{
			PrimaryKeyNames: []string{"email"},
			Columns: Columns{
				{
					Name: "email",
					Definition: &ColumnDefinition{
						Type:     "VARCHAR",
						Length:   DefaultString,
						Nullable: false,
						Default:  "",
					},
					BaseType: ptr(TypeString),
				},
				{
					Name: "comments",
					Definition: &ColumnDefinition{
						Type:     "NUMBER",
						Length:   "37",
						Default:  "100",
						Nullable: true,
					},
					BaseType: ptr(TypeNumeric),
				},
				{
					Name: "favorite_number",
					Definition: &ColumnDefinition{
						Type:     "NUMBER",
						Length:   "37",
						Nullable: true,
						Default:  "100",
					},
					BaseType: ptr(TypeNumeric),
				},
			},
		}

		// Create Table
		_, err = api.CreateTableDefinitionRequest(maxUseCaseTableKey, maxUseCaseTableDef).Send(ctx)
		require.NoError(t, err)

		maxUseCaseTable, err := api.GetTableRequest(maxUseCaseTableKey).Send(ctx)
		require.NoError(t, err)
		assert.Equal(t, Columns{
			{
				Name: "comments",
				Definition: &ColumnDefinition{
					Type:     "NUMBER",
					Length:   "37,0",
					Nullable: true,
					Default:  "100",
				},
				BaseType: ptr(TypeNumeric),
			},
			{
				Name: "email",
				Definition: &ColumnDefinition{
					Type:     "VARCHAR",
					Length:   DefaultString,
					Nullable: false,
				},
				BaseType: ptr(TypeString),
			},
			{
				Name: "favorite_number",
				Definition: &ColumnDefinition{
					Type:     "NUMBER",
					Length:   "37,0",
					Nullable: true,
					Default:  "100",
				},
				BaseType: ptr(TypeNumeric),
			},
		}, maxUseCaseTable.Definition.Columns)

		// Delete the table that was created in the CreateTableDefinitionRequest func
		_, err = api.DeleteTableRequest(maxUseCaseTable.TableKey).Send(ctx)
		require.NoError(t, err)
	}
}

func TestWithoutDefinition(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	project, api := APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	bucket, tableKey := createBucketAndTableKey(defBranch)

	// Create bucket
	resBucket, err := api.CreateBucketRequest(bucket).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, bucket, resBucket)

	// min use-case Create Table
	tableDef := TableDefinition{
		PrimaryKeyNames: []string{"name"},
		Columns: Columns{
			{
				Name: "name",
			},
			{
				Name: "age",
			},
			{
				Name: "time",
			},
		},
	}
	assert.Equal(t, []string{"name", "age", "time"}, tableDef.Columns.Names())

	// Create a new table
	newTable, err := api.CreateTableDefinitionRequest(tableKey, tableDef).Send(ctx)
	require.NoError(t, err)
	assert.Equal(t, tableKey.TableID.TableName, newTable.Name)

	// Get a list of the tables
	resTables, err := api.ListTablesRequest(defBranch.ID).Send(ctx)
	require.NoError(t, err)

	tableFound := false
	for _, table := range *resTables {
		if table.TableID == tableKey.TableID {
			tableFound = true
		}
	}
	assert.True(t, tableFound)

	// Get a specific table by tableID
	resTab, err := api.GetTableRequest(newTable.TableKey).Send(ctx)
	removeDynamicValueFromTable(resTab)
	resTab.Metadata = TableMetadata{}
	resTab.ColumnMetadata = ColumnsMetadata{}
	require.NoError(t, err)

	assert.Equal(t, &Table{
		TableKey:       newTable.TableKey,
		URI:            newTable.URI,
		Name:           newTable.Name,
		DisplayName:    newTable.DisplayName,
		SourceTable:    nil,
		PrimaryKey:     newTable.PrimaryKey,
		RowsCount:      0,
		DataSizeBytes:  0,
		Columns:        newTable.Columns,
		Metadata:       TableMetadata{},
		ColumnMetadata: ColumnsMetadata{},
		Bucket: &Bucket{
			BucketKey:   bucket.BucketKey,
			DisplayName: bucket.DisplayName,
			URI:         "https://" + project.StorageAPIHost() + "/v2/storage/buckets/" + tableKey.TableID.BucketID.String(),
		},
	}, resTab)
	assert.Equal(t, tableKey.TableID.TableName, resTab.Name)
	assert.Equal(t, len(newTable.Columns), len(resTab.Columns))

	// Delete the table that was created in the CreateTableDefinitionRequest func
	_, err = api.DeleteTableRequest(tableKey).Send(ctx)
	require.NoError(t, err)

	// Get a list of the tables
	res, err := api.ListTablesRequest(defBranch.ID).Send(ctx)
	require.NoError(t, err)
	assert.Empty(t, res)

	found := false
	for _, table := range *res {
		if table.TableID == tableKey.TableID {
			found = true
		}
	}
	assert.False(t, found)
}

func TestCreateTableDefinitionWithBigQuery(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	project, api := APIClientForAnEmptyProject(t, ctx, testproject.WithBigQueryBackend())

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	bucket, tableKey := createBucketAndTableKey(defBranch)

	// Create bucket
	resBucket, err := api.CreateBucketRequest(bucket).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, bucket, resBucket)

	// min use-case Create Table
	tableDef := TableDefinition{
		PrimaryKeyNames: []string{"id"},
		Columns: Columns{
			{
				Name:       "id",
				BaseType:   ptr(TypeInt),
				Definition: &ColumnDefinition{Type: TypeInt.String(), Nullable: false},
			},
			{
				Name:       "age",
				BaseType:   ptr(TypeInt),
				Definition: &ColumnDefinition{Type: TypeInt.String(), Nullable: true},
			},
			{
				Name:       "time",
				BaseType:   ptr(TypeTimestamp),
				Definition: &ColumnDefinition{Type: TypeTimestamp.String(), Nullable: false},
			},
		},
	}

	// Create a new table
	newTable, err := api.CreateTableDefinitionRequest(tableKey, tableDef).Send(ctx)
	require.NoError(t, err)
	assert.Equal(t, tableKey.TableID.TableName, newTable.Name)

	// Get table
	res, err := api.GetTableRequest(newTable.TableKey).Send(ctx)
	require.NoError(t, err)
	removeDynamicValueFromTable(res)
	res.Metadata = TableMetadata{}
	res.ColumnMetadata = ColumnsMetadata{}
	assert.Equal(t, &Table{
		TableKey:    newTable.TableKey,
		URI:         newTable.URI,
		Name:        newTable.Name,
		DisplayName: newTable.DisplayName,
		SourceTable: nil,
		PrimaryKey:  newTable.PrimaryKey,
		Definition: &TableDefinition{
			PrimaryKeyNames: tableDef.PrimaryKeyNames,
			Columns: Columns{
				{
					Name:       "age",
					BaseType:   ptr(TypeInt),
					Definition: &ColumnDefinition{Type: TypeInt.String(), Nullable: true},
				},
				{
					Name:       "id",
					BaseType:   ptr(TypeInt),
					Definition: &ColumnDefinition{Type: TypeInt.String(), Nullable: false},
				},
				{
					Name:       "time",
					BaseType:   ptr(TypeTimestamp),
					Definition: &ColumnDefinition{Type: TypeTimestamp.String(), Nullable: false},
				},
			},
		},

		RowsCount:      0,
		DataSizeBytes:  0,
		Columns:        newTable.Columns,
		Metadata:       TableMetadata{},
		ColumnMetadata: ColumnsMetadata{},
		Bucket: &Bucket{
			BucketKey:   bucket.BucketKey,
			DisplayName: bucket.DisplayName,
			URI:         "https://" + project.StorageAPIHost() + "/v2/storage/buckets/" + tableKey.TableID.BucketID.String(),
		},
	}, res)
}

// TestCreateTableDefinition_TimePartitioning tests special settings 'timePartitioning' to create a table for a BigQuery project.
func TestCreateTableDefinition_TimePartitioning(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	project, api := APIClientForAnEmptyProject(t, ctx, testproject.WithBigQueryBackend())

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	bucket, tableKey := createBucketAndTableKey(defBranch)

	// Create bucket
	resBucket, err := api.CreateBucketRequest(bucket).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, bucket, resBucket)

	// min use-case Create Table
	tableDef := TableDefinition{
		PrimaryKeyNames: []string{"id"},
		Columns: Columns{
			{
				Name:       "id",
				BaseType:   ptr(TypeInt),
				Definition: &ColumnDefinition{Type: TypeInt.String(), Nullable: false},
			},
			{
				Name:       "time",
				BaseType:   ptr(TypeTimestamp),
				Definition: &ColumnDefinition{Type: TypeTimestamp.String(), Nullable: false},
			},
		},
		TimePartitioning: &TimePartitioning{
			Type:         Day,
			ExpirationMs: "864000000",
			Field:        "time",
		},
	}

	// Create a new table
	newTable, err := api.CreateTableDefinitionRequest(tableKey, tableDef).Send(ctx)
	require.NoError(t, err)
	assert.Equal(t, tableKey.TableID.TableName, newTable.Name)

	// Get table
	res, err := api.GetTableRequest(newTable.TableKey).Send(ctx)
	require.NoError(t, err)
	removeDynamicValueFromTable(res)
	res.Metadata = TableMetadata{}
	res.ColumnMetadata = ColumnsMetadata{}
	assert.Equal(t, &Table{
		TableKey:    newTable.TableKey,
		URI:         newTable.URI,
		Name:        newTable.Name,
		DisplayName: newTable.DisplayName,
		SourceTable: nil,
		PrimaryKey:  newTable.PrimaryKey,
		Definition: &TableDefinition{
			PrimaryKeyNames: tableDef.PrimaryKeyNames,
			Columns: Columns{
				{
					Name:       "id",
					BaseType:   ptr(TypeInt),
					Definition: &ColumnDefinition{Type: TypeInt.String(), Nullable: false},
				},
				{
					Name:       "time",
					BaseType:   ptr(TypeTimestamp),
					Definition: &ColumnDefinition{Type: TypeTimestamp.String(), Nullable: false},
				},
			},
			TimePartitioning: tableDef.TimePartitioning,
		},

		RowsCount:      0,
		DataSizeBytes:  0,
		Columns:        newTable.Columns,
		Metadata:       TableMetadata{},
		ColumnMetadata: ColumnsMetadata{},
		Bucket: &Bucket{
			BucketKey:   bucket.BucketKey,
			DisplayName: bucket.DisplayName,
			URI:         "https://" + project.StorageAPIHost() + "/v2/storage/buckets/" + tableKey.TableID.BucketID.String(),
		},
	}, res)
}

// TestCreateTableDefinition_Clustering tests special settings 'clustering' to create a table for a BigQuery project.
func TestCreateTableDefinition_Clustering(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	project, api := APIClientForAnEmptyProject(t, ctx, testproject.WithBigQueryBackend())

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	bucket, tableKey := createBucketAndTableKey(defBranch)

	// Create bucket
	resBucket, err := api.CreateBucketRequest(bucket).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, bucket, resBucket)

	// min use-case Create Table
	tableDef := TableDefinition{
		PrimaryKeyNames: []string{"id"},
		Columns: Columns{
			{
				Name:       "id",
				BaseType:   ptr(TypeInt),
				Definition: &ColumnDefinition{Type: TypeInt.String(), Nullable: false},
			},
			{
				Name:       "time",
				BaseType:   ptr(TypeTimestamp),
				Definition: &ColumnDefinition{Type: TypeTimestamp.String(), Nullable: false},
			},
		},
		Clustering: &Clustering{
			Fields: []string{
				"id",
			},
		},
	}

	// Create a new table
	newTable, err := api.CreateTableDefinitionRequest(tableKey, tableDef).Send(ctx)
	require.NoError(t, err)
	assert.Equal(t, tableKey.TableID.TableName, newTable.Name)

	// Get table
	res, err := api.GetTableRequest(newTable.TableKey).Send(ctx)
	require.NoError(t, err)
	removeDynamicValueFromTable(res)
	res.Metadata = TableMetadata{}
	res.ColumnMetadata = ColumnsMetadata{}
	assert.Equal(t, &Table{
		TableKey:    newTable.TableKey,
		URI:         newTable.URI,
		Name:        newTable.Name,
		DisplayName: newTable.DisplayName,
		SourceTable: nil,
		PrimaryKey:  newTable.PrimaryKey,
		Definition: &TableDefinition{
			PrimaryKeyNames: tableDef.PrimaryKeyNames,
			Columns: Columns{
				{
					Name:       "id",
					BaseType:   ptr(TypeInt),
					Definition: &ColumnDefinition{Type: TypeInt.String(), Nullable: false},
				},
				{
					Name:       "time",
					BaseType:   ptr(TypeTimestamp),
					Definition: &ColumnDefinition{Type: TypeTimestamp.String(), Nullable: false},
				},
			},
			Clustering: tableDef.Clustering,
		},

		RowsCount:      0,
		DataSizeBytes:  0,
		Columns:        newTable.Columns,
		Metadata:       TableMetadata{},
		ColumnMetadata: ColumnsMetadata{},
		Bucket: &Bucket{
			BucketKey:   bucket.BucketKey,
			DisplayName: bucket.DisplayName,
			URI:         "https://" + project.StorageAPIHost() + "/v2/storage/buckets/" + tableKey.TableID.BucketID.String(),
		},
	}, res)
}

// TestCreateTableDefinition_RangePartitioning tests special settings 'rangePartitioning' to create a table for a BigQuery project.
func TestCreateTableDefinition_RangePartitioning(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	project, api := APIClientForAnEmptyProject(t, ctx, testproject.WithBigQueryBackend())

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	bucket, tableKey := createBucketAndTableKey(defBranch)

	// Create bucket
	resBucket, err := api.CreateBucketRequest(bucket).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, bucket, resBucket)

	// min use-case Create Table
	tableDef := TableDefinition{
		PrimaryKeyNames: []string{"id"},
		Columns: Columns{
			{
				Name:       "id",
				BaseType:   ptr(TypeInt),
				Definition: &ColumnDefinition{Type: TypeInt.String(), Nullable: false},
			},
			{
				Name:       "time",
				BaseType:   ptr(TypeTimestamp),
				Definition: &ColumnDefinition{Type: TypeTimestamp.String(), Nullable: false},
			},
		},
		RangePartitioning: &RangePartitioning{
			Field: "id",
			Range: Range{
				Start:    "0",
				End:      "10",
				Interval: "1",
			},
		},
	}

	// Create a new table
	newTable, err := api.CreateTableDefinitionRequest(tableKey, tableDef).Send(ctx)
	require.NoError(t, err)
	assert.Equal(t, tableKey.TableID.TableName, newTable.Name)

	// Get table
	res, err := api.GetTableRequest(newTable.TableKey).Send(ctx)
	require.NoError(t, err)
	removeDynamicValueFromTable(res)
	res.Metadata = TableMetadata{}
	res.ColumnMetadata = ColumnsMetadata{}
	assert.Equal(t, &Table{
		TableKey:    newTable.TableKey,
		URI:         newTable.URI,
		Name:        newTable.Name,
		DisplayName: newTable.DisplayName,
		SourceTable: nil,
		PrimaryKey:  newTable.PrimaryKey,
		Definition: &TableDefinition{
			PrimaryKeyNames: tableDef.PrimaryKeyNames,
			Columns: Columns{
				{
					Name:       "id",
					BaseType:   ptr(TypeInt),
					Definition: &ColumnDefinition{Type: TypeInt.String(), Nullable: false},
				},
				{
					Name:       "time",
					BaseType:   ptr(TypeTimestamp),
					Definition: &ColumnDefinition{Type: TypeTimestamp.String(), Nullable: false},
				},
			},
			RangePartitioning: tableDef.RangePartitioning,
		},

		RowsCount:      0,
		DataSizeBytes:  0,
		Columns:        newTable.Columns,
		Metadata:       TableMetadata{},
		ColumnMetadata: ColumnsMetadata{},
		Bucket: &Bucket{
			BucketKey:   bucket.BucketKey,
			DisplayName: bucket.DisplayName,
			URI:         "https://" + project.StorageAPIHost() + "/v2/storage/buckets/" + tableKey.TableID.BucketID.String(),
		},
	}, res)
}

func createBucketAndTableKey(branch *Branch) (*Bucket, TableKey) {
	bucket := &Bucket{
		BucketKey: BucketKey{
			BranchID: branch.ID,
			BucketID: BucketID{
				Stage:      BucketStageIn,
				BucketName: fmt.Sprintf("c-test_%d", rnd.Int()),
			},
		},
	}

	tableKey := TableKey{
		BranchID: branch.ID,
		TableID: TableID{
			BucketID:  bucket.BucketID,
			TableName: fmt.Sprintf("test_%d", rnd.Int()),
		},
	}
	return bucket, tableKey
}

func removeDynamicValueFromTable(table *Table) {
	table.Created = iso8601.Time{}
	table.LastImportDate = iso8601.Time{}
	table.LastChangeDate = nil
	table.Bucket.Created = iso8601.Time{}
	table.Bucket.LastChangeDate = nil
}

func ptr[T any](v T) *T {
	return &v
}
