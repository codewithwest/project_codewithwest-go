package mongoDB

import (
	"context"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"go_server/helper"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Session struct {
	ID        string    `json:"id" bson:"_id"`
	UserID    string    `json:"user_id" bson:"user_id"`
	Email     string    `json:"email" bson:"email"`
	Token     string    `json:"token" bson:"token"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	ExpiresAt time.Time `json:"expires_at" bson:"expires_at"`
}

func ResolveToken(r *http.Request, key string) (string, error) {
	token := r.Header.Get(key)
	if token == "" {
		return "", fmt.Errorf("no authorization token provided")
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}
	return token, nil
}

func GetSessionFromRequest(r *http.Request, key string) (*Session, error) {
	token, tokenError := ResolveToken(r, key)
	if tokenError != nil {
		return nil, tokenError
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
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("invalid or expired session")
		}
		return nil, fmt.Errorf("error fetching session: %w", err)
	}

	return &session, nil
}

func CreateSession(userID string, email string, year bool) (*Session, error) {
	collection, err := ConnectMongoDB("sessions")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to sessions collection: %w", err)
	}

	// Generate a secure random token
	token, err := GenerateSecureToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate session token: %w", err)
	}

	var ExpiresAt = time.Now().AddDate(0, 0, 14)
	if year {
		ExpiresAt = time.Now().AddDate(1, 0, 0)
	}

	var existingSession Session

	filter := bson.M{
		"$or": []bson.M{
			{"email": email},
			{"userID": userID},
		},
	}

	findSessionErr := collection.FindOne(context.Background(), filter).Decode(&existingSession)
	if findSessionErr == nil {
		// Session exists, update it
		update := bson.M{
			"$set": bson.M{
				"token":      token,
				"expires_at": ExpiresAt,
				"last_login": helper.GetCurrentDateTime(),
			},
		}

		updateResult, err := collection.UpdateOne(context.Background(), bson.M{"email": email}, update)
		if err != nil {
			return nil, fmt.Errorf("error updating session: %w", err)
		}

		if updateResult.ModifiedCount == 0 {
			return nil, fmt.Errorf("no session updated")
		}

		existingSession.UserID = userID
		existingSession.Token = token
		existingSession.ExpiresAt = ExpiresAt

		return &existingSession, nil

	} else if errors.Is(findSessionErr, mongo.ErrNoDocuments) {
		// Session doesn't exist, create a new one
		session := &Session{
			ID:        GenerateObjectID(),
			UserID:    userID,
			Email:     email,
			Token:     token,
			CreatedAt: time.Now(),
			ExpiresAt: ExpiresAt,
		}

		_, err := collection.InsertOne(context.Background(), session)
		if err != nil {
			return nil, fmt.Errorf("error inserting session: %w", err)
		}

		return session, nil

	} else {
		return nil, fmt.Errorf("error fetching session: %w", findSessionErr)
	}
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

func UserDataAccessIsAuthorized(params graphql.ResolveParams) (string, error) {
	requestContext := params.Context
	request, ok := requestContext.Value("http.Request").(*http.Request)

	if !ok {

		return "", fmt.Errorf("http.Request not found in context")
	}

	validSignature, isValidSignatureErr := GetSessionFromRequest(request, "Authorization")
	if isValidSignatureErr != nil || validSignature == nil {
		return "", isValidSignatureErr
	}
	userId := request.Header.Get("user_id")

	if userId == "" {
		return "", fmt.Errorf("userId header missing")
	}

	return userId, nil
}
