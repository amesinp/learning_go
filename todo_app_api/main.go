package main

import (
	"log"
	"net/http"
	"os"

	"github.com/amesinp/learning_go/todo_app_api/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverAddress := ":" + os.Getenv("PORT")
	log.Print("Server starting at " + serverAddress)

	log.Fatal(http.ListenAndServe(serverAddress, routes.ConfigureRouter()))
}
