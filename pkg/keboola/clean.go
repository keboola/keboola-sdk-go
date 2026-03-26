package keboola

import (
	"context"
	"sync"

	"github.com/hashicorp/go-multierror"
)

// CleanStorageWorkspaces deletes all storage workspaces in the project.
// This method lists all workspaces and deletes them one by one.
func (a *AuthorizedAPI) CleanStorageWorkspaces(ctx context.Context) error {
	// Get default branch
	defBranch, err := a.GetDefaultBranchRequest().Send(ctx)
	if err != nil {
		return err
	}

	// List all storage workspaces
	workspaces, err := a.StorageWorkspacesListRequest(defBranch.ID).Send(ctx)
	if err != nil {
		return err
	}

	// Delete each workspace using a waitgroup for concurrent deletion
	var wg sync.WaitGroup
	var mu sync.Mutex
	var result error

	for _, workspace := range *workspaces {
		wg.Add(1)
		go func(workspaceID uint64) {
			defer wg.Done()
			if _, deleteErr := a.StorageWorkspaceDeleteRequest(defBranch.ID, workspaceID).Send(ctx); deleteErr != nil {
				mu.Lock()
				result = multierror.Append(result, deleteErr)
				mu.Unlock()
			}
		}(workspace.ID)
	}

	// Wait for all workspace deletions to complete
	wg.Wait()

	return result
}

func CleanProject(
	ctx context.Context,
	api *AuthorizedAPI,
) error {
	var err error

	// Clean schedules first
	if e := api.CleanAllSchedulesRequest().SendOrErr(ctx); e != nil {
		err = multierror.Append(err, e)
	}

	// Clean notification subscriptions
	if e := api.CleanAllNotificationSubscriptionsRequest().SendOrErr(ctx); e != nil {
		err = multierror.Append(err, e)
	}

	if _, ok := api.Index().ServiceURLByID(ServiceID(DataScienceAPI)); ok {
		if e := api.CleanSandboxWorkspaceInstances(ctx); e != nil {
			err = multierror.Append(err, e)
		}
	}

	if _, ok := api.Index().ServiceURLByID(ServiceID(EditorAPI)); ok {
		if e := api.CleanEditorSessions(ctx); e != nil {
			err = multierror.Append(err, e)
		}
	}

	// Clean storage workspaces
	if e := api.CleanStorageWorkspaces(ctx); e != nil {
		err = multierror.Append(err, e)
	}

	// Clean the rest of the project
	if e := api.CleanProjectRequest().SendOrErr(ctx); e != nil {
		err = multierror.Append(err, e)
	}

	if err != nil {
		return err
	}

	return nil
}
