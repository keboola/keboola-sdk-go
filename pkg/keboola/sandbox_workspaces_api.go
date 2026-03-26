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

type SandboxWorkspaceDetails struct {
	Connection struct {
		Database  string `json:"database"`
		Schema    string `json:"schema"`
		Warehouse string `json:"warehouse"`
	} `json:"connection"`
}

// SandboxWorkspaceCredentials contains authentication credentials for the workspace.
type SandboxWorkspaceCredentials struct {
	Type                    string `json:"type"`                        // nolint: tagliatelle
	ProjectID               string `json:"project_id"`                  // nolint: tagliatelle
	PrivateKeyID            string `json:"private_key_id"`              // nolint: tagliatelle
	ClientEmail             string `json:"client_email"`                // nolint: tagliatelle
	ClientID                string `json:"client_id"`                   // nolint: tagliatelle
	AuthURI                 string `json:"auth_uri"`                    // nolint: tagliatelle
	TokenURI                string `json:"token_uri"`                   // nolint: tagliatelle
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"` // nolint: tagliatelle
	ClientX509CertURL       string `json:"client_x509_cert_url"`        // nolint: tagliatelle
	PrivateKey              string `json:"private_key"` //nolint:gosec,tagliatelle
}

type SandboxWorkspace struct {
	ID       SandboxWorkspaceID    `json:"id"`
	Type     SandboxWorkspaceType  `json:"type"`
	Size     string                `json:"size"` // Only exists for container workspaces (Python, R)
	Active   bool                  `json:"active"`
	Shared   bool                  `json:"shared"`
	User     string                `json:"user"`
	Host     string                `json:"host"`
	URL      string                `json:"url"`
	Password string                `json:"password"` //nolint:gosec
	Created  SandboxWorkspacesTime `json:"createdTimestamp"`
	Updated  SandboxWorkspacesTime `json:"updatedTimestamp"`
	Start    SandboxWorkspacesTime `json:"startTimestamp"`

	// SandboxWorkspace details - only exists for Snowflake workspaces
	Details *SandboxWorkspaceDetails `json:"workspaceDetails"`

	// SandboxWorkspace credentials - contains authentication information
	// Only exists for BigQuery workspaces
	Credentials *SandboxWorkspaceCredentials `json:"credentials"`
}

func (a *AuthorizedAPI) CleanSandboxWorkspaceInstances(ctx context.Context) error {
	// Use the /apps endpoint filtered to Python/R — Snowflake/BigQuery are managed by the
	// editor service and cleaned by CleanEditorSessions, so they are intentionally excluded.
	apps, err := a.ListDataScienceAppsRequest(
		WithDataScienceAppsType(DataScienceAppTypePython, DataScienceAppTypeR),
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
