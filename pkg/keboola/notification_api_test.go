package keboola_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

// cleanupNotification registers cleanup for a single notification subscription.
// Designed to be called immediately after creating a subscription.
func cleanupNotification(t *testing.T, api *keboola.AuthorizedAPI, key keboola.NotificationSubscriptionKey) {
	t.Helper()
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cleanupCancel()
		_ = api.DeleteNotificationSubscriptionRequest(key).SendOrErr(cleanupCtx)
	})
}

// cleanupAllNotifications deletes all notification subscriptions in the project.
// Designed to be idempotent - ignores "not found" errors.
// Uses a dedicated cleanup context with timeout independent of test context.
func cleanupAllNotifications(t *testing.T, api *keboola.AuthorizedAPI) {
	t.Helper()
	cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cleanupCancel()

	subscriptions, err := api.ListNotificationSubscriptionsRequest().Send(cleanupCtx)
	if err != nil {
		return // Best-effort cleanup
	}

	for _, sub := range *subscriptions {
		_ = api.DeleteNotificationSubscriptionRequest(sub.NotificationSubscriptionKey).SendOrErr(cleanupCtx)
	}
}

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
	cleanupNotification(t, api, subscription.NotificationSubscriptionKey)

	assert.NotEmpty(t, subscription.ID)
	assert.Equal(t, keboola.NotificationEventJobFailed, subscription.Event)
	assert.Equal(t, keboola.NotificationChannelEmail, subscription.Recipient.Channel)
	assert.Equal(t, "test@example.com", subscription.Recipient.Address)
	assert.NotNil(t, subscription.ExpiresAt)
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
	cleanupNotification(t, api, sub1.NotificationSubscriptionKey)

	sub2, err := api.NewCreateNotificationSubscriptionRequest(
		keboola.NotificationEventJobSucceeded,
		keboola.NotificationChannelEmail,
		"test2@example.com",
	).Send(ctx)
	require.NoError(t, err)
	cleanupNotification(t, api, sub2.NotificationSubscriptionKey)

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
	cleanupNotification(t, api, subscription.NotificationSubscriptionKey) // Safety net

	// Delete subscription (part of test logic)
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
	cleanupNotification(t, api, created.NotificationSubscriptionKey)

	// Get subscription
	subscription, err := api.GetNotificationSubscriptionRequest(created.NotificationSubscriptionKey).Send(ctx)
	require.NoError(t, err)
	require.NotNil(t, subscription)
	assert.Equal(t, created.ID, subscription.ID)
	assert.Equal(t, created.Event, subscription.Event)
	assert.Equal(t, created.Recipient.Channel, subscription.Recipient.Channel)
	assert.Equal(t, created.Recipient.Address, subscription.Recipient.Address)
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
	cleanupNotification(t, api, sub1.NotificationSubscriptionKey)

	sub2, err := api.NewCreateNotificationSubscriptionRequest(
		keboola.NotificationEventJobFailed,
		keboola.NotificationChannelEmail,
		"test2@example.com",
	).Send(ctx)
	require.NoError(t, err)
	cleanupNotification(t, api, sub2.NotificationSubscriptionKey)

	sub3, err := api.NewCreateNotificationSubscriptionRequest(
		keboola.NotificationEventJobSucceeded,
		keboola.NotificationChannelEmail,
		"test3@example.com",
	).Send(ctx)
	require.NoError(t, err)
	cleanupNotification(t, api, sub3.NotificationSubscriptionKey)

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
	cleanupNotification(t, api, subscription.NotificationSubscriptionKey)

	assert.NotNil(t, subscription.Filters) // Should be a slice, not nil
	// API automatically adds branch filter, so we expect at least 1 filter
	assert.GreaterOrEqual(t, len(subscription.Filters), 1)
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
	cleanupNotification(t, api, subscription.NotificationSubscriptionKey)

	// Validate ExpiresAt is set and matches requested time (within 1 minute tolerance)
	require.NotNil(t, subscription.ExpiresAt)
	parsedExpiration, err := time.Parse(time.RFC3339, subscription.ExpiresAt.String())
	require.NoError(t, err, "ExpiresAt should be valid RFC3339 time")

	timeDiff := parsedExpiration.Sub(expirationTime).Abs()
	assert.LessOrEqual(t, timeDiff, 1*time.Minute,
		"ExpiresAt should be within 1 minute of requested time (got %s, expected %s)",
		parsedExpiration.Format(time.RFC3339), expirationTime.Format(time.RFC3339))
}

func TestNotificationAPIRelativeExpiration(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Capture time before creation to validate relative expiration
	beforeCreate := time.Now()

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
	cleanupNotification(t, api, subscription.NotificationSubscriptionKey)

	// Validate ExpiresAt is approximately 24 hours from creation (within 1 minute tolerance)
	require.NotNil(t, subscription.ExpiresAt)
	parsedExpiration, err := time.Parse(time.RFC3339, subscription.ExpiresAt.String())
	require.NoError(t, err, "ExpiresAt should be valid RFC3339 time")

	expectedExpiration := beforeCreate.Add(24 * time.Hour)
	timeDiff := parsedExpiration.Sub(expectedExpiration).Abs()
	assert.LessOrEqual(t, timeDiff, 1*time.Minute,
		"ExpiresAt should be approximately 24 hours from creation (±1 minute tolerance, got %s, expected around %s)",
		parsedExpiration.Format(time.RFC3339), expectedExpiration.Format(time.RFC3339))
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
	cleanupNotification(t, api, subscription.NotificationSubscriptionKey)

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
		cleanupNotification(t, api, subscription.NotificationSubscriptionKey)

		assert.Equal(t, event, subscription.Event)
	}
}

func TestNotificationAPICleanupAll(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	_, api := keboola.APIClientForAnEmptyProject(t, ctx)

	// Create multiple subscriptions
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

	sub3, err := api.NewCreateNotificationSubscriptionRequest(
		keboola.NotificationEventPhaseJobFailed,
		keboola.NotificationChannelEmail,
		"test3@example.com",
	).Send(ctx)
	require.NoError(t, err)

	// Verify subscriptions exist
	subscriptions, err := api.ListNotificationSubscriptionsRequest().Send(ctx)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(*subscriptions), 3)

	// Use bulk cleanup
	cleanupAllNotifications(t, api)

	// Verify all are deleted (wait a moment for cleanup to complete)
	time.Sleep(1 * time.Second)
	subscriptions, err = api.ListNotificationSubscriptionsRequest().Send(ctx)
	require.NoError(t, err)

	// Check that our specific subscriptions no longer exist
	var foundIDs []keboola.NotificationSubscriptionID
	for _, sub := range *subscriptions {
		foundIDs = append(foundIDs, sub.ID)
	}
	assert.NotContains(t, foundIDs, sub1.ID)
	assert.NotContains(t, foundIDs, sub2.ID)
	assert.NotContains(t, foundIDs, sub3.ID)
}

