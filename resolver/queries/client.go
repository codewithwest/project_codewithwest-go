package queries

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"go_server/helper"
	"go_server/helper/clientReusables"
	"go_server/helper/mongoDB"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthenticateClient(params graphql.ResolveParams) (interface{}, error) {
	// Use type assertion with direct error checking
	inputArg, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input arguments")
	}

	// Validate input fields in one block
	email, isEmail := inputArg["email"].(string)
	password, isPassword := inputArg["password"].(string)
	if !isEmail || !isPassword || email == "" || password == "" {
		return nil, fmt.Errorf("invalid email or password")
	}

	// Use a shorter context timeout for DB operations
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	collection, err := mongoDB.ConnectMongoDB("clients")
	if err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}

	// Use a single database query with proper error handling
	var client clientReusables.ClientInputMongo
	if err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&client); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("invalid credentials")
		}
		return nil, fmt.Errorf("internal server error: %w", err)
	}

	// Validate password
	if client.Password == nil || !helper.ValidatePassword(password, *client.Password) {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Create session
	session, err := mongoDB.CreateSession(strconv.Itoa(client.ID), email, true)
	if err != nil {
		return nil, fmt.Errorf("session creation failed: %w", err)
	}

	return clientReusables.ClientInputReturnData(session, client), nil
}
