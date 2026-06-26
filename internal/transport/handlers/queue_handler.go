package handlers

import (
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
	"github.com/MaksimovYuriy/SupportPortal/internal/services"
	"github.com/MaksimovYuriy/SupportPortal/internal/transport/dto"
)

type QueueHandler struct {
	queueService *services.QueueService
}

func NewQueueHandler(queueService *services.QueueService) *QueueHandler {
	return &QueueHandler{queueService: queueService}
}

func (h *QueueHandler) Index(w http.ResponseWriter, r *http.Request) {
	queues, err := h.queueService.List(r.Context())
	if err != nil {
		handleError(w, err)
	}
	writeJSON(w, http.StatusOK, queues)
}

func (h *QueueHandler) Show(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		handleError(w, err)
	}
	queue, err := h.queueService.FindByID(r.Context(), id)
	if err != nil {
		handleError(w, err)
	}
	writeJSON(w, http.StatusOK, queue)
}

func (h *QueueHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request dto.QueueRequest
	if err := decodeJSONBody(r, &request); err != nil {
		handleError(w, err)
	}
	var queue = models.Queue{
		Name:     request.Name,
		IsActive: request.IsActive,
	}
	err := h.queueService.Create(r.Context(), &queue)
	if err != nil {
		handleError(w, err)
	}
	writeJSON(w, http.StatusOK, dto.NewQueueResponse(&queue))
}
