package keboola

import (
	"context"
	"sort"
	"strings"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

type listTablesConfig struct {
	include map[string]bool
}

func (v listTablesConfig) includeString() string {
	include := make([]string, 0, len(v.include))
	for k := range v.include {
		include = append(include, k)
	}
	sort.Strings(include)
	return strings.Join(include, ",")
}

type ListTableOption func(c *listTablesConfig)

func WithBuckets() ListTableOption {
	return func(c *listTablesConfig) {
		c.include["buckets"] = true
	}
}

func WithColumns() ListTableOption {
	return func(c *listTablesConfig) {
		c.include["columns"] = true
	}
}

func WithMetadata() ListTableOption {
	return func(c *listTablesConfig) {
		c.include["metadata"] = true
	}
}

func WithColumnMetadata() ListTableOption {
	return func(c *listTablesConfig) {
		c.include["columnMetadata"] = true
	}
}

// ListTablesRequest https://keboola.docs.apiary.io/#reference/tables/list-tables/list-all-tables
func (a *AuthorizedAPI) ListTablesRequest(branchID BranchID, opts ...ListTableOption) request.APIRequest[*[]*Table] {
	config := listTablesConfig{include: make(map[string]bool)}
	for _, opt := range opts {
		opt(&config)
	}

	result := make([]*Table, 0)
	req := a.
		newRequest(StorageAPI).
		WithResult(&result).
		WithGet("branch/{branchId}/tables").
		AndPathParam("branchId", branchID.String()).
		AndQueryParam("include", config.includeString()).
		WithOnSuccess(func(_ context.Context, _ request.HTTPResponse) error {
			for _, table := range result {
				table.BranchID = branchID
			}
			return nil
		})

	return request.NewAPIRequest(&result, req)
}
