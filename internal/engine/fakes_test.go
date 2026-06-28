package engine

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type fakeTicketService struct {
	newTickets    []*models.Ticket
	queuedTickets []*models.Ticket
	queueByTicket map[int64]int64
	startedIDs    []int64
	assignments   []assignment
}

type assignment struct {
	ticketID int64
	agentID  int64
}

func (s *fakeTicketService) ListNew(_ context.Context, _ int) ([]*models.Ticket, error) {
	return s.newTickets, nil
}

func (s *fakeTicketService) StartRoute(_ context.Context, ticketID int64) error {
	s.startedIDs = append(s.startedIDs, ticketID)
	return nil
}

func (s *fakeTicketService) ListInQueue(_ context.Context, _ int) ([]*models.Ticket, error) {
	return s.queuedTickets, nil
}

func (s *fakeTicketService) CurrentQueueID(_ context.Context, ticket *models.Ticket) (int64, error) {
	return s.queueByTicket[ticket.ID], nil
}

func (s *fakeTicketService) AssignToAgent(_ context.Context, ticketID int64, agentID int64) error {
	s.assignments = append(s.assignments, assignment{ticketID: ticketID, agentID: agentID})
	return nil
}

type fakeAgentService struct {
	availableByQueue map[int64]*models.Agent
}

func (s *fakeAgentService) FindAvailableForQueue(_ context.Context, queueID int64) (*models.Agent, error) {
	return s.availableByQueue[queueID], nil
}
