package keboola

import (
	"context"
	"slices"
	"sort"

	"github.com/keboola/go-utils/pkg/orderedmap"
	"github.com/relvacode/iso8601"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

// ConfigID represents an ID of a configuration in Storage API.
type ConfigID string

func (v ConfigID) String() string {
	return string(v)
}

// ConfigKey is a unique identifier of a configuration.
type ConfigKey struct {
	BranchID    BranchID    `json:"branchId"`
	ComponentID ComponentID `json:"componentId"`
	ID          ConfigID    `json:"id" writeas:"configurationId" writeoptional:"true"`
}

func (k ConfigKey) ObjectID() any {
	return k.ID
}

// Config https://keboola.docs.apiary.io/#reference/components-and-configurations/component-configurations/list-configurations
type Config struct {
	ConfigKey
	Name              string                 `json:"name"`
	Description       string                 `json:"description"`
	ChangeDescription string                 `json:"changeDescription"`
	IsDeleted         bool                   `json:"isDeleted" readonly:"true"`
	Created           iso8601.Time           `json:"created" readonly:"true"`
	Version           int                    `json:"version" readonly:"true"`
	State             *orderedmap.OrderedMap `json:"state" readonly:"true"`
	IsDisabled        bool                   `json:"isDisabled"`
	Content           *orderedmap.OrderedMap `json:"configuration"`
	RowsSortOrder     []string               `json:"rowsSortOrder,omitempty"`
}

// ConfigWithRows is a configuration with its configuration rows.
type ConfigWithRows struct {
	*Config
	Rows []*ConfigRow `json:"rows"`
}

// SortRows by name.
func (c *ConfigWithRows) SortRows() {
	sort.SliceStable(c.Rows, func(i, j int) bool {
		return c.Rows[i].Name < c.Rows[j].Name
	})
}

// ConfigMetadataItem is one item of configuration metadata.
type ConfigMetadataItem struct {
	BranchID    BranchID
	ComponentID ComponentID     `json:"idComponent"`
	ConfigID    ConfigID        `json:"configurationId"`
	Metadata    MetadataDetails `json:"metadata"`
}

// ConfigsMetadata slice.
type ConfigsMetadata []*ConfigMetadataItem

// ToMap converts slice to map.
func (v ConfigsMetadata) ToMap() map[ConfigKey]Metadata {
	out := make(map[ConfigKey]Metadata)
	for _, item := range v {
		key := ConfigKey{BranchID: item.BranchID, ComponentID: item.ComponentID, ID: item.ConfigID}
		out[key] = item.Metadata.ToMap()
	}
	return out
}

// ListConfigsAndRowsFrom https://keboola.docs.apiary.io/#reference/components-and-configurations/get-components/get-components
func (a *AuthorizedAPI) ListConfigsAndRowsFrom(branch BranchKey) request.APIRequest[*[]*ComponentWithConfigs] {
	result := make([]*ComponentWithConfigs, 0)
	req := a.
		newRequest(StorageAPI).
		WithResult(&result).
		WithGet("branch/{branchId}/components").
		AndPathParam("branchId", branch.ID.String()).
		AndQueryParam("include", "configuration,rows").
		WithOnSuccess(func(_ context.Context, _ request.HTTPResponse) error {
			// Add missing values
			for _, component := range result {
				component.BranchID = branch.ID

				// Set config IDs
				for _, config := range component.Configs {
					config.BranchID = branch.ID
					config.ComponentID = component.ID
					config.SortRows()

					// Set rows IDs
					for _, row := range config.Rows {
						row.BranchID = branch.ID
						row.ComponentID = component.ID
						row.ConfigID = config.ID
					}
				}
			}
			return nil
		})
	return request.NewAPIRequest(&result, req)
}

func (a *AuthorizedAPI) ListConfigRequest(branchID BranchID, componentID ComponentID) request.APIRequest[*[]*Config] {
	result := make([]*Config, 0)
	req := a.newRequest(StorageAPI).
		WithResult(&result).
		WithGet("branch/{branchId}/components/{componentId}/configs").
		AndPathParam("branchId", branchID.String()).
		AndPathParam("componentId", componentID.String()).
		WithOnSuccess(func(ctx context.Context, response request.HTTPResponse) error {
			for _, c := range result {
				c.BranchID = branchID
				c.ComponentID = componentID
			}
			return nil
		})
	return request.NewAPIRequest(&result, req)
}

// GetConfigRequest https://keboola.docs.apiary.io/#reference/components-and-configurations/manage-configurations/development-branch-configuration-detail
func (a *AuthorizedAPI) GetConfigRequest(key ConfigKey) request.APIRequest[*Config] {
	result := &Config{}
	result.BranchID = key.BranchID
	result.ComponentID = key.ComponentID
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithGet("branch/{branchId}/components/{componentId}/configs/{configId}").
		AndPathParam("branchId", key.BranchID.String()).
		AndPathParam("componentId", key.ComponentID.String()).
		AndPathParam("configId", key.ID.String())
	return request.NewAPIRequest(result, req)
}

// CreateConfigRequest https://keboola.docs.apiary.io/#reference/components-and-configurations/component-configurations/create-development-branch-configuration
func (a *AuthorizedAPI) CreateConfigRequest(config *ConfigWithRows, withAsyncRows bool) request.APIRequest[*ConfigWithRows] {
	// Create config
	req := a.
		newRequest(StorageAPI).
		WithResult(config).
		WithPost("branch/{branchId}/components/{componentId}/configs").
		AndPathParam("branchId", config.BranchID.String()).
		AndPathParam("componentId", string(config.ComponentID)).
		WithJSONBody(request.StructToMap(config.Config, nil)).
		WithOnError(ignoreResourceAlreadyExistsError(func(ctx context.Context) error {
			if result, err := a.GetConfigRequest(config.ConfigKey).Send(ctx); err == nil {
				*config.Config = *result
				return nil
			} else {
				return err
			}
		})).
		// Create config rows
		WithOnSuccess(func(ctx context.Context, _ request.HTTPResponse) error {
			wg := request.NewWaitGroup(ctx)
			for _, row := range config.Rows {
				row.BranchID = config.BranchID
				row.ComponentID = config.ComponentID
				row.ConfigID = config.ID
				if withAsyncRows {
					wg.Send(a.CreateConfigRowRequest(row))
					continue
				}

				_, err := a.CreateConfigRowRequest(row).Send(ctx)
				if err != nil {
					return err
				}
			}

			return wg.Wait()
		})
	return request.NewAPIRequest(config, req)
}

// UpdateConfigRequest https://keboola.docs.apiary.io/#reference/components-and-configurations/manage-configurations/update-configuration
func (a *AuthorizedAPI) UpdateConfigRequest(config *ConfigWithRows, changedFields []string) request.APIRequest[*ConfigWithRows] {
	// ID is required
	if config.ID == "" {
		panic("config id must be set")
	}

	// Update config
	tmpConfig := &ConfigWithRows{}
	req := a.
		newRequest(StorageAPI).
		WithResult(tmpConfig).
		WithPut("branch/{branchId}/components/{componentId}/configs/{configId}").
		AndPathParam("branchId", config.BranchID.String()).
		AndPathParam("componentId", string(config.ComponentID)).
		AndPathParam("configId", string(config.ID)).
		WithJSONBody(request.StructToMap(config, changedFields)).
		// Update config rows
		WithOnSuccess(func(ctx context.Context, resp request.HTTPResponse) error {
			// Update config fields from tmpConfig, preserving Rows
			tmpConfig.BranchID = config.BranchID
			tmpConfig.ComponentID = config.ComponentID
			tmpConfig.ID = config.ID
			*config.Config = *tmpConfig.Config

			// Synchronize rows
			return a.synchronizeConfigRows(ctx, config, changedFields)
		})

	return request.NewAPIRequest(config, req)
}

// DeleteConfigRequest https://keboola.docs.apiary.io/#reference/components-and-configurations/manage-configurations/delete-configuration
func (a *AuthorizedAPI) DeleteConfigRequest(config ConfigKey) request.APIRequest[request.NoResult] {
	req := a.
		newRequest(StorageAPI).
		WithDelete("branch/{branchId}/components/{componentId}/configs/{configId}").
		AndPathParam("branchId", config.BranchID.String()).
		AndPathParam("componentId", string(config.ComponentID)).
		AndPathParam("configId", string(config.ID)).
		WithOnError(ignoreResourceNotFoundError())
	return request.NewAPIRequest(request.NoResult{}, req)
}

// DeleteConfigsInBranchRequest lists all configs in branch and deletes them all.
func (a *AuthorizedAPI) DeleteConfigsInBranchRequest(branch BranchKey) request.APIRequest[request.NoResult] {
	req := a.
		ListConfigsAndRowsFrom(branch).
		WithOnSuccess(func(ctx context.Context, result *[]*ComponentWithConfigs) error {
			wg := request.NewWaitGroup(ctx)
			for _, component := range *result {
				for _, config := range component.Configs {
					config := config
					wg.Send(a.DeleteConfigRequest(config.ConfigKey))
				}
			}
			return wg.Wait()
		})
	return request.NewAPIRequest(request.NoResult{}, req)
}

// ListConfigMetadataRequest https://keboola.docs.apiary.io/#reference/search/search-components-configurations/search-component-configurations
func (a *AuthorizedAPI) ListConfigMetadataRequest(branchID BranchID) request.APIRequest[*ConfigsMetadata] {
	result := make(ConfigsMetadata, 0)
	req := a.
		newRequest(StorageAPI).
		WithResult(&result).
		WithGet("branch/{branchId}/search/component-configurations").
		AndPathParam("branchId", branchID.String()).
		AndQueryParam("include", "filteredMetadata").
		WithOnSuccess(func(_ context.Context, _ request.HTTPResponse) error {
			for _, item := range result {
				item.BranchID = branchID
			}
			return nil
		})
	return request.NewAPIRequest(&result, req)
}

// AppendConfigMetadataRequest https://keboola.docs.apiary.io/#reference/metadata/components-configurations-metadata/create-or-update
func (a *AuthorizedAPI) AppendConfigMetadataRequest(key ConfigKey, metadata Metadata) request.APIRequest[request.NoResult] {
	// Empty, we have nothing to append
	if len(metadata) == 0 {
		return request.NewNoOperationAPIRequest(request.NoResult{})
	}

	req := a.
		newRequest(StorageAPI).
		WithPost("branch/{branchId}/components/{componentId}/configs/{configId}/metadata").
		AndPathParam("branchId", key.BranchID.String()).
		AndPathParam("componentId", string(key.ComponentID)).
		AndPathParam("configId", string(key.ID)).
		WithJSONBody(metadata.ToPayload())
	return request.NewAPIRequest(request.NoResult{}, req)
}

// DeleteConfigMetadataRequest https://keboola.docs.apiary.io/#reference/metadata/components-configurations-metadata/delete
func (a *AuthorizedAPI) DeleteConfigMetadataRequest(key ConfigKey, metaID string) request.APIRequest[request.NoResult] {
	req := a.
		newRequest(StorageAPI).
		WithDelete("branch/{branchId}/components/{componentId}/configs/{configId}/metadata/{metadataId}").
		AndPathParam("branchId", key.BranchID.String()).
		AndPathParam("componentId", string(key.ComponentID)).
		AndPathParam("configId", string(key.ID)).
		AndPathParam("metadataId", metaID).
		WithOnError(ignoreResourceNotFoundError())
	return request.NewAPIRequest(request.NoResult{}, req)
}

// synchronizeConfigRows synchronizes configuration rows with the current state.
// It creates new rows, updates existing ones and deletes rows not present in the config.
// If onlyCreate is true, it will only create new rows (used for CreateConfigRequest).
func (a *AuthorizedAPI) synchronizeConfigRows(
	ctx context.Context,
	config *ConfigWithRows,
	changedFields []string,
) error {
	// Make sure all rows have correct parent identifiers
	for _, row := range config.Rows {
		row.BranchID = config.BranchID
		row.ComponentID = config.ComponentID
		row.ConfigID = config.ID
	}

	// For update, we need to synchronize rows with current state
	key := ConfigRowKey{BranchID: config.BranchID, ComponentID: config.ComponentID, ConfigID: config.ID}
	rows, err := a.ListConfigRowRequest(key).Send(ctx)
	if err != nil {
		return err
	}

	// Delete rows that aren't in the current config.Rows
	for _, row := range *rows {
		if slices.ContainsFunc(config.Rows, func(r *ConfigRow) bool { return r.ID == row.ID }) {
			continue
		}

		deleteKey := ConfigRowKey{
			BranchID:    config.BranchID,
			ComponentID: config.ComponentID,
			ConfigID:    config.ID,
			ID:          row.ID,
		}
		_, err := a.DeleteConfigRowRequest(deleteKey).Send(ctx)
		if err != nil {
			return err
		}
	}

	// Create new rows or update existing ones
	for _, row := range config.Rows {
		if row.ID == "" {
			// Create new row
			resp, err := a.CreateConfigRowRequest(row).Send(ctx)
			if err != nil {
				return err
			}

			*row = *resp
			continue
		}

		// Update existing row
		resp, err := a.UpdateConfigRowRequest(row, changedFields).Send(ctx)
		if err != nil {
			return err
		}

		*row = *resp
	}

	return nil
}
