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
	require.NoError(t, err)
	require.NotNil(t, branch)

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// Create workspace
	workspace, err := api.CreateSandboxWorkspace(
		ctx,
		branch.ID,
		"test",
		keboola.SandboxWorkspaceTypePython,
		keboola.WithExpireAfterHours(1),
		keboola.WithSize(keboola.SandboxWorkspaceSizeMedium),
	)
	require.NoError(t, err)
	require.NotNil(t, workspace)

	defer func() {
		// Delete workspace
		err = api.DeleteSandboxWorkspace(
			ctx,
			branch.ID,
			workspace.Config.ID,
			workspace.SandboxWorkspace.ID,
		)
		require.NoError(t, err)
	}()

	// List workspaces - try to find the one we just created
	workspaces, err := api.ListSandboxWorkspaces(ctx, branch.ID)
	require.NoError(t, err)
	foundInstance := false
	for _, v := range workspaces {
		if workspace.SandboxWorkspace.ID == v.SandboxWorkspace.ID {
			require.True(t, v.SandboxWorkspace.Type == keboola.SandboxWorkspaceTypePython)
			foundInstance = true
			break
		}
	}
	assert.True(t, foundInstance, "Workspace list did not find created workspace")
}

func TestWorkspacesCreateAndDeleteSnowflake(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, branch)

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// Create workspace
	workspace, err := api.CreateSandboxWorkspace(
		ctx,
		branch.ID,
		"test-snowflake",
		keboola.SandboxWorkspaceTypeSnowflake,
		keboola.WithExpireAfterHours(1),
		keboola.WithPublicKey(os.Getenv("TEST_SNOWFLAKE_PUBLIC_KEY")), //nolint: forbidigo
	)
	require.NoError(t, err)
	require.NotNil(t, workspace)

	defer func() {
		// Delete workspace
		err = api.DeleteSandboxWorkspace(
			ctx,
			branch.ID,
			workspace.Config.ID,
			workspace.SandboxWorkspace.ID,
		)
		require.NoError(t, err)
	}()

	// List workspaces - try to find the one we just created
	workspaces, err := api.ListSandboxWorkspaces(ctx, branch.ID)
	require.NoError(t, err)
	foundInstance := false
	for _, v := range workspaces {
		if workspace.SandboxWorkspace.ID == v.SandboxWorkspace.ID {
			require.True(t, v.SandboxWorkspace.Type == keboola.SandboxWorkspaceTypeSnowflake)
			foundInstance = true
			break
		}
	}
	assert.True(t, foundInstance, "Workspace list did not find created workspace")
}

func TestWorkspacesCreateAndDeleteBigQuery(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithBigQueryBackend())

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, branch)

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// Create workspace
	workspace, err := api.CreateSandboxWorkspace(
		ctx,
		branch.ID,
		"test-bigquery",
		keboola.SandboxWorkspaceTypeBigQuery,
		keboola.WithExpireAfterHours(1),
	)
	require.NoError(t, err)
	require.NotNil(t, workspace)

	defer func() {
		// Delete workspace
		err = api.DeleteSandboxWorkspace(
			ctx,
			branch.ID,
			workspace.Config.ID,
			workspace.SandboxWorkspace.ID,
		)
		require.NoError(t, err)
	}()

	// Verify that credentials are populated for BigQuery workspace
	require.NotNil(t, workspace.SandboxWorkspace.Credentials, "BigQuery workspace should have credentials")
	assert.NotEmpty(t, workspace.SandboxWorkspace.Credentials.Type, "Credentials type should not be empty")
	assert.NotEmpty(t, workspace.SandboxWorkspace.Credentials.ProjectID, "Credentials project_id should not be empty")
	assert.NotEmpty(t, workspace.SandboxWorkspace.Credentials.PrivateKeyID, "Credentials private_key_id should not be empty")
	assert.NotEmpty(t, workspace.SandboxWorkspace.Credentials.ClientEmail, "Credentials client_email should not be empty")
	assert.NotEmpty(t, workspace.SandboxWorkspace.Credentials.ClientID, "Credentials client_id should not be empty")
	assert.NotEmpty(t, workspace.SandboxWorkspace.Credentials.AuthURI, "Credentials auth_uri should not be empty")
	assert.NotEmpty(t, workspace.SandboxWorkspace.Credentials.TokenURI, "Credentials token_uri should not be empty")
	assert.NotEmpty(t, workspace.SandboxWorkspace.Credentials.AuthProviderX509CertURL, "Credentials auth_provider_x509_cert_url should not be empty")
	assert.NotEmpty(t, workspace.SandboxWorkspace.Credentials.ClientX509CertURL, "Credentials client_x509_cert_url should not be empty")
	assert.NotEmpty(t, workspace.SandboxWorkspace.Credentials.PrivateKey, "Credentials private_key should not be empty")

	// List workspaces - try to find the one we just created
	workspaces, err := api.ListSandboxWorkspaces(ctx, branch.ID)
	require.NoError(t, err)
	foundInstance := false
	for _, v := range workspaces {
		if workspace.SandboxWorkspace.ID == v.SandboxWorkspace.ID {
			require.True(t, v.SandboxWorkspace.Type == keboola.SandboxWorkspaceTypeBigQuery)
			foundInstance = true
			break
		}
	}
	assert.True(t, foundInstance, "Workspace list did not find created workspace")
}
