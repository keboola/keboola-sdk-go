package keboola_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/client"
	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

const dataScienceHost = "https://data-science.keboola.mock"

func mockedDataScienceAPI(t *testing.T) (*keboola.AuthorizedAPI, *httpmock.MockTransport) {
	t.Helper()
	c, transport := client.NewMockedClient()
	idx := &keboola.Index{
		Services: keboola.Services{
			{ID: "data-science", URL: dataScienceHost},
		},
	}
	publicAPI := keboola.NewPublicAPIFromIndex("https://connection.keboola.mock", idx, keboola.WithClient(&c))
	return publicAPI.NewAuthorizedAPI("test-token", time.Minute), transport
}

func TestListDataScienceAppsRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	api, transport := mockedDataScienceAPI(t)
	transport.RegisterResponder(http.MethodGet, dataScienceHost+"/apps",
		httpmock.NewJsonResponderOrPanic(http.StatusOK, []*keboola.DataScienceApp{
			{ID: "app-1", Type: keboola.DataScienceAppTypePython, Size: "small"},
			{ID: "app-2", Type: keboola.DataScienceAppTypeR, Size: "medium"},
		}),
	)

	apps, err := api.ListDataScienceAppsRequest().Send(ctx)
	require.NoError(t, err)
	require.Len(t, *apps, 2)
	assert.Equal(t, keboola.DataScienceAppID("app-1"), (*apps)[0].ID)
	assert.Equal(t, keboola.DataScienceAppTypePython, (*apps)[0].Type)
	assert.Equal(t, keboola.DataScienceAppID("app-2"), (*apps)[1].ID)
	assert.Equal(t, keboola.DataScienceAppTypeR, (*apps)[1].Type)
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET "+dataScienceHost+"/apps"])
}

func TestListDataScienceAppsRequest_TypeFilter(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	api, transport := mockedDataScienceAPI(t)
	transport.RegisterResponder(http.MethodGet, dataScienceHost+"/apps",
		func(req *http.Request) (*http.Response, error) {
			types := req.URL.Query()["type[]"]
			assert.ElementsMatch(t, []string{"python", "r"}, types)
			return httpmock.NewJsonResponderOrPanic(http.StatusOK, []*keboola.DataScienceApp{})(req)
		},
	)

	_, err := api.ListDataScienceAppsRequest(
		keboola.WithDataScienceAppsType(keboola.DataScienceAppTypePython, keboola.DataScienceAppTypeR),
	).Send(ctx)
	require.NoError(t, err)
}

func TestListDataScienceAppsRequest_PaginationParams(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	api, transport := mockedDataScienceAPI(t)
	transport.RegisterResponder(http.MethodGet, dataScienceHost+"/apps",
		func(req *http.Request) (*http.Response, error) {
			q := req.URL.Query()
			assert.Equal(t, "50", q.Get("limit"))
			assert.Equal(t, "10", q.Get("offset"))
			assert.Equal(t, "branch-123", q.Get("branchId"))
			return httpmock.NewJsonResponderOrPanic(http.StatusOK, []*keboola.DataScienceApp{})(req)
		},
	)

	_, err := api.ListDataScienceAppsRequest(
		keboola.WithDataScienceAppsLimit(50),
		keboola.WithDataScienceAppsOffset(10),
		keboola.WithDataScienceAppsBranchID("branch-123"),
	).Send(ctx)
	require.NoError(t, err)
}

func TestGetDataScienceAppRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	api, transport := mockedDataScienceAPI(t)
	transport.RegisterResponder(http.MethodGet, dataScienceHost+"/apps/app-42",
		httpmock.NewJsonResponderOrPanic(http.StatusOK, &keboola.DataScienceApp{
			ID:   "app-42",
			Type: keboola.DataScienceAppTypeStreamlit,
			Size: "large",
			URL:  "https://app.example.com",
		}),
	)

	app, err := api.GetDataScienceAppRequest(keboola.DataScienceAppID("app-42")).Send(ctx)
	require.NoError(t, err)
	assert.Equal(t, keboola.DataScienceAppID("app-42"), app.ID)
	assert.Equal(t, keboola.DataScienceAppTypeStreamlit, app.Type)
	assert.Equal(t, "large", app.Size)
	assert.Equal(t, "https://app.example.com", app.URL)
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET "+dataScienceHost+"/apps/app-42"])
}
