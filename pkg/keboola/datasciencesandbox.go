package keboola

import (
	"github.com/relvacode/iso8601"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

// DataScienceSandboxSize is the compute size of a Python/R sandbox.
type DataScienceSandboxSize string

const (
	DataScienceSandboxSizeMicro  DataScienceSandboxSize = "micro"
	DataScienceSandboxSizeSmall  DataScienceSandboxSize = "small"
	DataScienceSandboxSizeMedium DataScienceSandboxSize = "medium"
	DataScienceSandboxSizeLarge  DataScienceSandboxSize = "large"
)

// DataScienceSandboxSizesOrdered returns sandbox sizes in ascending order.
func DataScienceSandboxSizesOrdered() []DataScienceSandboxSize {
	return []DataScienceSandboxSize{
		DataScienceSandboxSizeMicro,
		DataScienceSandboxSizeSmall,
		DataScienceSandboxSizeMedium,
		DataScienceSandboxSizeLarge,
	}
}

// DataScienceSandboxSizesMap returns the set of valid sandbox sizes.
func DataScienceSandboxSizesMap() map[DataScienceSandboxSize]bool {
	return map[DataScienceSandboxSize]bool{
		DataScienceSandboxSizeMicro:  true,
		DataScienceSandboxSizeSmall:  true,
		DataScienceSandboxSizeMedium: true,
		DataScienceSandboxSizeLarge:  true,
	}
}

// DataScienceSandboxSupportsSizes reports whether the given app type supports compute size selection.
// Only Python and R sandboxes support sizes; Streamlit and PythonJS do not.
func DataScienceSandboxSupportsSizes(typ DataScienceAppType) bool {
	return typ == DataScienceAppTypePython || typ == DataScienceAppTypeR
}

// DataScienceSandboxPersistentStorage holds PVC details for a sandbox.
type DataScienceSandboxPersistentStorage struct {
	PVCName     string `json:"pvcName"`
	K8sManifest string `json:"k8sManifest"`
}

// DataScienceSandbox represents a Python/R sandbox returned by the data-science service /sandboxes endpoint.
type DataScienceSandbox struct {
	ID                      string                              `json:"id"`
	ProjectID               string                              `json:"projectId"`
	BranchID                string                              `json:"branchId"`
	ConfigurationID         string                              `json:"configurationId"`
	Type                    DataScienceAppType                  `json:"type"`
	Active                  bool                                `json:"active"`
	CreatedTimestamp        iso8601.Time                        `json:"createdTimestamp"`
	StartTimestamp          iso8601.Time                        `json:"startTimestamp"`
	UpdatedTimestamp        iso8601.Time                        `json:"updatedTimestamp"`
	PersistentStorage       DataScienceSandboxPersistentStorage `json:"persistentStorage"`
	Shared                  bool                                `json:"shared"`
	ComponentID             string                              `json:"componentId"`
	ConfigurationVersion    string                              `json:"configurationVersion"`
	Size                    DataScienceSandboxSize              `json:"size"`
	AutoSuspendAfterSeconds int                                 `json:"autoSuspendAfterSeconds"`
	URL                     string                              `json:"url"`
	Password                string                              `json:"password"` //nolint: gosec
	ImageVersion            string                              `json:"imageVersion"`
	Packages                []string                            `json:"packages"`
}

// CreateDataScienceSandboxPayload is the request body for POST /sandboxes.
type CreateDataScienceSandboxPayload struct {
	ComponentID             string                 `json:"componentId,omitempty"`
	ConfigurationID         string                 `json:"configurationId,omitempty"`
	ConfigurationVersion    string                 `json:"configurationVersion,omitempty"`
	Type                    DataScienceAppType     `json:"type"`
	PhysicalID              string                 `json:"physicalId,omitempty"`
	Host                    string                 `json:"host,omitempty"`
	Size                    DataScienceSandboxSize `json:"size,omitempty"`
	BranchID                string                 `json:"branchId,omitempty"`
	User                    string                 `json:"user,omitempty"`
	Password                string                 `json:"password,omitempty"` //nolint: gosec
	URL                     string                 `json:"url,omitempty"`
	Active                  bool                   `json:"active,omitempty"`
	Shared                  bool                   `json:"shared,omitempty"`
	ImageVersion            string                 `json:"imageVersion,omitempty"`
	AutoSuspendAfterSeconds int                    `json:"autoSuspendAfterSeconds,omitempty"`
}

// UpdateDataScienceSandboxPayload is the request body for PATCH /sandboxes/{appId}.
// Only set fields are serialised; omit a field to leave it unchanged.
type UpdateDataScienceSandboxPayload struct {
	Size                    DataScienceSandboxSize `json:"size,omitempty"`
	AutoSuspendAfterSeconds int                    `json:"autoSuspendAfterSeconds,omitempty"`
	Shared                  bool                   `json:"shared,omitempty"`
	ImageVersion            string                 `json:"imageVersion,omitempty"`
}

// CreateDataScienceSandboxRequest creates a new Python/R sandbox via POST /sandboxes.
// https://api.keboola.com/?service=sandboxes-service#post-/sandboxes
func (a *AuthorizedAPI) CreateDataScienceSandboxRequest(payload CreateDataScienceSandboxPayload) request.APIRequest[*DataScienceSandbox] {
	result := &DataScienceSandbox{}
	req := a.newRequest(DataScienceAPI).
		WithResult(result).
		WithPost(DataScienceAPISandboxes).
		WithJSONBody(payload)
	return request.NewAPIRequest(result, req)
}

// GetDataScienceSandboxRequest returns a single Python/R sandbox by ID via GET /sandboxes/{appId}.
// https://api.keboola.com/?service=sandboxes-service#get-/sandboxes/-appId-
func (a *AuthorizedAPI) GetDataScienceSandboxRequest(id DataScienceAppID) request.APIRequest[*DataScienceSandbox] {
	result := &DataScienceSandbox{}
	req := a.newRequest(DataScienceAPI).
		WithResult(result).
		WithGet(DataScienceAPISandbox).
		AndPathParam("appId", id.String())
	return request.NewAPIRequest(result, req)
}

// UpdateDataScienceSandboxRequest updates a Python/R sandbox via PATCH /sandboxes/{appId}.
// https://api.keboola.com/?service=sandboxes-service#patch-/sandboxes/-appId-
func (a *AuthorizedAPI) UpdateDataScienceSandboxRequest(id DataScienceAppID, payload UpdateDataScienceSandboxPayload) request.APIRequest[*DataScienceSandbox] {
	result := &DataScienceSandbox{}
	req := a.newRequest(DataScienceAPI).
		WithResult(result).
		WithPatch(DataScienceAPISandbox).
		AndPathParam("appId", id.String()).
		WithJSONBody(payload)
	return request.NewAPIRequest(result, req)
}

// DeleteDataScienceSandboxRequest deletes a Python/R sandbox via DELETE /sandboxes/{appId}.
// https://api.keboola.com/?service=sandboxes-service#delete-/sandboxes/-appId-
func (a *AuthorizedAPI) DeleteDataScienceSandboxRequest(id DataScienceAppID) request.APIRequest[request.NoResult] {
	req := a.newRequest(DataScienceAPI).
		WithDelete(DataScienceAPISandbox).
		AndPathParam("appId", id.String())
	return request.NewAPIRequest(request.NoResult{}, req)
}
