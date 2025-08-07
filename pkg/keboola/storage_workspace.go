package keboola

import (
	"context"
	"strconv"

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
	Backend           StorageWorkspaceBackend      `json:"backend"`
	Host              *string                      `json:"host"`
	Region            *string                      `json:"region"`
	Database          *string                      `json:"database"`
	Schema            *string                      `json:"schema"`
	Warehouse         *string                      `json:"warehouse"`
	User              *string                      `json:"user"`
	LoginType         *string                      `json:"loginType"`
	PrivateKey        *string                      `json:"privateKey"`
	SSOLoginAvailable *bool                        `json:"ssoLoginAvailable,omitempty"`
	Credentials       *StorageWorkspaceCredentials `json:"credentials,omitempty"`
}

// StorageWorkspaceCredentials contains authentication credentials for the workspace.
type StorageWorkspaceCredentials struct {
	ID                      uint64 `json:"id"`
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

type StorageWorkspacePayload struct {
	Backend               StorageWorkspaceBackend      `json:"backend"`
	BackendSize           *StorageWorkspaceBackendSize `json:"backendSize,omitempty" writeoptional:"true"`
	NetworkPolicy         *string                      `json:"networkPolicy,omitempty" writeoptional:"true"`
	ReadOnlyStorageAccess bool                         `json:"readOnlyStorageAccess"`
	LoginType             StorageWorkspaceLoginType    `json:"loginType"`
	PublicKey             *string                      `json:"publicKey,omitempty" writeoptional:"true"`
}

type StorageWorkspace struct {
	ID                      uint64                       `json:"id"`
	BackendSize             *StorageWorkspaceBackendSize `json:"backendSize"`
	Name                    string                       `json:"name"`
	Type                    string                       `json:"type"`
	Created                 StorageWorkspaceTime         `json:"created"`
	StorageWorkspaceDetails StorageWorkspaceDetails      `json:"connection"`
}

// StorageWorkspaceCreateRequest https://keboola.docs.apiary.io/#reference/workspaces/workspaces-collection/create-workspace
func (a *AuthorizedAPI) StorageWorkspaceCreateRequest(storageWorkspacePayload *StorageWorkspacePayload) request.APIRequest[*StorageWorkspace] {
	result := StorageWorkspace{}
	req := a.
		newRequest(StorageAPI).
		WithResult(&result).
		WithPost("workspaces").
		WithJSONBody(request.StructToMap(storageWorkspacePayload, nil)).
		WithOnError(ignoreResourceAlreadyExistsError(func(ctx context.Context) error {
			if workspaceResult, err := a.StorageWorkspaceDetailRequest(result.ID).Send(ctx); err == nil {
				result = *workspaceResult
				return nil
			} else {
				return err
			}
		}))
	return request.NewAPIRequest(&result, req)
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
func (a *AuthorizedAPI) StorageWorkspaceDetailRequest(workspaceID uint64) request.APIRequest[*StorageWorkspace] {
	result := &StorageWorkspace{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithGet("workspaces/{workspaceId}").
		AndPathParam("workspaceId", strconv.FormatUint(workspaceID, 10))
	return request.NewAPIRequest(result, req)
}

func (a *AuthorizedAPI) StorageWorkspaceCreateCredentialsRequest(workspaceID uint64) request.APIRequest[*StorageWorkspace] {
	result := &StorageWorkspace{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithPost("workspaces/{workspaceId}/credentials").
		AndPathParam("workspaceId", strconv.FormatUint(workspaceID, 10))
	return request.NewAPIRequest(result, req)
}

func (a *AuthorizedAPI) StorageWorkspaceFetchCredentialsRequest(workspaceID uint64, credentialID uint64) request.APIRequest[*StorageWorkspaceDetails] {
	result := &StorageWorkspaceDetails{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithGet("workspaces/{workspaceId}/credentials/{credentialId}").
		AndPathParam("workspaceId", strconv.FormatUint(workspaceID, 10)).
		AndPathParam("credentialId", strconv.FormatUint(credentialID, 10))
	return request.NewAPIRequest(result, req)
}

func (a *AuthorizedAPI) StorageWorkspaceDeleteCredentialsRequest(workspaceID uint64, credentialID uint64) request.APIRequest[*StorageWorkspace] {
	result := &StorageWorkspace{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithDelete("workspaces/{workspaceId}/credentials/{credentialId}").
		AndPathParam("workspaceId", strconv.FormatUint(workspaceID, 10)).
		AndPathParam("credentialId", strconv.FormatUint(credentialID, 10))
	return request.NewAPIRequest(result, req)
}

// StorageWorkspaceDeleteRequest https://keboola.docs.apiary.io/#reference/workspaces/manage-workspace/delete-workspace
func (a *AuthorizedAPI) StorageWorkspaceDeleteRequest(workspaceID uint64) request.APIRequest[*StorageWorkspace] {
	result := &StorageWorkspace{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithDelete("workspaces/{workspaceId}").
		AndPathParam("workspaceId", strconv.FormatUint(workspaceID, 10))
	return request.NewAPIRequest(result, req)
}
