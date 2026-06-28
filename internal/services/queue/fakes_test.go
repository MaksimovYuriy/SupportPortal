package queue

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type fakeQueueServiceRepository struct {
	created bool
	queues  []*models.Queue
}

func (r *fakeQueueServiceRepository) Create(_ context.Context, queue *models.Queue) error {
	r.created = true
	r.queues = append(r.queues, queue)
	return nil
}

func (r *fakeQueueServiceRepository) List(_ context.Context) ([]*models.Queue, error) {
	return r.queues, nil
}

func (r *fakeQueueServiceRepository) FindByID(_ context.Context, id int64) (*models.Queue, error) {
	for _, queue := range r.queues {
		if int64(queue.ID) == id {
			return queue, nil
		}
	}
	return nil, apperrors.ErrNotFound
}
