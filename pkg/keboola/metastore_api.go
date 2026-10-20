package keboola

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

const (
	// MetastoreAPI is the Metastore service. It is looked up in the index like every other service, so a caller that
	// reaches it by another address (for example an in-cluster one) passes an index that lists it there.
	//
	// A caller that must bound or harden the reads builds the API with a client of its own (client.WithMaxResponseBytes,
	// client.WithoutRedirects, its own headers): the client belongs to the whole API, so use a separate API for it.
	MetastoreAPI = ServiceType("metastore")

	MetastoreAPIObjects = "api/v1/repository/{objectType}"
	MetastoreAPIObject  = "api/v1/repository/{objectType}/{objectId}"

	// metastoreAccept is the media type the repository answers in. A client that does not ask for it may get a plain
	// JSON object back, and the SDK decodes only the JSON:API shape.
	metastoreAccept = "application/vnd.api+json"
)

// metastoreNamePattern is the shape of an object type, an object id and the field of a filter: no character that has a
// meaning in the path or in the filter syntax ("[in]", "[order]").
var metastoreNamePattern = regexp.MustCompile(`^[A-Za-z0-9._-]+$`)

// ListMetastoreObjectsOption is an optional parameter of ListMetastoreObjectsRequest.
type ListMetastoreObjectsOption func(c *listMetastoreObjectsConfig)

type metastoreQueryParam struct{ key, value string }

type listMetastoreObjectsConfig struct {
	query []metastoreQueryParam
	err   error
}

func (c *listMetastoreObjectsConfig) add(key, value string) {
	c.query = append(c.query, metastoreQueryParam{key: key, value: value})
}

// addFilter adds a filter on the attribute field. op is "" for an equality, or the operator in brackets.
func (c *listMetastoreObjectsConfig) addFilter(field, op, value string) {
	if err := validateMetastoreName("filter field", field); err != nil {
		c.err = err
		return
	}
	c.add(field+op, value)
}

// WithMetastoreBranch lists the objects of the branch.
func WithMetastoreBranch(branch string) ListMetastoreObjectsOption {
	return func(c *listMetastoreObjectsConfig) { c.add("branch", branch) }
}

// WithMetastoreLimit limits the number of objects of the answer.
func WithMetastoreLimit(limit int) ListMetastoreObjectsOption {
	return func(c *listMetastoreObjectsConfig) { c.add("limit", strconv.Itoa(limit)) }
}

// WithMetastoreOffset skips the first objects of the list.
func WithMetastoreOffset(offset int) ListMetastoreObjectsOption {
	return func(c *listMetastoreObjectsConfig) { c.add("offset", strconv.Itoa(offset)) }
}

// WithMetastoreFilter lists the objects whose attribute field equals value.
//
// The Metastore ignores a filter on a field its schema does not know, so a caller that depends on the filter must check
// the objects it gets back.
func WithMetastoreFilter(field, value string) ListMetastoreObjectsOption {
	return func(c *listMetastoreObjectsConfig) { c.addFilter(field, "", value) }
}

// WithMetastoreFilterIn lists the objects whose attribute field equals one of the values. One value is sent as an
// equality, because the Metastore refuses an "in" filter that has no pipe. A value must not contain "|", which separates
// the values of an "in" filter.
func WithMetastoreFilterIn(field string, values ...string) ListMetastoreObjectsOption {
	return func(c *listMetastoreObjectsConfig) {
		switch {
		case len(values) == 0:
			c.err = fmt.Errorf("the Metastore filter on %q needs at least one value", field)
		case slices.ContainsFunc(values, func(v string) bool { return strings.Contains(v, "|") }):
			c.err = fmt.Errorf(`a value of the Metastore filter on %q contains "|"`, field)
		case len(values) == 1:
			c.addFilter(field, "", values[0])
		default:
			c.addFilter(field, "[in]", strings.Join(values, "|"))
		}
	}
}

// ListMetastoreObjectsRequest returns a request to list the objects of a type from the Metastore repository, one page
// of at most the limit. The offset paging has no snapshot: a caller that reads several pages must expect rows to shift
// between them.
// See GET /api/v1/repository/{objectType} of the Metastore.
func (a *AuthorizedAPI) ListMetastoreObjectsRequest(
	objectType string,
	opts ...ListMetastoreObjectsOption,
) request.APIRequest[*MetastoreObjectList] {
	cfg := listMetastoreObjectsConfig{}
	for _, opt := range opts {
		opt(&cfg)
	}

	result := &MetastoreObjectList{}
	if err := validateMetastoreName("object type", objectType); err != nil {
		return request.NewAPIRequest(result, request.NewReqDefinitionError(err))
	}
	if cfg.err != nil {
		return request.NewAPIRequest(result, request.NewReqDefinitionError(cfg.err))
	}

	req := a.
		newRequest(MetastoreAPI).
		WithGet(MetastoreAPIObjects).
		AndPathParam("objectType", objectType).
		AndHeader("Accept", metastoreAccept)
	for _, p := range cfg.query {
		req = req.AndQueryParam(p.key, p.value)
	}
	req = req.WithResult(result).WithOnSuccess(keepMetastoreHeader(&result.Header))
	return request.NewAPIRequest(result, req)
}

// GetMetastoreObjectRequest returns a request to read one object of the Metastore repository by its id.
// See GET /api/v1/repository/{objectType}/{UUID} of the Metastore.
func (a *AuthorizedAPI) GetMetastoreObjectRequest(objectType, objectID string) request.APIRequest[*MetastoreObjectResponse] {
	result := &MetastoreObjectResponse{}
	if err := validateMetastoreName("object type", objectType); err != nil {
		return request.NewAPIRequest(result, request.NewReqDefinitionError(err))
	}
	if err := validateMetastoreName("object id", objectID); err != nil {
		return request.NewAPIRequest(result, request.NewReqDefinitionError(err))
	}

	req := a.
		newRequest(MetastoreAPI).
		WithGet(MetastoreAPIObject).
		AndPathParam("objectType", objectType).
		AndPathParam("objectId", objectID).
		AndHeader("Accept", metastoreAccept).
		WithResult(result).
		WithOnSuccess(keepMetastoreHeader(&result.Header))
	return request.NewAPIRequest(result, req)
}

// keepMetastoreHeader stores the headers of the answer where the result keeps them.
func keepMetastoreHeader(dst *http.Header) func(context.Context, request.HTTPResponse) error {
	return func(_ context.Context, response request.HTTPResponse) error {
		*dst = response.ResponseHeader()
		return nil
	}
}

// validateMetastoreName checks a value that becomes a segment of the request path or the name of a filter. "." and ".."
// match the pattern, but a client that resolves the URL would turn them into another endpoint of the service, with the
// caller's token.
func validateMetastoreName(kind, value string) error {
	if !metastoreNamePattern.MatchString(value) || value == "." || value == ".." {
		return fmt.Errorf("invalid Metastore %s %q", kind, value)
	}
	return nil
}
