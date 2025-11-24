package keboola

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

// CreateConfigWorkspaceRequestAsync https://keboola.docs.apiary.io/#reference/workspaces/configuration-workspaces-collection/create-configuration-workspace
// Creates a configuration workspace asynchronously and returns a storage job that can be monitored for completion.
func (a *AuthorizedAPI) CreateConfigWorkspaceRequestAsync(branchID BranchID, componentID ComponentID, configID ConfigID, payload *StorageWorkspacePayload) request.APIRequest[*StorageJob] {
	result := &StorageJob{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithPost("branch/{branchId}/components/{componentId}/configs/{configId}/workspaces").
		AndPathParam("branchId", branchID.String()).
		AndPathParam("componentId", componentID.String()).
		AndPathParam("configId", configID.String()).
		WithJSONBody(request.StructToMap(payload, nil)).
		AndQueryParam("async", "1")
	return request.NewAPIRequest(result, req)
}

// CreateConfigWorkspaceRequest https://keboola.docs.apiary.io/#reference/workspaces/configuration-workspaces-collection/create-configuration-workspace
func (a *AuthorizedAPI) CreateConfigWorkspaceRequest(branchID BranchID, componentID ComponentID, configID ConfigID, payload *StorageWorkspacePayload) request.APIRequest[*StorageWorkspace] {
	result := &StorageWorkspace{}
	req := a.
		CreateConfigWorkspaceRequestAsync(branchID, componentID, configID, payload).
		WithOnSuccess(func(ctx context.Context, job *StorageJob) error {
			// Wait for storage job
			waitCtx, waitCancelFn := context.WithTimeout(ctx, a.onSuccessTimeout)
			defer waitCancelFn()
			if err := a.WaitForStorageJob(waitCtx, job); err != nil {
				return err
			}

			// Map job results to workspace
			resultsBytes, err := json.Marshal(job.Results)
			if err != nil {
				return fmt.Errorf("cannot convert job.results to JSON: %w", err)
			}
			if err := json.Unmarshal(resultsBytes, result); err != nil {
				return fmt.Errorf("cannot map job.results to workspace: %w", err)
			}
			return nil
		})
	// Result is workspace, not job.
	return request.NewAPIRequest(result, req)
}

// ListConfigWorkspacesRequest https://keboola.docs.apiary.io/#reference/workspaces/configuration-workspaces-collection/list-configuration-workspaces
func (a *AuthorizedAPI) ListConfigWorkspacesRequest(branchID BranchID, componentID ComponentID, configID ConfigID) request.APIRequest[*[]*StorageWorkspace] {
	result := make([]*StorageWorkspace, 0)
	req := a.
		newRequest(StorageAPI).
		WithResult(&result).
		WithGet("branch/{branchId}/components/{componentId}/configs/{configId}/workspaces").
		AndPathParam("branchId", branchID.String()).
		AndPathParam("componentId", componentID.String()).
		AndPathParam("configId", configID.String())
	return request.NewAPIRequest(&result, req)
}
