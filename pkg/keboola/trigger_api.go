package keboola

// The file contains request definitions for triggers in the Storage API.
// Requests can be sent by any HTTP client that implements the client.Sender interface.

import (
	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

// ListTriggersRequest https://keboola.docs.apiary.io/#reference/triggers/create/list
func (a *AuthorizedAPI) ListTriggersRequest() request.APIRequest[*[]*Trigger] {
	result := make([]*Trigger, 0)
	req := a.newRequest(StorageAPI).
		WithResult(&result).
		WithGet(StorageAPITriggers)
	return request.NewAPIRequest(&result, req)
}

// GetTriggerRequest https://keboola.docs.apiary.io/#reference/triggers/get-update-or-delete/get-detail
func (a *AuthorizedAPI) GetTriggerRequest(key TriggerKey) request.APIRequest[*Trigger] {
	result := &Trigger{}
	req := a.newRequest(StorageAPI).
		WithResult(result).
		WithGet(StorageAPITrigger).
		AndPathParam("triggerId", key.ID.String())
	return request.NewAPIRequest(result, req)
}

// CreateTriggerRequest https://keboola.docs.apiary.io/#reference/triggers/create/create
func (a *AuthorizedAPI) CreateTriggerRequest(trigger *Trigger) request.APIRequest[*Trigger] {
	req := a.newRequest(StorageAPI).
		WithResult(trigger).
		WithPost(StorageAPITriggers).
		WithJSONBody(triggerWriteBody(trigger))
	return request.NewAPIRequest(trigger, req)
}

// UpdateTriggerRequest https://keboola.docs.apiary.io/#reference/triggers/get-update-or-delete/update
func (a *AuthorizedAPI) UpdateTriggerRequest(trigger *Trigger) request.APIRequest[*Trigger] {
	req := a.newRequest(StorageAPI).
		WithResult(trigger).
		WithPut(StorageAPITrigger).
		AndPathParam("triggerId", trigger.ID.String()).
		WithJSONBody(triggerWriteBody(trigger))
	return request.NewAPIRequest(trigger, req)
}

// DeleteTriggerRequest https://keboola.docs.apiary.io/#reference/triggers/get-update-or-delete/delete
func (a *AuthorizedAPI) DeleteTriggerRequest(key TriggerKey) request.APIRequest[request.NoResult] {
	req := a.newRequest(StorageAPI).
		WithDelete(StorageAPITrigger).
		AndPathParam("triggerId", key.ID.String())
	return request.NewAPIRequest(request.NoResult{}, req)
}

// triggerWriteBody constructs the JSON body for create and update trigger requests.
func triggerWriteBody(t *Trigger) map[string]any {
	tableIDs := make([]string, len(t.TableIDs))
	for i, id := range t.TableIDs {
		tableIDs[i] = id.String()
	}
	return map[string]any{
		"runWithTokenId":        t.RunWithTokenID,
		"component":             t.Component.String(),
		"configurationId":       t.ConfigID.String(),
		"coolDownPeriodMinutes": t.CoolDownPeriodMinutes,
		"tableIds":              tableIDs,
	}
}
