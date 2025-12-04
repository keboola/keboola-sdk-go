package keboola

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

type StorageWorkspaceBackend string
type StorageWorkspaceBackendSize string
type StorageWorkspaceLoginType string
type StorageWorkspaceUseCase string

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

	StorageWorkspaceUseCaseNormal StorageWorkspaceUseCase = "normal"
	StorageWorkspaceUseCaseReader StorageWorkspaceUseCase = "reader" // Use "reader" to create reader workspace.
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
	Account           *string                      `json:"account"`
	Role              *string                      `json:"role"`
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

// StorageWorkspaceCreateRequestAsync https://keboola.docs.apiary.io/#reference/workspaces/workspaces-collection/create-workspace
// Creates a workspace asynchronously and returns a storage job that can be monitored for completion.
func (a *AuthorizedAPI) StorageWorkspaceCreateRequestAsync(branchID BranchID, storageWorkspacePayload *StorageWorkspacePayload) request.APIRequest[*StorageJob] {
	result := &StorageJob{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithPost("branch/{branchId}/workspaces").
		AndPathParam("branchId", branchID.String()).
		WithJSONBody(request.StructToMap(storageWorkspacePayload, nil)).
		AndQueryParam("async", "1")
	return request.NewAPIRequest(result, req)
}

// StorageWorkspaceCreateRequest https://keboola.docs.apiary.io/#reference/workspaces/workspaces-collection/create-workspace
func (a *AuthorizedAPI) StorageWorkspaceCreateRequest(branchID BranchID, storageWorkspacePayload *StorageWorkspacePayload) request.APIRequest[*StorageWorkspace] {
	result := &StorageWorkspace{}
	req := a.
		StorageWorkspaceCreateRequestAsync(branchID, storageWorkspacePayload).
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

// StorageWorkspacesListRequest https://keboola.docs.apiary.io/#reference/workspaces/workspaces-collection/list-workspaces
func (a *AuthorizedAPI) StorageWorkspacesListRequest(branchID BranchID) request.APIRequest[*[]*StorageWorkspace] {
	result := make([]*StorageWorkspace, 0)
	req := a.
		newRequest(StorageAPI).
		WithResult(&result).
		WithGet("branch/{branchId}/workspaces").
		AndPathParam("branchId", branchID.String())
	return request.NewAPIRequest(&result, req)
}

// StorageWorkspaceDetailRequest https://keboola.docs.apiary.io/#reference/workspaces/manage-workspace/workspace-detail
func (a *AuthorizedAPI) StorageWorkspaceDetailRequest(branchID BranchID, workspaceID uint64) request.APIRequest[*StorageWorkspace] {
	result := &StorageWorkspace{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithGet("branch/{branchId}/workspaces/{workspaceId}").
		AndPathParam("branchId", branchID.String()).
		AndPathParam("workspaceId", strconv.FormatUint(workspaceID, 10))
	return request.NewAPIRequest(result, req)
}

func (a *AuthorizedAPI) StorageWorkspaceCreateCredentialsRequest(branchID BranchID, workspaceID uint64) request.APIRequest[*StorageWorkspace] {
	result := &StorageWorkspace{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithPost("branch/{branchId}/workspaces/{workspaceId}/credentials").
		AndPathParam("branchId", branchID.String()).
		AndPathParam("workspaceId", strconv.FormatUint(workspaceID, 10))
	return request.NewAPIRequest(result, req)
}

func (a *AuthorizedAPI) StorageWorkspaceFetchCredentialsRequest(branchID BranchID, workspaceID uint64, credentialID uint64) request.APIRequest[*StorageWorkspace] {
	result := &StorageWorkspace{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithGet("branch/{branchId}/workspaces/{workspaceId}/credentials/{credentialId}").
		AndPathParam("branchId", branchID.String()).
		AndPathParam("workspaceId", strconv.FormatUint(workspaceID, 10)).
		AndPathParam("credentialId", strconv.FormatUint(credentialID, 10))
	return request.NewAPIRequest(result, req)
}

func (a *AuthorizedAPI) StorageWorkspaceDeleteCredentialsRequest(branchID BranchID, workspaceID uint64, credentialID uint64) request.APIRequest[*StorageWorkspace] {
	result := &StorageWorkspace{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithDelete("branch/{branchId}/workspaces/{workspaceId}/credentials/{credentialId}").
		AndPathParam("branchId", branchID.String()).
		AndPathParam("workspaceId", strconv.FormatUint(workspaceID, 10)).
		AndPathParam("credentialId", strconv.FormatUint(credentialID, 10))
	return request.NewAPIRequest(result, req)
}

// StorageWorkspaceDeleteRequestAsync https://keboola.docs.apiary.io/#reference/workspaces/manage-workspace/delete-workspace
// Deletes a workspace asynchronously and returns a storage job that can be monitored for completion.
func (a *AuthorizedAPI) StorageWorkspaceDeleteRequestAsync(branchID BranchID, workspaceID uint64) request.APIRequest[*StorageJob] {
	result := &StorageJob{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithDelete("branch/{branchId}/workspaces/{workspaceId}").
		AndPathParam("branchId", branchID.String()).
		AndPathParam("workspaceId", strconv.FormatUint(workspaceID, 10)).
		AndQueryParam("async", "1")
	return request.NewAPIRequest(result, req)
}

// StorageWorkspaceDeleteRequest https://keboola.docs.apiary.io/#reference/workspaces/manage-workspace/delete-workspace
func (a *AuthorizedAPI) StorageWorkspaceDeleteRequest(branchID BranchID, workspaceID uint64) request.APIRequest[request.NoResult] {
	req := a.
		StorageWorkspaceDeleteRequestAsync(branchID, workspaceID).
		WithOnSuccess(func(ctx context.Context, job *StorageJob) error {
			// Wait for storage job
			waitCtx, waitCancelFn := context.WithTimeout(ctx, a.onSuccessTimeout)
			defer waitCancelFn()
			return a.WaitForStorageJob(waitCtx, job)
		})
	// Result is NoResult since delete doesn't return workspace data.
	return request.NewAPIRequest(request.NoResult{}, req)
}

// WorkspaceLoadDataInput represents a single table mapping for workspace load-data operation.
type WorkspaceLoadDataInput struct {
	Source      string `json:"source"`              // Full table identifier (e.g., in.c-bucket.mytable)
	Destination string `json:"destination"`         // Destination table name in workspace
	Overwrite   *bool  `json:"overwrite,omitempty"` // When preserve is true, duplicate tables will be overwritten
}

// WorkspaceLoadDataPayload represents the request body for workspace load-data operation.
type WorkspaceLoadDataPayload struct {
	Input    []WorkspaceLoadDataInput `json:"input"`              // Mappings of source tables with destinations
	Preserve *bool                    `json:"preserve,omitempty"` // Keep existing tables in workspace; otherwise purge before loading
}

// StorageWorkspaceLoadDataRequest https://keboola.docs.apiary.io/#reference/workspaces/load-data/load-data
// Loads data from Storage tables into a workspace and waits for the operation to complete.
func (a *AuthorizedAPI) StorageWorkspaceLoadDataRequest(branchID BranchID, workspaceID uint64, payload *WorkspaceLoadDataPayload) request.APIRequest[*StorageJob] {
	result := &StorageJob{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithPost("branch/{branchId}/workspaces/{workspaceId}/load").
		AndPathParam("branchId", branchID.String()).
		AndPathParam("workspaceId", strconv.FormatUint(workspaceID, 10)).
		WithJSONBody(request.StructToMap(payload, nil))
	return request.NewAPIRequest(result, req)
}
