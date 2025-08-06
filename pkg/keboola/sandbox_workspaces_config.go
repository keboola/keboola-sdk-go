package keboola

import (
	"fmt"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

func GetSandboxWorkspaceID(c *Config) (SandboxWorkspaceID, error) {
	id, found, err := c.Content.GetNested("parameters.id")
	if err != nil {
		return "", err
	}
	if !found {
		return "", fmt.Errorf("config is missing parameters.id")
	}

	out, ok := id.(string)
	if !ok {
		return "", fmt.Errorf("config.parameters.id is not a string")
	}

	return SandboxWorkspaceID(out), nil
}

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
