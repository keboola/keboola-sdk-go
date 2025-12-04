package keboola_test

import (
	"bytes"
	"context"
	"os"
	"slices"
	"testing"
	"time"

	"github.com/keboola/go-utils/pkg/testproject"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

func TestStorageWorkspacesCreateAndDeleteSnowflake(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

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
		_ = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).SendOrErr(cleanupCtx)
	})

	assert.Equal(t, keboola.StorageWorkspaceBackendSnowflake, createdWorkspace.StorageWorkspaceDetails.Backend)
	assert.Equal(t, keboola.StorageWorkspaceBackendSizeMedium, *createdWorkspace.BackendSize)
	assert.Equal(t, string(keboola.StorageWorkspaceLoginTypeSnowflakeServiceKeypair), *createdWorkspace.StorageWorkspaceDetails.LoginType)

	// Get workspace details
	retrievedWorkspace, err := api.StorageWorkspaceDetailRequest(defBranch.ID, createdWorkspace.ID).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, retrievedWorkspace)
	assert.Equal(t, createdWorkspace.ID, retrievedWorkspace.ID)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Backend, retrievedWorkspace.StorageWorkspaceDetails.Backend)
	assert.Equal(t, createdWorkspace.BackendSize, retrievedWorkspace.BackendSize)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.LoginType, retrievedWorkspace.StorageWorkspaceDetails.LoginType)

	// Create credentials
	credentials, err := api.StorageWorkspaceCreateCredentialsRequest(defBranch.ID, createdWorkspace.ID).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, credentials)
	assert.Equal(t, credentials.ID, credentials.ID)
	assert.Equal(t, credentials.StorageWorkspaceDetails.Backend, credentials.StorageWorkspaceDetails.Backend)
	assert.Equal(t, credentials.StorageWorkspaceDetails.Host, credentials.StorageWorkspaceDetails.Host)
	assert.Contains(t, *credentials.StorageWorkspaceDetails.User, "QS")
	assert.NotEmpty(t, *credentials.StorageWorkspaceDetails.Account)
	assert.Contains(t, *credentials.StorageWorkspaceDetails.Role, "WORKSPACE")
	assert.Contains(t, *credentials.StorageWorkspaceDetails.Schema, "WORKSPACE") //nolint: goconst
	assert.Contains(t, *credentials.StorageWorkspaceDetails.Warehouse, "KEBOOLA")

	// Fetch credentials
	fetchedCredentials, err := api.StorageWorkspaceFetchCredentialsRequest(defBranch.ID, createdWorkspace.ID, credentials.StorageWorkspaceDetails.Credentials.ID).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, fetchedCredentials)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Backend, fetchedCredentials.StorageWorkspaceDetails.Backend)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Host, fetchedCredentials.StorageWorkspaceDetails.Host)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Region, fetchedCredentials.StorageWorkspaceDetails.Region)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Database, fetchedCredentials.StorageWorkspaceDetails.Database)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.Schema, fetchedCredentials.StorageWorkspaceDetails.Schema)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.LoginType, fetchedCredentials.StorageWorkspaceDetails.LoginType)
	assert.Equal(t, createdWorkspace.StorageWorkspaceDetails.SSOLoginAvailable, fetchedCredentials.StorageWorkspaceDetails.SSOLoginAvailable)

	// Delete credentials
	deletedCredentials, err := api.StorageWorkspaceDeleteCredentialsRequest(defBranch.ID, createdWorkspace.ID, credentials.StorageWorkspaceDetails.Credentials.ID).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, deletedCredentials)

	// List workspaces - should contain the created workspace
	workspaces, err := api.StorageWorkspacesListRequest(defBranch.ID).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, workspaces)

	assert.True(t, slices.ContainsFunc(*workspaces, func(ws *keboola.StorageWorkspace) bool { return ws.ID == createdWorkspace.ID }))

	// Delete workspace
	err = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).SendOrErr(ctx)
	require.NoError(t, err)

	// List workspaces - should not contain the deleted workspace
	workspaces, err = api.StorageWorkspacesListRequest(defBranch.ID).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, workspaces)

	require.False(t, slices.ContainsFunc(*workspaces, func(ws *keboola.StorageWorkspace) bool { return ws.ID == createdWorkspace.ID }))
}

func TestStorageWorkspacesCreateWrongBigQuery(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithBigQueryBackend())

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// List workspaces - record initial count
	workspaces, err := api.StorageWorkspacesListRequest(defBranch.ID).Send(ctx)
	require.NoError(t, err)
	initialLen := len(*workspaces)

	// Create workspace - should fail
	workspace := &keboola.StorageWorkspacePayload{
		Backend:     keboola.StorageWorkspaceBackendBigQuery,
		BackendSize: ptr(keboola.StorageWorkspaceBackendSizeMedium),
		LoginType:   keboola.StorageWorkspaceLoginTypeDefault,
	}

	_, err = api.StorageWorkspaceCreateRequest(defBranch.ID, workspace).Send(ctx)
	assert.Error(t, err)

	// List again - no new workspace should have been created
	workspacesAfter, err := api.StorageWorkspacesListRequest(defBranch.ID).Send(ctx)
	require.NoError(t, err)
	assert.Len(t, *workspacesAfter, initialLen, "Workspace count should not increase when creation fails")
}

func TestStorageWorkspaceLoadData(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	// Create bucket and table
	bucket, tableKey := createBucketAndTableKey(defBranch)
	_, err = api.CreateBucketRequest(bucket).Send(ctx)
	require.NoError(t, err)

	// Create file with test data
	fileName := "test_data.csv"
	file, err := api.CreateFileResourceRequest(defBranch.ID, fileName).Send(ctx)
	require.NoError(t, err)

	// Upload file
	content := []byte("col1,col2\nval1,val2\nval3,val4\n")
	_, err = keboola.Upload(ctx, file, bytes.NewReader(content))
	require.NoError(t, err)

	// Create table from file
	_, err = api.CreateTableFromFileRequest(tableKey, file.FileKey).Send(ctx)
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
		_ = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).SendOrErr(cleanupCtx)
	})

	// Load data into workspace
	loadPayload := &keboola.WorkspaceLoadDataPayload{
		Input: []keboola.WorkspaceLoadDataInput{
			{
				Source:      tableKey.TableID.String(),
				Destination: "test_table",
			},
		},
	}

	job, err := api.StorageWorkspaceLoadDataRequest(defBranch.ID, createdWorkspace.ID, loadPayload).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, job)

	// Wait for job to complete
	waitCtx, waitCancelFn := context.WithTimeout(ctx, time.Minute*5)
	defer waitCancelFn()
	err = api.WaitForStorageJob(waitCtx, job)
	require.NoError(t, err)
	assert.Equal(t, "success", job.Status)
}

func TestStorageWorkspacesCreateAndDeleteBigQuery(t *testing.T) {
	t.Parallel()
	t.Skip("Skipping BigQuery test until we have a way to create a project with BigQuery backend")
	ctx := t.Context()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithBigQueryBackend())

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	// List workspaces - should be empty initially
	workspaces, err := api.StorageWorkspacesListRequest(defBranch.ID).Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *workspaces, 0, "Workspace list should be empty initially")

	// Create workspace
	workspace := &keboola.StorageWorkspacePayload{
		Backend:   keboola.StorageWorkspaceBackendBigQuery,
		LoginType: keboola.StorageWorkspaceLoginTypeDefault,
	}

	createdWorkspace, err := api.StorageWorkspaceCreateRequest(defBranch.ID, workspace).Send(ctx)
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
	retrievedWorkspace, err := api.StorageWorkspaceDetailRequest(defBranch.ID, createdWorkspace.ID).Send(ctx)
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
	workspaces, err = api.StorageWorkspacesListRequest(defBranch.ID).Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *workspaces, 1, "Workspace list should contain one workspace")
	assert.Equal(t, createdWorkspace.ID, (*workspaces)[0].ID)

	// Delete workspace
	err = api.StorageWorkspaceDeleteRequest(defBranch.ID, createdWorkspace.ID).SendOrErr(ctx)
	assert.NoError(t, err)

	// List workspaces - should be empty again
	workspaces, err = api.StorageWorkspacesListRequest(defBranch.ID).Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *workspaces, 0, "Workspace list should be empty after deletion")
}
