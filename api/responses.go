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

func RespondWithError(w http.ResponseWriter, httpCode int, message string) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": message,
	})
	w.WriteHeader(httpCode)
}

func RespondWithInternalServerError(w http.ResponseWriter) {
	RespondWithError(w, http.StatusInternalServerError, "An unknown error happened.")
}
