package keboola_test

import (
	"context"
	"testing"
	"time"

	"github.com/keboola/go-utils/pkg/testproject"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

// TestWorkspaceTableExportSuccess is in transfer/storage_workspace_export_test.go
// as it requires transfer.Upload.

func TestWorkspaceTableExport(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	ctx, cancelFn := context.WithTimeout(ctx, 5*time.Minute)
	t.Cleanup(cancelFn)

	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	createdWorkspace, err := api.StorageWorkspaceCreateRequest(defBranch.ID, &keboola.StorageWorkspacePayload{
		Backend:   keboola.StorageWorkspaceBackendSnowflake,
		LoginType: keboola.StorageWorkspaceLoginTypeDefault,
	}).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, createdWorkspace)

	t.Cleanup(func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		_, _ = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).Send(cleanupCtx)
	})

	runWorkspaceTableExportSubtests(t, ctx, api, defBranch.ID, createdWorkspace.ID)
}

func TestWorkspaceTableExportBigQuery(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithBigQueryBackend())

	ctx, cancelFn := context.WithTimeout(ctx, 5*time.Minute)
	t.Cleanup(cancelFn)

	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	createdWorkspace, err := api.StorageWorkspaceCreateRequest(defBranch.ID, &keboola.StorageWorkspacePayload{
		Backend:   keboola.StorageWorkspaceBackendBigQuery,
		LoginType: keboola.StorageWorkspaceLoginTypeDefault,
	}).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, createdWorkspace)

	t.Cleanup(func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		_, _ = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).Send(cleanupCtx)
	})

	runWorkspaceTableExportSubtests(t, ctx, api, defBranch.ID, createdWorkspace.ID)
}

// runWorkspaceTableExportSubtests runs error-path subtests that work for any backend.
// SendRequestWithAllOptions is intentionally omitted: Snowflake rejects synchronously while
// BigQuery queues the job and only fails on wait — SendAndWait covers that case for both.
func runWorkspaceTableExportSubtests(t *testing.T, ctx context.Context, api *keboola.AuthorizedAPI, branchID keboola.BranchID, workspaceID uint64) {
	t.Helper()

	t.Run("BuildRequest", func(t *testing.T) {
		t.Parallel()
		req := api.NewWorkspaceTableExportRequest(branchID, workspaceID, "test_table").
			WithFileName("exported_table.csv").
			WithFileType("csv").
			WithGzip(true).
			Build()
		require.NotNil(t, req)
	})

	t.Run("SendRequestMinimal", func(t *testing.T) {
		t.Parallel()
		_, err := api.NewWorkspaceTableExportRequest(branchID, workspaceID, "test_table").
			Send(ctx)
		require.Error(t, err)
	})

	t.Run("SendAndWait", func(t *testing.T) {
		t.Parallel()
		_, err := api.NewWorkspaceTableExportRequest(branchID, workspaceID, "test_table").
			WithFileName("exported_table.csv").
			WithFileType("csv").
			SendAndWait(ctx, 2*time.Minute)
		require.Error(t, err)
	})

	t.Run("Format does not match", func(t *testing.T) {
		t.Parallel()
		_, err := api.NewWorkspaceTableExportRequest(branchID, workspaceID, "test_table").
			WithFileName("exported_table.csv").
			WithFileType("fake_format").
			SendAndWait(ctx, 2*time.Minute)
		require.Error(t, err)
	})
}
