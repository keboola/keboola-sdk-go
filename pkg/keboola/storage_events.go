package keboola

import (
	jsonLib "encoding/json"
	"fmt"

	"github.com/keboola/keboola-sdk-go/v2/pkg/client"
	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

// EventID represents an ID of an event in Storage API.
type EventID string

func (v EventID) String() string {
	return string(v)
}

// Event https://keboola.docs.apiary.io/#reference/events/events/create-event
type Event struct {
	ID          EventID                `json:"uuid" readonly:"true"`
	ComponentID ComponentID            `json:"component"`
	Message     string                 `json:"message"`
	Type        string                 `json:"type"`
	Duration    client.DurationSeconds `json:"duration"`
	Params      JSONString             `json:"params"`
	Results     JSONString             `json:"results"`
	RunID       string                 `json:"runId"`
}

// CreateEventRequest https://keboola.docs.apiary.io/#reference/events/events/create-event
func (a *AuthorizedAPI) CreateEventRequest(event *Event) request.APIRequest[*Event] {
	// Params and results must be a JSON value encoded as string
	body := request.StructToMap(event, nil)
	pValue, err := jsonLib.Marshal(event.Params)
	if err != nil {
		return request.NewAPIRequest(event, request.NewReqDefinitionError(err))
	}
	rValue, err := jsonLib.Marshal(event.Results)
	if err != nil {
		return request.NewAPIRequest(event, request.NewReqDefinitionError(err))
	}
	body["params"] = string(pValue)
	body["results"] = string(rValue)
	req := a.
		newRequest(StorageAPI).
		WithResult(event).
		WithPost("events").
		WithJSONBody(body)
	return request.NewAPIRequest(event, req)
}

// JSONString is Json encoded as string, see CreateEventRequest.
type JSONString map[string]any

// UnmarshalJSON implements JSON decoding.
func (v *JSONString) UnmarshalJSON(data []byte) (err error) {
	out := make(map[string]any)
	err = jsonLib.Unmarshal(data, &out)
	*v = out
	return err
}

// MarshalJSON implements JSON encoding.
func (v *JSONString) MarshalJSON() ([]byte, error) {
	return jsonLib.Marshal(map[string]any(*v))
}

func (v JSONString) String() string {
	bytes, err := jsonLib.Marshal(map[string]any(v))
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// TableEvent represents an event from the table events API.
// This is different from Event which is used for creating events.
// https://keboola.docs.apiary.io/#reference/events/events/list-events
type TableEvent struct {
	ID        EventID           `json:"uuid"`
	Event     string            `json:"event"`
	Component string            `json:"component"`
	Message   string            `json:"message"`
	RunID     string            `json:"runId"`
	Created   string            `json:"created"`
	Context   TableEventContext `json:"context"`
}

// TableEventContext contains context information for a table event.
type TableEventContext struct {
	UserAgent string `json:"userAgent"`
}

// listTableEventsConfig holds configuration for listing table events.
type listTableEventsConfig struct {
	limit int
}

// ListTableEventsOption is an option for ListTableEventsRequest.
type ListTableEventsOption func(c *listTableEventsConfig)

// WithTableEventsLimit sets the maximum number of events to return.
// If not called, defaults to 50. Setting limit to 0 or negative
// will omit the limit parameter from the request, using server default.
func WithTableEventsLimit(limit int) ListTableEventsOption {
	return func(c *listTableEventsConfig) {
		c.limit = limit
	}
}

// ListTableEventsRequest lists events for a specific table.
// https://keboola.docs.apiary.io/#reference/events/events/list-events
func (a *AuthorizedAPI) ListTableEventsRequest(tableID TableID, opts ...ListTableEventsOption) request.APIRequest[*[]*TableEvent] {
	config := listTableEventsConfig{limit: 50} // default limit
	for _, opt := range opts {
		opt(&config)
	}

	result := make([]*TableEvent, 0)
	req := a.
		newRequest(StorageAPI).
		WithResult(&result).
		WithGet(fmt.Sprintf("tables/%s/events", tableID.String()))

	if config.limit > 0 {
		req = req.AndQueryParam("limit", fmt.Sprintf("%d", config.limit))
	}

	return request.NewAPIRequest(&result, req)
}
