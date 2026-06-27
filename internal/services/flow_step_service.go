package services

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type FlowStepService struct {
	flowStepRepo database.FlowStepRepository
}

func NewFlowStepService(flowStepRepo database.FlowStepRepository) *FlowStepService {
	return &FlowStepService{flowStepRepo: flowStepRepo}
}

func (s *FlowStepService) ListByFlowID(ctx context.Context, flowID int64) ([]*models.FlowStep, error) {
	steps, err := s.flowStepRepo.ListByFlowID(ctx, flowID)
	if err != nil {
		return nil, err
	}
	return steps, nil
}

func (s *FlowStepService) FindByID(ctx context.Context, id int64) (*models.FlowStep, error) {
	step, err := s.flowStepRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return step, nil
}
