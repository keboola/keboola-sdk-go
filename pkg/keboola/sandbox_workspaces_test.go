package keboola_test

import (
	"context"
	"testing"
	"time"

	"github.com/keboola/go-utils/pkg/testproject"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

func TestEditorSessionCreateAndDeleteSnowflake(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, branch)

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// Create editor session (replaces deprecated CreateSandboxWorkspace for Snowflake)
	workspace, err := api.CreateEditorSession(ctx, branch.ID, "test")
	require.NoError(t, err)
	require.NotNil(t, workspace)
	require.NotNil(t, workspace.Config)
	require.NotNil(t, workspace.EditorSession)
	assert.NotEmpty(t, workspace.EditorSession.ID)
	assert.Equal(t, keboola.EditorSessionStatusReady, workspace.EditorSession.Status)
	assert.Equal(t, keboola.EditorSessionBackendType("snowflake"), workspace.EditorSession.BackendType)

	defer func() {
		deleteErr := api.DeleteEditorSession(ctx, branch.ID, workspace.Config.ID, workspace.EditorSession.ID)
		require.NoError(t, deleteErr)
	}()

	// List sessions - verify the created session is present
	sessions, err := api.ListEditorSessionsRequest().Send(ctx)
	require.NoError(t, err)
	foundSession := false
	for _, s := range *sessions {
		if s.ID == workspace.EditorSession.ID {
			assert.Equal(t, keboola.EditorSessionBackendType("snowflake"), s.BackendType)
			foundSession = true
			break
		}
	}
	assert.True(t, foundSession, "Session list did not find created session")
}

func TestEditorSessionCreateAndDeleteBigQuery(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithBigQueryBackend())

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, branch)

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// Create editor session (replaces deprecated CreateSandboxWorkspace for BigQuery)
	workspace, err := api.CreateEditorSession(ctx, branch.ID, "test")
	require.NoError(t, err)
	require.NotNil(t, workspace)
	require.NotNil(t, workspace.Config)
	require.NotNil(t, workspace.EditorSession)
	assert.NotEmpty(t, workspace.EditorSession.ID)
	assert.Equal(t, keboola.EditorSessionStatusReady, workspace.EditorSession.Status)
	assert.Equal(t, keboola.EditorSessionBackendType("bigquery"), workspace.EditorSession.BackendType)
	assert.NotEmpty(t, workspace.EditorSession.WorkspaceID, "BigQuery session should have a workspace ID")

	defer func() {
		deleteErr := api.DeleteEditorSession(ctx, branch.ID, workspace.Config.ID, workspace.EditorSession.ID)
		require.NoError(t, deleteErr)
	}()

	// List sessions - verify the created session is present
	sessions, err := api.ListEditorSessionsRequest().Send(ctx)
	require.NoError(t, err)
	foundSession := false
	for _, s := range *sessions {
		if s.ID == workspace.EditorSession.ID {
			assert.Equal(t, keboola.EditorSessionBackendType("bigquery"), s.BackendType)
			foundSession = true
			break
		}
	}
	assert.True(t, foundSession, "Session list did not find created session")
}

func TestEditorSessionListAndGet(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, branch)

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// Create a session to list and get
	workspace, err := api.CreateEditorSession(ctx, branch.ID, "test")
	require.NoError(t, err)
	require.NotNil(t, workspace)

	defer func() {
		deleteErr := api.DeleteEditorSession(ctx, branch.ID, workspace.Config.ID, workspace.EditorSession.ID)
		require.NoError(t, deleteErr)
	}()

	// Get session by ID
	fetched, err := api.GetEditorSessionRequest(workspace.EditorSession.ID).Send(ctx)
	require.NoError(t, err)
	assert.Equal(t, workspace.EditorSession.ID, fetched.ID)
	assert.Equal(t, keboola.EditorSessionStatusReady, fetched.Status)

	// List sessions - verify it appears
	sessions, err := api.ListEditorSessionsRequest().Send(ctx)
	require.NoError(t, err)
	foundSession := false
	for _, s := range *sessions {
		if s.ID == workspace.EditorSession.ID {
			foundSession = true
			break
		}
	}
	assert.True(t, foundSession, "Session list did not include created session")
}

func TestResetEditorSessionCredentials(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := keboola.APIClientForAnEmptyProject(t, ctx, testproject.WithSnowflakeBackend())

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, branch)

	ctx, cancelFn := context.WithTimeout(ctx, time.Minute*10)
	defer cancelFn()

	// Create a session to reset credentials on
	workspace, err := api.CreateEditorSession(ctx, branch.ID, "test")
	require.NoError(t, err)
	require.NotNil(t, workspace)

	defer func() {
		deleteErr := api.DeleteEditorSession(ctx, branch.ID, workspace.Config.ID, workspace.EditorSession.ID)
		require.NoError(t, deleteErr)
	}()

	// Reset credentials
	creds, err := api.ResetEditorSessionCredentialsRequest(workspace.EditorSession.ID).Send(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, creds.PrivateKey)
	assert.NotEmpty(t, creds.PublicKey)
}
