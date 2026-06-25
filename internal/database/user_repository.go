package database

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	List(ctx context.Context) ([]*models.User, error)
	FindByID(ctx context.Context, id int64) (*models.User, error)
}
