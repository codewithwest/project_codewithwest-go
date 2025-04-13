package mutations

import (
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go_server/helper"
	"go_server/helper/clientReusables"
	"go_server/helper/mongoDB"
	"strconv"
	"time"
)

func CreateClient(params graphql.ResolveParams) (interface{}, error) {
	// Validate input
	clientInputData, err := clientReusables.ValidateClientInput(params)
	if err != nil {
		return nil, fmt.Errorf("input validation failed: %w", err)
	}

	// Get MongoDB collection with connection pooling
	collection, err := mongoDB.ConnectMongoDB("clients")
	if err != nil {
		// Don't use log.Fatal as it terminates the program
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Check if email exists (moved up for early validation)
	exists, err := mongoDB.EmailExist(collection, clientInputData.Email)
	if err != nil {
		return nil, fmt.Errorf("email check failed: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("email already exists")
	}

	// Check if username exists
	exists, err = mongoDB.UserNameExists(collection, clientInputData.UserName)
	if err != nil {
		return nil, fmt.Errorf("username check failed: %w", err)
	}

	if exists {
		return nil, fmt.Errorf("username already exists")
	}
	// Hash password
	hashedPassword, ok := helper.HashPassword(clientInputData.Password)
	if !ok {
		return nil, fmt.Errorf("password hashing failed")
	}
	clientInputData.Password = hashedPassword

	// Get next user ID
	userId, err := mongoDB.GetHighestIdInCollection(collection)
	if err != nil {
		return nil, fmt.Errorf("failed to generate user ID: %w", err)
	}

	// Create new client document
	client := clientReusables.CreateNewClientInput(userId, clientInputData)

	// Insert client with error handling
	result, err := collection.InsertOne(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("client creation failed: %w", err)
	}

	// Type assert and verify inserted ID
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("invalid inserted ID type")
	}

	// Retrieve created client
	var clientInput clientReusables.ClientInputMongo
	if err := collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&clientInput); err != nil {
		return nil, fmt.Errorf("failed to retrieve created client: %w", err)
	}

	// Create session
	session, err := mongoDB.CreateSession(
		strconv.Itoa(client.ID),
		clientInputData.Email,
		true,
	)
	if err != nil {
		return nil, fmt.Errorf("session creation failed: %w", err)
	}

	response := clientReusables.ClientInputReturnData(session, clientInput)

	return response, nil
}
