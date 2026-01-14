package mutations

import (
	"context"
	"fmt"
	"go_server/helper"
	"go_server/helper/contactReusables"
	"go_server/helper/mongoDB"
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

func CreateContactMessage(params graphql.ResolveParams) (interface{}, error) {
	inputArg, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input arguments")
	}

	name, _ := inputArg["name"].(string)
	email, _ := inputArg["email"].(string)
	message, _ := inputArg["message"].(string)

	if name == "" || email == "" || message == "" {
		return nil, fmt.Errorf("name, email, and message are required")
	}

	collection, err := mongoDB.ConnectMongoDB("contact_messages")
	if err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newMessage := contactReusables.ContactMessageMongo{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Message:   message,
		CreatedAt: helper.GetCurrentDateTime(),
	}

	_, err = collection.InsertOne(ctx, newMessage)
	if err != nil {
		return nil, fmt.Errorf("failed to insert message: %w", err)
	}

	return newMessage, nil
}
