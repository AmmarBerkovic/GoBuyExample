package main

import (
	"log"
	"net/http"

	"github.com/AmmarBerkovic/GoBuyExample/db"
	"github.com/AmmarBerkovic/GoBuyExample/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialize the database connection
	db.ConnectDatabase()

	// Create a new router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)    // Logs incoming requests
	r.Use(middleware.Recoverer) // Graceful panic recovery

	// Define routes
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.GetUsers)          // Fetch all users
		r.Post("/", handlers.CreateUser)       // Create a new user
		r.Put("/{id}", handlers.UpdateUser)    // Update user
		r.Delete("/{id}", handlers.DeleteUser) // Delete user
	})

	// Start the server
	log.Println("Server is running on :8080...")
	http.ListenAndServe(":8080", r)
}
