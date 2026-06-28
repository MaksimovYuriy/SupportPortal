package handlers

import (
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
	queueservice "github.com/MaksimovYuriy/SupportPortal/internal/services/queue"
	"github.com/MaksimovYuriy/SupportPortal/internal/transport/dto"
)

type QueueHandler struct {
	queueService *queueservice.QueueService
}

func NewQueueHandler(queueService *queueservice.QueueService) *QueueHandler {
	return &QueueHandler{queueService: queueService}
}

func (h *QueueHandler) Index(w http.ResponseWriter, r *http.Request) {
	queues, err := h.queueService.List(r.Context())
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.NewQueueListResponse(queues))
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
	writeJSON(w, http.StatusOK, dto.NewQueueResponse(queue))
}

func (h *QueueHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request dto.QueueRequest
	if err := decodeJSONBody(r, &request); err != nil {
		handleError(w, err)
		return
	}
	var queue = models.Queue{
		Name:     request.Name,
		IsActive: request.IsActive,
	}
	err := h.queueService.Create(r.Context(), &queue)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.NewQueueResponse(&queue))
}
