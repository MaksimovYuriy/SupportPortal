package agent

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type agentServiceTestEnv struct {
	agents      *fakeAgentServiceAgentRepository
	queues      *fakeAgentServiceQueueRepository
	agentQueues *fakeAgentServiceAgentQueueRepository
}

func newAgentServiceTestEnv() (*AgentService, *agentServiceTestEnv) {
	repos := &agentServiceTestEnv{
		agents: &fakeAgentServiceAgentRepository{
			agents: map[int64]*models.Agent{
				1: {ID: 1, Name: "agent@example.com", IsAvailable: true},
			},
			availableByQueue: map[int64]*models.Agent{
				10: {ID: 1, Name: "agent@example.com", IsAvailable: true},
			},
		},
		queues: &fakeAgentServiceQueueRepository{
			queues: map[int64]*models.Queue{
				10: {ID: 10, Name: "Support", IsActive: true},
			},
		},
		agentQueues: &fakeAgentServiceAgentQueueRepository{},
	}
	return NewAgentService(repos.agents, repos.queues, repos.agentQueues), repos
}

type fakeAgentServiceAgentRepository struct {
	agents           map[int64]*models.Agent
	availableByQueue map[int64]*models.Agent
}

func (r *fakeAgentServiceAgentRepository) List(_ context.Context) ([]*models.Agent, error) {
	return nil, nil
}

func (r *fakeAgentServiceAgentRepository) FindByID(_ context.Context, id int64) (*models.Agent, error) {
	agent, ok := r.agents[id]
	if !ok {
		return nil, apperrors.ErrNotFound
	}
	clone := *agent
	return &clone, nil
}

func (r *fakeAgentServiceAgentRepository) FindAvailableByQueueID(_ context.Context, queueID int64) (*models.Agent, error) {
	agent, ok := r.availableByQueue[queueID]
	if !ok || agent == nil {
		return nil, apperrors.ErrNotFound
	}
	clone := *agent
	return &clone, nil
}

func (r *fakeAgentServiceAgentRepository) CreateForUser(_ context.Context, _ *models.User) error {
	return nil
}

type fakeAgentServiceQueueRepository struct {
	queues map[int64]*models.Queue
}

func (r *fakeAgentServiceQueueRepository) Create(_ context.Context, _ *models.Queue) error {
	return nil
}

func (r *fakeAgentServiceQueueRepository) List(_ context.Context) ([]*models.Queue, error) {
	return nil, nil
}

func (r *fakeAgentServiceQueueRepository) FindByID(_ context.Context, id int64) (*models.Queue, error) {
	queue, ok := r.queues[id]
	if !ok {
		return nil, apperrors.ErrNotFound
	}
	clone := *queue
	return &clone, nil
}

type fakeAgentServiceAgentQueueRepository struct {
	created bool
}

func (r *fakeAgentServiceAgentQueueRepository) Create(_ context.Context, _ *models.AgentQueue) error {
	r.created = true
	return nil
}

func (r *fakeAgentServiceAgentQueueRepository) Delete(_ context.Context, _ *models.AgentQueue) error {
	return nil
}

func (r *fakeAgentServiceAgentQueueRepository) FindByAgentID(_ context.Context, _ int64) ([]*models.AgentQueue, error) {
	return nil, nil
}

func (r *fakeAgentServiceAgentQueueRepository) FindByQueueID(_ context.Context, _ int64) ([]*models.AgentQueue, error) {
	return nil, nil
}

func (r *fakeAgentServiceAgentQueueRepository) Exists(_ context.Context, _ int64, _ int64) (bool, error) {
	return false, nil
}
