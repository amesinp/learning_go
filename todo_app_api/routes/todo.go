package routes

import (
	"github.com/gorilla/mux"

	"github.com/amesinp/learning_go/todo_app_api/controllers"
)

func todoRoutes(router *mux.Router) {
	controller := controllers.TodoController{}

	todoRoutes := router.PathPrefix("/todos").Subrouter()
	todoRoutes.Use(AuthMiddleware)

	todoRoutes.HandleFunc("", controller.FetchTodos).Methods("GET")
	todoRoutes.HandleFunc("", controller.CreateTodo).Methods("POST")
	todoRoutes.HandleFunc("/{id}", controller.GetTodo).Methods("GET")
	todoRoutes.HandleFunc("/{id}", controller.UpdateTodo).Methods("PUT")
	todoRoutes.HandleFunc("/{id}", controller.DeleteTodo).Methods("DELETE")
}
