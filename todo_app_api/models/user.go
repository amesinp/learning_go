package models

import (
	"time"
)

// User model
type User struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	UserName     string     `json:"username"`
	PasswordHash string     `json:"-"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
}
