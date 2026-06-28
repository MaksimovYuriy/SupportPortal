package flow

import (
	"context"
	"errors"
	"testing"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

func TestFlowServiceCreatePersistsFlowWithSteps(t *testing.T) {
	service, repos := newFlowServiceTestEnv()
	flow := &models.Flow{Name: "Default", Description: "Main route", IsActive: true}
	steps := []*models.FlowStep{
		{QueueID: 1, Position: 1, Name: "First step"},
		{QueueID: 2, Position: 2, Name: "Second step"},
	}

	if err := service.Create(context.Background(), flow, steps); err != nil {
		t.Fatalf("Create returned error: %v", err)
	}

	if !repos.flows.createWithStepsCalled {
		t.Fatalf("expected CreateWithSteps to be called")
	}
	if repos.flows.createdFlow != flow {
		t.Fatalf("expected created flow to be passed through")
	}
	if len(repos.flows.createdSteps) != len(steps) {
		t.Fatalf("expected %d steps, got %d", len(steps), len(repos.flows.createdSteps))
	}
}

func TestFlowServiceCreateRejectsEmptySteps(t *testing.T) {
	service, _ := newFlowServiceTestEnv()
	flow := &models.Flow{Name: "Default", IsActive: true}

	err := service.Create(context.Background(), flow, nil)
	if !errors.Is(err, apperrors.ErrValidation) {
		t.Fatalf("expected validation error, got %v", err)
	}
}

func TestFlowServiceCreateRejectsDuplicatePositions(t *testing.T) {
	service, _ := newFlowServiceTestEnv()
	flow := &models.Flow{Name: "Default", IsActive: true}
	steps := []*models.FlowStep{
		{QueueID: 1, Position: 1, Name: "First step"},
		{QueueID: 2, Position: 1, Name: "Second step"},
	}

	err := service.Create(context.Background(), flow, steps)
	if !errors.Is(err, apperrors.ErrValidation) {
		t.Fatalf("expected validation error, got %v", err)
	}
}

func TestFlowServiceCreateRejectsInactiveQueue(t *testing.T) {
	service, repos := newFlowServiceTestEnv()
	repos.queues.queues[1] = &models.Queue{ID: 1, Name: "Inactive", IsActive: false}
	flow := &models.Flow{Name: "Default", IsActive: true}
	steps := []*models.FlowStep{
		{QueueID: 1, Position: 1, Name: "First step"},
	}

	err := service.Create(context.Background(), flow, steps)
	if !errors.Is(err, apperrors.ErrValidation) {
		t.Fatalf("expected validation error, got %v", err)
	}
	if repos.flows.createWithStepsCalled {
		t.Fatalf("expected CreateWithSteps not to be called")
	}
}
