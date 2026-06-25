package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/transport/dto"
)

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(value)
}

func writeNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func statusCodeName(status int) string {
	return strings.ToLower(
		strings.ReplaceAll(http.StatusText(status), " ", "_"),
	)
}

func writeError(w http.ResponseWriter, status int, message string, details ...string) {
	writeJSON(w, status, dto.NewErrorResponse(statusCodeName(status), message, details...))
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrBadRequest):
		writeError(w, http.StatusBadRequest, "Bad request")
	case errors.Is(err, apperrors.ErrNotFound):
		writeError(w, http.StatusNotFound, "Resource not found")
	default:
		writeError(w, http.StatusInternalServerError, "Internal server error")
	}
}
