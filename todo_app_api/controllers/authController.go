package controllers

import (
	"net/http"

	"github.com/amesinp/learning_go/todo_app_api/utils"
)

// AuthController to categorize controller functions
type AuthController struct{}

// Login function
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	utils.SendSuccessResponse(utils.ResponseParams{Writer: w, Message: "Login successful!"})
}
