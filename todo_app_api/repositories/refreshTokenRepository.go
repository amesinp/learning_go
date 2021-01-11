package repositories

import (
	"github.com/amesinp/learning_go/todo_app_api/config"
	"github.com/amesinp/learning_go/todo_app_api/models"
)

// RefreshTokenRepository struct is to categorize the user repository functions
type RefreshTokenRepository struct{}

// Get returns a single refresh token using the token id
func (r *RefreshTokenRepository) Get(id int) *models.RefreshToken {
	sqlQuery := "SELECT * FROM refresh_tokens WHERE id=$1"
	row := config.DB.QueryRow(sqlQuery, id)

	var token models.RefreshToken
	err := row.Scan(&token.ID, &token.TokenHash, &token.UserID, &token.UserAgent, &token.ExpiresAt, &token.CreatedAt, &token.IsUsed)
	if err != nil {
		return nil
	}

	return &token
}

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

// DeleteByUserID deletes all the refresh tokens for a particular user
func (r *RefreshTokenRepository) DeleteByUserID(userID int) error {
	sqlQuery := "DELETE FROM refresh_tokens WHERE user_id=$1"
	_, err := config.DB.Query(sqlQuery, userID)

	if err != nil {
		return err
	}

	return nil
}

// UpdateToUsed updates the status of a refresh token to indicate that it has been used
func (r *RefreshTokenRepository) UpdateToUsed(ID int) error {
	sqlQuery := "UPDATE refresh_tokens SET is_used=true WHERE id=$1"
	_, err := config.DB.Query(sqlQuery, ID)

	if err != nil {
		return err
	}

	return nil
}
