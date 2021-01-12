package models

import (
	"time"
)

// Todo model
type Todo struct {
	ID          int        `json:"id"`
	UserID      int        `json:"-"`
	Title       string     `json:"title"`
	Body        string     `json:"body"`
	IsCompleted bool       `json:"isCompleted"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}
