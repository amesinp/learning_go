package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/amesinp/learning_go/todo_app_api/dto"
	"github.com/amesinp/learning_go/todo_app_api/models"
	"github.com/amesinp/learning_go/todo_app_api/repositories"
	"github.com/amesinp/learning_go/todo_app_api/utils"
)

var todoRepository = repositories.TodoRepository{}

// TodoController to categorize todo controller functions
type TodoController struct{}

type fetchResponse struct {
	Count int            `json:"count"`
	Todos *[]models.Todo `json:"todos"`
}

// FetchTodos returns paginated todos for the authenticated user
func (c *TodoController) FetchTodos(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(utils.TokenClaimsKey).(utils.TokenClaims)

	skip := utils.ConvertToInt(r.URL.Query().Get("skip"), 0)
	limit := utils.ConvertToInt(r.URL.Query().Get("limit"), 10)

	todos, count := todoRepository.GetBy(
		&repositories.TodoFilter{UserID: claims.UserID},
		&repositories.ResponseParams{Skip: skip, Limit: limit, SortAscending: true, IncludeCount: true},
	)

	utils.SendSuccessResponse(utils.ResponseParams{Writer: w, Data: fetchResponse{count, todos}})
}

// CreateTodo allows the user to create a new todo
func (c *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(utils.TokenClaimsKey).(utils.TokenClaims)

	var createData *dto.CreateTodoDTO
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&createData)

	validationMsg := utils.ValidateDTO(createData)
	if validationMsg != "" {
		utils.SendErrorResponse(utils.ResponseParams{Writer: w, Message: validationMsg})
		return
	}

	todo := models.Todo{UserID: claims.UserID, Title: createData.Title, Body: createData.Body, IsCompleted: createData.IsCompleted}

	err := todoRepository.Create(&todo)
	if err != nil {
		utils.HandleServerError(w, err)
		return
	}

	utils.SendSuccessResponse(utils.ResponseParams{Writer: w, Message: "Todo added successfully!", Data: todo})
}

// GetTodo allows the user to create a single todo using the id
func (c *TodoController) GetTodo(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(utils.TokenClaimsKey).(utils.TokenClaims)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.SendErrorResponse(utils.ResponseParams{Writer: w, Message: "Invalid id provided"})
		return
	}

	todo := todoRepository.Get(&repositories.TodoFilter{ID: id, UserID: claims.UserID})
	if todo == nil {
		utils.SendErrorResponse(utils.ResponseParams{Writer: w, Message: "Todo not found", StatusCode: http.StatusNotFound})
		return
	}

	utils.SendSuccessResponse(utils.ResponseParams{Writer: w, Data: todo})
}

// UpdateTodo allows the user to update an existing todo
func (c *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(utils.TokenClaimsKey).(utils.TokenClaims)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.SendErrorResponse(utils.ResponseParams{Writer: w, Message: "Invalid id provided"})
		return
	}

	var updateData *dto.CreateTodoDTO
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&updateData)

	validationMsg := utils.ValidateDTO(updateData)
	if validationMsg != "" {
		utils.SendErrorResponse(utils.ResponseParams{Writer: w, Message: validationMsg})
		return
	}

	todo := models.Todo{ID: id, UserID: claims.UserID, Title: updateData.Title, Body: updateData.Body, IsCompleted: updateData.IsCompleted}

	rowsAffected, err := todoRepository.Update(&todo)
	if err != nil {
		utils.HandleServerError(w, err)
		return
	}

	if rowsAffected == 0 {
		utils.SendErrorResponse(utils.ResponseParams{Writer: w, Message: "Todo not found", StatusCode: http.StatusNotFound})
		return
	}

	utils.SendSuccessResponse(utils.ResponseParams{Writer: w, Message: "Todo updated successfully!", Data: todo})
}

// DeleteTodo allows the user to update an existing todo
func (c *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(utils.TokenClaimsKey).(utils.TokenClaims)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.SendErrorResponse(utils.ResponseParams{Writer: w, Message: "Invalid id provided"})
		return
	}

	rowsAffected, err := todoRepository.Delete(&repositories.TodoFilter{ID: id, UserID: claims.UserID})
	if rowsAffected == 0 {
		utils.SendErrorResponse(utils.ResponseParams{Writer: w, Message: "Todo not found", StatusCode: http.StatusNotFound})
		return
	}

	utils.SendSuccessResponse(utils.ResponseParams{Writer: w, Message: "Todo deleted successfully!"})
}
