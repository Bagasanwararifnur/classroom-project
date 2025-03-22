package main

import (
	"log"
	"net/http"

	"cls-project/db"
	"cls-project/handler"

	_ "cls-project/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title My API
// @version 1.0
// @description This is the API documentation for the app
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	// Initialize the database connection
	db.InitDB()

	// Initialize the router
	router := mux.NewRouter()

	// Swagger UI route
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// API Routes
	router.HandleFunc("/users", handler.GetUser).Methods("GET")
	router.HandleFunc("/users", handler.CreateUser).Methods("POST")
	router.HandleFunc("/details-user", handler.GetUserDetails).Methods("GET")
	router.HandleFunc("/user-thread", handler.GetUserThreads).Methods("GET")
	router.HandleFunc("/user-comment", handler.GetUserComments).Methods("GET")

	router.HandleFunc("/thread-create", handler.CreateThread).Methods("POST")
	router.HandleFunc("/thread-get", handler.GetThreads).Methods("GET")
	router.HandleFunc("/get-thread", handler.GetThreadByID).Methods("GET")
	router.HandleFunc("/thread-page", handler.GetThreadPage).Methods("GET")

	router.HandleFunc("/comment-create", handler.CreateComment).Methods("POST")

	// Start the server
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
