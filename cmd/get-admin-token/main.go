package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"go_server/helper"
	"go_server/helper/adminUserReusables"
	"go_server/helper/mongoDB"
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// Load environment variables
	if err := godotenv.Load("../../.env"); err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}

	// Define flags
	email := flag.String("email", "", "Admin email")
	password := flag.String("password", "", "Admin password")
	flag.Parse()

	if *email == "" || *password == "" {
		fmt.Println("Usage: go run main.go -email <email> -password <pass>")
		return
	}

	// Connect to MongoDB
	collection, err := mongoDB.ConnectMongoDB("admin_users")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find user
	var adminUser adminUserReusables.AdminUserInputMongo
	err = collection.FindOne(ctx, bson.M{"email": *email}).Decode(&adminUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Fatalf("Admin user not found with email: %s", *email)
		}
		log.Fatalf("Database error: %v", err)
	}

	// Validate password
	if !helper.ValidatePassword(*password, *adminUser.Password) {
		log.Fatal("Invalid password")
	}

	// Create session
	session, err := mongoDB.CreateSession(
		strconv.Itoa(adminUser.ID),
		*email,
		false, // isClient = false (it's an admin)
	)
	if err != nil {
		log.Fatalf("Session creation failed: %v", err)
	}

	fmt.Printf("\n--- Admin Token Retrieved ---\n")
	fmt.Printf("Email: %s\n", *email)
	fmt.Printf("ID:    %d\n", adminUser.ID)
	fmt.Printf("\n--- Token ---\n")
	fmt.Printf("%s\n", session.Token)
	fmt.Printf("-------------\n")
}
