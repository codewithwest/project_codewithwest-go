package mongoDB

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_server/helper/adminUserReusables"
)

//var RetrievedCollection *mongo.Collection

func ConnectMongoDB(uri, databaseName, collectionName string) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}
	var collection *mongo.Collection
	collection = client.Database(databaseName).Collection(collectionName)

	fmt.Println("Connected to MongoDB!")
	return collection, nil
}

func CreateIndex(collection *mongo.Collection, indexValue string) (*string, error) {
	// Create the index here, after establishing the connection:
	idIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: indexValue, Value: -1}},
		Options: options.Index().SetUnique(true), // Optional: If ID is always unique
	}

	_, indexError := collection.Indexes().CreateOne(context.Background(), idIndex)
	if indexError != nil {
		fmt.Printf("Warning: Failed to create index: %v\n", indexError)
		return nil, indexError
	}
	return &indexValue, nil
}

func GetHighestIdInCollection(collection *mongo.Collection) (int, error) {
	_, createIdIndexError := CreateIndex(collection, "id")
	if createIdIndexError != nil {
		return 0, createIdIndexError
	}

	opts := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})

	var result adminUserReusables.AdminUserInputMongo
	err := collection.FindOne(context.Background(), bson.D{}, opts).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return 0, nil // No documents found
		}
		return 0, fmt.Errorf("finding highest ID: %w", err)
	}

	fmt.Printf("Highest ID: %d\n", result.ID)
	return result.ID, nil
}

func EmailExist(collection *mongo.Collection, email string) (bool, error) {
	_, createEmailIndexError := CreateIndex(collection, "email")
	if createEmailIndexError != nil {
		return false, createEmailIndexError
	}

	filter := bson.M{"email": email} // Create a filter for the email

	var result struct {
		Email string `bson:"email"` // Only need the email field in the result
	}

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, fmt.Errorf("checking email existence: %w", err) // Other error
	}

	return true, nil
}

func NameExists(collection *mongo.Collection, name string) (bool, error) {
	_, createNameIndexError := CreateIndex(collection, "name")
	if createNameIndexError != nil {
		return false, createNameIndexError
	}

	filter := bson.M{"name": name} // Create a filter for the email

	var result struct {
		Name string `bson:"name"` // Only need the email field in the result
	}

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, fmt.Errorf("checking email existence: %w", err) // Other error
	}

	return true, nil
}
