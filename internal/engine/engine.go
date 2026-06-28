package engine

import (
	"context"
	"log/slog"
	"time"

	agentservice "github.com/MaksimovYuriy/SupportPortal/internal/services/agent"
	ticketservice "github.com/MaksimovYuriy/SupportPortal/internal/services/ticket"
)

type Engine struct {
	tickets  *ticketservice.TicketService
	agents   *agentservice.AgentService
	logger   *slog.Logger
	interval time.Duration
	limit    int
}

func NewEngine(tickets *ticketservice.TicketService, agents *agentservice.AgentService, logger *slog.Logger, interval time.Duration, limit int) *Engine {
	if logger == nil {
		logger = slog.Default()
	}
	return &Engine{
		tickets:  tickets,
		agents:   agents,
		logger:   logger,
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

	e.logger.Info("ticket engine started", "interval", e.interval.String(), "limit", e.limit)
	defer e.logger.Info("ticket engine stopped")

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
		e.logger.Error("ticket engine failed to list new tickets", "error", err)
		return
	}

	for _, ticket := range tickets {
		if err := e.tickets.StartRoute(ctx, ticket.ID); err != nil {
			e.logger.Error("ticket engine failed to start route", "ticket_id", ticket.ID, "error", err)
			continue
		}
		e.logger.Info("ticket moved to first queue", "ticket_id", ticket.ID)
	}
}

func (e *Engine) assignQueuedTickets(ctx context.Context) {
	tickets, err := e.tickets.ListInQueue(ctx, e.limit)
	if err != nil {
		e.logger.Error("ticket engine failed to list queued tickets", "error", err)
		return
	}

	for _, ticket := range tickets {
		queueID, err := e.tickets.CurrentQueueID(ctx, ticket)
		if err != nil {
			e.logger.Error("ticket engine failed to get current queue", "ticket_id", ticket.ID, "error", err)
			continue
		}

		agent, err := e.agents.FindAvailableForQueue(ctx, queueID)
		if err != nil {
			e.logger.Error("ticket engine failed to find available agent", "queue_id", queueID, "error", err)
			continue
		}
		if agent == nil {
			e.logger.Debug("no available agent for queue", "queue_id", queueID)
			continue
		}

		if err := e.tickets.AssignToAgent(ctx, ticket.ID, int64(agent.ID)); err != nil {
			e.logger.Error("ticket engine failed to assign ticket", "ticket_id", ticket.ID, "agent_id", agent.ID, "error", err)
			continue
		}
		e.logger.Info("ticket assigned", "ticket_id", ticket.ID, "agent_id", agent.ID)
	}
}
