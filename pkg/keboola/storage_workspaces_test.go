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
	workspace := &keboola.StorageWorkspacePayload{
		Backend:     keboola.StorageWorkspaceBackendSnowflake,
		BackendSize: ptr(keboola.StorageWorkspaceBackendSizeMedium),
		LoginType:   keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair,
		PublicKey:   ptr(os.Getenv("TEST_SNOWFLAKE_PUBLIC_KEY")), //nolint: forbidigo
	}

	createdWorkspace, err := api.StorageWorkspaceCreateRequest(workspace).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, createdWorkspace)
	assert.Equal(t, keboola.StorageWorkspaceBackendSnowflake, createdWorkspace.StorageWorkspaceDetails.Backend)
	assert.Equal(t, keboola.StorageWorkspaceBackendSizeMedium, *createdWorkspace.BackendSize)
	assert.Equal(t, string(keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair), *createdWorkspace.StorageWorkspaceDetails.LoginType)

	// Get workspace details
	retrievedWorkspace, err := api.StorageWorkspaceDetailRequest(createdWorkspace.ID).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedWorkspace)
	assert.Equal(t, createdWorkspace.ID, retrievedWorkspace.ID)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Backend, retrievedWorkspace.StorageWorkspaceDetails.Backend)
	assert.Equal(t, createdWorkspace.BackendSize, retrievedWorkspace.BackendSize)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.LoginType, retrievedWorkspace.StorageWorkspaceDetails.LoginType)

	/*createCredentials := &keboola.StorageWorkspacePayload{
		ID: createdWorkspace.ID,
	}
	// Create credentials
	credentials, err := api.StorageWorkspaceCreateCredentialsRequest(createCredentials).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, credentials)
	assert.Equal(t, createdWorkspace.ID, credentials.ID)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Backend, credentials.StorageWorkspaceDetails.Backend)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Host, credentials.StorageWorkspaceDetails.Host)*/

	// Fetch credentials
	/*fetchedCredentials, err := api.StorageWorkspaceFetchCredentialsRequest(createdWorkspace.ID, credentials.StorageWorkspaceDetails.Credentials.ID).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedCredentials)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Backend, fetchedCredentials.Backend)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Host, fetchedCredentials.Host)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Region, fetchedCredentials.Region)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Database, fetchedCredentials.Database)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Schema, fetchedCredentials.Schema)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Warehouse, fetchedCredentials.Warehouse)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.LoginType, fetchedCredentials.LoginType)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.SSOLoginAvailable, fetchedCredentials.SSOLoginAvailable)*/

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

func TestStorageWorkspacesCreateWrongBigQuery(t *testing.T) {
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
	workspace := &keboola.StorageWorkspacePayload{
		Backend:     keboola.StorageWorkspaceBackendBigQuery,
		BackendSize: ptr(keboola.StorageWorkspaceBackendSizeMedium),
		LoginType:   keboola.StorageWorkspaceLoginTypeDefault,
	}

	_, err = api.StorageWorkspaceCreateRequest(workspace).Send(ctx)
	assert.Error(t, err)
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
	workspace := &keboola.StorageWorkspacePayload{
		Backend:   keboola.StorageWorkspaceBackendBigQuery,
		LoginType: keboola.StorageWorkspaceLoginTypeDefault,
	}

	createdWorkspace, err := api.StorageWorkspaceCreateRequest(workspace).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, createdWorkspace)
	assert.Equal(t, keboola.StorageWorkspaceBackendBigQuery, createdWorkspace.StorageWorkspaceDetails.Backend)
	assert.Equal(t, string(keboola.StorageWorkspaceLoginTypeDefault), *createdWorkspace.StorageWorkspaceDetails.LoginType)

	// Verify that credentials are populated for BigQuery workspace
	assert.NotNil(t, createdWorkspace.StorageWorkspaceDetails.Credentials, "BigQuery workspace should have credentials")
	credentials := createdWorkspace.StorageWorkspaceDetails.Credentials
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
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Backend, retrievedWorkspace.StorageWorkspaceDetails.Backend)
	assert.Equal(t, createdWorkspace.BackendSize, retrievedWorkspace.BackendSize)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.LoginType, retrievedWorkspace.StorageWorkspaceDetails.LoginType)

	// Create credentials
	/*createCredentials := &keboola.StorageWorkspacePayload{
		ID: createdWorkspace.ID,
	}
	createdCredentials, err := api.StorageWorkspaceCreateCredentialsRequest(createCredentials).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, createdCredentials)

	// Fetch credentials
	fetchedCredentials, err := api.StorageWorkspaceFetchCredentialsRequest(createdWorkspace.ID, createdCredentials.StorageWorkspaceDetails.Credentials.ID).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedCredentials)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Credentials.Type, fetchedCredentials.Credentials.Type)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Credentials.ProjectID, fetchedCredentials.Credentials.ProjectID)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Credentials.PrivateKeyID, fetchedCredentials.Credentials.PrivateKeyID)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Credentials.ClientEmail, fetchedCredentials.Credentials.ClientEmail)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Credentials.ClientID, fetchedCredentials.Credentials.ClientID)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Credentials.AuthURI, fetchedCredentials.Credentials.AuthURI)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Credentials.TokenURI, fetchedCredentials.Credentials.TokenURI)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Credentials.AuthProviderX509CertURL, fetchedCredentials.Credentials.AuthProviderX509CertURL)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Credentials.ClientX509CertURL, fetchedCredentials.Credentials.ClientX509CertURL)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Credentials.PrivateKey, fetchedCredentials.Credentials.PrivateKey)*/

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
