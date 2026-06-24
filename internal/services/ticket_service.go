package services

import (
	"context"
	"fmt"

	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type TicketService struct {
	ticket_repo database.TicketRepository
	queue_repo  database.QueueRepository
}

func NewTicketService(ticket_repo database.TicketRepository, queue_repo database.QueueRepository) *TicketService {
	return &TicketService{
		ticket_repo: ticket_repo,
		queue_repo:  queue_repo,
	}
}

func (s *TicketService) List(ctx context.Context) ([]models.Ticket, error) {
	return s.ticket_repo.List(ctx)
}

func (s *TicketService) Create(ctx context.Context, ticket *models.Ticket) error {
	distributor_queue, err := s.queue_repo.FindByName(ctx, "Distributor")
	if err != nil {
		return err
	}
	if !distributor_queue.IsActive {
		return fmt.Errorf("distributor queue is inactive")
	}

	ticket.QueueID = &distributor_queue.ID
	err = s.ticket_repo.Create(ctx, ticket)
	if err != nil {
		return err
	}

	return nil
}

func (s *TicketService) FindByID(ctx context.Context, id int64) (*models.Ticket, error) {
	return s.ticket_repo.FindByID(ctx, id)
}

func (s *TicketService) Update(ctx context.Context, ticket *models.Ticket) error {
	return s.ticket_repo.Update(ctx, ticket)
}

func (s *TicketService) Delete(ctx context.Context, id int64) error {
	return s.ticket_repo.Delete(ctx, id)
}
