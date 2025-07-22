package keboola_test

import (
	"context"
	"testing"

	"github.com/keboola/go-utils/pkg/testproject"
	"github.com/stretchr/testify/assert"

	"github.com/keboola/keboola-sdk-go/v2/pkg/client"
	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

func TestEncryptRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	project, _ := testproject.GetTestProjectForTest(t)
	c := client.NewTestClient()
	api, err := keboola.NewPublicAPI(ctx, project.StorageAPIHost(), keboola.WithClient(&c))
	assert.NoError(t, err)

	mapToEncrypt := map[string]string{"#keyToEncrypt": "value"}
	componentID := keboola.ComponentID("keboola.ex-generic-v2")
	encryptedMapPtr, err := api.EncryptRequest(
		1234,
		&componentID,
		nil,
		nil,
		mapToEncrypt,
	).Send(ctx)
	assert.NoError(t, err)
	encryptedMap := *encryptedMapPtr
	assert.NotEmpty(t, encryptedMap)
	assert.True(t, keboola.IsEncrypted(encryptedMap["#keyToEncrypt"]))
}

func TestEncryptRequestProjectScope(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	project, _ := testproject.GetTestProjectForTest(t)
	c := client.NewTestClient()
	api, err := keboola.NewPublicAPI(ctx, project.StorageAPIHost(), keboola.WithClient(&c))
	assert.NoError(t, err)

	mapToEncrypt := map[string]string{"#keyToEncrypt": "value"}
	_, err = api.EncryptRequest(1234, nil, nil, nil, mapToEncrypt).Send(ctx)
	assert.NoError(t, err)
}
