package keboola_test

import (
	"context"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
)

func TestGenerateNewId(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := APIClientForRandomProject(t, ctx)

	ticket, err := api.GenerateIDRequest().Send(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, ticket)
	assert.NotEmpty(t, ticket.ID)
}

func TestTicketProvider(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	_, api := APIClientForRandomProject(t, ctx)
	provider := NewTicketProvider(ctx, api)
	tickets := make([]string, 0)

	// Request 3 tickets
	for range 3 {
		provider.Request(func(ticket *Ticket) {
			tickets = append(tickets, ticket.ID)
		})
	}

	// Get tickets
	assert.NoError(t, provider.Resolve())

	// Assert order
	expected := make([]string, len(tickets))
	copy(expected, tickets)
	sort.Strings(expected)
	assert.Equal(t, expected, tickets)
}
