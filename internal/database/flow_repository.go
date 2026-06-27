package database

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type FlowRepository interface {
	Create(ctx context.Context, flow *models.Flow) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*models.Flow, error)
	FindByID(ctx context.Context, id int64) (*models.Flow, error)
}
