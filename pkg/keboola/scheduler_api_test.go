package keboola_test

import (
	"context"
	"testing"

	"github.com/keboola/go-utils/pkg/orderedmap"
	"github.com/keboola/go-utils/pkg/testproject"
	"github.com/stretchr/testify/assert"

	"github.com/keboola/keboola-sdk-go/v2/pkg/client"
	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

func TestSchedulerApiCalls(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	project, _ := testproject.GetTestProjectForTest(t)
	c := client.NewTestClient()
	api, err := keboola.NewAuthorizedAPI(ctx, project.StorageAPIHost(), project.StorageAPIToken(), keboola.WithClient(&c))
	assert.NoError(t, err)

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, branch)

	// Create a config to schedule
	targetConfig := &keboola.ConfigWithRows{
		Config: &keboola.Config{
			ConfigKey: keboola.ConfigKey{
				BranchID:    branch.ID,
				ComponentID: "ex-generic-v2",
			},
			Name:              "Test",
			Description:       "Test description",
			ChangeDescription: "My test",
			Content: orderedmap.FromPairs([]orderedmap.Pair{
				{
					Key: "foo",
					Value: orderedmap.FromPairs([]orderedmap.Pair{
						{Key: "bar", Value: "baz"},
					}),
				},
			}),
		},
	}
	_, err = api.CreateConfigRequest(targetConfig, true).Send(ctx)
	assert.NoError(t, err)

	// Create scheduler config
	schedulerConfig := &keboola.ConfigWithRows{
		Config: &keboola.Config{
			ConfigKey: keboola.ConfigKey{
				BranchID:    branch.ID,
				ComponentID: "keboola.scheduler",
			},
			Name:              "Test",
			Description:       "Test description",
			ChangeDescription: "My test",
			Content: orderedmap.FromPairs([]orderedmap.Pair{
				{
					Key: "schedule",
					Value: orderedmap.FromPairs([]orderedmap.Pair{
						{Key: "cronTab", Value: "*/2 * * * *"},
						{Key: "timezone", Value: "UTC"},
						{Key: "state", Value: "disabled"},
					}),
				},
				{
					Key: "target",
					Value: orderedmap.FromPairs([]orderedmap.Pair{
						{Key: "componentId", Value: "ex-generic-v2"},
						{Key: "configurationId", Value: targetConfig.ID},
						{Key: "mode", Value: "run"},
					}),
				},
			}),
		},
	}
	_, err = api.CreateConfigRequest(schedulerConfig, true).Send(ctx)
	assert.NoError(t, err)

	// List should return no schedule
	schedules, err := api.ListSchedulesRequest().Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *schedules, 0)

	// Activate
	schedule, err := api.ActivateScheduleRequest(schedulerConfig.ID, "").Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, schedule)
	assert.NotEmpty(t, schedule.ID)

	// List should return one schedule
	schedules, err = api.ListSchedulesRequest().Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *schedules, 1)

	schedule, err = api.GetScheduleRequest(schedule.ScheduleKey).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, &keboola.Schedule{
		ScheduleKey: keboola.ScheduleKey{
			ID: schedule.ID,
		},
		ConfigID:               schedulerConfig.ID,
		TokenID:                schedule.TokenID,
		ConfigurationVersionID: "1",
		ScheduleCron: keboola.ScheduleCron{
			CronTab:  "*/2 * * * *",
			Timezone: "UTC",
			State:    "disabled",
		},
		ScheduleTarget: keboola.ScheduleTarget{
			ComponentID:     "ex-generic-v2",
			ConfigurationID: targetConfig.ID,
			Mode:            "run",
		},
		Executions: []keboola.ScheduleExecution{},
	}, schedule)

	// Delete
	_, err = api.DeleteScheduleRequest(schedule.ScheduleKey).Send(ctx)
	assert.NoError(t, err)

	// List should return no schedule
	schedules, err = api.ListSchedulesRequest().Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *schedules, 0)

	// Activate again
	schedule, err = api.ActivateScheduleRequest(schedulerConfig.ID, "").Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, schedule)
	assert.NotEmpty(t, schedule.ID)

	// List should return one schedule
	schedules, err = api.ListSchedulesRequest().Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *schedules, 1)

	tokenDetailBeforeRefresh, err := api.TokenDetailRequest(schedule.TokenID).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, tokenDetailBeforeRefresh)
	assert.Equal(t, tokenDetailBeforeRefresh.ID, schedule.TokenID)

	// Store the refresh time to compare later
	refreshTimeBeforeRefresh := tokenDetailBeforeRefresh.Refreshed.Time.String()

	// Create a variable to store the refreshed token
	var tokenDetailAfterRefresh *keboola.Token

	// Refresh token with a callback to check the token details immediately after refresh
	_, err = api.RefreshScheduleTokenRequest(schedule.ID).
		WithOnSuccess(func(ctx context.Context, _ request.NoResult) error {
			// Request token details right after the refresh operation
			var fetchErr error
			tokenDetailAfterRefresh, fetchErr = api.TokenDetailRequest(tokenDetailBeforeRefresh.ID).Send(ctx)
			return fetchErr
		}).
		Send(ctx)
	assert.NoError(t, err)

	// Verify token was properly fetched in the callback
	assert.NotNil(t, tokenDetailAfterRefresh)
	assert.Equal(t, tokenDetailAfterRefresh.ID, schedule.TokenID)
	
	// Verify the token was actually refreshed by comparing refresh times
	assert.NotEqual(t, refreshTimeBeforeRefresh, tokenDetailAfterRefresh.Refreshed.Time.String())

	// Delete for configuration
	_, err = api.DeleteSchedulesForConfigurationRequest(schedulerConfig.ID).Send(ctx)
	assert.NoError(t, err)

	// List should return no schedule
	schedules, err = api.ListSchedulesRequest().Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *schedules, 0)
}
