package keboola_test

import (
	"context"
	"testing"

	"github.com/keboola/go-utils/pkg/testproject"
	"github.com/stretchr/testify/assert"

	"github.com/keboola/keboola-sdk-go/v2/pkg/client"
	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

func TestVariablesApiCalls(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	project, _ := testproject.GetTestProjectForTest(t)
	c := client.NewTestClient()
	api, err := keboola.NewAuthorizedAPI(ctx, project.StorageAPIHost(), project.StorageAPIToken(), keboola.WithClient(&c))
	assert.NoError(t, err)

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, branch)

	createPayload := &keboola.VaultVariableCreatePayload{
		Key:   "test_variable",
		Value: "test_value",
		Attributes: map[string]interface{}{
			"branchId": branch.ID.String(),
		},
	}
	variable, err := api.CreateVariableRequest(createPayload).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, variable)
	assert.Equal(t, "test_variable", variable.Key)
	assert.Equal(t, "test_value", variable.Value)
	assert.NotEmpty(t, variable.Hash)

	listOpts := &keboola.VaultVariableListOptions{
		Attributes: map[string]interface{}{
			"branchId": branch.ID.String(),
		},
	}
	variables, err := api.ListVariablesRequest(listOpts).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, variables)
	assert.GreaterOrEqual(t, len(*variables), 1)

	found := false
	for _, v := range *variables {
		if v.Hash == variable.Hash {
			found = true
			assert.Equal(t, "test_variable", v.Key)
			assert.Equal(t, "test_value", v.Value)
			break
		}
	}
	assert.True(t, found, "Created variable should be in the list")

	scopedVariables, err := api.ListVariablesScopedRequest(branch.ID, nil).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, scopedVariables)
	assert.GreaterOrEqual(t, len(*scopedVariables), 1)

	found = false
	for _, v := range *scopedVariables {
		if v.Hash == variable.Hash {
			found = true
			assert.Equal(t, "test_variable", v.Key)
			assert.Equal(t, "test_value", v.Value)
			break
		}
	}
	assert.True(t, found, "Created variable should be in the scoped list")

	err = api.DeleteVariableRequest(variable.Hash).SendOrErr(ctx)
	assert.NoError(t, err)

	variables, err = api.ListVariablesRequest(listOpts).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, variables)

	found = false
	for _, v := range *variables {
		if v.Hash == variable.Hash {
			found = true
			break
		}
	}
	assert.False(t, found, "Deleted variable should not be in the list")
}

func TestVariablesWithEncryption(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	project, _ := testproject.GetTestProjectForTest(t)
	c := client.NewTestClient()
	api, err := keboola.NewAuthorizedAPI(ctx, project.StorageAPIHost(), project.StorageAPIToken(), keboola.WithClient(&c))
	assert.NoError(t, err)

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, branch)

	createPayload := &keboola.VaultVariableCreatePayload{
		Key:   "test_encrypted_variable",
		Value: "secret_value",
		Flags: []string{"encrypted"},
		Attributes: map[string]interface{}{
			"branchId": branch.ID.String(),
		},
	}
	variable, err := api.CreateVariableRequest(createPayload).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, variable)
	assert.Equal(t, "test_encrypted_variable", variable.Key)
	assert.NotEmpty(t, variable.Hash)
	assert.Contains(t, variable.Flags, "encrypted")

	listOpts := &keboola.VaultVariableListOptions{
		Key: "test_encrypted_variable",
		Attributes: map[string]interface{}{
			"branchId": branch.ID.String(),
		},
	}
	variables, err := api.ListVariablesRequest(listOpts).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, variables)
	assert.GreaterOrEqual(t, len(*variables), 1)

	found := false
	for _, v := range *variables {
		if v.Hash == variable.Hash {
			found = true
			assert.Equal(t, "test_encrypted_variable", v.Key)
			assert.Contains(t, v.Flags, "encrypted")
			break
		}
	}
	assert.True(t, found, "Created encrypted variable should be in the list")

	err = api.DeleteVariableRequest(variable.Hash).SendOrErr(ctx)
	assert.NoError(t, err)
}

func TestVariablesWithGroup(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	project, _ := testproject.GetTestProjectForTest(t)
	c := client.NewTestClient()
	api, err := keboola.NewAuthorizedAPI(ctx, project.StorageAPIHost(), project.StorageAPIToken(), keboola.WithClient(&c))
	assert.NoError(t, err)

	// Get default branch
	branch, err := api.GetDefaultBranchRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, branch)

	createPayload := &keboola.VaultVariableCreatePayload{
		Key:   "test_grouped_variable",
		Value: "grouped_value",
		Group: "Test Group",
		Attributes: map[string]interface{}{
			"branchId": branch.ID.String(),
		},
	}
	variable, err := api.CreateVariableRequest(createPayload).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, variable)
	assert.Equal(t, "test_grouped_variable", variable.Key)
	assert.Equal(t, "Test Group", variable.Group)
	assert.NotEmpty(t, variable.Hash)

	listOpts := &keboola.VaultVariableListOptions{
		Key: "test_grouped_variable",
	}
	scopedVariables, err := api.ListVariablesScopedRequest(branch.ID, listOpts).Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, scopedVariables)
	assert.GreaterOrEqual(t, len(*scopedVariables), 1)

	found := false
	for _, v := range *scopedVariables {
		if v.Hash == variable.Hash {
			found = true
			assert.Equal(t, "test_grouped_variable", v.Key)
			assert.Equal(t, "Test Group", v.Group)
			break
		}
	}
	assert.True(t, found, "Created grouped variable should be in the scoped list")

	err = api.DeleteVariableRequest(variable.Hash).SendOrErr(ctx)
	assert.NoError(t, err)
}
