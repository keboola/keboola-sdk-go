package keboola_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

const tableListV2FirstPage = `{
  "data": [
    {
      "id": "in.c-main.typed",
      "bucketId": "in.c-main",
      "name": "typed",
      "displayName": "typed",
      "branchId": 1234,
      "primaryKey": ["id"],
      "isAlias": false,
      "isTyped": true,
      "tableType": "table",
      "rowsCount": 10,
      "dataSizeBytes": 2048,
      "created": "2026-10-22T10:00:00+0200",
      "lastImportDate": "2026-10-22T10:05:00+0200",
      "lastChangeDate": null,
      "description": "Native table text",
      "isDescriptionSystemManaged": false,
      "columns": ["id", "note"],
      "definition": {
        "primaryKeysNames": ["id"],
        "description": "Native table text",
        "columns": [
          {"name": "id", "basetype": "INTEGER", "definition": {"type": "INTEGER", "nullable": false, "description": "Native column text"}},
          {"name": "note", "basetype": "STRING", "definition": {"type": "VARCHAR", "length": "255", "nullable": true}}
        ]
      }
    },
    {
      "id": "in.c-main.legacy",
      "bucketId": "in.c-main",
      "name": "legacy",
      "displayName": "legacy",
      "branchId": 1234,
      "primaryKey": [],
      "isAlias": false,
      "isTyped": true,
      "tableType": "table",
      "rowsCount": null,
      "dataSizeBytes": null,
      "created": null,
      "lastImportDate": null,
      "lastChangeDate": null,
      "description": null,
      "isDescriptionSystemManaged": false,
      "definition": null
    }
  ],
  "meta": {"total": 3, "limit": 2},
  "links": {
    "self": "https://connection.keboola.mock/v2/storage/branch/1234/table-list?include=definition&page[limit]=2",
    "first": "https://connection.keboola.mock/v2/storage/branch/1234/table-list?include=definition&page[limit]=2",
    "next": "https://connection.keboola.mock/v2/storage/branch/1234/table-list?include=definition&page[limit]=2&page[after]=eyJhIjoiYiJ9"
  }
}`

const tableListV2LastPage = `{
  "data": [
    {
      "id": "in.c-main.untyped",
      "bucketId": "in.c-main",
      "name": "untyped",
      "displayName": "untyped",
      "branchId": 1234,
      "primaryKey": [],
      "isAlias": false,
      "isTyped": false,
      "tableType": "table",
      "rowsCount": 0,
      "dataSizeBytes": 0,
      "created": null,
      "lastImportDate": null,
      "lastChangeDate": null,
      "description": "Untyped text",
      "isDescriptionSystemManaged": false,
      "definition": {
        "primaryKeysNames": [],
        "columns": [
          {"name": "a", "definition": {"description": "Untyped column text"}},
          {"name": "b", "definition": {}}
        ]
      }
    }
  ],
  "meta": {"total": 3, "limit": 2},
  "links": {
    "self": "https://connection.keboola.mock/v2/storage/branch/1234/table-list?include=definition&page[limit]=2&page[after]=eyJhIjoiYiJ9",
    "first": "https://connection.keboola.mock/v2/storage/branch/1234/table-list?include=definition&page[limit]=2",
    "next": null
  }
}`

func jsonResponse(status int, body string) *http.Response {
	response := httpmock.NewStringResponse(status, body)
	response.Header.Set("Content-Type", "application/json")
	return response
}

func authorizedAPIWithMockedTransport(t *testing.T) (*keboola.AuthorizedAPI, *httpmock.MockTransport) {
	t.Helper()

	c, transport := mockedClient()
	api, err := keboola.NewPublicAPI(t.Context(), "https://connection.keboola.mock", keboola.WithClient(&c))
	require.NoError(t, err)
	return api.NewAuthorizedAPI("my-token", 1*time.Minute), transport
}

func TestListTablesV2Request_Query(t *testing.T) {
	t.Parallel()

	api, transport := authorizedAPIWithMockedTransport(t)
	transport.RegisterResponder(http.MethodGet, "/v2/storage/branch/1234/table-list", func(request *http.Request) (*http.Response, error) {
		assert.Equal(t, "my-token", request.Header.Get("X-StorageApi-Token"))
		query := request.URL.Query()
		assert.Equal(t, "columns,definition,metadata", query.Get("include"))
		assert.Equal(t, "in.c-main", query.Get("filter[bucketId]"))
		assert.Equal(t, "50", query.Get("page[limit]"))
		assert.Equal(t, "cursor", query.Get("page[after]"))
		return jsonResponse(http.StatusOK, `{"data":[],"meta":{"total":0,"limit":50},"links":{"next":null}}`), nil
	})

	page, err := api.ListTablesV2Request(
		1234,
		keboola.WithTableListV2Definition(),
		keboola.WithTableListV2Columns(),
		keboola.WithTableListV2Metadata(),
		keboola.WithTableListV2Bucket(keboola.BucketID{Stage: keboola.BucketStageIn, BucketName: "c-main"}),
		keboola.WithTableListV2Limit(50),
		keboola.WithTableListV2After("cursor"),
	).Send(t.Context())
	require.NoError(t, err)
	assert.Empty(t, page.Data)
	assert.Empty(t, page.Links.Next)
}

func TestListTablesV2Request_OmitsUnsetParameters(t *testing.T) {
	t.Parallel()

	api, transport := authorizedAPIWithMockedTransport(t)
	transport.RegisterResponder(http.MethodGet, "/v2/storage/branch/1234/table-list", func(request *http.Request) (*http.Response, error) {
		assert.Empty(t, request.URL.RawQuery)
		return jsonResponse(http.StatusOK, `{"data":[],"meta":{},"links":{"next":null}}`), nil
	})

	_, err := api.ListTablesV2Request(1234).Send(t.Context())
	require.NoError(t, err)
}

func TestListTablesV2_FollowsPagesAndDecodesDefinitions(t *testing.T) {
	t.Parallel()

	api, transport := authorizedAPIWithMockedTransport(t)
	transport.RegisterResponder(http.MethodGet, "/v2/storage/branch/1234/table-list", func(request *http.Request) (*http.Response, error) {
		assert.Equal(t, "definition", request.URL.Query().Get("include"))
		switch request.URL.Query().Get("page[after]") {
		case "":
			return jsonResponse(http.StatusOK, tableListV2FirstPage), nil
		case "eyJhIjoiYiJ9":
			return jsonResponse(http.StatusOK, tableListV2LastPage), nil
		default:
			t.Errorf("unexpected cursor %q", request.URL.Query().Get("page[after]"))
			return jsonResponse(http.StatusBadRequest, "{}"), nil
		}
	})

	tables, err := api.ListTablesV2(t.Context(), 1234, keboola.WithTableListV2Definition())
	require.NoError(t, err)
	assert.Equal(t, 2, transport.GetCallCountInfo()["GET /v2/storage/branch/1234/table-list"])
	require.Len(t, tables, 3)

	typed := tables[0]
	assert.Equal(t, "in.c-main.typed", typed.TableID.String())
	assert.Equal(t, keboola.BranchID(1234), typed.BranchID)
	assert.True(t, typed.IsTyped)
	assert.Equal(t, "Native table text", typed.Description)
	require.NotNil(t, typed.Definition)
	assert.Equal(t, "Native table text", typed.Definition.Description)
	require.Len(t, typed.Definition.Columns, 2)
	require.NotNil(t, typed.Definition.Columns[0].Definition)
	assert.Equal(t, "Native column text", typed.Definition.Columns[0].Definition.Description)
	assert.Equal(t, "INTEGER", typed.Definition.Columns[0].Definition.Type)
	assert.Equal(t, "255", typed.Definition.Columns[1].Definition.Length)
	require.NotNil(t, typed.RowsCount)
	assert.Equal(t, uint64(10), *typed.RowsCount)
	assert.Nil(t, typed.LastChangeDate)

	// A typed table whose types are not all native has no definition in the list.
	legacy := tables[1]
	assert.Nil(t, legacy.Definition)
	assert.Nil(t, legacy.RowsCount)
	assert.Empty(t, legacy.Description)

	// An untyped column keeps a description-only definition and drops an empty one.
	untyped := tables[2]
	require.NotNil(t, untyped.Definition)
	assert.Equal(t, "Untyped column text", untyped.Definition.Columns[0].Definition.Description)
	assert.Empty(t, untyped.Definition.Columns[0].Definition.Type)
	assert.Nil(t, untyped.Definition.Columns[1].Definition)
}

func TestListTablesV2_Error(t *testing.T) {
	t.Parallel()

	api, transport := authorizedAPIWithMockedTransport(t)
	transport.RegisterResponder(http.MethodGet, "/v2/storage/branch/1234/table-list", func(*http.Request) (*http.Response, error) {
		return jsonResponse(http.StatusNotFound, `{"error":"Not found"}`), nil
	})

	_, err := api.ListTablesV2(t.Context(), 1234)
	require.Error(t, err)
}

func TestTableListV2Links_After(t *testing.T) {
	t.Parallel()

	after, err := keboola.TableListV2Links{Next: "https://connection.keboola.mock/v2/storage/branch/default/table-list?include=definition&page[after]=abc"}.After()
	require.NoError(t, err)
	assert.Equal(t, "abc", after)

	after, err = keboola.TableListV2Links{}.After()
	require.NoError(t, err)
	assert.Empty(t, after)
}

func TestListTablesV2_StartsAfterGivenCursorAndStopsOnRepeatedCursor(t *testing.T) {
	t.Parallel()

	api, transport := authorizedAPIWithMockedTransport(t)
	transport.RegisterResponder(http.MethodGet, "/v2/storage/branch/1234/table-list", func(request *http.Request) (*http.Response, error) {
		assert.Equal(t, "start", request.URL.Query().Get("page[after]"))
		return jsonResponse(http.StatusOK, `{"data":[],"meta":{},"links":{"next":"https://connection.keboola.mock/v2/storage/branch/1234/table-list?page[after]=start"}}`), nil
	})

	_, err := api.ListTablesV2(t.Context(), 1234, keboola.WithTableListV2After("start"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "again")
	assert.Equal(t, 1, transport.GetCallCountInfo()["GET /v2/storage/branch/1234/table-list"])
}
