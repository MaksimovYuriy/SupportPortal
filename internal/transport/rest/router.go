package rest

import (
	"encoding/json"
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/transport/handlers"
)

func NewRouter(ticketHandler *handlers.TicketHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)

	mux.HandleFunc("GET /tickets", ticketHandler.Index)
	mux.HandleFunc("GET /tickets/{id}", ticketHandler.Show)
	mux.HandleFunc("POST /tickets", ticketHandler.Create)
	mux.HandleFunc("PUT /tickets/{id}", ticketHandler.Update)
	mux.HandleFunc("DELETE /tickets/{id}", ticketHandler.Delete)

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
