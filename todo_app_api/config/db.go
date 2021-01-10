package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	// Postgres driver import
	_ "github.com/lib/pq"
)

// DB is the sql.DB instance for the database
var DB *sql.DB

// InitializeDb creates the database connection
func InitializeDb() {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Database port is not valid")
	}

	dbConnString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		dbPort,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("postgres", dbConnString)
	if err != nil {
		log.Fatal("Could not connect to db- ", err)
	}
}
