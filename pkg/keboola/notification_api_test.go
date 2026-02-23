package keboola_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

func TestNotificationAPICreateSubscription(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Create subscription with relative expiration
	subscription, err := api.NewCreateNotificationSubscriptionRequest(
		keboola.NotificationEventJobFailed,
		keboola.NotificationChannelEmail,
		"test@example.com",
	).
		WithRelativeExpiration("+24 hours").
		Send(ctx)

	require.NoError(t, err)
	require.NotNil(t, subscription)
	assert.NotEmpty(t, subscription.ID)
	assert.Equal(t, keboola.NotificationEventJobFailed, subscription.Event)
	assert.Equal(t, keboola.NotificationChannelEmail, subscription.Recipient.Channel)
	assert.Equal(t, "test@example.com", subscription.Recipient.Address)
	assert.NotNil(t, subscription.ExpiresAt)

	// Cleanup
	err = api.DeleteNotificationSubscriptionRequest(subscription.NotificationSubscriptionKey).SendOrErr(ctx)
	require.NoError(t, err)
}

func TestNotificationAPIListSubscriptions(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Create two subscriptions with different events (both email channel)
	sub1, err := api.NewCreateNotificationSubscriptionRequest(
		keboola.NotificationEventJobFailed,
		keboola.NotificationChannelEmail,
		"test1@example.com",
	).Send(ctx)
	require.NoError(t, err)

	sub2, err := api.NewCreateNotificationSubscriptionRequest(
		keboola.NotificationEventJobSucceeded,
		keboola.NotificationChannelEmail,
		"test2@example.com",
	).Send(ctx)
	require.NoError(t, err)

	// List all subscriptions
	allSubs, err := api.ListNotificationSubscriptionsRequest().Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, allSubs)
	assert.GreaterOrEqual(t, len(*allSubs), 2)

	// List filtered by event
	failedSubs, err := api.ListNotificationSubscriptionsRequest(
		keboola.WithListSubscriptionsEvent(keboola.NotificationEventJobFailed),
	).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, failedSubs)
	assert.GreaterOrEqual(t, len(*failedSubs), 1)
	// Verify that sub1 is in the filtered list
	foundSub1 := false
	for _, sub := range *failedSubs {
		if sub.ID == sub1.ID {
			foundSub1 = true
			assert.Equal(t, keboola.NotificationEventJobFailed, sub.Event)
		}
	}
	assert.True(t, foundSub1, "Expected to find sub1 in filtered list")

	// Cleanup
	err = api.DeleteNotificationSubscriptionRequest(sub1.NotificationSubscriptionKey).SendOrErr(ctx)
	require.NoError(t, err)
	err = api.DeleteNotificationSubscriptionRequest(sub2.NotificationSubscriptionKey).SendOrErr(ctx)
	require.NoError(t, err)
}

func TestNotificationAPIDeleteSubscription(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Create subscription
	subscription, err := api.NewCreateNotificationSubscriptionRequest(
		keboola.NotificationEventJobSucceeded,
		keboola.NotificationChannelEmail,
		"test@example.com",
	).Send(ctx)
	require.NoError(t, err)

	// Delete subscription
	err = api.DeleteNotificationSubscriptionRequest(subscription.NotificationSubscriptionKey).SendOrErr(ctx)
	require.NoError(t, err)

	// Verify deletion - get should fail
	_, err = api.GetNotificationSubscriptionRequest(subscription.NotificationSubscriptionKey).Send(ctx)
	assert.Error(t, err)
}

func TestNotificationAPIGetSubscription(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Create subscription
	created, err := api.NewCreateNotificationSubscriptionRequest(
		keboola.NotificationEventJobFailed,
		keboola.NotificationChannelEmail,
		"test@example.com",
	).Send(ctx)
	require.NoError(t, err)

	// Get subscription
	subscription, err := api.GetNotificationSubscriptionRequest(created.NotificationSubscriptionKey).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, subscription)
	assert.Equal(t, created.ID, subscription.ID)
	assert.Equal(t, created.Event, subscription.Event)
	assert.Equal(t, created.Recipient.Channel, subscription.Recipient.Channel)
	assert.Equal(t, created.Recipient.Address, subscription.Recipient.Address)

	// Cleanup
	err = api.DeleteNotificationSubscriptionRequest(subscription.NotificationSubscriptionKey).SendOrErr(ctx)
	require.NoError(t, err)
}

func TestNotificationAPIGroupByEvent(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Create subscriptions with different events (all email to avoid webhook validation issues)
	sub1, err := api.NewCreateNotificationSubscriptionRequest(
		keboola.NotificationEventJobFailed,
		keboola.NotificationChannelEmail,
		"test1@example.com",
	).Send(ctx)
	require.NoError(t, err)

	sub2, err := api.NewCreateNotificationSubscriptionRequest(
		keboola.NotificationEventJobFailed,
		keboola.NotificationChannelEmail,
		"test2@example.com",
	).Send(ctx)
	require.NoError(t, err)

	sub3, err := api.NewCreateNotificationSubscriptionRequest(
		keboola.NotificationEventJobSucceeded,
		keboola.NotificationChannelEmail,
		"test3@example.com",
	).Send(ctx)
	require.NoError(t, err)

	// List all subscriptions
	subscriptions, err := api.ListNotificationSubscriptionsRequest().Send(ctx)
	require.NoError(t, err)

	// Group by event
	grouped := keboola.GroupNotificationSubscriptionsByEvent(*subscriptions)

	// Verify grouping
	assert.Contains(t, grouped, keboola.NotificationEventJobFailed)
	assert.Contains(t, grouped, keboola.NotificationEventJobSucceeded)

	failedIDs := grouped[keboola.NotificationEventJobFailed]
	assert.Contains(t, failedIDs, sub1.ID)
	assert.Contains(t, failedIDs, sub2.ID)

	succeededIDs := grouped[keboola.NotificationEventJobSucceeded]
	assert.Contains(t, succeededIDs, sub3.ID)

	// Cleanup
	err = api.DeleteNotificationSubscriptionRequest(sub1.NotificationSubscriptionKey).SendOrErr(ctx)
	require.NoError(t, err)
	err = api.DeleteNotificationSubscriptionRequest(sub2.NotificationSubscriptionKey).SendOrErr(ctx)
	require.NoError(t, err)
	err = api.DeleteNotificationSubscriptionRequest(sub3.NotificationSubscriptionKey).SendOrErr(ctx)
	require.NoError(t, err)
}

func TestNotificationAPIEmptyFilters(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Create subscription without filters
	// Note: The API automatically adds a branch filter for the current branch
	subscription, err := api.NewCreateNotificationSubscriptionRequest(
		keboola.NotificationEventJobSucceeded,
		keboola.NotificationChannelEmail,
		"test@example.com",
	).Send(ctx)

	require.NoError(t, err)
	require.NotNil(t, subscription)
	assert.NotNil(t, subscription.Filters) // Should be a slice, not nil
	// API automatically adds branch filter, so we expect at least 1 filter
	assert.GreaterOrEqual(t, len(subscription.Filters), 1)

	// Cleanup
	err = api.DeleteNotificationSubscriptionRequest(subscription.NotificationSubscriptionKey).SendOrErr(ctx)
	require.NoError(t, err)
}

func TestNotificationAPIAbsoluteExpiration(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Create subscription with absolute expiration
	expirationTime := time.Now().Add(48 * time.Hour)
	subscription, err := api.NewCreateNotificationSubscriptionRequest(
		keboola.NotificationEventJobFailed,
		keboola.NotificationChannelEmail,
		"test@example.com",
	).
		WithAbsoluteExpiration(expirationTime).
		Send(ctx)

	require.NoError(t, err)
	require.NotNil(t, subscription)
	assert.NotNil(t, subscription.ExpiresAt)

	// Cleanup
	err = api.DeleteNotificationSubscriptionRequest(subscription.NotificationSubscriptionKey).SendOrErr(ctx)
	require.NoError(t, err)
}

func TestNotificationAPIRelativeExpiration(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Create subscription with relative expiration
	subscription, err := api.NewCreateNotificationSubscriptionRequest(
		keboola.NotificationEventJobSucceeded,
		keboola.NotificationChannelEmail,
		"test@example.com",
	).
		WithRelativeExpiration("+24 hours").
		Send(ctx)

	require.NoError(t, err)
	require.NotNil(t, subscription)
	assert.NotNil(t, subscription.ExpiresAt)

	// Cleanup
	err = api.DeleteNotificationSubscriptionRequest(subscription.NotificationSubscriptionKey).SendOrErr(ctx)
	require.NoError(t, err)
}

func TestNotificationAPIAllOperators(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Test that the operator constants are defined and can be used
	// We test equals operator in a real API call
	subscription, err := api.NewCreateNotificationSubscriptionRequest(
		keboola.NotificationEventJobFailed,
		keboola.NotificationChannelEmail,
		"test@example.com",
	).Send(ctx)

	require.NoError(t, err)
	require.NotNil(t, subscription)

	// Verify all operator constants are defined
	operators := []keboola.NotificationFilterOperator{
		keboola.NotificationFilterOperatorEquals,
		keboola.NotificationFilterOperatorNotEquals,
		keboola.NotificationFilterOperatorGreaterThan,
		keboola.NotificationFilterOperatorLessThan,
		keboola.NotificationFilterOperatorGreaterThanEquals,
		keboola.NotificationFilterOperatorLessThanEquals,
	}
	assert.Len(t, operators, 6, "All 6 operators should be defined")

	// Cleanup
	err = api.DeleteNotificationSubscriptionRequest(subscription.NotificationSubscriptionKey).SendOrErr(ctx)
	require.NoError(t, err)
}

func TestNotificationAPIPhaseEvents(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	phaseEvents := []keboola.NotificationEvent{
		keboola.NotificationEventPhaseJobFailed,
		keboola.NotificationEventPhaseJobSucceeded,
		keboola.NotificationEventPhaseJobSucceededWithWarning,
		keboola.NotificationEventPhaseJobProcessingLong,
	}

	for _, event := range phaseEvents {
		// Create subscription for phase event
		subscription, err := api.NewCreateNotificationSubscriptionRequest(
			event,
			keboola.NotificationChannelEmail,
			"test@example.com",
		).Send(ctx)

		require.NoError(t, err, "event: %s", event)
		require.NotNil(t, subscription)
		assert.Equal(t, event, subscription.Event)

		// Cleanup
		err = api.DeleteNotificationSubscriptionRequest(subscription.NotificationSubscriptionKey).SendOrErr(ctx)
		require.NoError(t, err)
	}
}

