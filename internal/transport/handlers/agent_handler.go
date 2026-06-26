package handlers

import (
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
	"github.com/MaksimovYuriy/SupportPortal/internal/services"
	"github.com/MaksimovYuriy/SupportPortal/internal/transport/dto"
)

type AgentHandler struct {
	agentService *services.AgentService
}

func NewAgentHandler(agentService *services.AgentService) *AgentHandler {
	return &AgentHandler{agentService: agentService}
}

func (h *AgentHandler) Index(w http.ResponseWriter, r *http.Request) {
	agents, err := h.agentService.ListAgents(r.Context())
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.NewAgentListResponse(agents))
}

func (h *AgentHandler) Show(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		handleError(w, err)
		return
	}
	agent, err := h.agentService.FindByID(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.NewAgentResponse(agent))
}

func (h *AgentHandler) AssignToQueue(w http.ResponseWriter, r *http.Request) {
	var request dto.AgentQueueRequest
	if err := decodeJSONBody(r, &request); err != nil {
		handleError(w, err)
		return
	}
	agentQueue := models.AgentQueue{
		AgentID: request.AgentID,
		QueueID: request.QueueID,
	}
	err := h.agentService.AssignToQueue(r.Context(), &agentQueue)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.NewAgentQueueResponse(&agentQueue))
}
