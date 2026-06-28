package flow

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type flowServiceTestEnv struct {
	flows  *fakeFlowServiceFlowRepository
	steps  *fakeFlowServiceFlowStepRepository
	queues *fakeFlowServiceQueueRepository
}

func newFlowServiceTestEnv() (*FlowService, *flowServiceTestEnv) {
	repos := &flowServiceTestEnv{
		flows: &fakeFlowServiceFlowRepository{},
		steps: &fakeFlowServiceFlowStepRepository{},
		queues: &fakeFlowServiceQueueRepository{
			queues: map[int64]*models.Queue{
				1: {ID: 1, Name: "Sales", IsActive: true},
				2: {ID: 2, Name: "Support", IsActive: true},
			},
		},
	}

	return NewFlowService(repos.flows, repos.steps, repos.queues), repos
}

type fakeFlowServiceFlowRepository struct {
	createWithStepsCalled bool
	createdFlow           *models.Flow
	createdSteps          []*models.FlowStep
}

func (r *fakeFlowServiceFlowRepository) Create(_ context.Context, _ *models.Flow) error {
	return nil
}

func (r *fakeFlowServiceFlowRepository) CreateWithSteps(_ context.Context, flow *models.Flow, steps []*models.FlowStep) error {
	r.createWithStepsCalled = true
	r.createdFlow = flow
	r.createdSteps = steps
	return nil
}

func (r *fakeFlowServiceFlowRepository) Delete(_ context.Context, _ int64) error {
	return nil
}

func (r *fakeFlowServiceFlowRepository) List(_ context.Context) ([]*models.Flow, error) {
	return nil, nil
}

func (r *fakeFlowServiceFlowRepository) FindByID(_ context.Context, _ int64) (*models.Flow, error) {
	return nil, nil
}

type fakeFlowServiceFlowStepRepository struct{}

func (r *fakeFlowServiceFlowStepRepository) ListByFlowID(_ context.Context, _ int64) ([]*models.FlowStep, error) {
	return nil, nil
}

func (r *fakeFlowServiceFlowStepRepository) FindByID(_ context.Context, _ int64) (*models.FlowStep, error) {
	return nil, nil
}

type fakeFlowServiceQueueRepository struct {
	queues map[int64]*models.Queue
}

func (r *fakeFlowServiceQueueRepository) Create(_ context.Context, _ *models.Queue) error {
	return nil
}

func (r *fakeFlowServiceQueueRepository) List(_ context.Context) ([]*models.Queue, error) {
	return nil, nil
}

func (r *fakeFlowServiceQueueRepository) FindByID(_ context.Context, id int64) (*models.Queue, error) {
	queue, ok := r.queues[id]
	if !ok {
		return nil, apperrors.ErrNotFound
	}
	clone := *queue
	return &clone, nil
}
