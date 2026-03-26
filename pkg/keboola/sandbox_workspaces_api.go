package keboola

import (
	"context"
	"sync"

	"github.com/hashicorp/go-multierror"
)

// CleanSandboxWorkspaceInstances deletes all Python/R/Streamlit/PythonJS data-science sandboxes
// in the project using the direct DELETE /sandboxes/{appId} endpoint.
// Snowflake/BigQuery are managed by the editor service and cleaned by CleanEditorSessions.
func (a *AuthorizedAPI) CleanSandboxWorkspaceInstances(ctx context.Context) error {
	apps, err := a.ListDataScienceAppsRequest(
		WithDataScienceAppsType(
			DataScienceAppTypePython,
			DataScienceAppTypeR,
			DataScienceAppTypeStreamlit,
			DataScienceAppTypePythonJS,
		),
		WithDataScienceAppsLimit(500),
	).Send(ctx)
	if err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}

	for _, app := range *apps {
		wg.Go(func() {
			if e := a.DeleteDataScienceSandboxRequest(app.ID).SendOrErr(ctx); e != nil {
				m.Lock()
				defer m.Unlock()
				err = multierror.Append(err, e)
			}
		})
	}

	wg.Wait()
	return err
}
