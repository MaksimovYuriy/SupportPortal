package handlers

import (
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
	"github.com/MaksimovYuriy/SupportPortal/internal/services"
	"github.com/MaksimovYuriy/SupportPortal/internal/transport/dto"
)

type TicketHandler struct {
	ticketService *services.TicketService
}

func NewTicketHandler(ticketService *services.TicketService) *TicketHandler {
	return &TicketHandler{ticketService: ticketService}
}

func (h *TicketHandler) Index(w http.ResponseWriter, r *http.Request) {
	tickets, err := h.ticketService.List(r.Context())
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.NewTicketListResponse(tickets))
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
