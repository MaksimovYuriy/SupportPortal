package rest

import (
	"encoding/json"
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/transport/handlers"
)

type Handlers struct {
	UserHandler *handlers.UserHandler
}

func NewRouter(handlers *Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)

	mux.HandleFunc("GET /users", handlers.UserHandler.Index)
	mux.HandleFunc("GET /users/{id}", handlers.UserHandler.Show)
	mux.HandleFunc("POST /users", handlers.UserHandler.Create)

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
