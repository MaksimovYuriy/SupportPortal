package services

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type FlowService struct {
	flowRepo database.FlowRepository
}

func NewFlowService(flowRepo database.FlowRepository) *FlowService {
	return &FlowService{flowRepo: flowRepo}
}

func (s *FlowService) Create(ctx context.Context, flow *models.Flow) error {
	if err := s.flowRepo.Create(ctx, flow); err != nil {
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
