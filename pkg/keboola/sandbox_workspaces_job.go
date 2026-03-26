package keboola

import (
	"context"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

func (a *AuthorizedAPI) DeleteSandboxWorkspaceJobRequest(workspaceID SandboxWorkspaceID) request.APIRequest[request.NoResult] {
	configData := map[string]any{
		"parameters": map[string]any{
			"task": "delete",
			"id":   workspaceID.String(),
		},
	}
	req := a.NewCreateJobRequest(SandboxWorkspacesComponent).
		WithConfigData(configData).
		Build().WithOnSuccess(func(ctx context.Context, result *QueueJob) error {
		return a.WaitForQueueJob(ctx, result.ID)
	})
	return request.NewAPIRequest(request.NoResult{}, req)
}
