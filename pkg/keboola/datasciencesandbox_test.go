package keboola_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

func TestCreateDataScienceSandboxRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	api, transport := mockedDataScienceAPI(t)
	transport.RegisterResponder(http.MethodPost, dataScienceHost+"/sandboxes",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponderOrPanic(http.StatusCreated, &keboola.DataScienceSandbox{
				ID:           "sandbox-1",
				Type:         keboola.DataScienceAppTypePython,
				Size:         keboola.DataScienceSandboxSizeSmall,
				Active:       true,
				BranchID:     "123",
				ComponentID:  keboola.SandboxWorkspacesComponent,
				URL:          "https://sandbox.example.com",
				Password:     "secret",
				ImageVersion: "1.0.0",
				Packages:     []string{"numpy", "pandas"},
			})(req)
		},
	)

	payload := keboola.CreateDataScienceSandboxPayload{
		Type:        keboola.DataScienceAppTypePython,
		Size:        keboola.DataScienceSandboxSizeSmall,
		BranchID:    "123",
		ComponentID: keboola.SandboxWorkspacesComponent,
	}
	sandbox, err := api.CreateDataScienceSandboxRequest(payload).Send(ctx)
	require.NoError(t, err)
	assert.Equal(t, keboola.DataScienceAppID("sandbox-1"), keboola.DataScienceAppID(sandbox.ID))
	assert.Equal(t, keboola.DataScienceAppTypePython, sandbox.Type)
	assert.Equal(t, keboola.DataScienceSandboxSizeSmall, sandbox.Size)
	assert.True(t, sandbox.Active)
	assert.Equal(t, "https://sandbox.example.com", sandbox.URL)
	assert.Equal(t, []string{"numpy", "pandas"}, sandbox.Packages)
	assert.Equal(t, 1, transport.GetCallCountInfo()["POST "+dataScienceHost+"/sandboxes"])
}

func TestGetDataScienceSandboxRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	api, transport := mockedDataScienceAPI(t)
	transport.RegisterResponder(http.MethodGet, dataScienceHost+"/sandboxes/sandbox-42",
		httpmock.NewJsonResponderOrPanic(http.StatusOK, &keboola.DataScienceSandbox{
			ID:     "sandbox-42",
			Type:   keboola.DataScienceAppTypeR,
			Size:   keboola.DataScienceSandboxSizeMedium,
			Active: true,
			URL:    "https://r-sandbox.example.com",
		}),
	)

	sandbox, err := api.GetDataScienceSandboxRequest(keboola.DataScienceAppID("sandbox-42")).Send(ctx)
	require.NoError(t, err)
	assert.Equal(t, "sandbox-42", sandbox.ID)
	assert.Equal(t, keboola.DataScienceAppTypeR, sandbox.Type)
	assert.Equal(t, keboola.DataScienceSandboxSizeMedium, sandbox.Size)
	assert.True(t, sandbox.Active)
	assert.Equal(t, "https://r-sandbox.example.com", sandbox.URL)
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET "+dataScienceHost+"/sandboxes/sandbox-42"])
}

func TestUpdateDataScienceSandboxRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	api, transport := mockedDataScienceAPI(t)
	transport.RegisterResponder(http.MethodPatch, dataScienceHost+"/sandboxes/sandbox-42",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponderOrPanic(http.StatusOK, &keboola.DataScienceSandbox{
				ID:                      "sandbox-42",
				Type:                    keboola.DataScienceAppTypePython,
				Size:                    keboola.DataScienceSandboxSizeLarge,
				AutoSuspendAfterSeconds: 3600,
			})(req)
		},
	)

	payload := keboola.UpdateDataScienceSandboxPayload{
		Size:                    keboola.DataScienceSandboxSizeLarge,
		AutoSuspendAfterSeconds: 3600,
	}
	sandbox, err := api.UpdateDataScienceSandboxRequest(keboola.DataScienceAppID("sandbox-42"), payload).Send(ctx)
	require.NoError(t, err)
	assert.Equal(t, "sandbox-42", sandbox.ID)
	assert.Equal(t, keboola.DataScienceSandboxSizeLarge, sandbox.Size)
	assert.Equal(t, 3600, sandbox.AutoSuspendAfterSeconds)
	assert.Equal(t, 1, transport.GetCallCountInfo()["PATCH "+dataScienceHost+"/sandboxes/sandbox-42"])
}

func TestDeleteDataScienceSandboxRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	api, transport := mockedDataScienceAPI(t)
	transport.RegisterResponder(http.MethodDelete, dataScienceHost+"/sandboxes/sandbox-42",
		httpmock.NewStringResponder(http.StatusNoContent, ""),
	)

	err := api.DeleteDataScienceSandboxRequest(keboola.DataScienceAppID("sandbox-42")).SendOrErr(ctx)
	require.NoError(t, err)
	assert.Equal(t, 1, transport.GetCallCountInfo()["DELETE "+dataScienceHost+"/sandboxes/sandbox-42"])
}
