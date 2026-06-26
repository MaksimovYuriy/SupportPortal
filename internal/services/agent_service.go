package services

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type AgentService struct {
	agentRepo database.AgentRepository
}

func NewAgentService(agentRepo database.AgentRepository) *AgentService {
	return &AgentService{agentRepo: agentRepo}
}

func (s *AgentService) ListAgents(ctx context.Context) ([]*models.Agent, error) {
	agents, err := s.agentRepo.List(ctx)

	if err != nil {
		return nil, err
	}
	return agents, nil
}

func (s *AgentService) FindByID(ctx context.Context, id int64) (*models.Agent, error) {
	agent, err := s.agentRepo.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}
	return agent, nil
}

func (s *AgentService) CreateAgentForUser(ctx context.Context, user *models.User) error {
	err := s.agentRepo.CreateForUser(ctx, user)

	if err != nil {
		return err
	}
	return nil
}
