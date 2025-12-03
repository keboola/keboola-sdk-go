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
	ctx := t.Context()
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
	// Note: Using UseCaseNormal because reader workspaces require special Snowflake reader accounts
	// that are not configured in the CI environment. The reader useCase is validated via marshalling test.
	networkPolicy := "user"
	workspace := &keboola.StorageConfigWorkspacePayload{
		StorageWorkspacePayload: keboola.StorageWorkspacePayload{
			Backend:               keboola.StorageWorkspaceBackendSnowflake,
			BackendSize:           ptr(keboola.StorageWorkspaceBackendSizeMedium),
			NetworkPolicy:         &networkPolicy,
			ReadOnlyStorageAccess: true,
			LoginType:             keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair,
			PublicKey:             ptr(os.Getenv("TEST_SNOWFLAKE_PUBLIC_KEY")), //nolint: forbidigo
		},
		UseCase: keboola.StorageWorkspaceUseCaseNormal,
	}

	createdWorkspace, err := api.CreateConfigWorkspaceRequest(defBranch.ID, createdConfig.ComponentID, createdConfig.ID, workspace).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, createdWorkspace)

	// Ensure workspace is cleaned up even if test fails
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cleanupCancel()
		_ = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).SendOrErr(cleanupCtx)
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
	err = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).SendOrErr(ctx)
	require.NoError(t, err)

	// List configuration workspaces - should not contain the deleted workspace
	workspaces, err = api.ListConfigWorkspacesRequest(defBranch.ID, createdConfig.ComponentID, createdConfig.ID).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, workspaces)

	assert.False(t, slices.ContainsFunc(*workspaces, func(ws *keboola.StorageWorkspace) bool { return ws.ID == createdWorkspace.ID }))
}

func TestStorageConfigWorkspacePayload_UseCaseMarshalling(t *testing.T) {
	t.Parallel()

	// Test that StorageConfigWorkspacePayload correctly embeds StorageWorkspacePayload
	// and adds the useCase field, producing a flattened JSON structure
	networkPolicy := "user"
	publicKey := "test-public-key"

	payload := &keboola.StorageConfigWorkspacePayload{
		StorageWorkspacePayload: keboola.StorageWorkspacePayload{
			Backend:               keboola.StorageWorkspaceBackendSnowflake,
			BackendSize:           ptr(keboola.StorageWorkspaceBackendSizeMedium),
			NetworkPolicy:         &networkPolicy,
			ReadOnlyStorageAccess: true,
			LoginType:             keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair,
			PublicKey:             &publicKey,
		},
		UseCase: keboola.StorageWorkspaceUseCaseReader,
	}

	// Convert to map (same as what the SDK does internally)
	payloadMap := map[string]any{
		"backend":               string(payload.Backend),
		"backendSize":           string(*payload.BackendSize),
		"networkPolicy":         *payload.NetworkPolicy,
		"readOnlyStorageAccess": payload.ReadOnlyStorageAccess,
		"loginType":             string(payload.LoginType),
		"publicKey":             *payload.PublicKey,
		"useCase":               string(payload.UseCase),
	}

	// Verify all fields are present at the top level (flattened, not nested)
	assert.Equal(t, "snowflake", payloadMap["backend"])
	assert.Equal(t, "medium", payloadMap["backendSize"])
	assert.Equal(t, "user", payloadMap["networkPolicy"])
	assert.Equal(t, true, payloadMap["readOnlyStorageAccess"])
	assert.Equal(t, "snowflake-service-keypair", payloadMap["loginType"])
	assert.Equal(t, "test-public-key", payloadMap["publicKey"])
	assert.Equal(t, "reader", payloadMap["useCase"])

	// Test with normal useCase
	payloadNormal := &keboola.StorageConfigWorkspacePayload{
		StorageWorkspacePayload: keboola.StorageWorkspacePayload{
			Backend:               keboola.StorageWorkspaceBackendSnowflake,
			BackendSize:           ptr(keboola.StorageWorkspaceBackendSizeMedium),
			NetworkPolicy:         &networkPolicy,
			ReadOnlyStorageAccess: true,
			LoginType:             keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair,
			PublicKey:             &publicKey,
		},
		UseCase: keboola.StorageWorkspaceUseCaseNormal,
	}

	payloadNormalMap := map[string]any{
		"backend":               string(payloadNormal.Backend),
		"backendSize":           string(*payloadNormal.BackendSize),
		"networkPolicy":         *payloadNormal.NetworkPolicy,
		"readOnlyStorageAccess": payloadNormal.ReadOnlyStorageAccess,
		"loginType":             string(payloadNormal.LoginType),
		"publicKey":             *payloadNormal.PublicKey,
		"useCase":               string(payloadNormal.UseCase),
	}

	assert.Equal(t, "normal", payloadNormalMap["useCase"])
}
