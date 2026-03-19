package transfer

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

// CreateTable creates a new table with the given columns by uploading a CSV header file.
// It creates a temporary file resource, uploads a CSV with just the header row,
// then creates the table from that file.
func CreateTable(ctx context.Context, api *keboola.AuthorizedAPI, k keboola.TableKey, columns []string, opts ...keboola.CreateTableOption) (*keboola.Table, error) {
	file, err := api.CreateFileResourceRequest(k.BranchID, fmt.Sprintf("table_header_%s", k.TableID.String())).Send(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot create file resource: %w", err)
	}

	var buf bytes.Buffer
	csvW := csv.NewWriter(&buf)
	if err := csvW.Write(columns); err != nil {
		return nil, fmt.Errorf("cannot write CSV header: %w", err)
	}
	csvW.Flush()
	if err := csvW.Error(); err != nil {
		return nil, fmt.Errorf("cannot flush CSV header: %w", err)
	}

	if _, err := Upload(ctx, file, &buf); err != nil {
		return nil, fmt.Errorf("cannot upload table header: %w", err)
	}

	table, err := api.CreateTableFromFileRequest(k, file.FileKey, opts...).Send(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot create table from file: %w", err)
	}
	return table, nil
}
