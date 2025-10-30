package keboola

import (
	"context"
	jsonLib "encoding/json"
	"strconv"
	"time"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

// WorkspaceTableExportRequestBuilder builds requests for exporting tables from workspaces.
// This operation exports existing tables from a workspace to a file in the storage backend.
// Currently, only Snowflake backend is supported.
type WorkspaceTableExportRequestBuilder struct {
	branchID    BranchID
	workspaceID uint64
	config      workspaceTableExportConfig
	api         *AuthorizedAPI
}

type workspaceTableExportConfig struct {
	TableName string `json:"tableName"`
	FileName  string `json:"fileName"`
	Format    string `json:"format"`
}

// NewWorkspaceTableExportRequest creates a new workspace table export request builder.
// The table will be exported from the workspace to a file in the storage backend.
func (a *AuthorizedAPI) NewWorkspaceTableExportRequest(branchID BranchID, workspaceID uint64, tableName string) *WorkspaceTableExportRequestBuilder {
	return &WorkspaceTableExportRequestBuilder{
		branchID:    branchID,
		workspaceID: workspaceID,
		config: workspaceTableExportConfig{
			TableName: tableName,
		},
		api: a,
	}
}

func (b *WorkspaceTableExportRequestBuilder) Build() request.APIRequest[*StorageJob] {
	result := &StorageJob{}
	req := b.api.newRequest(StorageAPI).
		WithResult(result).
		WithPost("branch/{branchId}/workspaces/{workspaceId}/table-export").
		AndPathParam("branchId", b.branchID.String()).
		AndPathParam("workspaceId", strconv.FormatUint(b.workspaceID, 10)).
		WithJSONBody(request.StructToMap(b.config, nil))
	return request.NewAPIRequest(result, req)
}

func (b *WorkspaceTableExportRequestBuilder) Send(ctx context.Context) (*StorageJob, error) {
	return b.Build().Send(ctx)
}

// WorkspaceTableExportJobResult contains the result of a workspace table export job.
type WorkspaceTableExportJobResult struct {
	File ExportedFile `json:"file"`
}

type ExportedFile struct {
	FileID FileID `json:"id"`
}

// SendAndWait executes the request and waits for the resulting storage job to finish.
// Once the job finishes, this returns its results object, which contains the exported file information.
func (b *WorkspaceTableExportRequestBuilder) SendAndWait(ctx context.Context, timeout time.Duration) (*WorkspaceTableExportJobResult, error) {
	// send request
	job, err := b.Send(ctx)
	if err != nil {
		return nil, err
	}

	// wait for job
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err = b.api.WaitForStorageJob(timeoutCtx, job)
	if err != nil {
		return nil, err
	}

	// parse result
	result := &WorkspaceTableExportJobResult{}
	data, err := jsonLib.Marshal(job.Results)
	if err != nil {
		return nil, err
	}

	err = jsonLib.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// WithFileName sets the name of the file to export the table to.
// If not specified, a default file name will be generated.
func (b *WorkspaceTableExportRequestBuilder) WithFileName(fileName string) *WorkspaceTableExportRequestBuilder {
	b.config.FileName = fileName
	return b
}

// WithFormat sets the export format (e.g., "csv", "json").
// If not specified, the default format will be used.
func (b *WorkspaceTableExportRequestBuilder) WithFormat(format string) *WorkspaceTableExportRequestBuilder {
	b.config.Format = format
	return b
}
