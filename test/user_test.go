// user_test.go
package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/henok/blog-app/internal/handler" // Importing the handler package
)

// TestLoginUser tests the LoginUser handler
func TestLoginUser(t *testing.T) {
	// Create a new HTTP POST request for the login endpoint
	req, err := http.NewRequest("POST", "/api/auth/login", nil)
	if err != nil {
		t.Fatal(err) // If the request creation fails, terminate the test
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Use the LoginUser handler to handle the request
	handler := http.HandlerFunc(handler.LoginUser)
	handler.ServeHTTP(rr, req)

	// Check the status code returned by the handler
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
