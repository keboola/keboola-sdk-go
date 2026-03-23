package keboola_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/keboola/go-utils/pkg/testproject"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

// TestWorkspaceTableExportSuccess is in transfer/storage_workspace_export_test.go
// as it requires transfer.Upload.

func TestWorkspaceTableExport(t *testing.T) {
	t.Skip()
	t.Parallel()
	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	ctx, cancelFn := context.WithTimeout(ctx, time.Second*30)
	t.Cleanup(cancelFn)

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	// Create workspace
	networkPolicy := "user"
	workspace := &keboola.StorageWorkspacePayload{
		Backend:       keboola.StorageWorkspaceBackendSnowflake,
		BackendSize:   ptr(keboola.StorageWorkspaceBackendSizeMedium),
		NetworkPolicy: &networkPolicy,
		LoginType:     keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair,
		PublicKey:     ptr(os.Getenv("TEST_SNOWFLAKE_PUBLIC_KEY")), //nolint: forbidigo
	}

	createdWorkspace, err := api.StorageWorkspaceCreateRequest(defBranch.ID, workspace).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, createdWorkspace)

	// Ensure workspace is cleaned up even if test fails
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cleanupCancel()
		_, _ = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).Send(cleanupCtx)
	})

	// Test the workspace table export functionality
	// Note: A full end-to-end test would require:
	// 1. Creating a workspace (done above)
	// 2. Loading a Storage table into the workspace
	// 3. Exporting the table from the workspace
	// 4. Verifying the exported file
	//
	// Since workspace data loading is complex and may not be implemented in all backends,
	// these tests focus on verifying the API structure and error handling.

	t.Run("BuildRequest", func(t *testing.T) {
		t.Parallel()
		// Test building the request without sending
		req := api.NewWorkspaceTableExportRequest(defBranch.ID, createdWorkspace.ID, "test_table").
			WithFileName("exported_table.csv").
			WithFileType("csv").
			Build()

		require.NotNil(t, req)
	})

	t.Run("SendRequestWithAllOptions", func(t *testing.T) {
		t.Parallel()
		// Test sending a request with all options
		// This will fail because the table doesn't exist in the workspace
		_, err := api.NewWorkspaceTableExportRequest(defBranch.ID, createdWorkspace.ID, "test_table").
			WithFileName("exported_table.csv").
			WithFileType("csv").
			Send(ctx)
		require.Error(t, err)
	})

	t.Run("SendRequestMinimal", func(t *testing.T) {
		t.Parallel()
		// Test sending a minimal request with only required parameters
		// This will fail because the table doesn't exist in the workspace
		_, err := api.NewWorkspaceTableExportRequest(defBranch.ID, createdWorkspace.ID, "test_table").
			Send(ctx)
		require.Error(t, err)
	})

	t.Run("SendAndWait", func(t *testing.T) {
		t.Parallel()
		// Test SendAndWait method
		// This will fail because the table doesn't exist in the workspace
		_, err := api.NewWorkspaceTableExportRequest(defBranch.ID, createdWorkspace.ID, "test_table").
			WithFileName("exported_table.csv").
			WithFileType("csv").
			SendAndWait(ctx, time.Minute*2)
		require.Error(t, err)
	})

	t.Run("Format does not match", func(t *testing.T) {
		t.Parallel()
		_, err := api.NewWorkspaceTableExportRequest(defBranch.ID, createdWorkspace.ID, "test_table").
			WithFileName("exported_table.csv").
			WithFileType("fake_format").
			SendAndWait(ctx, time.Minute*2)
		require.Error(t, err)
	})
}
