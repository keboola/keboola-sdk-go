package keboola_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/keboola/go-utils/pkg/testproject"
	"github.com/stretchr/testify/assert"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

func TestStorageWorkspacesCreateAndDeleteSnowflake(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// List workspaces - should be empty initially
	workspaces, err := api.StorageWorkspacesListRequest().Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *workspaces, 0, "Workspace list should be empty initially")

	// Create workspace
	workspace := &keboola.StorageWorkspace{
		Backend:     keboola.StorageWorkspaceBackendSnowflake,
		BackendSize: keboola.StorageWorkspaceBackendSizeMedium,
		LoginType:   keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair,
		PublicKey:   os.Getenv("TEST_SNOWFLAKE_PUBLIC_KEY"), //nolint: forbidigo
	}

	createdWorkspace, err := api.StorageWorkspaceCreateRequest("test-snowflake-workspace", workspace).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, createdWorkspace)
	assert.Equal(t, "test-snowflake-workspace", createdWorkspace.ID)
	assert.Equal(t, keboola.StorageWorkspaceBackendSnowflake, createdWorkspace.Backend)
	assert.Equal(t, keboola.StorageWorkspaceBackendSizeMedium, createdWorkspace.BackendSize)
	assert.Equal(t, keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair, createdWorkspace.LoginType)

	// Get workspace details
	retrievedWorkspace, err := api.StorageWorkspaceDetailRequest(createdWorkspace.ID).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedWorkspace)
	assert.Equal(t, createdWorkspace.ID, retrievedWorkspace.ID)
	assert.Equal(t, createdWorkspace.Backend, retrievedWorkspace.Backend)
	assert.Equal(t, createdWorkspace.BackendSize, retrievedWorkspace.BackendSize)
	assert.Equal(t, createdWorkspace.LoginType, retrievedWorkspace.LoginType)

	// List workspaces - should contain the created workspace
	workspaces, err = api.StorageWorkspacesListRequest().Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *workspaces, 1, "Workspace list should contain one workspace")
	assert.Equal(t, createdWorkspace.ID, (*workspaces)[0].ID)

	// Delete workspace
	deletedWorkspace, err := api.StorageWorkspaceDeleteRequest(createdWorkspace.ID).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, deletedWorkspace)

	// List workspaces - should be empty again
	workspaces, err = api.StorageWorkspacesListRequest().Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *workspaces, 0, "Workspace list should be empty after deletion")
}

func TestStorageWorkspacesCreateAndDeleteBigQuery(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithBigQueryBackend())

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// List workspaces - should be empty initially
	workspaces, err := api.StorageWorkspacesListRequest().Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *workspaces, 0, "Workspace list should be empty initially")

	// Create workspace
	workspace := &keboola.StorageWorkspace{
		Backend:     keboola.StorageWorkspaceBackendBigQuery,
		BackendSize: keboola.StorageWorkspaceBackendSizeMedium,
		LoginType:   keboola.StorageWorkspaceLoginTypeDefault,
	}

	createdWorkspace, err := api.StorageWorkspaceCreateRequest("test-bigquery-workspace", workspace).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, createdWorkspace)
	assert.Equal(t, "test-bigquery-workspace", createdWorkspace.ID)
	assert.Equal(t, keboola.StorageWorkspaceBackendBigQuery, createdWorkspace.Backend)
	assert.Equal(t, keboola.StorageWorkspaceBackendSizeMedium, createdWorkspace.BackendSize)
	assert.Equal(t, keboola.StorageWorkspaceLoginTypeDefault, createdWorkspace.LoginType)

	// Verify that credentials are populated for BigQuery workspace
	assert.NotNil(t, createdWorkspace.Details.Connection.Credentials, "BigQuery workspace should have credentials")
	credentials := createdWorkspace.Details.Connection.Credentials
	assert.NotEmpty(t, credentials.Type, "Credentials type should not be empty")
	assert.NotEmpty(t, credentials.ProjectID, "Credentials project_id should not be empty")
	assert.NotEmpty(t, credentials.PrivateKeyID, "Credentials private_key_id should not be empty")
	assert.NotEmpty(t, credentials.ClientEmail, "Credentials client_email should not be empty")
	assert.NotEmpty(t, credentials.ClientID, "Credentials client_id should not be empty")
	assert.NotEmpty(t, credentials.AuthURI, "Credentials auth_uri should not be empty")
	assert.NotEmpty(t, credentials.TokenURI, "Credentials token_uri should not be empty")
	assert.NotEmpty(t, credentials.AuthProviderX509CertURL, "Credentials auth_provider_x509_cert_url should not be empty")
	assert.NotEmpty(t, credentials.ClientX509CertURL, "Credentials client_x509_cert_url should not be empty")
	assert.NotEmpty(t, credentials.PrivateKey, "Credentials private_key should not be empty")

	// Get workspace details
	retrievedWorkspace, err := api.StorageWorkspaceDetailRequest(createdWorkspace.ID).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedWorkspace)
	assert.Equal(t, createdWorkspace.ID, retrievedWorkspace.ID)
	assert.Equal(t, createdWorkspace.Backend, retrievedWorkspace.Backend)
	assert.Equal(t, createdWorkspace.BackendSize, retrievedWorkspace.BackendSize)
	assert.Equal(t, createdWorkspace.LoginType, retrievedWorkspace.LoginType)

	// List workspaces - should contain the created workspace
	workspaces, err = api.StorageWorkspacesListRequest().Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *workspaces, 1, "Workspace list should contain one workspace")
	assert.Equal(t, createdWorkspace.ID, (*workspaces)[0].ID)

	// Delete workspace
	deletedWorkspace, err := api.StorageWorkspaceDeleteRequest(createdWorkspace.ID).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, deletedWorkspace)

	// List workspaces - should be empty again
	workspaces, err = api.StorageWorkspacesListRequest().Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *workspaces, 0, "Workspace list should be empty after deletion")
}
