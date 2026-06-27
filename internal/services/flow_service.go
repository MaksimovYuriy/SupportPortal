package services

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type FlowService struct {
	flowRepo     database.FlowRepository
	flowStepRepo database.FlowStepRepository
	queueRepo    database.QueueRepository
}

func NewFlowService(flowRepo database.FlowRepository, flowStepRepo database.FlowStepRepository, queueRepo database.QueueRepository) *FlowService {
	return &FlowService{flowRepo: flowRepo, flowStepRepo: flowStepRepo, queueRepo: queueRepo}
}

func (s *FlowService) Create(ctx context.Context, flow *models.Flow, steps []*models.FlowStep) error {
	if err := validateFlow(flow, steps); err != nil {
		return err
	}
	for _, step := range steps {
		if _, err := s.queueRepo.FindByID(ctx, int64(step.QueueID)); err != nil {
			return err
		}
	}
	if err := s.flowRepo.CreateWithSteps(ctx, flow, steps); err != nil {
		return err
	}
	return nil
}

func (s *FlowService) Delete(ctx context.Context, id int64) error {
	if err := s.flowRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *FlowService) List(ctx context.Context) ([]*models.Flow, error) {
	flows, err := s.flowRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return flows, nil
}

func (s *FlowService) FindByID(ctx context.Context, id int64) (*models.Flow, error) {
	flow, err := s.flowRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return flow, nil
}

func (s *FlowService) FindStepsByFlowID(ctx context.Context, flowID int64) ([]*models.FlowStep, error) {
	steps, err := s.flowStepRepo.ListByFlowID(ctx, flowID)
	if err != nil {
		return nil, err
	}
	return steps, nil
}

func validateFlow(flow *models.Flow, steps []*models.FlowStep) error {
	if err := flow.Validate(); err != nil {
		return err
	}
	if len(steps) == 0 {
		return apperrors.ErrValidation
	}

	positions := make(map[int]struct{}, len(steps))
	for _, step := range steps {
		if err := step.Validate(); err != nil {
			return err
		}
		if _, ok := positions[step.Position]; ok {
			return apperrors.ErrValidation
		}
		positions[step.Position] = struct{}{}
	}

	return nil
}
