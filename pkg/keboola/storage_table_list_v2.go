package keboola

import (
	"context"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/relvacode/iso8601"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

const (
	tableListV2IncludeColumns    = "columns"
	tableListV2IncludeMetadata   = "metadata"
	tableListV2IncludeDefinition = "definition"
	tableListV2QueryAfter        = "page[after]"
	tableListV2QueryLimit        = "page[limit]"
	tableListV2QueryBucketID     = "filter[bucketId]"
)

// TableListV2Item is one table of the table list V2. It is lighter than Table: the optional fields are
// filled only when requested with WithTableListV2Columns, WithTableListV2Metadata or WithTableListV2Definition.
type TableListV2Item struct {
	TableKey
	// TableBranchID is the branch the table belongs to; an inherited table carries the ID of the default branch.
	// It is also copied to TableKey.BranchID, the branch the table detail is read through.
	TableBranchID BranchID `json:"branchId"`
	Name          string   `json:"name"`
	DisplayName   string   `json:"displayName"`
	PrimaryKey    []string `json:"primaryKey"`
	IsAlias       bool     `json:"isAlias"`
	IsTyped       bool     `json:"isTyped"`
	TableType     string   `json:"tableType"`
	// RowsCount and DataSizeBytes are nil for a filtered alias.
	RowsCount      *uint64       `json:"rowsCount"`
	DataSizeBytes  *uint64       `json:"dataSizeBytes"`
	Created        *iso8601.Time `json:"created"`
	LastImportDate *iso8601.Time `json:"lastImportDate"`
	LastChangeDate *iso8601.Time `json:"lastChangeDate"`
	// Description is the native table description (empty when none is set). An alias that syncs its columns
	// falls back to the description of its source chain.
	Description                string `json:"description"`
	IsDescriptionSystemManaged bool   `json:"isDescriptionSystemManaged"`
	// Columns is filled with WithTableListV2Columns.
	Columns []string `json:"columns,omitempty"`
	// Metadata is filled with WithTableListV2Metadata.
	Metadata TableMetadata `json:"metadata,omitempty"`
	// Definition is filled with WithTableListV2Definition. It is nil when it was not requested and for a typed
	// table whose column types are not all native yet (and for an alias of such a table): read the table detail
	// for those.
	Definition *TableDefinition `json:"definition,omitempty"`
}

// TableListV2Page is one page of the cursor-paginated table list V2.
type TableListV2Page struct {
	Data  []*TableListV2Item `json:"data"`
	Meta  TableListV2Meta    `json:"meta"`
	Links TableListV2Links   `json:"links"`
}

// TableListV2Meta describes the listing regardless of the page.
type TableListV2Meta struct {
	// Total is the number of tables matching the filters.
	Total int `json:"total"`
	// Limit is the applied page size.
	Limit int `json:"limit"`
}

// TableListV2Links are absolute URLs of the listing.
type TableListV2Links struct {
	Self  string `json:"self"`
	First string `json:"first"`
	// Next is empty on the last page.
	Next string `json:"next"`
}

// After returns the cursor of the next page, or an empty string on the last page.
func (v TableListV2Links) After() (string, error) {
	if v.Next == "" {
		return "", nil
	}
	u, err := url.Parse(v.Next)
	if err != nil {
		return "", err
	}
	return u.Query().Get(tableListV2QueryAfter), nil
}

type listTablesV2Config struct {
	include  map[string]bool
	bucketID string
	limit    int
	after    string
}

func (v listTablesV2Config) includeString() string {
	include := make([]string, 0, len(v.include))
	for k := range v.include {
		include = append(include, k)
	}
	sort.Strings(include)
	return strings.Join(include, ",")
}

// ListTablesV2Option is an option of ListTablesV2Request and ListTablesV2.
type ListTablesV2Option func(c *listTablesV2Config)

// WithTableListV2Columns adds the column names of each table.
func WithTableListV2Columns() ListTablesV2Option {
	return func(c *listTablesV2Config) {
		c.include[tableListV2IncludeColumns] = true
	}
}

// WithTableListV2Metadata adds the table-level metadata of each table.
func WithTableListV2Metadata() ListTablesV2Option {
	return func(c *listTablesV2Config) {
		c.include[tableListV2IncludeMetadata] = true
	}
}

// WithTableListV2Definition adds the native table definition (types and descriptions) of each table.
func WithTableListV2Definition() ListTablesV2Option {
	return func(c *listTablesV2Config) {
		c.include[tableListV2IncludeDefinition] = true
	}
}

// WithTableListV2Bucket lists only the tables of one bucket.
func WithTableListV2Bucket(bucketID BucketID) ListTablesV2Option {
	return func(c *listTablesV2Config) {
		c.bucketID = bucketID.String()
	}
}

// WithTableListV2Limit sets the page size; the server rejects a value above its maximum.
func WithTableListV2Limit(limit int) ListTablesV2Option {
	return func(c *listTablesV2Config) {
		c.limit = limit
	}
}

// WithTableListV2After starts the page after the given cursor (see TableListV2Links.After).
// ListTablesV2 sets it itself.
func WithTableListV2After(cursor string) ListTablesV2Option {
	return func(c *listTablesV2Config) {
		c.after = cursor
	}
}

func newListTablesV2Config(opts []ListTablesV2Option) listTablesV2Config {
	config := listTablesV2Config{include: make(map[string]bool)}
	for _, opt := range opts {
		opt(&config)
	}
	return config
}

// ListTablesV2Request lists one page of the tables of the branch, `GET /v2/storage/branch/{branchId}/table-list`.
// Use ListTablesV2 to read all pages.
func (a *AuthorizedAPI) ListTablesV2Request(branchID BranchID, opts ...ListTablesV2Option) request.APIRequest[*TableListV2Page] {
	config := newListTablesV2Config(opts)

	result := &TableListV2Page{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithGet("branch/{branchId}/table-list").
		AndPathParam("branchId", branchID.String())
	if include := config.includeString(); include != "" {
		req = req.AndQueryParam("include", include)
	}
	if config.bucketID != "" {
		req = req.AndQueryParam(tableListV2QueryBucketID, config.bucketID)
	}
	if config.limit > 0 {
		req = req.AndQueryParam(tableListV2QueryLimit, strconv.Itoa(config.limit))
	}
	if config.after != "" {
		req = req.AndQueryParam(tableListV2QueryAfter, config.after)
	}
	req = req.WithOnSuccess(func(_ context.Context, _ request.HTTPResponse) error {
		for _, table := range result.Data {
			table.BranchID = table.TableBranchID
		}
		return nil
	})

	return request.NewAPIRequest(result, req)
}

// ListTablesV2 lists all tables of the branch, following the cursor of each page.
func (a *AuthorizedAPI) ListTablesV2(ctx context.Context, branchID BranchID, opts ...ListTablesV2Option) ([]*TableListV2Item, error) {
	var tables []*TableListV2Item
	after := ""
	for {
		pageOpts := append(append([]ListTablesV2Option{}, opts...), WithTableListV2After(after))
		page, err := a.ListTablesV2Request(branchID, pageOpts...).Send(ctx)
		if err != nil {
			return nil, err
		}
		tables = append(tables, page.Data...)

		after, err = page.Links.After()
		if err != nil {
			return nil, err
		}
		if after == "" {
			return tables, nil
		}
	}
}
