package keboola

import (
	"context"
	"net/url"
	"time"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

// CreateNotificationSubscriptionRequestBuilder builds a request to create a notification subscription.
// See https://keboolaconnection.docs.apiary.io/#reference/notification-api
type CreateNotificationSubscriptionRequestBuilder struct {
	payload createSubscriptionRequest
	api     *AuthorizedAPI
}

// NewCreateNotificationSubscriptionRequest creates a builder for creating a notification subscription.
// Required parameters: event (type of event to subscribe to), channel (delivery method), address (recipient address).
func (a *AuthorizedAPI) NewCreateNotificationSubscriptionRequest(
	event NotificationEvent,
	channel NotificationChannel,
	address string,
) *CreateNotificationSubscriptionRequestBuilder {
	return &CreateNotificationSubscriptionRequestBuilder{
		payload: createSubscriptionRequest{
			Event: event,
			Recipient: NotificationRecipient{
				Channel: channel,
				Address: address,
			},
			Filters: []NotificationFilter{}, // Initialize as empty slice to marshal as [] not null
		},
		api: a,
	}
}

// WithFilter adds a single filter to the subscription.
// If operator is empty, it defaults to "==".
func (b *CreateNotificationSubscriptionRequestBuilder) WithFilter(
	field, value string,
	operator NotificationFilterOperator,
) *CreateNotificationSubscriptionRequestBuilder {
	filter := NotificationFilter{
		Field:    field,
		Value:    value,
		Operator: operator,
	}
	if err := filter.Validate(); err != nil {
		panic(err)
	}
	b.payload.Filters = append(b.payload.Filters, filter)
	return b
}

// WithFilters adds multiple filters to the subscription.
// Empty operators default to "==".
func (b *CreateNotificationSubscriptionRequestBuilder) WithFilters(
	filters []NotificationFilter,
) *CreateNotificationSubscriptionRequestBuilder {
	for _, filter := range filters {
		if err := filter.Validate(); err != nil {
			panic(err)
		}
		b.payload.Filters = append(b.payload.Filters, filter)
	}
	return b
}

// WithAbsoluteExpiration sets an absolute expiration timestamp for the subscription.
func (b *CreateNotificationSubscriptionRequestBuilder) WithAbsoluteExpiration(
	t time.Time,
) *CreateNotificationSubscriptionRequestBuilder {
	b.payload.ExpiresAt = NewAbsoluteExpiration(t)
	return b
}

// WithRelativeExpiration sets a relative expiration duration for the subscription (e.g., "+24 hours").
func (b *CreateNotificationSubscriptionRequestBuilder) WithRelativeExpiration(
	duration string,
) *CreateNotificationSubscriptionRequestBuilder {
	b.payload.ExpiresAt = NewRelativeExpiration(duration)
	return b
}

// Build constructs the API request without sending it.
func (b *CreateNotificationSubscriptionRequestBuilder) Build() request.APIRequest[*NotificationSubscription] {
	result := &NotificationSubscription{}
	req := b.api.
		newRequest(NotificationAPI).
		WithPost(NotificationAPIProjectSubscriptions).
		WithJSONBody(b.payload).
		WithResult(result)
	return request.NewAPIRequest(result, req)
}

// Send constructs and sends the API request.
func (b *CreateNotificationSubscriptionRequestBuilder) Send(ctx context.Context) (*NotificationSubscription, error) {
	return b.Build().Send(ctx)
}

// ListNotificationSubscriptionsConfig holds optional configuration for listing subscriptions.
type ListNotificationSubscriptionsConfig struct {
	Event NotificationEvent
}

// ListNotificationSubscriptionsOption is a functional option for configuring list requests.
type ListNotificationSubscriptionsOption func(c *ListNotificationSubscriptionsConfig)

// WithListSubscriptionsEvent filters subscriptions by event type.
func WithListSubscriptionsEvent(event NotificationEvent) ListNotificationSubscriptionsOption {
	return func(c *ListNotificationSubscriptionsConfig) {
		c.Event = event
	}
}

// ListNotificationSubscriptionsRequest returns a request to list notification subscriptions.
// Optionally filter by event type using WithListSubscriptionsEvent.
// See https://keboolaconnection.docs.apiary.io/#reference/notification-api
func (a *AuthorizedAPI) ListNotificationSubscriptionsRequest(
	opts ...ListNotificationSubscriptionsOption,
) request.APIRequest[*[]*NotificationSubscription] {
	config := &ListNotificationSubscriptionsConfig{}
	for _, opt := range opts {
		opt(config)
	}

	result := &[]*NotificationSubscription{}
	req := a.
		newRequest(NotificationAPI).
		WithGet(NotificationAPIProjectSubscriptions).
		WithResult(result)

	// Add optional event filter
	if config.Event != "" {
		req = req.AndQueryParam("event", string(config.Event))
	}

	return request.NewAPIRequest(result, req)
}

// GetNotificationSubscriptionRequest returns a request to get a single notification subscription.
// See https://keboolaconnection.docs.apiary.io/#reference/notification-api
func (a *AuthorizedAPI) GetNotificationSubscriptionRequest(
	key NotificationSubscriptionKey,
) request.APIRequest[*NotificationSubscription] {
	result := &NotificationSubscription{}
	req := a.
		newRequest(NotificationAPI).
		WithGet(NotificationAPIProjectSubscription).
		AndPathParam("subscriptionId", url.PathEscape(string(key.ID))).
		WithResult(result)
	return request.NewAPIRequest(result, req)
}

// DeleteNotificationSubscriptionRequest returns a request to delete a notification subscription.
// See https://keboolaconnection.docs.apiary.io/#reference/notification-api
func (a *AuthorizedAPI) DeleteNotificationSubscriptionRequest(
	key NotificationSubscriptionKey,
) request.APIRequest[request.NoResult] {
	req := a.
		newRequest(NotificationAPI).
		WithDelete(NotificationAPIProjectSubscription).
		AndPathParam("subscriptionId", url.PathEscape(string(key.ID)))
	return request.NewAPIRequest(request.NoResult{}, req)
}
