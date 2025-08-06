package keboola

import (
	"context"
	"time"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

type StorageWorkspaceBackend string
type StorageWorkspaceBackendSize string
type StorageWorkspaceLoginType string

const (
	StorageWorkspaceBackendRedshift  StorageWorkspaceBackend = "redshift"
	StorageWorkspaceBackendSnowflake StorageWorkspaceBackend = "snowflake"
	StorageWorkspaceBackendBigQuery  StorageWorkspaceBackend = "bigquery"

	StorageWorkspaceBackendSizeSmall  StorageWorkspaceBackendSize = "small"
	StorageWorkspaceBackendSizeMedium StorageWorkspaceBackendSize = "medium"
	StorageWorkspaceBackendSizeLarge  StorageWorkspaceBackendSize = "large"

	StorageWorkspaceLoginTypeDefault                 StorageWorkspaceLoginType = "default"
	StorageWorkspaceLoginTypeSnowflakePersonSSO      StorageWorkspaceLoginType = "snowflake-person-sso"
	StorageWorkspaceLoginTypeSnowflakePersonKeypair  StorageWorkspaceLoginType = "snowflake-person-keypair"
	StorageWorkspaceLoginTypeSnowflakeServiceKeypair StorageWorkspaceLoginType = "snowflake-service-keypair"
)

type StorageWorkspaceDetails struct {
	Connection struct {
		Host        *string                      `json:"host"`
		Region      *string                      `json:"region"`
		Database    *string                      `json:"database"`
		Schema      *string                      `json:"schema"`
		Warehouse   *string                      `json:"warehouse"`
		LoginType   *string                      `json:"loginType"`
		Credentials *SandboxWorkspaceCredentials `json:"credentials"`
	} `json:"connection"`
}

// SandboxWorkspaceCredentials contains authentication credentials for the workspace.
type StorageWorkspaceCredentials struct {
	Type                    string `json:"type"`                        // nolint: tagliatelle
	ProjectID               string `json:"project_id"`                  // nolint: tagliatelle
	PrivateKeyID            string `json:"private_key_id"`              // nolint: tagliatelle
	PrivateKey              string `json:"private_key"`                 // nolint: tagliatelle
	ClientEmail             string `json:"client_email"`                // nolint: tagliatelle
	ClientID                string `json:"client_id"`                   // nolint: tagliatelle
	AuthURI                 string `json:"auth_uri"`                    // nolint: tagliatelle
	TokenURI                string `json:"token_uri"`                   // nolint: tagliatelle
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"` // nolint: tagliatelle
	ClientX509CertURL       string `json:"client_x509_cert_url"`        // nolint: tagliatelle
}

type StorageWorkspace struct {
	ID                    string                      `json:"id" readonly:"true"`
	Backend               StorageWorkspaceBackend     `json:"backend"`
	BackendSize           StorageWorkspaceBackendSize `json:"backendSize"`
	NetworkPolicy         string                      `json:"networkPolicy"`
	ReadOnlyStorageAccess bool                        `json:"readOnlyStorageAccess"`
	LoginType             StorageWorkspaceLoginType   `json:"loginType"`
	PublicKey             string                      `json:"publicKey"`
	Name                  string                      `json:"name" readonly:"true"`
	Type                  string                      `json:"type" readonly:"true"`
	Created               time.Time                   `json:"created" readonly:"true"`
	Details               StorageWorkspaceDetails     `json:"details" readonly:"true"`
}

// StorageWorkspaceCreateRequest https://keboola.docs.apiary.io/#reference/workspaces/workspaces-collection/create-workspace
func (a *AuthorizedAPI) StorageWorkspaceCreateRequest(workspaceID string, storageWorkspace *StorageWorkspace) request.APIRequest[*StorageWorkspace] {
	req := a.
		newRequest(StorageAPI).
		WithResult(storageWorkspace).
		WithPost("workspaces").
		WithJSONBody(request.StructToMap(storageWorkspace, nil)).
		WithOnError(ignoreResourceAlreadyExistsError(func(ctx context.Context) error {
			if result, err := a.StorageWorkspaceDetailRequest(storageWorkspace.ID).Send(ctx); err == nil {
				*storageWorkspace = *result
				return nil
			} else {
				return err
			}
		}))
	return request.NewAPIRequest(storageWorkspace, req)
}

// StorageWorkspacesListRequest https://keboola.docs.apiary.io/#reference/workspaces/workspaces-collection/list-workspaces
func (a *AuthorizedAPI) StorageWorkspacesListRequest() request.APIRequest[*[]*StorageWorkspace] {
	result := make([]*StorageWorkspace, 0)
	req := a.
		newRequest(StorageAPI).
		WithResult(&result).
		WithGet("workspaces")
	return request.NewAPIRequest(&result, req)
}

// StorageWorkspaceDetailRequest https://keboola.docs.apiary.io/#reference/workspaces/manage-workspace/workspace-detail
func (a *AuthorizedAPI) StorageWorkspaceDetailRequest(workspaceID string) request.APIRequest[*StorageWorkspace] {
	result := &StorageWorkspace{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithGet("workspaces/{workspaceId}").
		AndPathParam("workspaceId", workspaceID)
	return request.NewAPIRequest(result, req)
}

// StorageWorkspaceDeleteRequest https://keboola.docs.apiary.io/#reference/workspaces/manage-workspace/delete-workspace
func (a *AuthorizedAPI) StorageWorkspaceDeleteRequest(workspaceID string) request.APIRequest[*StorageWorkspace] {
	result := &StorageWorkspace{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithDelete("workspaces/{workspaceId}").
		AndPathParam("workspaceId", workspaceID)
	return request.NewAPIRequest(result, req)
}
