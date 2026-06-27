package database

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type FlowStepRepository interface {
	ListByFlowID(ctx context.Context, flowID int64) ([]*models.FlowStep, error)
	FindByID(ctx context.Context, id int64) (*models.FlowStep, error)
}
