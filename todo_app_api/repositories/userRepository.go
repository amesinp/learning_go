package repositories

import (
	"fmt"

	"github.com/amesinp/learning_go/todo_app_api/config"
	"github.com/amesinp/learning_go/todo_app_api/models"
)

// UserRepository struct is to categorize the user repository functions
type UserRepository struct{}

// GetBy returns an array of users
func (r *UserRepository) GetBy(skip int, limit int, includeCount bool) []models.User {
	var sqlQuery string
	if includeCount {
		sqlQuery = fmt.Sprintf("SELECT *, COUNT(*) OVER() AS total_count FROM users LIMIT %d OFFSET %d", limit, skip)
	} else {
		sqlQuery = fmt.Sprintf("SELECT * FROM users LIMIT %d OFFSET %d", limit, skip)
	}

	rows, err := config.DB.Query(sqlQuery)
	if err != nil {
		return nil
	}
	defer rows.Close()

	result := []models.User{}
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name, &user.UserName, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			continue
		}
		result = append(result, user)
	}

	return result
}

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
