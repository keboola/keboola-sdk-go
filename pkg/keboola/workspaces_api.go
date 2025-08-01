package keboola

import (
	"context"
	"sync"

	"github.com/hashicorp/go-multierror"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

type WorkspaceID string

func (v WorkspaceID) String() string {
	return string(v)
}

type WorkspaceDetails struct {
	Connection struct {
		Database  string `json:"database"`
		Schema    string `json:"schema"`
		Warehouse string `json:"warehouse"`
	} `json:"connection"`
}

// WorkspaceCredentials contains authentication credentials for the workspace
type WorkspaceCredentials struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
	PrivateKey              string `json:"private_key"`
}

type Workspace struct {
	ID       WorkspaceID    `json:"id"`
	Type     string         `json:"type"`
	Size     string         `json:"size"` // Only exists for container workspaces (Python, R)
	Active   bool           `json:"active"`
	Shared   bool           `json:"shared"`
	User     string         `json:"user"`
	Host     string         `json:"host"`
	URL      string         `json:"url"`
	Password string         `json:"password"`
	Created  WorkspacesTime `json:"createdTimestamp"`
	Updated  WorkspacesTime `json:"updatedTimestamp"`
	Start    WorkspacesTime `json:"startTimestamp"`

	// Workspace details - only exists for Snowflake workspaces
	Details *WorkspaceDetails `json:"workspaceDetails"`

	// Workspace credentials - contains authentication information
	// Only exists for BigQuery workspaces
	Credentials *WorkspaceCredentials `json:"credentials"`
}

// GetWorkspaceInstanceRequest retrieves a workspace by its ID
// https://sandboxes.keboola.com/documentation
func (a *AuthorizedAPI) GetWorkspaceInstanceRequest(workspaceID WorkspaceID) request.APIRequest[*Workspace] {
	result := &Workspace{}
	req := a.newRequest(WorkspacesAPI).
		WithResult(&result).
		WithGet(WorkspacesAPISandbox).
		AndPathParam("sandboxId", workspaceID.String())
	return request.NewAPIRequest(result, req)
}

// ListWorkspaceInstancesRequest returns a list of all workspaces
// https://sandboxes.keboola.com/documentation
func (a *AuthorizedAPI) ListWorkspaceInstancesRequest() request.APIRequest[*[]*Workspace] {
	result := make([]*Workspace, 0)
	req := a.newRequest(WorkspacesAPI).
		WithResult(&result).
		WithGet(WorkspacesAPISandboxes)
	return request.NewAPIRequest(&result, req)
}

func (a *AuthorizedAPI) CleanWorkspaceInstances(ctx context.Context) error {
	instances, err := a.ListWorkspaceInstancesRequest().Send(ctx)
	if err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}

	for _, s := range *instances {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if e := a.DeleteWorkspaceJobRequest(s.ID).SendOrErr(ctx); e != nil {
				m.Lock()
				defer m.Unlock()
				err = multierror.Append(err, e)
			}
		}()
	}

	wg.Wait()
	if err != nil {
		return err
	}

	return nil
}
