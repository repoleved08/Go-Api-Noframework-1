package main

import (
	"log"
	"net/http"

	"github.com/repoleved08/go-api-nof-1/database"
	"github.com/repoleved08/go-api-nof-1/handlers"
)

func main() {
	// Initialize database connection
	database.InitDB()
	// Create new ServerMux
	router := http.NewServeMux()
	// Register Routes
	router.HandleFunc("/users/create", handlers.CreateUser)
	router.HandleFunc("/users", handlers.GetUsers)
	router.HandleFunc("/users/get", handlers.GetUserByID) 
	router.HandleFunc("/users/update", handlers.UpdateUser)
	router.HandleFunc("/users/delete", handlers.DeleteUser)

	// Serve
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
