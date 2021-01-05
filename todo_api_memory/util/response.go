package util

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type successResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// SendErrorResponse send a json payload for error responses
func SendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	response := errorResponse{Success: false, Message: message}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// SendSuccessResponse send a json payload for success responses
func SendSuccessResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	response := successResponse{Success: true, Data: data}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
