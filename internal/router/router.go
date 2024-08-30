package router

import (
	"github.com/gorilla/mux"                     // Import the Gorilla Mux router package
	"github.com/henok/blog-app/internal/handler" // Import the handler package
)

func New() *mux.Router {
	r := mux.NewRouter() // Create a new router instance

	// Define routes and their corresponding handlers
	r.HandleFunc("/api/auth/login", handler.LoginUser).Methods("POST")
	r.HandleFunc("/api/auth/register", handler.RegisterUser).Methods("POST")
	r.HandleFunc("/api/blogs", handler.ListAllBlogs).Methods("GET")
	r.HandleFunc("/api/blogs", handler.CreateBlog).Methods("POST")
	r.HandleFunc("/api/blogs/{id}", handler.GetBlog).Methods("GET")
	r.HandleFunc("/api/blogs/{id}", handler.DeleteBlog).Methods("DELETE")
	r.HandleFunc("/api/users/{user_id}/blogs", handler.ListUserBlogs).Methods("GET")

	return r // Return the configured router
}
