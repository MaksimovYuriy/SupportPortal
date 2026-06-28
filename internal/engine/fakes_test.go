package engine

import (
	"context"
	"errors"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

var errEngineTest = errors.New("engine test error")

type fakeTicketService struct {
	newTickets    []*models.Ticket
	queuedTickets []*models.Ticket
	queueByTicket map[int64]int64
	startErrors   map[int64]error
	queueErrors   map[int64]error
	assignErrors  map[int64]error
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
	if err := s.startErrors[ticketID]; err != nil {
		return err
	}
	s.startedIDs = append(s.startedIDs, ticketID)
	return nil
}

func (s *fakeTicketService) ListInQueue(_ context.Context, _ int) ([]*models.Ticket, error) {
	return s.queuedTickets, nil
}

func (s *fakeTicketService) CurrentQueueID(_ context.Context, ticket *models.Ticket) (int64, error) {
	if err := s.queueErrors[ticket.ID]; err != nil {
		return 0, err
	}
	return s.queueByTicket[ticket.ID], nil
}

func (s *fakeTicketService) AssignToAgent(_ context.Context, ticketID int64, agentID int64) error {
	if err := s.assignErrors[ticketID]; err != nil {
		return err
	}
	s.assignments = append(s.assignments, assignment{ticketID: ticketID, agentID: agentID})
	return nil
}

type fakeAgentService struct {
	availableByQueue map[int64]*models.Agent
	errorsByQueue    map[int64]error
}

func (s *fakeAgentService) FindAvailableForQueue(_ context.Context, queueID int64) (*models.Agent, error) {
	if err := s.errorsByQueue[queueID]; err != nil {
		return nil, err
	}
	return s.availableByQueue[queueID], nil
}
