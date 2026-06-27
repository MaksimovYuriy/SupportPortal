package handlers

import (
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
	"github.com/MaksimovYuriy/SupportPortal/internal/services"
	"github.com/MaksimovYuriy/SupportPortal/internal/transport/dto"
)

type FlowHandler struct {
	flowService *services.FlowService
}

func NewFlowHandler(flowService *services.FlowService) *FlowHandler {
	return &FlowHandler{flowService: flowService}
}

func (h *FlowHandler) Index(w http.ResponseWriter, r *http.Request) {
	flows, err := h.flowService.List(r.Context())
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.NewFlowListResponse(flows))
}

func (h *FlowHandler) Show(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		handleError(w, err)
		return
	}
	flow, err := h.flowService.FindByID(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}
	steps, err := h.flowService.FindStepsByFlowID(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.NewFlowResponseWithSteps(flow, steps))
}

func (h *FlowHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request dto.FlowRequest
	if err := decodeJSONBody(r, &request); err != nil {
		handleError(w, err)
		return
	}
	flow := models.Flow{
		Name:        request.Name,
		Description: request.Description,
		IsActive:    request.IsActive,
	}
	steps := make([]*models.FlowStep, len(request.Steps))
	for i, stepRequest := range request.Steps {
		steps[i] = &models.FlowStep{
			QueueID:  stepRequest.QueueID,
			Position: stepRequest.Position,
			Name:     stepRequest.Name,
		}
	}
	if err := h.flowService.Create(r.Context(), &flow, steps); err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, dto.NewFlowResponseWithSteps(&flow, steps))
}

func (h *FlowHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		handleError(w, err)
		return
	}
	if err := h.flowService.Delete(r.Context(), id); err != nil {
		handleError(w, err)
		return
	}
	writeNoContent(w)
}
