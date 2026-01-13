package main

import (
	"context"
	"flag"
	"fmt"
	"go_server/helper"
	"go_server/helper/adminUserReusables"
	"go_server/helper/mongoDB"
	"log"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// Load environment variables
	if err := godotenv.Load("../../.env"); err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}

	// Define flags
	username := flag.String("username", "", "Admin username")
	email := flag.String("email", "", "Admin email")
	password := flag.String("password", "", "Admin password")
	flag.Parse()

	if *username == "" || *email == "" || *password == "" {
		fmt.Println("Usage: go run main.go -username <name> -email <email> -password <pass>")
		return
	}

	// Validate email
	if _, err := helper.ValidateEmailAddress(*email); err != nil {
		log.Fatalf("Invalid email address: %v", err)
	}

	// Connect to MongoDB
	collection, err := mongoDB.ConnectMongoDB("admin_users")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if user already exists
	var existingUser adminUserReusables.AdminUserInputMongo
	err = collection.FindOne(ctx, bson.M{"email": *email}).Decode(&existingUser)
	if err == nil {
		log.Fatalf("User with email %s already exists", *email)
	}

	// Hash password
	hashedPassword, ok := helper.HashPassword(*password)
	if !ok {
		log.Fatal("Failed to hash password")
	}

	// Get highest ID
	highestId, err := mongoDB.GetHighestIdInCollection(collection)
	if err != nil {
		log.Fatalf("Failed to get highest ID: %v", err)
	}

	// Create user
	userInput := &adminUserReusables.AdminUserInputData{
		UserName: *username,
		Email:    *email,
		Password: hashedPassword,
	}

	user := adminUserReusables.NewAdminUser(highestId, userInput)

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	fmt.Printf("Successfully created admin user: %s (%s)\n", *username, *email)
}
