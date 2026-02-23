package keboola

import (
	"encoding/json"
	"fmt"
	"time"
)

// NotificationSubscriptionID represents a unique identifier for a notification subscription.
type NotificationSubscriptionID string

// NotificationSubscriptionKey identifies a notification subscription.
type NotificationSubscriptionKey struct {
	ID NotificationSubscriptionID `json:"id" validate:"required"`
}

// NotificationEvent represents the type of event that triggers a notification.
type NotificationEvent string

const (
	// Job-level events.
	NotificationEventJobFailed               NotificationEvent = "job-failed"
	NotificationEventJobSucceeded            NotificationEvent = "job-succeeded"
	NotificationEventJobSucceededWithWarning NotificationEvent = "job-succeeded-with-warning"
	NotificationEventJobProcessingLong       NotificationEvent = "job-processing-long"

	// Phase-level events.
	NotificationEventPhaseJobFailed               NotificationEvent = "phase-job-failed"
	NotificationEventPhaseJobSucceeded            NotificationEvent = "phase-job-succeeded"
	NotificationEventPhaseJobSucceededWithWarning NotificationEvent = "phase-job-succeeded-with-warning"
	NotificationEventPhaseJobProcessingLong       NotificationEvent = "phase-job-processing-long"
)

// NotificationFilterOperator represents the comparison operator for a filter.
type NotificationFilterOperator string

const (
	NotificationFilterOperatorEquals            NotificationFilterOperator = "=="
	NotificationFilterOperatorNotEquals         NotificationFilterOperator = "!="
	NotificationFilterOperatorGreaterThan       NotificationFilterOperator = ">"
	NotificationFilterOperatorLessThan          NotificationFilterOperator = "<"
	NotificationFilterOperatorGreaterThanEquals NotificationFilterOperator = ">="
	NotificationFilterOperatorLessThanEquals    NotificationFilterOperator = "<="
)

// NotificationFilter defines a condition for when a notification should be triggered.
type NotificationFilter struct {
	Field    string                     `json:"field"`
	Value    string                     `json:"value"`
	Operator NotificationFilterOperator `json:"operator,omitempty"` // defaults to "=="
}

// NotificationChannel represents the delivery method for a notification.
type NotificationChannel string

const (
	NotificationChannelEmail   NotificationChannel = "email"
	NotificationChannelWebhook NotificationChannel = "webhook"
)

// NotificationRecipient defines where a notification should be delivered.
type NotificationRecipient struct {
	Channel NotificationChannel `json:"channel"`
	Address string              `json:"address"`
}

// NotificationExpiration represents when a subscription should expire.
// It can be either an absolute timestamp (RFC3339) or a relative duration (e.g., "+24 hours").
type NotificationExpiration struct {
	value string
}

// NewAbsoluteExpiration creates an expiration with an absolute timestamp.
func NewAbsoluteExpiration(t time.Time) *NotificationExpiration {
	return &NotificationExpiration{value: t.Format(time.RFC3339)}
}

// NewRelativeExpiration creates an expiration with a relative duration (e.g., "+24 hours").
func NewRelativeExpiration(duration string) *NotificationExpiration {
	return &NotificationExpiration{value: duration}
}

// MarshalJSON implements json.Marshaler.
func (e NotificationExpiration) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NotificationExpiration) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &e.value)
}

// String returns the expiration value as a string.
func (e NotificationExpiration) String() string {
	return e.value
}

// NotificationSubscription represents a notification subscription.
// See https://keboolaconnection.docs.apiary.io/#reference/notification-api
type NotificationSubscription struct {
	NotificationSubscriptionKey
	Event     NotificationEvent         `json:"event"`
	Filters   []NotificationFilter      `json:"filters,omitempty"`
	Recipient NotificationRecipient     `json:"recipient"`
	ExpiresAt *NotificationExpiration   `json:"expiresAt,omitempty"`
	CreatedAt *time.Time                `json:"createdAt,omitempty"`
}

// GroupNotificationSubscriptionsByEvent groups subscriptions by their event type.
// Returns a map where keys are NotificationEvent types and values are slices of subscription IDs.
func GroupNotificationSubscriptionsByEvent(subscriptions []*NotificationSubscription) map[NotificationEvent][]NotificationSubscriptionID {
	grouped := make(map[NotificationEvent][]NotificationSubscriptionID)
	for _, sub := range subscriptions {
		grouped[sub.Event] = append(grouped[sub.Event], sub.ID)
	}
	return grouped
}

// createSubscriptionRequest represents the payload for creating a notification subscription.
type createSubscriptionRequest struct {
	Event     NotificationEvent         `json:"event"`
	Filters   []NotificationFilter      `json:"filters"`
	Recipient NotificationRecipient     `json:"recipient"`
	ExpiresAt *NotificationExpiration   `json:"expiresAt,omitempty"`
}

// Validate ensures the filter operator is valid, defaulting to "==" if empty.
func (f *NotificationFilter) Validate() error {
	if f.Operator == "" {
		f.Operator = NotificationFilterOperatorEquals
	}

	validOperators := map[NotificationFilterOperator]bool{
		NotificationFilterOperatorEquals:            true,
		NotificationFilterOperatorNotEquals:         true,
		NotificationFilterOperatorGreaterThan:       true,
		NotificationFilterOperatorLessThan:          true,
		NotificationFilterOperatorGreaterThanEquals: true,
		NotificationFilterOperatorLessThanEquals:    true,
	}

	if !validOperators[f.Operator] {
		return fmt.Errorf(`invalid filter operator "%s"`, f.Operator)
	}

	return nil
}
