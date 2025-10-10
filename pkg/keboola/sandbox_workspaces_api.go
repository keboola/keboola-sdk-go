package keboola

import (
	"context"
	"sync"

	"github.com/hashicorp/go-multierror"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
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
	PrivateKey              string `json:"private_key"`                 // nolint: tagliatelle
}

type SandboxWorkspace struct {
	ID       SandboxWorkspaceID    `json:"id"`
	Type     string                `json:"type"`
	Size     string                `json:"size"` // Only exists for container workspaces (Python, R)
	Active   bool                  `json:"active"`
	Shared   bool                  `json:"shared"`
	User     string                `json:"user"`
	Host     string                `json:"host"`
	URL      string                `json:"url"`
	Password string                `json:"password"`
	Created  SandboxWorkspacesTime `json:"createdTimestamp"`
	Updated  SandboxWorkspacesTime `json:"updatedTimestamp"`
	Start    SandboxWorkspacesTime `json:"startTimestamp"`

	// SandboxWorkspace details - only exists for Snowflake workspaces
	Details *SandboxWorkspaceDetails `json:"workspaceDetails"`

	// SandboxWorkspace credentials - contains authentication information
	// Only exists for BigQuery workspaces
	Credentials *SandboxWorkspaceCredentials `json:"credentials"`
}

// GetSandboxWorkspaceInstanceRequest retrieves a workspace by its ID
// https://sandboxes.keboola.com/documentation
func (a *AuthorizedAPI) GetSandboxWorkspaceInstanceRequest(workspaceID SandboxWorkspaceID) request.APIRequest[*SandboxWorkspace] {
	result := &SandboxWorkspace{}
	req := a.newRequest(WorkspacesAPI).
		WithResult(&result).
		WithGet(WorkspacesAPISandbox).
		AndPathParam("sandboxId", workspaceID.String())
	return request.NewAPIRequest(result, req)
}

// ListSandboxWorkspaceInstancesRequest returns a list of all workspaces
// https://sandboxes.keboola.com/documentation
func (a *AuthorizedAPI) ListSandboxWorkspaceInstancesRequest() request.APIRequest[*[]*SandboxWorkspace] {
	result := make([]*SandboxWorkspace, 0)
	req := a.newRequest(WorkspacesAPI).
		WithResult(&result).
		WithGet(WorkspacesAPISandboxes)
	return request.NewAPIRequest(&result, req)
}

func (a *AuthorizedAPI) CleanSandboxWorkspaceInstances(ctx context.Context) error {
	instances, err := a.ListSandboxWorkspaceInstancesRequest().Send(ctx)
	if err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}

	for _, s := range *instances {
		wg.Go(func() {
			if e := a.DeleteSandboxWorkspaceJobRequest(s.ID).SendOrErr(ctx); e != nil {
				m.Lock()
				defer m.Unlock()
				err = multierror.Append(err, e)
			}
		})
	}

	wg.Wait()
	if err != nil {
		return err
	}

	return nil
}
