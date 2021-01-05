package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/amesinp/learning_go/todo_api_memory/util"
	"github.com/gorilla/mux"
)

type todoItem struct {
	ID          int    `json:"id"`
	Body        string `json:"body"`
	IsCompleted bool   `json:"isCompleted"`
}

type createTodoItem struct {
	Body string `json:"body"`
}

type updateTodoItem struct {
	Body        string `json:"body"`
	IsCompleted bool   `json:"isCompleted"`
}

var todos []todoItem

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang TODO API")
}

func fetchTodos(w http.ResponseWriter, r *http.Request) {
	util.SendSuccessResponse(w, http.StatusOK, todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var createData createTodoItem
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&createData); err != nil {
		util.SendErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	defer r.Body.Close()

	var todoID int
	if len(todos) == 0 {
		todoID = 1
	} else {
		todoID = todos[len(todos)-1].ID + 1
	}

	todo := todoItem{todoID, createData.Body, false}
	todos = append(todos, todo)

	util.SendSuccessResponse(w, http.StatusCreated, todo)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.SendErrorResponse(w, http.StatusBadRequest, "Invalid todo ID")
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			util.SendSuccessResponse(w, http.StatusOK, todo)
			return
		}
	}

	util.SendErrorResponse(w, http.StatusNotFound, "Invalid todo ID")
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.SendErrorResponse(w, http.StatusBadRequest, "Invalid todo ID")
		return
	}

	var updateData updateTodoItem
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&updateData); err != nil {
		util.SendErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todo.Body = updateData.Body
			todo.IsCompleted = updateData.IsCompleted
			todos[i] = todo

			util.SendSuccessResponse(w, http.StatusOK, todo)
			return
		}
	}

	util.SendErrorResponse(w, http.StatusNotFound, "Invalid todo ID")
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.SendErrorResponse(w, http.StatusBadRequest, "Invalid todo ID")
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)

			util.SendSuccessResponse(w, http.StatusOK, todo)
			return
		}
	}

	util.SendErrorResponse(w, http.StatusNotFound, "Invalid todo ID")
}

func configureRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", index)

	router.HandleFunc("/todos", fetchTodos).Methods("GET")
	router.HandleFunc("/todos", createTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	return router
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Initialize todos
	todos = []todoItem{
		todoItem{1, "Sweep the room", false},
		todoItem{2, "Clean the floor", false},
	}

	router := configureRoutes()
	router.Use(jsonMiddleware)

	const serverAddress = ":3000"
	fmt.Println("Server starting on " + serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, router))
}
