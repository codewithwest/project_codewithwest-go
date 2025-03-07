package mongoDB

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Session struct {
	ID        string    `bson:"_id"`
	UserID    string    `bson:"user_id"`
	Token     string    `bson:"token"`
	CreatedAt time.Time `bson:"created_at"`
	ExpiresAt time.Time `bson:"expires_at"`
}

func GetSessionFromRequest(r *http.Request) (*Session, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, fmt.Errorf("no authorization token provided")
	}

	// Remove "Bearer " prefix if present
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	collection, err := ConnectMongoDB("sessions")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to sessions collection: %w", err)
	}

	var session Session
	err = collection.FindOne(context.Background(), bson.M{
		"token":      token,
		"expires_at": bson.M{"$gt": time.Now()},
	}).Decode(&session)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("invalid or expired session")
		}
		return nil, fmt.Errorf("error fetching session: %w", err)
	}

	return &session, nil
}

func CreateSession(userID string) (*Session, error) {
	collection, err := ConnectMongoDB("sessions")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to sessions collection: %w", err)
	}

	// Generate a secure random token
	token, err := GenerateSecureToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate session token: %w", err)
	}

	session := &Session{
		ID:        GenerateObjectID(),
		UserID:    userID,
		Token:     token,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour), // Sessions expire after 24 hours
	}

	_, err = collection.InsertOne(context.Background(), session)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return session, nil
}

func InvalidateSession(token string) error {
	collection, err := ConnectMongoDB("sessions")
	if err != nil {
		return fmt.Errorf("failed to connect to sessions collection: %w", err)
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"token": token})
	if err != nil {
		return fmt.Errorf("failed to invalidate session: %w", err)
	}

	return nil
}
