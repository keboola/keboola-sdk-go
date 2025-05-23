package keboola_test

import (
	"context"
	"testing"

	"github.com/keboola/go-utils/pkg/testproject"
	"github.com/stretchr/testify/assert"

	"github.com/keboola/keboola-sdk-go/v2/pkg/client"
	. "github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

func TestIndexRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	project, _ := testproject.GetTestProjectForTest(t)
	c := client.NewTestClient()
	api, err := NewPublicAPI(ctx, project.StorageAPIHost(), WithClient(&c))
	assert.NoError(t, err)
	result, err := api.IndexRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, result.Features)
	assert.NotEmpty(t, result.Services)
	_, found := result.AllServices().URLByID("encryption")
	assert.True(t, found)
	_, found = result.AllServices().URLByID("scheduler")
	assert.True(t, found)
}

func TestIndexRequest_WithoutToken(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	project, _ := testproject.GetTestProjectForTest(t)
	c := client.NewTestClient()
	api, err := NewPublicAPI(ctx, project.StorageAPIHost(), WithClient(&c))
	assert.NoError(t, err)
	result, err := api.IndexRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, result.Features)
	assert.NotEmpty(t, result.Services)
	_, found := result.AllServices().URLByID("encryption")
	assert.True(t, found)
	_, found = result.AllServices().URLByID("scheduler")
	assert.True(t, found)
}

func TestIndexComponents(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	project, _ := testproject.GetTestProjectForTest(t)
	c := client.NewTestClient()
	api, err := NewPublicAPI(ctx, project.StorageAPIHost(), WithClient(&c))
	assert.NoError(t, err)
	result, err := api.IndexComponentsRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, result.Components)
	assert.NotEmpty(t, result.Features)
	assert.NotEmpty(t, result.Services)
	_, found := result.AllServices().URLByID("encryption")
	assert.True(t, found)
	_, found = result.AllServices().URLByID("scheduler")
	assert.True(t, found)
}

func TestIndexComponents_WithoutToken(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	c := client.NewTestClient()
	api, err := NewPublicAPI(ctx, `connection.keboola.com`, WithClient(&c))
	assert.NoError(t, err)
	result, err := api.IndexComponentsRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, result.Components)
	assert.NotEmpty(t, result.Features)
	assert.NotEmpty(t, result.Services)
	// Services
	url, found := result.AllServices().URLByID("encryption")
	assert.True(t, found)
	assert.Equal(t, ServiceURL("https://encryption.keboola.com"), url)
	url, found = result.AllServices().URLByID("scheduler")
	assert.True(t, found)
	assert.Equal(t, ServiceURL("https://scheduler.keboola.com"), url)
}
