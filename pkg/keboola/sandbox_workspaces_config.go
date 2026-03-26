package keboola

import (
	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

func (a *AuthorizedAPI) GetSandboxWorkspaceConfigRequest(branchID BranchID, configID ConfigID) request.APIRequest[*Config] {
	key := ConfigKey{
		BranchID:    branchID,
		ComponentID: SandboxWorkspacesComponent,
		ID:          configID,
	}
	return a.GetConfigRequest(key)
}

func (a *AuthorizedAPI) ListSandboxWorkspaceConfigRequest(branchID BranchID) request.APIRequest[*[]*Config] {
	return a.ListConfigRequest(branchID, SandboxWorkspacesComponent)
}

func (a *AuthorizedAPI) CreateSandboxWorkspaceConfigRequest(branchID BranchID, name string) request.APIRequest[*ConfigWithRows] {
	config := &ConfigWithRows{
		Config: &Config{
			ConfigKey: ConfigKey{
				BranchID:    branchID,
				ComponentID: SandboxWorkspacesComponent,
			},
			Name: name,
		},
	}
	return a.CreateConfigRequest(config, true)
}

func (a *AuthorizedAPI) DeleteSandboxWorkspaceConfigRequest(branchID BranchID, configID ConfigID) request.APIRequest[request.NoResult] {
	req := a.DeleteConfigRequest(ConfigKey{
		BranchID:    branchID,
		ComponentID: SandboxWorkspacesComponent,
		ID:          configID,
	})
	return request.NewAPIRequest(request.NoResult{}, req)
}
