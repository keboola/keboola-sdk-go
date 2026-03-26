package keboola

import (
	"fmt"

	"github.com/relvacode/iso8601"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

// DataScienceAppID is the unique identifier of a data science app.
type DataScienceAppID string

func (v DataScienceAppID) String() string {
	return string(v)
}

// DataScienceAppType is the type of a data science app.
// Snowflake and BigQuery are intentionally omitted — use editor service for those.
type DataScienceAppType string

const (
	DataScienceAppTypePython    DataScienceAppType = "python"
	DataScienceAppTypeR         DataScienceAppType = "r"
	DataScienceAppTypeStreamlit DataScienceAppType = "streamlit"
	DataScienceAppTypePythonJS  DataScienceAppType = "python-js"
)

// DataScienceAppState is the current lifecycle state of a data science app.
type DataScienceAppState string

// DataScienceAppDesiredState is the desired lifecycle state of a data science app.
type DataScienceAppDesiredState string

// DataScienceAppProvisioningStrategy is the provisioning strategy of a data science app.
type DataScienceAppProvisioningStrategy string

// DataScienceApp represents a data science app returned by the data-science service /apps endpoint.
// It replaces the deprecated SandboxWorkspace for listing and status checks.
// Field mapping from old SandboxWorkspace: ID→ID, Type→Type (narrowed), Active(bool)→State+DesiredState,
// Size→Size, URL→URL, BranchID→BranchID, ConfigurationID→ConfigID.
// Connection details (User/Password/Host/Credentials) are not exposed here.
type DataScienceApp struct {
	ID                      DataScienceAppID                   `json:"id"`
	Type                    DataScienceAppType                 `json:"type"`
	ProjectID               string                             `json:"projectId"`
	ComponentID             ComponentID                        `json:"componentId"`
	BranchID                string                             `json:"branchId"`
	ConfigID                string                             `json:"configId"`
	ConfigVersion           string                             `json:"configVersion"`
	State                   DataScienceAppState                `json:"state"`
	DesiredState            DataScienceAppDesiredState         `json:"desiredState"`
	LastRequestTimestamp    iso8601.Time                       `json:"lastRequestTimestamp"`
	LastStartTimestamp      iso8601.Time                       `json:"lastStartTimestamp"`
	URL                     string                             `json:"url"`
	AutoSuspendAfterSeconds int                                `json:"autoSuspendAfterSeconds"`
	AutoRestartEnabled      bool                               `json:"autoRestartEnabled"`
	Size                    string                             `json:"size"`
	ProvisioningStrategy    DataScienceAppProvisioningStrategy `json:"provisioningStrategy"`
}

type listDataScienceAppsConfig struct {
	Offset      int
	Limit       int
	ComponentID ComponentID
	Types       []DataScienceAppType
	BranchID    string
}

// ListDataScienceAppsOption is a functional option for ListDataScienceAppsRequest.
type ListDataScienceAppsOption func(*listDataScienceAppsConfig)

// WithDataScienceAppsOffset sets the offset for listing data science apps.
func WithDataScienceAppsOffset(v int) ListDataScienceAppsOption {
	return func(c *listDataScienceAppsConfig) { c.Offset = v }
}

// WithDataScienceAppsLimit sets the maximum number of apps to return (1–500, default 100).
func WithDataScienceAppsLimit(v int) ListDataScienceAppsOption {
	return func(c *listDataScienceAppsConfig) { c.Limit = v }
}

// WithDataScienceAppsComponentID filters apps by component ID.
func WithDataScienceAppsComponentID(v ComponentID) ListDataScienceAppsOption {
	return func(c *listDataScienceAppsConfig) { c.ComponentID = v }
}

// WithDataScienceAppsType filters apps by type (can be specified multiple times).
func WithDataScienceAppsType(v ...DataScienceAppType) ListDataScienceAppsOption {
	return func(c *listDataScienceAppsConfig) { c.Types = append(c.Types, v...) }
}

// WithDataScienceAppsBranchID filters apps by branch ID.
func WithDataScienceAppsBranchID(v string) ListDataScienceAppsOption {
	return func(c *listDataScienceAppsConfig) { c.BranchID = v }
}

// ListDataScienceAppsRequest returns a list of data science apps from the data-science service.
// https://api.keboola.com/?service=sandboxes-service#get-/apps
func (a *AuthorizedAPI) ListDataScienceAppsRequest(opts ...ListDataScienceAppsOption) request.APIRequest[*[]*DataScienceApp] {
	cfg := &listDataScienceAppsConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	result := make([]*DataScienceApp, 0)
	req := a.newRequest(DataScienceAPI).
		WithResult(&result).
		WithGet(DataScienceAPIApps)

	if cfg.Offset > 0 {
		req = req.AndQueryParam("offset", fmt.Sprintf("%d", cfg.Offset))
	}
	if cfg.Limit > 0 {
		req = req.AndQueryParam("limit", fmt.Sprintf("%d", cfg.Limit))
	}
	if cfg.ComponentID != "" {
		req = req.AndQueryParam("componentId", cfg.ComponentID.String())
	}
	if len(cfg.Types) > 0 {
		typeStrs := make([]string, len(cfg.Types))
		for i, t := range cfg.Types {
			typeStrs[i] = string(t)
		}
		req = req.AndQueryParamValues("type[]", typeStrs)
	}
	if cfg.BranchID != "" {
		req = req.AndQueryParam("branchId", cfg.BranchID)
	}

	return request.NewAPIRequest(&result, req)
}

// GetDataScienceAppRequest returns a single data science app by ID from the data-science service.
// https://api.keboola.com/?service=sandboxes-service#get-/apps/-appId-
func (a *AuthorizedAPI) GetDataScienceAppRequest(id DataScienceAppID) request.APIRequest[*DataScienceApp] {
	result := &DataScienceApp{}
	req := a.newRequest(DataScienceAPI).
		WithResult(result).
		WithGet(DataScienceAPIApp).
		AndPathParam("appId", id.String())
	return request.NewAPIRequest(result, req)
}
