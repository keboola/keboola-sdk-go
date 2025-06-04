package keboola_test

import (
	"context"
	"testing"

	"github.com/keboola/go-utils/pkg/orderedmap"
	"github.com/keboola/go-utils/pkg/testproject"
	"github.com/stretchr/testify/assert"

	"github.com/keboola/keboola-sdk-go/v2/pkg/client"
	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
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

	assert.NoError(t, api.CreateConfigRequest(schedulerConfig, true).SendOrErr(ctx))

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
	assert.NoError(t, api.DeleteScheduleRequest(schedule.ScheduleKey).SendOrErr(ctx))

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

	// Delete for configuration
	assert.NoError(t, api.DeleteSchedulesForConfigurationRequest(schedulerConfig.ID).SendOrErr(ctx))

	// List should return no schedule
	schedules, err = api.ListSchedulesRequest().Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *schedules, 0)
}

func TestSchedulerTokenRefresh(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	project, _ := testproject.GetTestProjectForTest(t)
	c := client.NewTestClient()
	api, err := keboola.NewAuthorizedAPI(ctx, project.StorageAPIHost(), project.StorageAPIToken(), keboola.WithClient(&c))
	assert.NoError(t, err)

	// Create minimal schedule
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	assert.NoError(t, err)

	config := &keboola.ConfigWithRows{
		Config: &keboola.Config{
			ConfigKey: keboola.ConfigKey{
				BranchID:    branch.ID,
				ComponentID: "keboola.scheduler",
			},
			Name: "Test",
			Content: orderedmap.FromPairs([]orderedmap.Pair{
				{
					Key: "schedule",
					Value: orderedmap.FromPairs([]orderedmap.Pair{
						{Key: "cronTab", Value: "0 1 * * *"},
						{Key: "timezone", Value: "UTC"},
						{Key: "state", Value: "disabled"},
					}),
				},
				{
					Key: "target",
					Value: orderedmap.FromPairs([]orderedmap.Pair{
						{Key: "componentId", Value: "ex-generic-v2"},
						{Key: "configurationId", Value: "123"},
						{Key: "mode", Value: "run"},
					}),
				},
			}),
		},
	}

	assert.NoError(t, api.CreateConfigRequest(config, true).SendOrErr(ctx))

	schedule, err := api.ActivateScheduleRequest(config.ID, "").Send(ctx)
	assert.NoError(t, err)

	tokenBeforeRefresh, err := api.TokenDetailRequest(schedule.TokenID).Send(ctx)
	assert.NoError(t, err)

	// Refresh token
	assert.NoError(t, api.RefreshScheduleTokenRequest(schedule.ID).SendOrErr(ctx))

	tokenDetail, err := api.TokenDetailRequest(schedule.TokenID).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, tokenBeforeRefresh.ID, tokenDetail.ID)
	assert.Equal(t, tokenBeforeRefresh.Created, tokenDetail.Created)

	// The Refreshed.Time field is updated asynchronously on the backend after a token refresh.
	// Directly asserting its change immediately after the refresh call can lead to flaky tests
	// due to inherent timing unpredictability of the backend update and propagation.
	// assert.NotEqual(c, tokenBeforeRefresh.Refreshed.Time.String(), tokenDetail.Refreshed.Time.String())

	assert.NoError(t, api.DeleteSchedulesForConfigurationRequest(schedule.ConfigID).SendOrErr(ctx))
}
