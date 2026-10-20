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
	ID            string                           `json:"id"`
	Type          string                           `json:"type"`
	Attributes    json.RawMessage                  `json:"attributes"`
	Meta          MetastoreObjectMeta              `json:"meta"`
	Relationships map[string]MetastoreRelationship `json:"relationships,omitempty"`
}

// MetastoreObjectMeta is the metadata of a Metastore object.
type MetastoreObjectMeta struct {
	Name                      string     `json:"name,omitempty"`
	Branch                    string     `json:"branch,omitempty"`
	SchemaVersion             string     `json:"schemaVersion,omitempty"`
	Scope                     string     `json:"scope,omitempty"`
	Revision                  int        `json:"revision,omitempty"`
	OrganizationID            string     `json:"organizationId,omitempty"`
	ProjectID                 int        `json:"projectId,omitempty"`
	SourceProjectID           int        `json:"sourceProjectId,omitempty"`
	CreatedAt                 *time.Time `json:"createdAt,omitempty"`
	LastUpdated               *time.Time `json:"lastUpdated,omitempty"`
	RevisionCreatedAt         *time.Time `json:"revisionCreatedAt,omitempty"`
	DeletedAt                 *time.Time `json:"deletedAt,omitempty"`
	AutomigratedAt            *time.Time `json:"automigratedAt,omitempty"`
	ScopeElevationRequestedAt *time.Time `json:"scopeElevationRequestedAt,omitempty"`
}

// MetastoreRelationship is a reference of an object to another object.
type MetastoreRelationship struct {
	Data MetastoreRelationshipData `json:"data"`
	Meta MetastoreRelationshipMeta `json:"meta"`
}

// MetastoreRelationshipData identifies the referred object.
type MetastoreRelationshipData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// MetastoreRelationshipMeta says which attribute of the object holds the reference.
type MetastoreRelationshipMeta struct {
	SourcePath string `json:"sourcePath,omitempty"`
}

// MetastoreObjectList is the answer to a list of Metastore objects.
//
// Data is nil when the answer has no "data" key, an empty non-nil slice for an empty list: a caller that must tell a
// broken answer from an empty list checks for nil.
type MetastoreObjectList struct {
	Data     []*MetastoreObject `json:"data"`
	Included []*MetastoreObject `json:"included,omitempty"`
	// Header is the headers of the HTTP answer, for a caller that relies on one the Metastore sets (for example the
	// marker of a verified security-policy read). It is not part of the JSON.
	Header http.Header `json:"-"`
}

// MetastoreObjectResponse is the answer to a read of one Metastore object.
type MetastoreObjectResponse struct {
	Data     *MetastoreObject   `json:"data"`
	Included []*MetastoreObject `json:"included,omitempty"`
	// Header is the headers of the HTTP answer, see MetastoreObjectList.Header.
	Header http.Header `json:"-"`
}
