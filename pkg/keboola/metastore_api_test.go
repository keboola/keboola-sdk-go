package keboola_test

import (
	"errors"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/client"
	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

const (
	metastoreTestHost = "https://metastore.test"
	metastoreListURL  = metastoreTestHost + "/api/v1/repository/rls-policy"
)

const metastoreListBody = `{
	"data": [
		{
			"id": "11111111-1111-1111-1111-111111111111",
			"type": "rls-policy",
			"attributes": {"table": "in.c-main.orders", "big": 9007199254740993},
			"meta": {"schemaVersion": "1.1.0", "branch": "main", "scope": "organization", "revision": 3,
				"projectId": 123, "createdAt": "2026-10-01T10:00:00Z", "deletedAt": null}
		},
		{
			"id": "22222222-2222-2222-2222-222222222222",
			"type": "rls-policy",
			"attributes": {"table": "in.c-main.users"},
			"meta": {"schemaVersion": "1.0.0", "branch": "main", "scope": "project", "revision": 1,
				"deletedAt": "2026-10-02T10:00:00Z"}
		}
	]
}`

// newMetastoreAPI returns an API whose Metastore requests go to metastoreTestHost, as in-cluster callers configure it, on
// top of an index that does not know the service.
func newMetastoreAPI(t *testing.T) (*keboola.AuthorizedAPI, *httpmock.MockTransport) {
	t.Helper()

	c, transport := client.NewMockedClient()
	c = c.WithHeader("X-Kubernetes-Authorization", "Bearer sa-token")
	api := keboola.NewPublicAPIFromIndex(
		"https://connection.test",
		&keboola.Index{},
		keboola.WithClient(&c),
		keboola.WithServiceURL(keboola.MetastoreAPI, metastoreTestHost),
	).NewAuthorizedAPI("my-token", time.Minute)
	return api, transport
}

func metastoreResponder(status int, body string, header http.Header, onRequest func(*http.Request)) httpmock.Responder {
	return func(req *http.Request) (*http.Response, error) {
		if onRequest != nil {
			onRequest(req)
		}
		resp := httpmock.NewStringResponse(status, body)
		resp.Header.Set("Content-Type", "application/vnd.api+json")
		for k, values := range header {
			for _, v := range values {
				resp.Header.Add(k, v)
			}
		}
		return resp, nil
	}
}

func TestListMetastoreObjectsRequest(t *testing.T) {
	t.Parallel()

	api, transport := newMetastoreAPI(t)
	var got *http.Request
	transport.RegisterResponder(http.MethodGet, `=~^`+metastoreListURL+`\?`,
		metastoreResponder(http.StatusOK, metastoreListBody, http.Header{"X-Verified-Read": {"true"}}, func(req *http.Request) { got = req }))

	list, err := api.ListMetastoreObjectsRequest("rls-policy",
		keboola.WithMetastoreBranch("main"),
		keboola.WithMetastoreFilterIn("table", "in.c-main.orders", "in.c-main.users"),
		keboola.WithMetastoreLimit(50),
		keboola.WithMetastoreOffset(10),
	).Send(t.Context())
	require.NoError(t, err)

	// What was sent.
	require.NotNil(t, got)
	assert.Equal(t, "my-token", got.Header.Get("X-StorageApi-Token"))
	assert.Equal(t, "Bearer sa-token", got.Header.Get("X-Kubernetes-Authorization"), "the client's header reaches the Metastore")
	assert.Equal(t, "application/vnd.api+json", got.Header.Get("Accept"))
	assert.Equal(t, url.Values{
		"branch":    {"main"},
		"table[in]": {"in.c-main.orders|in.c-main.users"},
		"limit":     {"50"},
		"offset":    {"10"},
	}, got.URL.Query())

	// What was decoded.
	require.Len(t, list.Data, 2)
	first := list.Data[0]
	assert.Equal(t, "11111111-1111-1111-1111-111111111111", first.ID)
	assert.Equal(t, "rls-policy", first.Type)
	assert.JSONEq(t, `{"table": "in.c-main.orders", "big": 9007199254740993}`, string(first.Attributes))
	assert.Contains(t, string(first.Attributes), "9007199254740993", "attributes stay raw, a large number is not rounded")
	assert.Equal(t, "1.1.0", first.Meta.SchemaVersion)
	assert.Equal(t, "organization", first.Meta.Scope)
	assert.Equal(t, 3, first.Meta.Revision)
	assert.Equal(t, 123, first.Meta.ProjectID)
	assert.Nil(t, first.Meta.DeletedAt, "a null deletedAt is not set")
	assert.NotNil(t, list.Data[1].Meta.DeletedAt, "a deletedAt is kept")
	assert.Equal(t, "true", list.Header.Get("X-Verified-Read"), "the headers of the answer are exposed")
}

// One value is an equality: the Metastore refuses an "in" filter without a pipe.
func TestListMetastoreObjectsRequest_FilterShape(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		opt    keboola.ListMetastoreObjectsOption
		expect url.Values
	}{
		{name: "one value", opt: keboola.WithMetastoreFilterIn("table", "in.c-main.orders"), expect: url.Values{"table": {"in.c-main.orders"}}},
		{name: "two values", opt: keboola.WithMetastoreFilterIn("table", "a.b", "c.d"), expect: url.Values{"table[in]": {"a.b|c.d"}}},
		{name: "equality", opt: keboola.WithMetastoreFilter("kid", "KID:x"), expect: url.Values{"kid": {"KID:x"}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			api, transport := newMetastoreAPI(t)
			var got url.Values
			transport.RegisterResponder(http.MethodGet, `=~^`+metastoreListURL,
				metastoreResponder(http.StatusOK, `{"data": []}`, nil, func(req *http.Request) { got = req.URL.Query() }))

			_, err := api.ListMetastoreObjectsRequest("rls-policy", tc.opt).Send(t.Context())
			require.NoError(t, err)
			assert.Equal(t, tc.expect, got)
		})
	}
}

// A request that cannot be formed is never sent.
func TestListMetastoreObjectsRequest_Invalid(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name       string
		objectType string
		opts       []keboola.ListMetastoreObjectsOption
	}{
		{name: "a value with a pipe", objectType: "rls-policy", opts: []keboola.ListMetastoreObjectsOption{keboola.WithMetastoreFilterIn("table", "a.b|c.d")}},
		{name: "no value", objectType: "rls-policy", opts: []keboola.ListMetastoreObjectsOption{keboola.WithMetastoreFilterIn("table")}},
		{name: "a field with a bracket", objectType: "rls-policy", opts: []keboola.ListMetastoreObjectsOption{keboola.WithMetastoreFilter("table[in]", "a.b")}},
		{name: "an empty object type", objectType: ""},
		{name: "an object type with a slash", objectType: "rls-policy/x"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			api, transport := newMetastoreAPI(t)
			_, err := api.ListMetastoreObjectsRequest(tc.objectType, tc.opts...).Send(t.Context())
			require.Error(t, err)
			assert.Zero(t, transport.GetTotalCallCount(), "nothing is sent")
		})
	}
}

// The list answer tells a missing "data" from an empty list, so a caller can refuse a broken answer.
func TestListMetastoreObjectsRequest_EmptyAndMissingData(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name  string
		body  string
		isNil bool
	}{
		{name: "empty list", body: `{"data": []}`, isNil: false},
		{name: "no data key", body: `{}`, isNil: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			api, transport := newMetastoreAPI(t)
			transport.RegisterResponder(http.MethodGet, `=~^`+metastoreListURL, metastoreResponder(http.StatusOK, tc.body, nil, nil))

			list, err := api.ListMetastoreObjectsRequest("rls-policy").Send(t.Context())
			require.NoError(t, err)
			assert.Equal(t, tc.isNil, list.Data == nil)
			assert.Empty(t, list.Data)
		})
	}
}

func TestListMetastoreObjectsRequest_Error(t *testing.T) {
	t.Parallel()

	api, transport := newMetastoreAPI(t)
	transport.RegisterResponder(http.MethodGet, `=~^`+metastoreListURL, metastoreResponder(http.StatusForbidden,
		`{"error": 403, "code": "403", "exception": "Access denied", "exceptionId": "metastore-abc", "status": "error"}`,
		http.Header{"X-Refused": {"true"}}, nil))

	_, err := api.ListMetastoreObjectsRequest("rls-policy").Send(t.Context())
	require.Error(t, err)

	var metastoreErr *keboola.MetastoreError
	require.True(t, errors.As(err, &metastoreErr), "the error is a MetastoreError: %v", err)
	assert.Equal(t, http.StatusForbidden, metastoreErr.StatusCode())
	assert.Equal(t, "Access denied", metastoreErr.ErrorUserMessage())
	assert.Equal(t, "metastore-abc", metastoreErr.ErrorExceptionID())
	assert.Equal(t, "true", metastoreErr.ResponseHeader().Get("X-Refused"), "a refusal marker header is readable")
}

func TestGetMetastoreObjectRequest(t *testing.T) {
	t.Parallel()

	api, transport := newMetastoreAPI(t)
	var got *http.Request
	transport.RegisterResponder(http.MethodGet, metastoreListURL+"/11111111-1111-1111-1111-111111111111",
		metastoreResponder(http.StatusOK, `{"data": {"id": "11111111-1111-1111-1111-111111111111", "type": "rls-policy", "attributes": {"table": "a.b"}, "meta": {"revision": 2}}}`,
			nil, func(req *http.Request) { got = req }))

	object, err := api.GetMetastoreObjectRequest("rls-policy", "11111111-1111-1111-1111-111111111111").Send(t.Context())
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, "my-token", got.Header.Get("X-StorageApi-Token"))
	require.NotNil(t, object.Data)
	assert.Equal(t, 2, object.Data.Meta.Revision)

	_, err = api.GetMetastoreObjectRequest("rls-policy", "").Send(t.Context())
	require.Error(t, err, "an empty id is never sent")
}

// Without an override the service is looked up in the index, like every other service.
func TestMetastoreAPI_ServiceFromIndex(t *testing.T) {
	t.Parallel()

	c, transport := client.NewMockedClient()
	api := keboola.NewPublicAPIFromIndex(
		"https://connection.test",
		&keboola.Index{Services: keboola.Services{{ID: "metastore", URL: "https://metastore.from-index.test"}}},
		keboola.WithClient(&c),
	).NewAuthorizedAPI("my-token", time.Minute)
	transport.RegisterResponder(http.MethodGet, `=~^https://metastore\.from-index\.test/api/v1/repository/rls-policy`,
		metastoreResponder(http.StatusOK, `{"data": []}`, nil, nil))

	_, err := api.ListMetastoreObjectsRequest("rls-policy").Send(t.Context())
	require.NoError(t, err)
	assert.Equal(t, 1, transport.GetTotalCallCount())
}
