package handlers

import (
	"encoding/json"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]string{"error": message}
	json.NewEncoder(w).Encode(response)
}
