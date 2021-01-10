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
