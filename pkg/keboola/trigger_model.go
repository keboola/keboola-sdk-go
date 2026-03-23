package keboola

import (
	"strconv"

	"github.com/relvacode/iso8601"
)

// TriggerID is an ID of a trigger in Storage API.
type TriggerID int

func (v TriggerID) String() string {
	return strconv.Itoa(int(v))
}

// TriggerKey is a unique identifier of a trigger.
type TriggerKey struct {
	ID TriggerID `json:"id"`
}

// Trigger https://keboola.docs.apiary.io/#reference/triggers/get-update-or-delete/get-detail
type Trigger struct {
	TriggerKey
	RunWithTokenID        string          `json:"runWithTokenId"`
	Component             ComponentID     `json:"component"`
	ConfigID              ConfigID        `json:"configurationId"`
	CoolDownPeriodMinutes int             `json:"coolDownPeriodMinutes"`
	TableIDs              []TableID       `json:"tableIds"`
	LastRun               *TriggerLastRun `json:"lastRun,omitempty"`
	CreatedBy             *TriggerToken   `json:"createdBy,omitempty"`
}

// TriggerLastRun contains information about the last trigger execution.
type TriggerLastRun struct {
	Date  iso8601.Time `json:"date"`
	JobID string       `json:"jobId"`
}

// TriggerToken contains information about the token that created a trigger.
type TriggerToken struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}
