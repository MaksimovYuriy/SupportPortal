package services

import (
	"context"
	"fmt"

	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type TicketService struct {
	ticketRepo database.TicketRepository
	queueRepo  database.QueueRepository
}

func NewTicketService(ticketRepo database.TicketRepository, queueRepo database.QueueRepository) *TicketService {
	return &TicketService{
		ticketRepo: ticketRepo,
		queueRepo:  queueRepo,
	}
}

func (s *TicketService) List(ctx context.Context) ([]models.Ticket, error) {
	return s.ticketRepo.List(ctx)
}

func (s *TicketService) Create(ctx context.Context, ticket *models.Ticket) error {
	distributorQueue, err := s.queueRepo.FindByName(ctx, "Distribution")
	if err != nil {
		return err
	}
	if !distributorQueue.IsActive {
		return fmt.Errorf("distributor queue is inactive")
	}

	ticket.QueueID = &distributorQueue.ID
	err = s.ticketRepo.Create(ctx, ticket)
	if err != nil {
		return err
	}

	return nil
}

func (s *TicketService) FindByID(ctx context.Context, id int64) (*models.Ticket, error) {
	return s.ticketRepo.FindByID(ctx, id)
}

func (s *TicketService) Update(ctx context.Context, ticket *models.Ticket) error {
	return s.ticketRepo.Update(ctx, ticket)
}

func (s *TicketService) Delete(ctx context.Context, id int64) error {
	return s.ticketRepo.Delete(ctx, id)
}
