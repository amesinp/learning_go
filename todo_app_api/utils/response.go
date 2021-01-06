package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ResponseParams struct is used for responses
type ResponseParams struct {
	Writer     http.ResponseWriter
	Message    string
	StatusCode int
	Data       interface{}
}

type jsonResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// SendSuccessResponse sends a success json payload
func SendSuccessResponse(params ResponseParams) {
	fmt.Println(params)

	params.Writer.WriteHeader(getStatusCode(params.StatusCode, http.StatusOK))
	json.NewEncoder(params.Writer).Encode(jsonResponse{Success: true, Message: params.Message, Data: params.Data})
}

// SendErrorResponse sends an error json payload
func SendErrorResponse(params ResponseParams) {
	if params.Message == "" {
		params.Message = "An error occurred"
	}

	params.Writer.WriteHeader(getStatusCode(params.StatusCode, http.StatusInternalServerError))
	json.NewEncoder(params.Writer).Encode(jsonResponse{Success: false, Error: params.Message, Data: params.Data})
}

func getStatusCode(statusCode int, fallbackStatusCode int) int {
	if statusCode != 0 {
		return statusCode
	}
	return fallbackStatusCode
}
