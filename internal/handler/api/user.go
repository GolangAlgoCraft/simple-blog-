// user.go
package handler

import (
	"encoding/json" // Package for JSON encoding and decoding
	"net/http"      // Package for HTTP server

	"github.com/henok/blog-app/internal/auth" // Importing the auth package
)

// LoginUser handles the user login request
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var request struct {
		IDToken string `json:"idToken"` // Expecting an ID token in the request body
	}

	// Decode the JSON request body into the request struct
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Return a bad request error if decoding fails
		return
	}

	// Verify the ID token with Firebase
	token, err := auth.VerifyIDToken(request.IDToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized) // Return unauthorized error if token verification fails
		return
	}

	// Respond with the verified token
	json.NewEncoder(w).Encode(token)
}

// RegisterUser handles the user registration request
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Custom registration logic (for email/password)
}
