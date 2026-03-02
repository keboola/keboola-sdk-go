package keboola

import (
	"context"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

// CleanAllNotificationSubscriptionsRequest cleans all notification subscriptions in whole project.
// Useful for E2E tests.
func (a *AuthorizedAPI) CleanAllNotificationSubscriptionsRequest() request.APIRequest[request.NoResult] {
	req := a.ListNotificationSubscriptionsRequest().
		WithOnSuccess(func(ctx context.Context, result *[]*NotificationSubscription) error {
			wg := request.NewWaitGroup(ctx)
			for _, subscription := range *result {
				wg.Send(a.DeleteNotificationSubscriptionRequest(subscription.NotificationSubscriptionKey))
			}
			return wg.Wait()
		})
	return request.NewAPIRequest(request.NoResult{}, req)
}
