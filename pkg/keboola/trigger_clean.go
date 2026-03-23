package keboola

import (
	"context"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

// CleanAllTriggersRequest cleans all triggers in the project.
// Useful for E2E tests.
func (a *AuthorizedAPI) CleanAllTriggersRequest() request.APIRequest[request.NoResult] {
	req := a.ListTriggersRequest().
		WithOnSuccess(func(ctx context.Context, result *[]*Trigger) error {
			wg := request.NewWaitGroup(ctx)
			for _, trigger := range *result {
				wg.Send(a.DeleteTriggerRequest(trigger.TriggerKey))
			}
			return wg.Wait()
		})
	return request.NewAPIRequest(request.NoResult{}, req)
}
