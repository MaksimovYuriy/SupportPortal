package rest

import (
	"encoding/json"
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/transport/handlers"
)

type Handlers struct {
	UserHandler   *handlers.UserHandler
	AgentHandler  *handlers.AgentHandler
	QueueHandler  *handlers.QueueHandler
	FlowHandler   *handlers.FlowHandler
	TicketHandler *handlers.TicketHandler
}

func NewRouter(handlers *Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)

	mux.HandleFunc("GET /users", handlers.UserHandler.Index)
	mux.HandleFunc("GET /users/{id}", handlers.UserHandler.Show)
	mux.HandleFunc("POST /users", handlers.UserHandler.Create)

	mux.HandleFunc("GET /agents", handlers.AgentHandler.Index)
	mux.HandleFunc("GET /agents/{id}", handlers.AgentHandler.Show)
	mux.HandleFunc("POST /agents/assign_to_queue", handlers.AgentHandler.AssignToQueue)

	mux.HandleFunc("GET /queues", handlers.QueueHandler.Index)
	mux.HandleFunc("GET /queues/{id}", handlers.QueueHandler.Show)
	mux.HandleFunc("POST /queues", handlers.QueueHandler.Create)

	mux.HandleFunc("GET /flows", handlers.FlowHandler.Index)
	mux.HandleFunc("GET /flows/{id}", handlers.FlowHandler.Show)
	mux.HandleFunc("POST /flows", handlers.FlowHandler.Create)
	mux.HandleFunc("DELETE /flows/{id}", handlers.FlowHandler.Delete)

	mux.HandleFunc("GET /tickets", handlers.TicketHandler.Index)
	mux.HandleFunc("GET /tickets/{id}", handlers.TicketHandler.Show)
	mux.HandleFunc("POST /tickets", handlers.TicketHandler.Create)

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
