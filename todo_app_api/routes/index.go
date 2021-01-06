package routes

import (
	"net/http"

	"github.com/amesinp/learning_go/todo_app_api/utils"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	utils.SendSuccessResponse(utils.ResponseParams{Writer: w, Message: "Todo API v1"})
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// ConfigureRouter sets up the application routes
func ConfigureRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(jsonMiddleware)

	authRoutes(router)

	router.HandleFunc("/", index)

	return router
}
