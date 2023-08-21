package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorHandler struct {
	Writer http.ResponseWriter
}

func NewErrorHandler(w http.ResponseWriter) *ErrorHandler {
	return &ErrorHandler{Writer: w}
}

func (eh *ErrorHandler) HandleError(err error, statusCode int, message string) {
	response := struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}{
		Error:   err.Error(),
		Message: message,
	}

	eh.Writer.Header().Set("Content-Type", "application/json")
	eh.Writer.WriteHeader(statusCode)

	errResponse, _ := json.Marshal(response)
	eh.Writer.Write(errResponse)
	log.Println(response.Error)
}
