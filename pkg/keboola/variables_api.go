package keboola

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

func (a *AuthorizedAPI) CreateVariableRequest(payload *VariableCreatePayload) request.APIRequest[*Variable] {
	result := &Variable{}
	req := a.newRequest(VariablesAPI).
		WithResult(result).
		WithMethod(http.MethodPost).
		WithURL(VariablesAPIVariables).
		WithJSONBody(payload)
	return request.NewAPIRequest(result, req)
}

func (a *AuthorizedAPI) DeleteVariableRequest(hash VariableHash) request.APIRequest[request.NoResult] {
	req := a.newRequest(VariablesAPI).
		WithMethod(http.MethodDelete).
		WithURL(VariablesAPIVariable).
		AndPathParam("hash", hash.String())
	return request.NewAPIRequest(request.NoResult{}, req)
}

func (a *AuthorizedAPI) ListVariablesRequest(opts *VariableListOptions) request.APIRequest[*[]*Variable] {
	result := make([]*Variable, 0)
	req := a.newRequest(VariablesAPI).
		WithResult(&result).
		WithMethod(http.MethodGet).
		WithURL(VariablesAPIVariables)

	if opts != nil {
		if opts.Key != "" {
			req = req.AndQueryParam("key", opts.Key)
		}
		if opts.Attributes != nil && len(opts.Attributes) > 0 {
			for k, v := range opts.Attributes {
				req = req.AndQueryParam("attributes["+k+"]", toString(v))
			}
		}
		if opts.Offset > 0 {
			req = req.AndQueryParam("offset", strconv.Itoa(opts.Offset))
		}
		if opts.Limit > 0 {
			req = req.AndQueryParam("limit", strconv.Itoa(opts.Limit))
		}
	}

	return request.NewAPIRequest(&result, req)
}

func (a *AuthorizedAPI) ListVariablesScopedRequest(branchID BranchID, opts *VariableListOptions) request.APIRequest[*[]*Variable] {
	result := make([]*Variable, 0)
	req := a.newRequest(VariablesAPI).
		WithResult(&result).
		WithMethod(http.MethodGet).
		WithURL(VariablesAPIScopedBranch).
		AndPathParam("branchId", branchID.String())

	if opts != nil {
		if opts.Key != "" {
			req = req.AndQueryParam("key", opts.Key)
		}
		if opts.Attributes != nil && len(opts.Attributes) > 0 {
			for k, v := range opts.Attributes {
				req = req.AndQueryParam("attributes["+k+"]", toString(v))
			}
		}
	}

	return request.NewAPIRequest(&result, req)
}

func toString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	default:
		return url.QueryEscape(string([]byte(v.(string))))
	}
}
