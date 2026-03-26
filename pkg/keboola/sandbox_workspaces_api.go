package keboola

import (
	"context"
	"sync"

	"github.com/hashicorp/go-multierror"
)

type SandboxWorkspaceID string

func (v SandboxWorkspaceID) String() string {
	return string(v)
}

func (a *AuthorizedAPI) CleanSandboxWorkspaceInstances(ctx context.Context) error {
	// Use the /apps endpoint for all non-editor app types. Snowflake/BigQuery are managed
	// by the editor service and cleaned by CleanEditorSessions, so they are intentionally excluded.
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
			if e := a.DeleteSandboxWorkspaceJobRequest(SandboxWorkspaceID(app.ID)).SendOrErr(ctx); e != nil {
				m.Lock()
				defer m.Unlock()
				err = multierror.Append(err, e)
			}
		})
	}

	wg.Wait()
	return err
}
