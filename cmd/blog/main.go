// main.go
package main

import (
	"log"      // Package for logging
	"net/http" // Package for HTTP server

	"github.com/henok/blog-app/internal/router" // Importing our router package
	"github.com/henok/blog-app/pkg/database"    // Importing our database package
)

func main() {
	// Initialize the database connection
	err := database.InitMySQL("user:password@tcp(db:3306)/blog")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close() // Ensure the database connection is closed when the app stops

	// Set up the router
	r := router.New()

	// Start the HTTP server on port 8080
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
