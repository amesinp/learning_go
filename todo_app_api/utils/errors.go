package utils

import (
	"log"
	"net/http"
)

// HandleServerError logs server errors and send a 500 response
func HandleServerError(w http.ResponseWriter, err error) {
	log.Println("Error: ", err)

	SendErrorResponse(ResponseParams{Writer: w, Message: "An error occurred. Please try again", StatusCode: http.StatusInternalServerError})
}

// HandleAuthenticationError sends a 401 response
func HandleAuthenticationError(w http.ResponseWriter) {
	SendErrorResponse(ResponseParams{Writer: w, Message: "Authentication failed!", StatusCode: http.StatusUnauthorized})
}
