package keboola

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

func (a *AuthorizedAPI) CreateVariableRequest(payload *VaultVariableCreatePayload) request.APIRequest[*VaultVariable] {
	result := &VaultVariable{}
	req := a.newRequest(VaultVariablesAPI).
		WithResult(result).
		WithMethod(http.MethodPost).
		WithURL(VaultAPIVariables).
		WithJSONBody(payload)
	return request.NewAPIRequest(result, req)
}

func (a *AuthorizedAPI) DeleteVariableRequest(hash VaultVariableHash) request.APIRequest[request.NoResult] {
	req := a.newRequest(VaultVariablesAPI).
		WithMethod(http.MethodDelete).
		WithURL(VaultAPIVariable).
		AndPathParam("hash", hash.String())
	return request.NewAPIRequest(request.NoResult{}, req)
}

func (a *AuthorizedAPI) ListVariablesRequest(opts *VaultVariableListOptions) request.APIRequest[*[]*VaultVariable] {
	result := make([]*VaultVariable, 0)
	req := a.newRequest(VaultVariablesAPI).
		WithResult(&result).
		WithMethod(http.MethodGet).
		WithURL(VaultAPIVariables)

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

func (a *AuthorizedAPI) ListVariablesScopedRequest(branchID BranchID, opts *VaultVariableListOptions) request.APIRequest[*[]*VaultVariable] {
	result := make([]*VaultVariable, 0)
	req := a.newRequest(VaultVariablesAPI).
		WithResult(&result).
		WithMethod(http.MethodGet).
		WithURL(VaultAPIScopedBranch).
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
		if opts.Offset > 0 {
			req = req.AndQueryParam("offset", strconv.Itoa(opts.Offset))
		}
		if opts.Limit > 0 {
			req = req.AndQueryParam("limit", strconv.Itoa(opts.Limit))
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
		return fmt.Sprint(v)
	}
}
