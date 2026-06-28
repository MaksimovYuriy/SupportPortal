package handlers

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

func TestParseTicketFilter(t *testing.T) {
	request := httptest.NewRequest("GET", "/tickets?status=in_progress&assigned_agent_id=7&flow_id=3&queue_id=9", nil)

	filter, err := parseTicketFilter(request)
	if err != nil {
		t.Fatalf("parseTicketFilter returned error: %v", err)
	}

	if filter.Status != models.TicketStatusInProgress {
		t.Fatalf("expected status %q, got %q", models.TicketStatusInProgress, filter.Status)
	}
	if filter.AssignedAgentID == nil || *filter.AssignedAgentID != 7 {
		t.Fatalf("expected assigned_agent_id 7, got %v", filter.AssignedAgentID)
	}
	if filter.FlowID == nil || *filter.FlowID != 3 {
		t.Fatalf("expected flow_id 3, got %v", filter.FlowID)
	}
	if filter.QueueID == nil || *filter.QueueID != 9 {
		t.Fatalf("expected queue_id 9, got %v", filter.QueueID)
	}
}

func TestParseTicketFilterRejectsInvalidIDs(t *testing.T) {
	tests := []string{
		"/tickets?assigned_agent_id=bad",
		"/tickets?assigned_agent_id=0",
		"/tickets?flow_id=bad",
		"/tickets?queue_id=bad",
		"/tickets?status=bad",
	}

	for _, url := range tests {
		request := httptest.NewRequest("GET", url, nil)

		_, err := parseTicketFilter(request)
		if !errors.Is(err, ErrBadRequest) {
			t.Fatalf("expected bad request error for %s, got %v", url, err)
		}
	}
}
