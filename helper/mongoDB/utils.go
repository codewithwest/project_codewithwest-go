package mongoDB

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GenerateSecureToken creates a cryptographically secure random token
func GenerateSecureToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("error generating random token: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// GenerateObjectID generates a new MongoDB ObjectID as a string
func GenerateObjectID() string {
	return primitive.NewObjectID().Hex()
}
