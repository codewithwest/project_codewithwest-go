package main

import (
	"context"
	"flag"
	"fmt"
	"go_server/helper"
	"go_server/helper/clientReusables"
	"go_server/helper/mongoDB"
	"log"
	"strconv"
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
	username := flag.String("username", "", "Client username")
	email := flag.String("email", "", "Client email")
	password := flag.String("password", "", "Client password")
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
	collection, err := mongoDB.ConnectMongoDB("clients")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if user already exists
	var existingUser clientReusables.ClientInputMongo
	err = collection.FindOne(ctx, bson.M{"email": *email}).Decode(&existingUser)
	if err == nil {
		log.Fatalf("Client with email %s already exists", *email)
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

	// Create user input
	userInput := &clientReusables.ClientInputData{
		UserName: *username,
		Email:    *email,
		Password: hashedPassword,
	}

	client := clientReusables.CreateNewClientInput(highestId, userInput)

	_, err = collection.InsertOne(ctx, client)
	if err != nil {
		log.Fatalf("Failed to insert client: %v", err)
	}

	// Generate Token
	session, err := mongoDB.CreateSession(
		strconv.Itoa(client.ID),
		*email,
		true, // isClient = true
	)
	if err != nil {
		log.Fatalf("Failed to create session/token: %v", err)
	}

	fmt.Printf("\n--- Client Created Successfully ---\n")
	fmt.Printf("Username: %s\n", *username)
	fmt.Printf("Email:    %s\n", *email)
	fmt.Printf("ID:       %d\n", client.ID)
	fmt.Printf("\n--- Session Token ---\n")
	fmt.Printf("%s\n", session.Token)
	fmt.Printf("-------------------------\n")
}
