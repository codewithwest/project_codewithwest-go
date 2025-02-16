package helper

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var RetrievedCollection *mongo.Collection

func ConnectMongoDB(uri, databaseName, collectionName string) error {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	RetrievedCollection = client.Database(databaseName).Collection(collectionName)
	fmt.Println("Connected to MongoDB!")
	return nil
}
