package services

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type AgentService struct {
	agentRepo      database.AgentRepository
	queueRepo      database.QueueRepository
	agentQueueRepo database.AgentQueueRepository
}

func NewAgentService(agentRepo database.AgentRepository, queueRepo database.QueueRepository, agentQueueRepo database.AgentQueueRepository) *AgentService {
	return &AgentService{agentRepo: agentRepo, queueRepo: queueRepo, agentQueueRepo: agentQueueRepo}
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

func (s *AgentService) AssignToQueue(ctx context.Context, agentQueue *models.AgentQueue) error {
	if err := agentQueue.Validate(); err != nil {
		return err
	}
	if _, err := s.agentRepo.FindByID(ctx, agentQueue.AgentID); err != nil {
		return err
	}
	if _, err := s.queueRepo.FindByID(ctx, agentQueue.QueueID); err != nil {
		return err
	}
	err := s.agentQueueRepo.Create(ctx, agentQueue)
	if err != nil {
		return err
	}
	return nil
}
