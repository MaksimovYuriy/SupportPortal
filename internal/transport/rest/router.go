package rest

import (
	"encoding/json"
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/transport/handlers"
)

type Handlers struct {
	TicketHandler *handlers.TicketHandler
	QueueHandler  *handlers.QueueHandler
}

func NewRouter(handlers *Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)

	mux.HandleFunc("GET /tickets", handlers.TicketHandler.Index)
	mux.HandleFunc("GET /tickets/{id}", handlers.TicketHandler.Show)
	mux.HandleFunc("POST /tickets", handlers.TicketHandler.Create)
	mux.HandleFunc("PUT /tickets/{id}", handlers.TicketHandler.Update)
	mux.HandleFunc("DELETE /tickets/{id}", handlers.TicketHandler.Delete)

	mux.HandleFunc("GET /queues", handlers.QueueHandler.Index)
	mux.HandleFunc("GET /queues/{id}", handlers.QueueHandler.Show)
	mux.HandleFunc("POST /queues", handlers.QueueHandler.Create)
	mux.HandleFunc("PUT /queues/{id}", handlers.QueueHandler.Update)
	mux.HandleFunc("DELETE /queues/{id}", handlers.QueueHandler.Delete)

	return mux
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"status":  "ok",
		"service": "support_portal",
	}

	_ = json.NewEncoder(w).Encode(response)
}
