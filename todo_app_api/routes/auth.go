package routes

import (
	"net/http"

	"github.com/amesinp/learning_go/todo_app_api/controllers"
	"github.com/gorilla/mux"
)

// RefreshTokenURI stores the uri of the refresh token enpdoint
const RefreshTokenURI = "/auth/refresh"

func authRoutes(router *mux.Router) {
	controller := controllers.AuthController{}

	router.HandleFunc("/auth/login", controller.Login).Methods("POST")
	router.HandleFunc("/auth/register", controller.Register).Methods("POST")
	router.Handle(RefreshTokenURI, AuthMiddleware(http.HandlerFunc(controller.RefreshToken))).Methods("POST")
	router.Handle("/auth/logout", AuthMiddleware(http.HandlerFunc(controller.Logout))).Methods("POST")
}
