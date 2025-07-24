package keboola

// The file contains request definitions for the Encryption API.
// Requests can be sent by any HTTP client that implements the client.Sender interface.

import (
	"net/http"

	"github.com/spf13/cast"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

// EncryptRequest - https://keboolaencryption.docs.apiary.io/#reference/encrypt/encryption/encrypt-data
func (a *PublicAPI) EncryptRequest(
	projectID int,
	componentID *ComponentID,
	configID *ConfigID,
	branchType *BranchType,
	data map[string]string,
) request.APIRequest[*map[string]string] {
	result := make(map[string]string)
	req := a.newRequest(EncryptionAPI).
		WithResult(&result).
		WithMethod(http.MethodPost).
		WithURL(EncryptionAPIEncrypt).
		AndQueryParam("projectId", cast.ToString(projectID)).
		WithJSONBody(data)
	if componentID != nil {
		req = req.AndQueryParam("componentId", componentID.String())
	}
	if configID != nil {
		req = req.AndQueryParam("configId", configID.String())
	}
	if branchType != nil {
		req = req.AndQueryParam("branchType", string(*branchType))
	}

	return request.NewAPIRequest(&result, req)
}
