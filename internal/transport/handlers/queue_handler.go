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
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, queues)
}

func (h *QueueHandler) Create(w http.ResponseWriter, r *http.Request) {
	var queue models.Queue
	if err := decodeJSONBody(r, &queue); err != nil {
		handleError(w, err)
		return
	}

	if err := h.queueService.Create(r.Context(), &queue); err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, queue)
}

func (h *QueueHandler) Show(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		handleError(w, err)
		return
	}
	queue, err := h.queueService.FindByID(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, queue)
}

func (h *QueueHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		handleError(w, err)
		return
	}

	var queue models.Queue
	if err := decodeJSONBody(r, &queue); err != nil {
		handleError(w, err)
		return
	}
	queue.ID = id

	if err := h.queueService.Update(r.Context(), &queue); err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, queue)
}

func (h *QueueHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		handleError(w, err)
		return
	}

	if err := h.queueService.Delete(r.Context(), id); err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
