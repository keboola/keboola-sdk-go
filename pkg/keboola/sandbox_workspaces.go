package keboola

import (
	"context"
	"fmt"
	"sync"

	"github.com/hashicorp/go-multierror"
)

const SandboxWorkspacesComponent = "keboola.sandboxes"

const (
	SandboxWorkspaceSizeSmall  = "small"
	SandboxWorkspaceSizeMedium = "medium"
	SandboxWorkspaceSizeLarge  = "large"
)
const (
	SandboxWorkspaceTypeSnowflake = SandboxWorkspaceType("snowflake")
	SandboxWorkspaceTypeBigQuery  = SandboxWorkspaceType("bigquery")
	SandboxWorkspaceTypePython    = SandboxWorkspaceType("python")
	SandboxWorkspaceTypeR         = SandboxWorkspaceType("r")
)

type SandboxWorkspaceType string

type SandboxWorkspaceWithConfig struct {
	SandboxWorkspace *SandboxWorkspace
	Config           *Config
}

func (v SandboxWorkspaceWithConfig) String() string {
	if SandboxWorkspaceSupportsSizes(SandboxWorkspaceType(v.SandboxWorkspace.Type)) {
		return fmt.Sprintf("ID: %s, Type: %s, Size: %s, Name: %s", v.SandboxWorkspace.ID, v.SandboxWorkspace.Type, v.SandboxWorkspace.Size, v.Config.Name)
	} else {
		return fmt.Sprintf("ID: %s, Type: %s, Name: %s", v.SandboxWorkspace.ID, v.SandboxWorkspace.Type, v.Config.Name)
	}
}

func SandboxWorkspaceSizesOrdered() []string {
	return []string{
		SandboxWorkspaceSizeSmall,
		SandboxWorkspaceSizeMedium,
		SandboxWorkspaceSizeLarge,
	}
}

func SandboxWorkspaceSizesMap() map[string]bool {
	return map[string]bool{
		SandboxWorkspaceSizeSmall:  true,
		SandboxWorkspaceSizeMedium: true,
		SandboxWorkspaceSizeLarge:  true,
	}
}

func SandboxWorkspaceTypesOrdered() []SandboxWorkspaceType {
	return []SandboxWorkspaceType{
		SandboxWorkspaceTypeSnowflake,
		SandboxWorkspaceTypeBigQuery,
		SandboxWorkspaceTypePython,
		SandboxWorkspaceTypeR,
	}
}

func SandboxWorkspaceTypesMap() map[SandboxWorkspaceType]bool {
	return map[SandboxWorkspaceType]bool{
		SandboxWorkspaceTypeSnowflake: true,
		SandboxWorkspaceTypeBigQuery:  true,
		SandboxWorkspaceTypePython:    true,
		SandboxWorkspaceTypeR:         true,
	}
}

func SandboxWorkspaceSupportsSizes(typ SandboxWorkspaceType) bool {
	switch typ {
	case SandboxWorkspaceTypePython:
		return true
	case SandboxWorkspaceTypeR:
		return true
	default:
		return false
	}
}

func (a *AuthorizedAPI) CreateSandboxWorkspace(
	ctx context.Context,
	branchID BranchID,
	workspaceName string,
	workspaceType SandboxWorkspaceType,
	opts ...CreateSandboxWorkspaceOption,
) (*SandboxWorkspaceWithConfig, error) {
	// Create config
	emptyConfig, err := a.CreateSandboxWorkspaceConfigRequest(branchID, workspaceName).Send(ctx)
	if err != nil {
		return nil, err
	}

	// Create workspace from config
	_, err = a.CreateSandboxWorkspaceJobRequest(emptyConfig.ID, workspaceType, opts...).Send(ctx)
	if err != nil {
		return nil, err
	}

	// Get workspace
	workspace, err := a.GetSandboxWorkspace(ctx, branchID, emptyConfig.ID)
	if err != nil {
		return nil, err
	}

	return workspace, nil
}

func (a *AuthorizedAPI) DeleteSandboxWorkspace(
	ctx context.Context,
	branchID BranchID,
	configID ConfigID,
	workspaceID SandboxWorkspaceID,
) error {
	// Delete workspace (this stops the instance and deletes it)
	_, err := a.DeleteSandboxWorkspaceJobRequest(workspaceID).Send(ctx)
	if err != nil {
		return err
	}

	// Delete workspace config (so it is no longer visible in UI)
	_, err = a.DeleteSandboxWorkspaceConfigRequest(branchID, configID).Send(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthorizedAPI) GetSandboxWorkspace(
	ctx context.Context,
	branchID BranchID,
	configID ConfigID,
) (*SandboxWorkspaceWithConfig, error) {
	config, err := a.GetSandboxWorkspaceConfigRequest(branchID, configID).Send(ctx)
	if err != nil {
		return nil, err
	}

	workspaceID, err := GetSandboxWorkspaceID(config)
	if err != nil {
		return nil, err
	}

	workspace, err := a.GetSandboxWorkspaceInstanceRequest(workspaceID).Send(ctx)
	if err != nil {
		return nil, err
	}

	out := &SandboxWorkspaceWithConfig{
		SandboxWorkspace: workspace,
		Config:           config,
	}
	return out, nil
}

func (a *AuthorizedAPI) ListSandboxWorkspaces(
	ctx context.Context,
	branchID BranchID,
) ([]*SandboxWorkspaceWithConfig, error) {
	// List configs and instances in parallel
	var configs []*Config
	var instances map[string]*SandboxWorkspace
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		data, e := a.ListSandboxWorkspaceConfigRequest(branchID).Send(ctx)
		if e != nil {
			m.Lock()
			defer m.Unlock()
			err = multierror.Append(err, e)
			return
		}
		configs = *data
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		data, e := a.ListSandboxWorkspaceInstancesRequest().Send(ctx)
		if e != nil {
			m.Lock()
			defer m.Unlock()
			err = multierror.Append(err, e)
			return
		}
		m := make(map[string]*SandboxWorkspace, len(*data))
		for _, workspace := range *data {
			m[workspace.ID.String()] = workspace
		}
		instances = m
	}()

	wg.Wait()
	if err != nil {
		return nil, err
	}

	// Combine config and instance lists
	out := make([]*SandboxWorkspaceWithConfig, 0)
	for _, config := range configs {
		workspaceID, err := GetSandboxWorkspaceID(config)
		if err != nil {
			// invalid configurations are ignored
			continue
		}

		instance, found := instances[workspaceID.String()]
		if !found {
			continue
		}

		out = append(out, &SandboxWorkspaceWithConfig{
			SandboxWorkspace: instance,
			Config:           config,
		})
	}
	return out, nil
}
