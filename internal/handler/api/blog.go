package handler

import (
    "encoding/json"  // Package for JSON encoding and decoding
    "net/http"       // Package for HTTP server
    "strconv"        // Package for converting strings to other types

    "github.com/gorilla/mux"                   // Import the Gorilla Mux package
    "github.com/henok/blog-app/internal/repository" // Import the repository package
)

// CreateBlog handles the creation of a new blog post
func CreateBlog(w http.ResponseWriter, r *http.Request) {
    var blog repository.Blog  // Define a variable to hold the blog data
    err := json.NewDecoder(r.Body).Decode(&blog)  // Decode the request body into the blog variable
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)  // Return a 400 error if the request body is invalid
        return
    }

    id, err := repository.CreateBlog(blog)  // Call the repository function to save the blog in the database
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerErrorCertainly! Let’s continue and ensure that we include detailed comments throughout the implementation process.

### 6. **Implement Handlers and Business Logic (Continued)**

**1. Blog Handlers:**

Continue with the `CreateBlog` handler function and implement other handlers as well.

go
package handler

import (
    "encoding/json"  // Package for JSON encoding and decoding
    "net/http"       // Package for HTTP server
    "strconv"        // Package for converting strings to other types

    "github.com/gorilla/mux"                   // Import the Gorilla Mux package
    "github.com/henok/blog-app/internal/repository" // Import the repository package
)

// CreateBlog handles the creation of a new blog post
func CreateBlog(w http.ResponseWriter, r *http.Request) {
    var blog repository.Blog  // Define a variable to hold the blog data
    err := json.NewDecoder(r.Body).Decode(&blog)  // Decode the request body into the blog variable
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)  // Return a 400 error if the request body is invalid
        return
    }

    id, err := repository.CreateBlog(blog)  // Call the repository function to save the blog in the database
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)  // Return a 500 error if saving fails
        return
    }

    w.WriteHeader(http.StatusCreated)  // Return a 201 status code indicating the blog was created
    json.NewEncoder(w).Encode(map[string]int64{"id": id})  // Return the ID of the newly created blog
}

// DeleteBlog handles the deletion of a blog post by its ID
func DeleteBlog(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)  // Extract route variables (e.g., blog ID) from the request
    blogID, err := strconv.ParseInt(vars["id"], 10, 64)  // Convert the blog ID from string to int64
    if err != nil {
        http.Error(w, "Invalid blog ID", http.StatusBadRequest)  // Return a 400 error if the ID is invalid
        return
    }

    err = repository.DeleteBlog(blogID)  // Call the repository function to delete the blog
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)  // Return a 500 error if deletion fails
        return
    }

    w.WriteHeader(http.StatusNoContent)  // Return a 204 status code indicating the blog was deleted
}

// ListAllBlogs handles listing all blog posts
func ListAllBlogs(w http.ResponseWriter, r *http.Request) {
    blogs, err := repository.GetAllBlogs()  // Call the repository function to retrieve all blogs
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)  // Return a 500 error if retrieval fails
        return
    }

    json.NewEncoder(w).Encode(blogs)  // Encode the list of blogs as JSON and send it in the response
}

// ListUserBlogs handles listing all blogs for a specific user
func ListUserBlogs(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)  // Extract route variables (e.g., user ID) from the request
    userID, err := strconv.ParseInt(vars["user_id"], 10, 64)  // Convert the user ID from string to int64
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)  // Return a 400 error if the ID is invalid
        return
    }

    blogs, err := repository.GetBlogsByUser(userID)  // Call the repository function to retrieve the user's blogs
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)  // Return a 500 error if retrieval fails
        return
    }

    json.NewEncoder(w).Encode(blogs)  // Encode the list of blogs as JSON and send it in the response
}

// GetBlog handles retrieving a single blog post by its ID
func GetBlog(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)  // Extract route variables (e.g., blog ID) from the request
    blogID, err := strconv.ParseInt(vars["id"], 10, 64)  // Convert the blog ID from string to int64
    if err != nil {
        http.Error(w, "Invalid blog ID", http.StatusBadRequest)  // Return a 400 error if the ID is invalid
        return
    }

    blog, err := repository.GetBlogByID(blogID)  // Call the repository function to retrieve the blog
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)  // Return a 500 error if retrieval fails
        return
    }

    json.NewEncoder(w).Encode(blog)  // Encode the blog as JSON and send it in the response
}
