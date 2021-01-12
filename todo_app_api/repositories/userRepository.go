package repositories

import (
	"github.com/amesinp/learning_go/todo_app_api/config"
	"github.com/amesinp/learning_go/todo_app_api/models"
)

// UserRepository struct is to categorize the user repository functions
type UserRepository struct{}

// Get returns a single user using the user id
func (r *UserRepository) Get(id int) *models.User {
	sqlQuery := "SELECT * FROM users WHERE id=$1"
	row := config.DB.QueryRow(sqlQuery, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.UserName, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil
	}

	return &user
}

// GetByUsername returns a single user using the username
func (r *UserRepository) GetByUsername(username string) *models.User {
	sqlQuery := "SELECT * FROM users WHERE username=$1"
	row := config.DB.QueryRow(sqlQuery, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.UserName, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil
	}

	return &user
}

// Create creates a new user
func (r *UserRepository) Create(user *models.User) error {
	sqlQuery := "INSERT INTO users(name, username, password_hash) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"
	row, err := config.DB.Query(sqlQuery, user.Name, user.UserName, user.PasswordHash)

	if err != nil {
		return err
	}

	row.Next()
	row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	return nil
}

// Delete deletes a user
func (r *UserRepository) Delete(ID int) error {
	sqlQuery := "DELETE FROM users WHERE id=$1"
	_, err := config.DB.Query(sqlQuery, ID)

	if err != nil {
		return err
	}

	return nil
}
