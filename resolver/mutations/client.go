package mutations

import (
	"context"
	"fmt"
	"go_server/helper"
	"go_server/helper/clientReusables"
	"go_server/helper/mongoDB"
	"log"
	"time"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateClient(params graphql.ResolveParams) (interface{}, error) {

	userValues, validationError := clientReusables.ValidateClientInput(
		params)
	if validationError != nil {
		return nil, validationError
	}
	email := userValues[0]
	password := userValues[2]

	collection, err := mongoDB.ConnectMongoDB("clients")
	if err != nil {
		log.Fatal(err)
	}

	hashedPassword, isPassword := helper.HashPassword(password)
	if !isPassword {
		return nil, fmt.Errorf("oops! something went wrong on our side while creating your password! Please contact support")
	}

	emailExist, isEmailExists := mongoDB.EmailExist(
		collection, email)

	if isEmailExists != nil {
		return nil, fmt.Errorf("failed to convert inserted ID to ObjectID", isEmailExists)
	}
	if emailExist {
		return nil, fmt.Errorf("email already exists")
	}

	userId, userIdError := mongoDB.GetHighestIdInCollection(collection)
	if userIdError != nil {
		return nil, userIdError
	}

	client := clientReusables.NewClient(userId, userValues[2], email, userValues[3], hashedPassword)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	insertedID := result.InsertedID // No type assertion here yet

	// Convert ObjectID to string
	objectID, ok := insertedID.(primitive.ObjectID) // Type assertion to primitive.ObjectID
	if !ok {
		return nil, fmt.Errorf("failed to convert inserted ID to ObjectID")
	}

	// Adjust type assertion if needed
	var clientInput clientReusables.ClientInputMongo
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&clientInput)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve created user: %w", err)
	}

	session, err := mongoDB.CreateSession(string(rune(clientInput.ID)))
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"token": session.Token,
	}, nil

}
