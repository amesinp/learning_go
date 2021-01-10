package models

import (
	"time"
)

// RefreshToken model
type RefreshToken struct {
	ID         int
	UserID     int
	ClearToken string `json:"-"`
	TokenHash  string `json:"-"`
	UserAgent  string
	ExpiresAt  *time.Time
	CreatedAt  *time.Time
}
