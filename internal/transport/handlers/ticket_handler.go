package handlers

import (
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
	"github.com/MaksimovYuriy/SupportPortal/internal/services"
	"github.com/MaksimovYuriy/SupportPortal/internal/transport/dto"
)

type TicketHandler struct {
	service *services.TicketService
}

func NewTicketHandler(service *services.TicketService) *TicketHandler {
	return &TicketHandler{
		service: service,
	}
}

func (h *TicketHandler) Index(w http.ResponseWriter, r *http.Request) {
	tickets, err := h.service.List(r.Context())
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.NewListTicketsResponse(tickets))
}

func (h *TicketHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateTicketRequest
	if err := decodeJSONBody(r, &request); err != nil {
		handleError(w, err)
		return
	}
	ticket := models.Ticket{
		Title:       request.Title,
		Description: request.Description,
	}
	if err := h.service.Create(r.Context(), &ticket); err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, dto.NewTicketResponse(ticket))
}

func (h *TicketHandler) Show(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		handleError(w, err)
		return
	}
	ticket, err := h.service.FindByID(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.NewTicketResponse(*ticket))
}

func (h *TicketHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		handleError(w, err)
		return
	}
	var request dto.UpdateTicketRequest
	if err := decodeJSONBody(r, &request); err != nil {
		handleError(w, err)
		return
	}
	ticket := models.Ticket{
		Title:       request.Title,
		Description: request.Description,
	}
	ticket.ID = id

	if err := h.service.Update(r.Context(), &ticket); err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, dto.NewTicketResponse(ticket))
}

func (h *TicketHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		handleError(w, err)
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		handleError(w, err)
		return
	}

	writeNoContent(w)
}
