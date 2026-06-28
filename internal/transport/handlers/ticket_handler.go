package handlers

import (
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
	ticketservice "github.com/MaksimovYuriy/SupportPortal/internal/services/ticket"
	"github.com/MaksimovYuriy/SupportPortal/internal/transport/dto"
)

type TicketHandler struct {
	ticketService *ticketservice.TicketService
}

func NewTicketHandler(ticketService *ticketservice.TicketService) *TicketHandler {
	return &TicketHandler{ticketService: ticketService}
}

func (h *TicketHandler) Index(w http.ResponseWriter, r *http.Request) {
	filter, err := parseTicketFilter(r)
	if err != nil {
		handleError(w, err)
		return
	}
	tickets, err := h.ticketService.List(r.Context(), filter)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.NewTicketListResponse(tickets))
}

func parseTicketFilter(r *http.Request) (models.TicketFilter, error) {
	query := r.URL.Query()
	filter := models.TicketFilter{
		Status: query.Get("status"),
	}
	if filter.Status != "" && !models.IsValidTicketStatus(filter.Status) {
		return models.TicketFilter{}, ErrBadRequest
	}
	if value := query.Get("assigned_agent_id"); value != "" {
		id, err := parsePositiveInt64(value)
		if err != nil {
			return models.TicketFilter{}, ErrBadRequest
		}
		filter.AssignedAgentID = &id
	}
	if value := query.Get("flow_id"); value != "" {
		id, err := parsePositiveInt64(value)
		if err != nil {
			return models.TicketFilter{}, ErrBadRequest
		}
		filter.FlowID = &id
	}
	if value := query.Get("queue_id"); value != "" {
		id, err := parsePositiveInt64(value)
		if err != nil {
			return models.TicketFilter{}, ErrBadRequest
		}
		filter.QueueID = &id
	}
	return filter, nil
}

func (h *TicketHandler) Show(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		handleError(w, err)
		return
	}
	ticket, err := h.ticketService.FindByID(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.NewTicketResponse(ticket))
}

func (h *TicketHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request dto.TicketRequest
	if err := decodeJSONBody(r, &request); err != nil {
		handleError(w, err)
		return
	}
	ticket := models.Ticket{
		FlowID:      request.FlowID,
		Title:       request.Title,
		Description: request.Description,
	}
	if err := h.ticketService.Create(r.Context(), &ticket); err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, dto.NewTicketResponse(&ticket))
}

func (h *TicketHandler) CompleteCurrentStep(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		handleError(w, err)
		return
	}
	if err := h.ticketService.CompleteCurrentStep(r.Context(), id); err != nil {
		handleError(w, err)
		return
	}
	ticket, err := h.ticketService.FindByID(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.NewTicketResponse(ticket))
}
