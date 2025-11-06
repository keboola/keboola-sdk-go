package keboola_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/keboola/go-utils/pkg/orderedmap"
	"github.com/keboola/go-utils/pkg/wildcards"
	"github.com/stretchr/testify/assert"

	. "github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

func TestConfigApiCalls(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := APIClientForAnEmptyProject(t, ctx)

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, branch)

	// List - no component/config
	components, err := api.ListConfigsAndRowsFrom(branch.BranchKey).Send(ctx)
	assert.NoError(t, err)
	assert.Empty(t, components)

	// Create config with rows
	row1 := &ConfigRow{
		Name:              "Row1",
		Description:       "Row1 description",
		ChangeDescription: "Row1 test",
		IsDisabled:        false,
		Content: orderedmap.FromPairs([]orderedmap.Pair{
			{Key: "row1", Value: "value1"},
		}),
	}
	row2 := &ConfigRow{
		Name:              "Row2",
		Description:       "Row2 description",
		ChangeDescription: "Row2 test",
		IsDisabled:        true,
		Content: orderedmap.FromPairs([]orderedmap.Pair{
			{Key: "row2", Value: "value2"},
		}),
	}
	config := &ConfigWithRows{
		Config: &Config{
			ConfigKey: ConfigKey{
				BranchID:    branch.ID,
				ComponentID: "ex-generic-v2",
			},
			Name:              "Test",
			Description:       "Test description",
			ChangeDescription: "My test",
			Content: orderedmap.FromPairs([]orderedmap.Pair{
				{
					Key: "foo",
					Value: orderedmap.FromPairs([]orderedmap.Pair{
						{Key: "bar", Value: "baz"},
					}),
				},
			}),
			RowsSortOrder: []string{},
		},
		Rows: []*ConfigRow{row1, row2},
	}
	resConfig, err := api.CreateConfigRequest(config, true).Send(ctx)
	assert.NoError(t, err)
	assert.Same(t, config, resConfig)
	assert.NotEmpty(t, config.ID)
	assert.Equal(t, config.ID, row1.ConfigID)
	assert.Equal(t, ComponentID("ex-generic-v2"), row1.ComponentID)
	assert.Equal(t, branch.ID, row1.BranchID)
	assert.Equal(t, config.ID, row2.ConfigID)
	assert.Equal(t, ComponentID("ex-generic-v2"), row2.ComponentID)
	assert.Equal(t, branch.ID, row2.BranchID)

	// Get config
	resultConfig, err := api.GetConfigRequest(config.ConfigKey).Send(ctx)
	assert.NoError(t, err)

	// Change description and version differs, because rows have been created after the configuration has been created.
	config.ChangeDescription = resultConfig.ChangeDescription
	config.Version = resultConfig.Version
	assert.Equal(t, config.Config, resultConfig)

	// List configs (should contain 1)
	configList, err := api.ListConfigRequest(config.BranchID, config.ComponentID).Send(ctx)
	assert.NoError(t, err)
	assert.Len(t, *configList, 1)
	assert.Equal(t, config.Config, (*configList)[0])

	// Create a new row (row3) and add it to the existing configuration
	row3 := &ConfigRow{
		Name:              "Row3",
		Description:       "Row3 description",
		ChangeDescription: "Row3 test",
		IsDisabled:        false,
		Content: orderedmap.FromPairs([]orderedmap.Pair{
			{Key: "row3", Value: "value3"},
		}),
		ConfigRowKey: ConfigRowKey{
			BranchID:    config.BranchID,
			ComponentID: config.ComponentID,
			ConfigID:    config.ID,
		},
	}

	config.Rows = append(config.Rows, row3)

	// Update config
	config.Name = "Test modified +++úěš!@#"
	config.Description = "Test description modified"
	config.ChangeDescription = "updated"
	config.Content = orderedmap.FromPairs([]orderedmap.Pair{
		{
			Key: "foo",
			Value: orderedmap.FromPairs([]orderedmap.Pair{
				{Key: "bar", Value: "modified"},
			}),
		},
	})
	resConfig, err = api.UpdateConfigRequest(config, []string{"name", "description", "changeDescription", "configuration", "rows"}).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, *config.Config, *resConfig.Config)

	// List components
	components, err = api.ListConfigsAndRowsFrom(branch.BranchKey).Send(ctx)
	assert.NotEmpty(t, components)
	assert.NoError(t, err)
	componentsJSON, err := json.MarshalIndent(components, "", "  ")
	assert.NoError(t, err)
	wildcards.Assert(t, expectedComponentsConfigTest(), string(componentsJSON), "Unexpected components")

	// Test configuration change without rows - verify rows are preserved
	configOnly := &ConfigWithRows{
		Config: &Config{
			ConfigKey:         config.ConfigKey,
			Name:              "Test config only update",
			Description:       "Test description config only",
			ChangeDescription: "config only update",
			Content: orderedmap.FromPairs([]orderedmap.Pair{
				{
					Key: "config_only",
					Value: orderedmap.FromPairs([]orderedmap.Pair{
						{Key: "test", Value: "config_only_value"},
					}),
				},
			}),
			RowsSortOrder: []string{},
		},
	}

	// Update only the configuration (without rows)
	resConfigOnly, err := api.UpdateConfigRequest(configOnly, []string{"name", "description", "changeDescription", "configuration"}).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, *configOnly.Config, *resConfigOnly.Config)

	// List components and verify rows are still there
	componentsAfterConfigOnly, err := api.ListConfigsAndRowsFrom(branch.BranchKey).Send(ctx)
	assert.NotEmpty(t, componentsAfterConfigOnly)
	assert.NoError(t, err)
	componentsAfterConfigOnlyJSON, err := json.MarshalIndent(componentsAfterConfigOnly, "", "  ")
	assert.NoError(t, err)
	wildcards.Assert(t, expectedComponentsAfterConfigOnlyUpdate(), string(componentsAfterConfigOnlyJSON), "Unexpected components after config-only update")

	// Update metadata
	metadata := map[string]string{"KBC.KaC.meta1": "value"}
	_, err = api.AppendConfigMetadataRequest(config.ConfigKey, metadata).Send(ctx)
	assert.NoError(t, err)

	// List metadata
	configsMetadata, err := api.ListConfigMetadataRequest(branch.ID).Send(ctx)
	assert.NoError(t, err)
	assert.Equal(t, map[ConfigKey]Metadata{
		config.ConfigKey: map[string]string{"KBC.KaC.meta1": "value"},
	}, configsMetadata.ToMap())

	// Delete metadata
	_, err = api.DeleteConfigMetadataRequest(config.ConfigKey, (*configsMetadata)[0].Metadata[0].ID).Send(ctx)
	assert.NoError(t, err)

	// Check that metadata is deleted
	configsMetadata, err = api.ListConfigMetadataRequest(branch.ID).Send(ctx)
	assert.NoError(t, err)
	assert.Empty(t, configsMetadata)

	// Delete configuration
	_, err = api.DeleteConfigRequest(config.ConfigKey).Send(ctx)
	assert.NoError(t, err)

	// List components - no component
	components, err = api.ListConfigsAndRowsFrom(branch.BranchKey).Send(ctx)
	assert.NoError(t, err)
	assert.Empty(t, components)
}

func expectedComponentsConfigTest() string {
	return `[
  {
    "branchId": %s,
    "id": "ex-generic-v2",
    "type": "extractor",
    "name": "Generic",
    %A,
    "configurations": [
      {
        "branchId": %s,
        "componentId": "ex-generic-v2",
        "id": "%s",
        "name": "Test modified +++úěš!@#",
        "description": "Test description modified",
        "changeDescription": "Row%d test",
        "isDeleted": false,
        "created": "%s",
        "version": 7,
        "state": null,
        "isDisabled": false,
        "configuration": {
          "foo": {
            "bar": "modified"
          }
        },
        "rows": [
          {
            "id": "%s",
            "name": "Row1",
            "description": "Row1 description",
            "changeDescription": "Row1 test",
            "isDisabled": false,
            "version": 2,
            "state": null,
            "configuration": {
              "row1": "value1"
            }
          },
          {
            "id": "%s",
            "name": "Row2",
            "description": "Row2 description",
            "changeDescription": "Row2 test",
            "isDisabled": true,
            "version": 2,
            "state": null,
            "configuration": {
              "row2": "value2"
            }
          },
          {
            "id": "%s",
            "name": "Row3",
            "description": "Row3 description",
            "changeDescription": "Row3 test",
            "isDisabled": false,
            "version": 1,
            "state": null,
            "configuration": {
              "row3": "value3"
            }
          }
        ]
      }
    ]
  }
]
`
}

func expectedComponentsAfterConfigOnlyUpdate() string {
	return `[
  {
    "branchId": %s,
    "id": "ex-generic-v2",
    "type": "extractor",
    "name": "Generic",
    %A,
    "configurations": [
      {
        "branchId": %s,
        "componentId": "ex-generic-v2",
        "id": "%s",
        "name": "Test config only update",
        "description": "Test description config only",
        "changeDescription": "config only update",
        "isDeleted": false,
        "created": "%s",
        "version": 8,
        "state": null,
        "isDisabled": false,
        "configuration": {
          "config_only": {
            "test": "config_only_value"
          }
        },
        "rows": [
          {
            "id": "%s",
            "name": "Row1",
            "description": "Row1 description",
            "changeDescription": "Row1 test",
            "isDisabled": false,
            "version": 2,
            "state": null,
            "configuration": {
              "row1": "value1"
            }
          },
          {
            "id": "%s",
            "name": "Row2",
            "description": "Row2 description",
            "changeDescription": "Row2 test",
            "isDisabled": true,
            "version": 2,
            "state": null,
            "configuration": {
              "row2": "value2"
            }
          },
          {
            "id": "%s",
            "name": "Row3",
            "description": "Row3 description",
            "changeDescription": "Row3 test",
            "isDisabled": false,
            "version": 1,
            "state": null,
            "configuration": {
              "row3": "value3"
            }
          }
        ]
      }
    ]
  }
]
`
}

// TestConfigWithArrayOfStringConfiguration tests real-world case where API returns
// configuration as array containing JSON string: ["{\...}"]
// This was seen in keboola.variables component configurations.
func TestConfigWithArrayOfStringConfiguration(t *testing.T) {
	t.Parallel()

	// Real example from keboola.variables component
	configJSON := `{
	"id": "01k4yhwsmgfdh1pk4r8kv09ske",
	"name": "Test configuration",
	"description": "",
	"created": "2025-09-12T10:50:26+0200",
	"creatorToken": {
		"id": 165618,
		"description": "queue test"
	},
	"version": 3,
	"changeDescription": "Row row2 added",
	"isDisabled": false,
	"isDeleted": false,
	"configuration": {
	"key1":["val1","val2"]
},
	"rowsSortOrder": ["1"],
	"currentVersion": {
		"created": "2025-09-12T10:50:26+0200",
		"creatorToken": {
			"id": 165618,
			"description": "queue test"
		},
		"changeDescription": "Row row2 added",
		"versionIdentifier": "01K4YHWSWRK0T35KZFYF4QWR5R"
	}
}`

	var config Config
	err := json.Unmarshal([]byte(configJSON), &config)
	assert.NoError(t, err)
}
