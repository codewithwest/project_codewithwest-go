package main

import (
	"context"
	"fmt"
	"go_server/helper/integrationReusables"
	"go_server/helper/mongoDB"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/create-integration/main.go <name>")
		return
	}

	name := os.Args[1]
	token, err := mongoDB.GenerateSecureToken()
	if err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}

	newIntegration := integrationReusables.NewIntegration(name, token)

	collection, err := mongoDB.ConnectMongoDB("integrations")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if already exists
	var existing integrationReusables.Integration
	err = collection.FindOne(ctx, bson.M{"name": name}).Decode(&existing)
	if err == nil {
		log.Fatalf("Integration with name %s already exists", name)
	}

	_, err = collection.InsertOne(ctx, newIntegration)
	if err != nil {
		log.Fatalf("Failed to insert integration: %v", err)
	}

	fmt.Printf("Successfully created integration: %s\n", name)
	fmt.Printf("Token: %s\n", token)
	fmt.Println("Please save this token. It will not be shown again.")
}
