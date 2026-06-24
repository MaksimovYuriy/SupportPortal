package handlers

import (
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
	"github.com/MaksimovYuriy/SupportPortal/internal/services"
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
	writeJSON(w, http.StatusOK, tickets)
}

func (h *TicketHandler) Create(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	if err := decodeJSONBody(r, &ticket); err != nil {
		handleError(w, err)
		return
	}

	if err := h.service.Create(r.Context(), &ticket); err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, ticket)
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
	writeJSON(w, http.StatusOK, ticket)
}

func (h *TicketHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		handleError(w, err)
		return
	}

	var ticket models.Ticket
	if err := decodeJSONBody(r, &ticket); err != nil {
		handleError(w, err)
		return
	}
	ticket.ID = id

	if err := h.service.Update(r.Context(), &ticket); err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, ticket)
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
