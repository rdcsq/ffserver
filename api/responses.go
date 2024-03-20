package api

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, httpCode int, data any) {
	if data == nil {
		return
	}

	json.NewEncoder(w).Encode(data)
	w.WriteHeader(httpCode)
}

func RespondWithError(w http.ResponseWriter, httpCode int, code string) {
	json.NewEncoder(w).Encode(map[string]string{
		"error_code": code,
	})
	w.WriteHeader(httpCode)
}

func RespondWithInternalServerError(w http.ResponseWriter) {
	RespondWithError(w, http.StatusInternalServerError, "unknown")
}
