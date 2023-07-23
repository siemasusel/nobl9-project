package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"stddevapi"

	"golang.org/x/exp/slog"
)

func respondJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		respondWithError(err, w, r)
		return
	}
}

func respondWithError(err error, w http.ResponseWriter, r *http.Request) {
	var valErr *stddevapi.ValidationError
	if !errors.As(err, &valErr) {
		httpRespondWithError(err, "internal server error", w, r, 500)
		return
	}

	httpRespondWithError(err, valErr.Message, w, r, 422)
}

func httpRespondWithError(err error, msg string, w http.ResponseWriter, r *http.Request, status int) {
	resp := ErrorResponse{Error: msg}

	slog.ErrorContext(r.Context(), "Http request faild", "error", err)

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.ErrorContext(r.Context(), "Unable to send response")
	}
}

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}
