package repositories

import (
	"github.com/amesinp/learning_go/todo_app_api/config"
	"github.com/amesinp/learning_go/todo_app_api/models"
)

// RefreshTokenRepository struct is to categorize the user repository functions
type RefreshTokenRepository struct{}

// Create creates a new refresh token
func (r *RefreshTokenRepository) Create(token *models.RefreshToken) error {
	sqlQuery := "INSERT INTO refresh_tokens(token_hash, user_id, user_agent, expires_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at"
	row, err := config.DB.Query(sqlQuery, token.TokenHash, token.UserID, token.UserAgent, token.ExpiresAt)

	if err != nil {
		return err
	}

	row.Next()
	row.Scan(&token.ID, &token.CreatedAt)
	return nil
}

// Delete deletes a refresh token
func (r *RefreshTokenRepository) Delete(ID int) error {
	sqlQuery := "DELETE FROM refresh_tokens WHERE id=$1"
	_, err := config.DB.Query(sqlQuery, ID)

	if err != nil {
		return err
	}

	return nil
}
