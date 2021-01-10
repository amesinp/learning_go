package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/amesinp/learning_go/todo_app_api/config"
	"github.com/amesinp/learning_go/todo_app_api/routes"
	"github.com/amesinp/learning_go/todo_app_api/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	config.InitializeDb()

	err = config.DB.Ping()
	if err != nil {
		log.Fatal("Could not ping db- ", err)
	}

	// Configure validator
	utils.SetupValidator()

	serverAddress := ":" + os.Getenv("PORT")
	log.Print("Server starting at " + serverAddress)

	log.Fatal(http.ListenAndServe(serverAddress, routes.ConfigureRouter()))
}
