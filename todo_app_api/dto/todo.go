package dto

// CreateTodoDTO defines the parameters for the create todo endpoint
type CreateTodoDTO struct {
	Title       string `json:"title" validate:"required"`
	Body        string `json:"body"`
	IsCompleted bool   `json:"isCompleted"`
}
