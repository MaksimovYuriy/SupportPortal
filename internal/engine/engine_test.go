package engine

import (
	"context"
	"io"
	"log/slog"
	"reflect"
	"testing"
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

func TestEngineTickStartsNewTickets(t *testing.T) {
	tickets := &fakeTicketService{
		newTickets: []*models.Ticket{
			{ID: 1},
			{ID: 2},
		},
	}
	agents := &fakeAgentService{}
	engine := NewEngine(tickets, agents, discardLogger(), time.Second, 100)

	engine.tick(context.Background())

	expected := []int64{1, 2}
	if !reflect.DeepEqual(tickets.startedIDs, expected) {
		t.Fatalf("expected started ids %v, got %v", expected, tickets.startedIDs)
	}
}

func TestEngineTickAssignsQueuedTickets(t *testing.T) {
	ticket := &models.Ticket{ID: 10}
	tickets := &fakeTicketService{
		queuedTickets: []*models.Ticket{ticket},
		queueByTicket: map[int64]int64{10: 20},
	}
	agents := &fakeAgentService{
		availableByQueue: map[int64]*models.Agent{
			20: {ID: 30},
		},
	}
	engine := NewEngine(tickets, agents, discardLogger(), time.Second, 100)

	engine.tick(context.Background())

	expected := []assignment{{ticketID: 10, agentID: 30}}
	if !reflect.DeepEqual(tickets.assignments, expected) {
		t.Fatalf("expected assignments %v, got %v", expected, tickets.assignments)
	}
}

func TestEngineTickSkipsQueuedTicketWhenNoAgentAvailable(t *testing.T) {
	tickets := &fakeTicketService{
		queuedTickets: []*models.Ticket{{ID: 10}},
		queueByTicket: map[int64]int64{10: 20},
	}
	agents := &fakeAgentService{
		availableByQueue: map[int64]*models.Agent{20: nil},
	}
	engine := NewEngine(tickets, agents, discardLogger(), time.Second, 100)

	engine.tick(context.Background())

	if len(tickets.assignments) != 0 {
		t.Fatalf("expected no assignments, got %v", tickets.assignments)
	}
}

func TestEngineTickContinuesWhenStartRouteFails(t *testing.T) {
	tickets := &fakeTicketService{
		newTickets: []*models.Ticket{{ID: 1}, {ID: 2}},
		startErrors: map[int64]error{
			1: errEngineTest,
		},
	}
	agents := &fakeAgentService{}
	engine := NewEngine(tickets, agents, discardLogger(), time.Second, 100)

	engine.tick(context.Background())

	expected := []int64{2}
	if !reflect.DeepEqual(tickets.startedIDs, expected) {
		t.Fatalf("expected started ids %v, got %v", expected, tickets.startedIDs)
	}
}

func TestEngineTickSkipsAssignmentWhenAgentLookupFails(t *testing.T) {
	tickets := &fakeTicketService{
		queuedTickets: []*models.Ticket{{ID: 10}},
		queueByTicket: map[int64]int64{10: 20},
	}
	agents := &fakeAgentService{
		errorsByQueue: map[int64]error{20: errEngineTest},
	}
	engine := NewEngine(tickets, agents, discardLogger(), time.Second, 100)

	engine.tick(context.Background())

	if len(tickets.assignments) != 0 {
		t.Fatalf("expected no assignments, got %v", tickets.assignments)
	}
}

func discardLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}
