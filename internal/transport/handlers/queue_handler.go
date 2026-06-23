package handlers

import (
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
	"github.com/MaksimovYuriy/SupportPortal/internal/services"
)

type QueueHandler struct {
	queueService *services.QueueService
}

func NewQueueHandler(queueService *services.QueueService) *QueueHandler {
	return &QueueHandler{
		queueService: queueService,
	}
}

func (h *QueueHandler) Index(w http.ResponseWriter, r *http.Request) {
	queues, err := h.queueService.List(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusOK, queues)
}

func (h *QueueHandler) Create(w http.ResponseWriter, r *http.Request) {
	var queue models.Queue
	if err := decodeJSONBody(r, &queue); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.queueService.Create(r.Context(), &queue); err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusCreated, queue)
}

func (h *QueueHandler) Show(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid path params")
		return
	}
	queue, err := h.queueService.FindByID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusOK, queue)
}

func (h *QueueHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid path params")
		return
	}

	var queue models.Queue
	if err := decodeJSONBody(r, &queue); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	queue.ID = id

	if err := h.queueService.Update(r.Context(), &queue); err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	writeJSON(w, http.StatusOK, queue)
}

func (h *QueueHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid path params")
		return
	}

	if err := h.queueService.Delete(r.Context(), id); err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
