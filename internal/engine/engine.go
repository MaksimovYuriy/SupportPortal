package engine

import (
	"context"
	"log"
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/services"
)

type Engine struct {
	tickets  *services.TicketService
	agents   *services.AgentService
	interval time.Duration
	limit    int
}

func NewEngine(tickets *services.TicketService, agents *services.AgentService, interval time.Duration, limit int) *Engine {
	return &Engine{
		tickets:  tickets,
		agents:   agents,
		interval: interval,
		limit:    limit,
	}
}

func (e *Engine) Run(ctx context.Context) {
	if e.interval <= 0 {
		e.interval = 5 * time.Second
	}
	if e.limit <= 0 {
		e.limit = 100
	}

	log.Printf("ticket engine started: interval=%s limit=%d", e.interval, e.limit)
	defer log.Printf("ticket engine stopped")

	e.tick(ctx)

	ticker := time.NewTicker(e.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			e.tick(ctx)
		}
	}
}

func (e *Engine) tick(ctx context.Context) {
	e.startNewTickets(ctx)
	e.assignQueuedTickets(ctx)
}

func (e *Engine) startNewTickets(ctx context.Context) {
	tickets, err := e.tickets.ListNew(ctx, e.limit)
	if err != nil {
		log.Printf("ticket engine: failed to list new tickets: %v", err)
		return
	}

	for _, ticket := range tickets {
		if err := e.tickets.StartRoute(ctx, ticket.ID); err != nil {
			log.Printf("ticket engine: failed to start route for ticket %d: %v", ticket.ID, err)
			continue
		}
		log.Printf("ticket engine: ticket %d moved to first queue", ticket.ID)
	}
}

func (e *Engine) assignQueuedTickets(ctx context.Context) {
	tickets, err := e.tickets.ListInQueue(ctx, e.limit)
	if err != nil {
		log.Printf("ticket engine: failed to list queued tickets: %v", err)
		return
	}

	for _, ticket := range tickets {
		queueID, err := e.tickets.CurrentQueueID(ctx, ticket)
		if err != nil {
			log.Printf("ticket engine: failed to get current queue for ticket %d: %v", ticket.ID, err)
			continue
		}

		agent, err := e.agents.FindAvailableForQueue(ctx, queueID)
		if err != nil {
			log.Printf("ticket engine: failed to find available agent for queue %d: %v", queueID, err)
			continue
		}
		if agent == nil {
			continue
		}

		if err := e.tickets.AssignToAgent(ctx, ticket.ID, int64(agent.ID)); err != nil {
			log.Printf("ticket engine: failed to assign ticket %d to agent %d: %v", ticket.ID, agent.ID, err)
			continue
		}
		log.Printf("ticket engine: ticket %d assigned to agent %d", ticket.ID, agent.ID)
	}
}
