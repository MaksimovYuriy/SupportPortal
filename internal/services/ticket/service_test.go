package ticket

import (
	"context"
	"errors"
	"testing"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

func TestTicketServiceCreateSetsInitialState(t *testing.T) {
	service, repos := newTicketServiceTestEnv()
	ticket := &models.Ticket{
		FlowID:      1,
		Title:       "Need help",
		Description: "Customer question",
	}

	if err := service.Create(context.Background(), ticket); err != nil {
		t.Fatalf("Create returned error: %v", err)
	}

	if ticket.Status != models.TicketStatusNew {
		t.Fatalf("expected status %q, got %q", models.TicketStatusNew, ticket.Status)
	}
	if ticket.CurrentFlowStepID != nil {
		t.Fatalf("expected nil current step, got %v", *ticket.CurrentFlowStepID)
	}
	if ticket.AssignedAgentID != nil {
		t.Fatalf("expected nil assigned agent, got %v", *ticket.AssignedAgentID)
	}
	if _, ok := repos.tickets.tickets[ticket.ID]; !ok {
		t.Fatalf("expected ticket to be stored")
	}
}

func TestTicketServiceStartRouteMovesTicketToFirstStep(t *testing.T) {
	service, repos := newTicketServiceTestEnv()
	ticket := repos.tickets.add(&models.Ticket{
		FlowID: 1,
		Title:  "Need help",
		Status: models.TicketStatusNew,
	})

	if err := service.StartRoute(context.Background(), ticket.ID); err != nil {
		t.Fatalf("StartRoute returned error: %v", err)
	}

	updated := repos.tickets.tickets[ticket.ID]
	if updated.Status != models.TicketStatusInQueue {
		t.Fatalf("expected status %q, got %q", models.TicketStatusInQueue, updated.Status)
	}
	if updated.CurrentFlowStepID == nil || *updated.CurrentFlowStepID != 1 {
		t.Fatalf("expected current step 1, got %v", updated.CurrentFlowStepID)
	}
}

func TestTicketServiceAssignToAgentValidatesQueueAndAvailability(t *testing.T) {
	tests := []struct {
		name     string
		agent    *models.Agent
		hasQueue bool
		busy     bool
	}{
		{
			name:     "unavailable agent",
			agent:    &models.Agent{ID: 1, IsAvailable: false},
			hasQueue: true,
		},
		{
			name:     "agent from another queue",
			agent:    &models.Agent{ID: 1, IsAvailable: true},
			hasQueue: false,
		},
		{
			name:     "busy agent",
			agent:    &models.Agent{ID: 1, IsAvailable: true},
			hasQueue: true,
			busy:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, repos := newTicketServiceTestEnv()
			repos.agents.agents[1] = tt.agent
			repos.agentQueues.exists[[2]int64{1, 10}] = tt.hasQueue
			repos.tickets.hasInProgress[1] = tt.busy

			stepID := int64(1)
			ticket := repos.tickets.add(&models.Ticket{
				FlowID:            1,
				CurrentFlowStepID: &stepID,
				Title:             "Need help",
				Status:            models.TicketStatusInQueue,
			})

			err := service.AssignToAgent(context.Background(), ticket.ID, 1)
			if !errors.Is(err, apperrors.ErrValidation) {
				t.Fatalf("expected validation error, got %v", err)
			}
		})
	}
}

func TestTicketServiceCompleteCurrentStepClosesOneStepFlow(t *testing.T) {
	service, repos := newTicketServiceTestEnv()
	stepID := int64(1)
	agentID := int64(1)
	ticket := repos.tickets.add(&models.Ticket{
		FlowID:            1,
		CurrentFlowStepID: &stepID,
		AssignedAgentID:   &agentID,
		Title:             "Need help",
		Status:            models.TicketStatusInProgress,
	})

	if err := service.CompleteCurrentStep(context.Background(), ticket.ID); err != nil {
		t.Fatalf("CompleteCurrentStep returned error: %v", err)
	}

	updated := repos.tickets.tickets[ticket.ID]
	if updated.Status != models.TicketStatusClosed {
		t.Fatalf("expected status %q, got %q", models.TicketStatusClosed, updated.Status)
	}
	if updated.CurrentFlowStepID != nil {
		t.Fatalf("expected nil current step, got %v", *updated.CurrentFlowStepID)
	}
	if updated.AssignedAgentID == nil || *updated.AssignedAgentID != agentID {
		t.Fatalf("expected last assigned agent to remain")
	}
}

func TestTicketServiceCompleteCurrentStepMovesToNextStep(t *testing.T) {
	service, repos := newTicketServiceTestEnv()
	repos.steps.stepsByFlow[1] = []*models.FlowStep{
		{ID: 1, FlowID: 1, QueueID: 10, Position: 1, Name: "First"},
		{ID: 2, FlowID: 1, QueueID: 20, Position: 2, Name: "Second"},
	}
	repos.steps.stepsByID[2] = repos.steps.stepsByFlow[1][1]

	stepID := int64(1)
	agentID := int64(1)
	ticket := repos.tickets.add(&models.Ticket{
		FlowID:            1,
		CurrentFlowStepID: &stepID,
		AssignedAgentID:   &agentID,
		Title:             "Need help",
		Status:            models.TicketStatusInProgress,
	})

	if err := service.CompleteCurrentStep(context.Background(), ticket.ID); err != nil {
		t.Fatalf("CompleteCurrentStep returned error: %v", err)
	}

	updated := repos.tickets.tickets[ticket.ID]
	if updated.Status != models.TicketStatusInQueue {
		t.Fatalf("expected status %q, got %q", models.TicketStatusInQueue, updated.Status)
	}
	if updated.CurrentFlowStepID == nil || *updated.CurrentFlowStepID != 2 {
		t.Fatalf("expected current step 2, got %v", updated.CurrentFlowStepID)
	}
	if updated.AssignedAgentID != nil {
		t.Fatalf("expected assigned agent to be reset")
	}
}
