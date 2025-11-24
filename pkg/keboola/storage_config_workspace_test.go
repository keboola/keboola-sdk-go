package keboola_test

import (
	"context"
	"os"
	"slices"
	"testing"
	"time"

	"github.com/keboola/go-utils/pkg/orderedmap"
	"github.com/keboola/go-utils/pkg/testproject"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

func TestConfigWorkspacesCreateAndListSnowflake(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	// Create a test configuration first
	config := &keboola.ConfigWithRows{
		Config: &keboola.Config{
			ConfigKey: keboola.ConfigKey{
				BranchID:    defBranch.ID,
				ComponentID: "keboola.app-data-gateway",
			},
			Name:        "Test Data Gateway Config",
			Description: "Test configuration for workspace creation",
			Content: orderedmap.FromPairs([]orderedmap.Pair{
				{Key: "parameters", Value: orderedmap.FromPairs([]orderedmap.Pair{
					{Key: "db", Value: orderedmap.New()},
				})},
			}),
		},
		Rows: []*keboola.ConfigRow{},
	}
	createdConfig, err := api.CreateConfigRequest(config, true).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, createdConfig)
	require.NotEmpty(t, createdConfig.ID)

	// Cleanup config at the end
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cleanupCancel()
		_, _ = api.DeleteConfigRequest(createdConfig.ConfigKey).Send(cleanupCtx)
	})

	// Create workspace for the configuration
	networkPolicy := "user"
	workspace := &keboola.StorageWorkspacePayload{
		Backend:               keboola.StorageWorkspaceBackendSnowflake,
		BackendSize:           ptr(keboola.StorageWorkspaceBackendSizeMedium),
		NetworkPolicy:         &networkPolicy,
		ReadOnlyStorageAccess: true,
		LoginType:             keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair,
		PublicKey:             ptr(os.Getenv("TEST_SNOWFLAKE_PUBLIC_KEY")), //nolint: forbidigo
	}

	createdWorkspace, err := api.CreateConfigWorkspaceRequest(defBranch.ID, createdConfig.ComponentID, createdConfig.ID, workspace).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, createdWorkspace)

	// Ensure workspace is cleaned up even if test fails
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cleanupCancel()
		_, _ = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).Send(cleanupCtx)
	})

	assert.Equal(t, keboola.StorageWorkspaceBackendSnowflake, createdWorkspace.StorageWorkspaceDetails.Backend)
	assert.Equal(t, keboola.StorageWorkspaceBackendSizeMedium, *createdWorkspace.BackendSize)
	assert.Equal(t, string(keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair), *createdWorkspace.StorageWorkspaceDetails.LoginType)

	// List configuration workspaces - should contain the created workspace
	workspaces, err := api.ListConfigWorkspacesRequest(defBranch.ID, createdConfig.ComponentID, createdConfig.ID).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, workspaces)

	assert.True(t, slices.ContainsFunc(*workspaces, func(ws *keboola.StorageWorkspace) bool { return ws.ID == createdWorkspace.ID }))

	// Delete workspace
	deletedWorkspace, err := api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, deletedWorkspace)

	// List configuration workspaces - should not contain the deleted workspace
	workspaces, err = api.ListConfigWorkspacesRequest(defBranch.ID, createdConfig.ComponentID, createdConfig.ID).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, workspaces)

	assert.False(t, slices.ContainsFunc(*workspaces, func(ws *keboola.StorageWorkspace) bool { return ws.ID == deletedWorkspace.ID }))
}
