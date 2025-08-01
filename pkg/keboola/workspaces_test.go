package keboola_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/keboola/go-utils/pkg/testproject"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

func TestWorkspacesCreateAndDeletePython(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, branch)

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// Create workspace
	workspace, err := api.CreateWorkspace(
		ctx,
		branch.ID,
		"test",
		keboola.WorkspaceTypePython,
		keboola.WithExpireAfterHours(1),
		keboola.WithSize(keboola.WorkspaceSizeMedium),
	)
	assert.NoError(t, err)
	assert.NotNil(t, workspace)

	// List workspaces - try to find the one we just created
	workspaces, err := api.ListWorkspaces(ctx, branch.ID)
	assert.NoError(t, err)
	foundInstance := false
	for _, v := range workspaces {
		if workspace.Workspace.ID == v.Workspace.ID {
			require.True(t, v.Workspace.Type == string(keboola.WorkspaceTypePython))
			foundInstance = true
			break
		}
	}
	assert.True(t, foundInstance, "Workspace list did not find created workspace")

	// Delete workspace
	err = api.DeleteWorkspace(
		ctx,
		branch.ID,
		workspace.Config.ID,
		workspace.Workspace.ID,
	)
	assert.NoError(t, err)
}

func TestWorkspacesCreateAndDeleteSnowflake(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, branch)

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// Create workspace
	workspace, err := api.CreateWorkspace(
		ctx,
		branch.ID,
		"test-snowflake",
		keboola.WorkspaceTypeSnowflake,
		keboola.WithExpireAfterHours(1),
		keboola.WithPublicKey(os.Getenv("TEST_SNOWFLAKE_PUBLIC_KEY")), //nolint: forbidigo
	)
	assert.NoError(t, err)
	assert.NotNil(t, workspace)

	// List workspaces - try to find the one we just created
	workspaces, err := api.ListWorkspaces(ctx, branch.ID)
	assert.NoError(t, err)
	foundInstance := false
	for _, v := range workspaces {
		if workspace.Workspace.ID == v.Workspace.ID {
			require.True(t, v.Workspace.Type == string(keboola.WorkspaceTypeSnowflake))
			foundInstance = true
			break
		}
	}
	assert.True(t, foundInstance, "Workspace list did not find created workspace")

	// Delete workspace
	err = api.DeleteWorkspace(
		ctx,
		branch.ID,
		workspace.Config.ID,
		workspace.Workspace.ID,
	)
	assert.NoError(t, err)
}

func TestWorkspacesCreateAndDeleteBigQuery(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithBigQueryBackend())

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, branch)

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// Create workspace
	workspace, err := api.CreateWorkspace(
		ctx,
		branch.ID,
		"test-bigquery",
		keboola.WorkspaceTypeBigQuery,
		keboola.WithExpireAfterHours(1),
	)
	assert.NoError(t, err)
	assert.NotNil(t, workspace)

	// List workspaces - try to find the one we just created
	workspaces, err := api.ListWorkspaces(ctx, branch.ID)
	assert.NoError(t, err)
	foundInstance := false
	for _, v := range workspaces {
		if workspace.Workspace.ID == v.Workspace.ID {
			require.True(t, v.Workspace.Type == string(keboola.WorkspaceTypeBigQuery))
			foundInstance = true
			break
		}
	}
	assert.True(t, foundInstance, "Workspace list did not find created workspace")

	// Delete workspace
	err = api.DeleteWorkspace(
		ctx,
		branch.ID,
		workspace.Config.ID,
		workspace.Workspace.ID,
	)
	assert.NoError(t, err)
}
