package queries

import (
	"context"
	"fmt"
	"go_server/helper"
	"go_server/helper/contactReusables"
	"go_server/helper/mongoDB"
	"math"
	"time"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetContactMessages(params graphql.ResolveParams) (interface{}, error) {
	// Authorization check (Admin only)
	_, err := mongoDB.UserDataAccessIsAuthorized(params)
	if err != nil {
		return nil, fmt.Errorf("not authorized: %w", err)
	}
	// We could further check if the user is an admin here if needed

	limit := int32(10)
	if limitArg, ok := params.Args["limit"].(int); ok && limitArg > 0 {
		limit = int32(limitArg)
	}

	page := int32(1)
	if pageArg, ok := params.Args["page"].(int); ok && pageArg > 0 {
		page = int32(pageArg)
	}

	skip := (page - 1) * limit

	collection, err := mongoDB.ConnectMongoDB("contact_messages")
	if err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	totalItems, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error counting messages: %w", err)
	}

	findOptions := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(skip)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("error fetching messages: %w", err)
	}
	defer cursor.Close(ctx)

	var messages []contactReusables.ContactMessageMongo
	if err := cursor.All(ctx, &messages); err != nil {
		return nil, fmt.Errorf("error decoding messages: %w", err)
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	return map[string]interface{}{
		"data": messages,
		"pagination": helper.Pagination{
			CurrentPage: int(page),
			PerPage:     int(limit),
			Count:       len(messages),
			TotalPages:  totalPages,
			TotalItems:  int(totalItems),
		},
	}, nil
}
