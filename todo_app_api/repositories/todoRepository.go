package repositories

import (
	"fmt"
	"strings"
	"time"

	"github.com/amesinp/learning_go/todo_app_api/config"
	"github.com/amesinp/learning_go/todo_app_api/models"
)

// TodoFilter defines parameters to filter a todo
type TodoFilter struct {
	ID     int
	UserID int
	Title  string
}

func (f *TodoFilter) getQueryAndParams() (string, []interface{}) {
	var sqlParameters []interface{}
	var filterParams []string
	var indexTracker = 1

	if f.ID != 0 {
		filterParams = append(filterParams, fmt.Sprintf("id=$%d", indexTracker))
		sqlParameters = append(sqlParameters, f.ID)
		indexTracker++
	}
	if f.UserID != 0 {
		filterParams = append(filterParams, fmt.Sprintf("user_id=$%d", indexTracker))
		sqlParameters = append(sqlParameters, f.UserID)
		indexTracker++
	}
	if f.Title != "" {
		filterParams = append(filterParams, fmt.Sprintf("title=$%d", indexTracker))
		sqlParameters = append(sqlParameters, f.Title)
		indexTracker++
	}

	var filterQuery string
	if len(filterParams) > 0 {
		filterQuery = "WHERE " + strings.Join(filterParams, " AND ")
	} else {
		filterQuery = ""
	}

	return filterQuery, sqlParameters
}

// TodoRepository struct is to categorize the user repository functions
type TodoRepository struct{}

// GetBy returns an array of todos
func (r *TodoRepository) GetBy(filter *TodoFilter, params *ResponseParams) (*[]models.Todo, int) {
	result := []models.Todo{}
	totalCount := 0

	filterQuery, filterParams := filter.getQueryAndParams()

	sqlQuery := fmt.Sprintf(
		"SELECT * %s FROM todos %s %s",
		params.getCountQuery(),
		filterQuery,
		params.getQuerySuffix(),
	)

	rows, err := config.DB.Query(sqlQuery, filterParams...)
	if err != nil {
		return &result, totalCount
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo

		if params.IncludeCount {
			err = rows.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Body, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt, &totalCount)
		} else {
			err = rows.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Body, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt)
		}

		if err != nil {
			continue
		}
		result = append(result, todo)
	}

	return &result, totalCount
}

// Get returns a single todo
func (r *TodoRepository) Get(filter *TodoFilter) *models.Todo {
	filterQuery, filterParams := filter.getQueryAndParams()

	sqlQuery := fmt.Sprintf("SELECT * FROM todos %s", filterQuery)
	rows := config.DB.QueryRow(sqlQuery, filterParams...)

	var todo models.Todo
	err := rows.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Body, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil
	}

	return &todo
}

// Create creates a new todo
func (r *TodoRepository) Create(todo *models.Todo) error {
	sqlQuery := "INSERT INTO todos(user_id, title, body, is_completed) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at"
	row, err := config.DB.Query(sqlQuery, todo.UserID, todo.Title, todo.Body, todo.IsCompleted)

	if err != nil {
		return err
	}

	row.Next()
	row.Scan(&todo.ID, &todo.CreatedAt, &todo.UpdatedAt)
	return nil
}

// Update updates the todo
func (r *TodoRepository) Update(todo *models.Todo) (int, error) {
	currentTime := time.Now()
	todo.UpdatedAt = &currentTime

	sqlQuery := "UPDATE todos SET title=$1,body=$2,is_completed=$3,updated_at=$4 WHERE id=$5 AND user_id=$6"
	result, err := config.DB.Exec(sqlQuery, todo.Title, todo.Body, todo.IsCompleted, todo.UpdatedAt, todo.ID, todo.UserID)

	if err != nil {
		return 0, err
	}

	rowsAffected, _ := result.RowsAffected()
	return int(rowsAffected), nil
}

// Delete deletes a todo
func (r *TodoRepository) Delete(filter *TodoFilter) (int, error) {
	filterQuery, filterParams := filter.getQueryAndParams()

	sqlQuery := fmt.Sprintf("DELETE FROM todos %s", filterQuery)
	result, err := config.DB.Exec(sqlQuery, filterParams...)

	if err != nil {
		return 0, err
	}

	rowsAffected, _ := result.RowsAffected()
	return int(rowsAffected), nil
}
