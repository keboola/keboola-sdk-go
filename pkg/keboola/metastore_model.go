package keboola

import (
	"encoding/json"
	"net/http"
	"time"
)

// MetastoreObject is one object of the Metastore repository as the API serves it (JSON:API resource).
//
// Attributes are the object's data, kept raw: their shape is defined by the JSON schema of the object type, and decoding
// numbers into float64 would change a value above 2^53. A caller decodes them into its own type.
type MetastoreObject struct {
	ID         string              `json:"id"`
	Type       string              `json:"type"`
	Attributes json.RawMessage     `json:"attributes"`
	Meta       MetastoreObjectMeta `json:"meta"`
}

// MetastoreObjectMeta is the metadata of a Metastore object.
type MetastoreObjectMeta struct {
	Name          string     `json:"name,omitempty"`
	Branch        string     `json:"branch,omitempty"`
	SchemaVersion string     `json:"schemaVersion,omitempty"`
	Scope         string     `json:"scope,omitempty"`
	Revision      int        `json:"revision,omitempty"`
	ProjectID     int        `json:"projectId,omitempty"`
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	LastUpdated   *time.Time `json:"lastUpdated,omitempty"`
	DeletedAt     *time.Time `json:"deletedAt,omitempty"`
}

// MetastoreObjectList is the answer to a list of Metastore objects.
//
// Data is nil when the answer has no "data" key, an empty non-nil slice for an empty list: a caller that must tell a
// broken answer from an empty list checks for nil.
type MetastoreObjectList struct {
	Data []*MetastoreObject `json:"data"`
	// StatusCode and Header are those of the HTTP answer. They are set whenever an answer arrived, also when the request
	// returned an error (a refusal, an answer that could not be decoded), so a caller can tell a refusal from a failure to
	// get an answer: StatusCode is 0 when there was none. They are not part of the JSON.
	StatusCode int         `json:"-"`
	Header     http.Header `json:"-"`
}

// MetastoreObjectResponse is the answer to a read of one Metastore object.
type MetastoreObjectResponse struct {
	Data *MetastoreObject `json:"data"`
	// StatusCode and Header are those of the HTTP answer, see MetastoreObjectList.
	StatusCode int         `json:"-"`
	Header     http.Header `json:"-"`
}
