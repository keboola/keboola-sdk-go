package keboola_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

func TestTriggerApiCalls(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	project, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Get default branch
	defBranch, err := api.GetDefaultBranchRequest().Send(ctx)
	require.NoError(t, err)

	// Get current token ID for runWithTokenId
	tokenInfo, err := api.VerifyTokenRequest(project.StorageAPIToken()).Send(ctx)
	require.NoError(t, err)
	runWithTokenID := fmt.Sprint(tokenInfo.ID)

	// Create a config to attach the trigger to
	config := &keboola.ConfigWithRows{
		Config: &keboola.Config{
			ConfigKey: keboola.ConfigKey{
				BranchID:    defBranch.ID,
				ComponentID: "ex-generic-v2",
			},
			Name:              "Trigger test config",
			ChangeDescription: "test",
		},
	}
	_, err = api.CreateConfigRequest(config, false).Send(ctx)
	require.NoError(t, err)
	t.Cleanup(func() {
		cleanCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		_ = api.DeleteConfigRequest(config.ConfigKey).SendOrErr(cleanCtx)
	})

	// Create a bucket and table for the trigger to watch
	bucket, tableKey := createBucketAndTableKey(defBranch)
	_, err = api.CreateBucketRequest(bucket).Send(ctx)
	require.NoError(t, err)

	_, err = api.CreateTableRequest(tableKey, []string{"id", "name"}).Send(ctx)
	require.NoError(t, err)

	tableID := tableKey.TableID

	// --- Create trigger ---
	trigger := &keboola.Trigger{
		RunWithTokenID:        runWithTokenID,
		Component:             config.ComponentID,
		ConfigID:              config.ID,
		CoolDownPeriodMinutes: 5,
		TableIDs:              []keboola.TableID{tableID},
	}
	created, err := api.CreateTriggerRequest(trigger).Send(ctx)
	require.NoError(t, err)
	assert.NotZero(t, created.ID)
	assert.Equal(t, config.ComponentID, created.Component)
	assert.Equal(t, config.ID, created.ConfigID)
	assert.Equal(t, 5, created.CoolDownPeriodMinutes)
	assert.Equal(t, []keboola.TableID{tableID}, created.TableIDs)

	t.Cleanup(func() {
		cleanCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		_ = api.DeleteTriggerRequest(created.TriggerKey).SendOrErr(cleanCtx)
	})

	// --- Get trigger ---
	got, err := api.GetTriggerRequest(created.TriggerKey).Send(ctx)
	require.NoError(t, err)
	assert.Equal(t, created.ID, got.ID)
	assert.Equal(t, created.Component, got.Component)
	assert.Equal(t, created.ConfigID, got.ConfigID)

	// --- List triggers ---
	list, err := api.ListTriggersRequest().Send(ctx)
	require.NoError(t, err)
	found := false
	for _, tr := range *list {
		if tr.ID == created.ID {
			found = true
			break
		}
	}
	assert.True(t, found, "created trigger not found in list")

	// --- Update trigger ---
	created.CoolDownPeriodMinutes = 10
	updated, err := api.UpdateTriggerRequest(created).Send(ctx)
	require.NoError(t, err)
	assert.Equal(t, 10, updated.CoolDownPeriodMinutes)

	// --- Delete trigger ---
	err = api.DeleteTriggerRequest(created.TriggerKey).SendOrErr(ctx)
	require.NoError(t, err)

	// Verify deletion
	list, err = api.ListTriggersRequest().Send(ctx)
	require.NoError(t, err)
	for _, tr := range *list {
		assert.NotEqual(t, created.ID, tr.ID, "deleted trigger still in list")
	}
}
