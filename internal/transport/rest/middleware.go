package rest

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/transport/dto"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *statusRecorder) Write(data []byte) (int, error) {
	if r.status == 0 {
		r.status = http.StatusOK
	}
	return r.ResponseWriter.Write(data)
}

func withLogging(next http.Handler, logger *slog.Logger) http.Handler {
	if logger == nil {
		logger = slog.Default()
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startedAt := time.Now()
		recorder := &statusRecorder{ResponseWriter: w}

		next.ServeHTTP(recorder, r)

		status := recorder.status
		if status == 0 {
			status = http.StatusOK
		}

		attrs := []any{
			"method", r.Method,
			"path", r.URL.Path,
			"status", status,
			"duration", time.Since(startedAt).String(),
		}
		if status >= http.StatusInternalServerError {
			logger.Error("http request completed", attrs...)
			return
		}
		logger.Info("http request completed", attrs...)
	})
}

func withRecovery(next http.Handler, logger *slog.Logger) http.Handler {
	if logger == nil {
		logger = slog.Default()
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if value := recover(); value != nil {
				logger.Error("panic recovered from http handler", "panic", value, "method", r.Method, "path", r.URL.Path)
				writeRecoveryError(w)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func writeRecoveryError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	response := dto.NewErrorResponse("internal_server_error", "Internal server error")
	_ = json.NewEncoder(w).Encode(response)
}
