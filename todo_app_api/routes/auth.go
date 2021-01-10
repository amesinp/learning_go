package routes

import (
	"github.com/amesinp/learning_go/todo_app_api/controllers"
	"github.com/gorilla/mux"
)

func authRoutes(router *mux.Router) {
	controller := controllers.AuthController{}

	router.HandleFunc("/auth/login", controller.Login).Methods("POST")
	router.HandleFunc("/auth/register", controller.Register).Methods("POST")
}
