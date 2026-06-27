package services

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type TicketService struct {
	ticketRepo   database.TicketRepository
	flowRepo     database.FlowRepository
	flowStepRepo database.FlowStepRepository
	agentRepo    database.AgentRepository
}

func NewTicketService(
	ticketRepo database.TicketRepository,
	flowRepo database.FlowRepository,
	flowStepRepo database.FlowStepRepository,
	agentRepo database.AgentRepository,
) *TicketService {
	return &TicketService{
		ticketRepo:   ticketRepo,
		flowRepo:     flowRepo,
		flowStepRepo: flowStepRepo,
		agentRepo:    agentRepo,
	}
}

func (s *TicketService) Create(ctx context.Context, ticket *models.Ticket) error {
	if err := ticket.Validate(); err != nil {
		return err
	}
	flow, err := s.flowRepo.FindByID(ctx, ticket.FlowID)
	if err != nil {
		return err
	}
	if !flow.IsActive {
		return apperrors.ErrValidation
	}

	ticket.Status = models.TicketStatusNew
	ticket.CurrentFlowStepID = nil
	ticket.AssignedAgentID = nil
	if err := s.ticketRepo.Create(ctx, ticket); err != nil {
		return err
	}
	return nil
}

func (s *TicketService) List(ctx context.Context) ([]*models.Ticket, error) {
	tickets, err := s.ticketRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (s *TicketService) FindByID(ctx context.Context, id int64) (*models.Ticket, error) {
	ticket, err := s.ticketRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (s *TicketService) ListNew(ctx context.Context, limit int) ([]*models.Ticket, error) {
	return s.listByStatus(ctx, models.TicketStatusNew, limit)
}

func (s *TicketService) ListInQueue(ctx context.Context, limit int) ([]*models.Ticket, error) {
	return s.listByStatus(ctx, models.TicketStatusInQueue, limit)
}

func (s *TicketService) StartRoute(ctx context.Context, ticketID int64) error {
	ticket, err := s.ticketRepo.FindByID(ctx, ticketID)
	if err != nil {
		return err
	}
	if ticket.Status != models.TicketStatusNew || ticket.CurrentFlowStepID != nil {
		return apperrors.ErrValidation
	}

	firstStep, err := s.findFirstFlowStep(ctx, ticket.FlowID)
	if err != nil {
		return err
	}
	stepID := int64(firstStep.ID)
	ticket.CurrentFlowStepID = &stepID
	ticket.AssignedAgentID = nil
	ticket.Status = models.TicketStatusInQueue
	return s.ticketRepo.UpdateState(ctx, ticket)
}

func (s *TicketService) AssignToAgent(ctx context.Context, ticketID int64, agentID int64) error {
	ticket, err := s.ticketRepo.FindByID(ctx, ticketID)
	if err != nil {
		return err
	}
	if ticket.Status != models.TicketStatusInQueue || ticket.CurrentFlowStepID == nil {
		return apperrors.ErrValidation
	}
	if _, err := s.agentRepo.FindByID(ctx, agentID); err != nil {
		return err
	}

	ticket.AssignedAgentID = &agentID
	ticket.Status = models.TicketStatusInProgress
	return s.ticketRepo.UpdateState(ctx, ticket)
}

func (s *TicketService) CompleteCurrentStep(ctx context.Context, ticketID int64) error {
	ticket, err := s.ticketRepo.FindByID(ctx, ticketID)
	if err != nil {
		return err
	}
	if ticket.Status != models.TicketStatusInProgress || ticket.CurrentFlowStepID == nil {
		return apperrors.ErrValidation
	}

	currentStep, err := s.flowStepRepo.FindByID(ctx, *ticket.CurrentFlowStepID)
	if err != nil {
		return err
	}
	nextStep, err := s.findNextFlowStep(ctx, ticket.FlowID, currentStep.Position)
	if err != nil {
		return err
	}
	if nextStep == nil {
		ticket.CurrentFlowStepID = nil
		ticket.Status = models.TicketStatusClosed
		return s.ticketRepo.UpdateState(ctx, ticket)
	}

	nextStepID := int64(nextStep.ID)
	ticket.CurrentFlowStepID = &nextStepID
	ticket.AssignedAgentID = nil
	ticket.Status = models.TicketStatusInQueue
	return s.ticketRepo.UpdateState(ctx, ticket)
}

func (s *TicketService) listByStatus(ctx context.Context, status string, limit int) ([]*models.Ticket, error) {
	if limit <= 0 {
		limit = 100
	}
	tickets, err := s.ticketRepo.ListByStatus(ctx, status, limit)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (s *TicketService) findFirstFlowStep(ctx context.Context, flowID int64) (*models.FlowStep, error) {
	steps, err := s.flowStepRepo.ListByFlowID(ctx, flowID)
	if err != nil {
		return nil, err
	}
	if len(steps) == 0 {
		return nil, apperrors.ErrValidation
	}
	return steps[0], nil
}

func (s *TicketService) findNextFlowStep(ctx context.Context, flowID int64, currentPosition int) (*models.FlowStep, error) {
	steps, err := s.flowStepRepo.ListByFlowID(ctx, flowID)
	if err != nil {
		return nil, err
	}
	for _, step := range steps {
		if step.Position > currentPosition {
			return step, nil
		}
	}
	return nil, nil
}
