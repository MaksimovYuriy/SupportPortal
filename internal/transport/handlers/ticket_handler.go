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
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusOK, tickets)
}

func (h *TicketHandler) Create(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	if err := decodeJSONBody(r, &ticket); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.Create(r.Context(), &ticket); err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	writeJSON(w, http.StatusCreated, ticket)
}

func (h *TicketHandler) Show(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid path params")
		return
	}
	ticket, err := h.service.FindByID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusOK, ticket)
}

func (h *TicketHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid path params")
		return
	}

	var ticket models.Ticket
	if err := decodeJSONBody(r, &ticket); err != nil {
		writeError(w, http.StatusBadRequest, "invalid body parameters")
		return
	}
	ticket.ID = id

	if err := h.service.Update(r.Context(), &ticket); err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	writeJSON(w, http.StatusOK, ticket)
}

func (h *TicketHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid path params")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	writeNoContent(w)
}
