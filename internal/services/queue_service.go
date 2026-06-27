package services

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type QueueService struct {
	queueRepo database.QueueRepository
}

func NewQueueService(queueRepo database.QueueRepository) *QueueService {
	return &QueueService{queueRepo: queueRepo}
}

func (s *QueueService) Create(ctx context.Context, queue *models.Queue) error {
	if err := queue.Validate(); err != nil {
		return err
	}
	err := s.queueRepo.Create(ctx, queue)
	if err != nil {
		return err
	}
	return nil
}

func (s *QueueService) List(ctx context.Context) ([]*models.Queue, error) {
	queues, err := s.queueRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return queues, nil
}

func (s *QueueService) FindByID(ctx context.Context, id int64) (*models.Queue, error) {
	queue, err := s.queueRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return queue, nil
}
