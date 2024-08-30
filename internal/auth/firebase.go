// firebase.go
package auth

import (
	"context"
	"log"

	firebase "firebase.google.com/go" // Importing Firebase SDK
	"firebase.google.com/go/auth"     // Importing Firebase Auth package
	"google.golang.org/api/option"    // Importing Google API options package
)

var client *auth.Client // Declare a Firebase Auth client

// InitializeFirebase sets up the Firebase authentication client
func InitializeFirebase() {
	// Path to your Firebase service account key JSON file
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")

	// Initialize a Firebase app with the credentials
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Initialize Firebase Auth client
	client, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing auth client: %v\n", err)
	}
}

// VerifyIDToken verifies the Firebase ID token provided by the client
func VerifyIDToken(idToken string) (*auth.Token, error) {
	return client.VerifyIDToken(context.Background(), idToken)
}
