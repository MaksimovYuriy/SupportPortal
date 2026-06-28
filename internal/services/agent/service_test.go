package agent

import (
	"context"
	"errors"
	"testing"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

func TestAgentServiceAssignToQueue(t *testing.T) {
	t.Run("creates assignment when agent and queue exist", func(t *testing.T) {
		service, repos := newAgentServiceTestEnv()
		agentQueue := &models.AgentQueue{AgentID: 1, QueueID: 10}

		if err := service.AssignToQueue(context.Background(), agentQueue); err != nil {
			t.Fatalf("AssignToQueue returned error: %v", err)
		}

		if !repos.agentQueues.created {
			t.Fatalf("expected assignment to be created")
		}
	})

	t.Run("rejects invalid ids", func(t *testing.T) {
		service, repos := newAgentServiceTestEnv()
		err := service.AssignToQueue(context.Background(), &models.AgentQueue{AgentID: 0, QueueID: 10})

		if !errors.Is(err, apperrors.ErrValidation) {
			t.Fatalf("expected validation error, got %v", err)
		}
		if repos.agentQueues.created {
			t.Fatalf("expected assignment not to be created")
		}
	})

	t.Run("returns not found for missing queue", func(t *testing.T) {
		service, _ := newAgentServiceTestEnv()
		err := service.AssignToQueue(context.Background(), &models.AgentQueue{AgentID: 1, QueueID: 999})

		if !errors.Is(err, apperrors.ErrNotFound) {
			t.Fatalf("expected not found error, got %v", err)
		}
	})
}

func TestAgentServiceFindAvailableForQueue(t *testing.T) {
	t.Run("returns available agent", func(t *testing.T) {
		service, _ := newAgentServiceTestEnv()

		agent, err := service.FindAvailableForQueue(context.Background(), 10)
		if err != nil {
			t.Fatalf("FindAvailableForQueue returned error: %v", err)
		}
		if agent == nil || agent.ID != 1 {
			t.Fatalf("expected agent 1, got %#v", agent)
		}
	})

	t.Run("rejects inactive queue", func(t *testing.T) {
		service, repos := newAgentServiceTestEnv()
		repos.queues.queues[10] = &models.Queue{ID: 10, Name: "Support", IsActive: false}

		agent, err := service.FindAvailableForQueue(context.Background(), 10)
		if !errors.Is(err, apperrors.ErrValidation) {
			t.Fatalf("expected validation error, got %v", err)
		}
		if agent != nil {
			t.Fatalf("expected nil agent, got %#v", agent)
		}
	})

	t.Run("returns nil when no agent is available", func(t *testing.T) {
		service, repos := newAgentServiceTestEnv()
		repos.agents.availableByQueue[10] = nil

		agent, err := service.FindAvailableForQueue(context.Background(), 10)
		if err != nil {
			t.Fatalf("FindAvailableForQueue returned error: %v", err)
		}
		if agent != nil {
			t.Fatalf("expected nil agent, got %#v", agent)
		}
	})
}
