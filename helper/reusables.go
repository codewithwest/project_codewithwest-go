package helper

import (
	"fmt"
	"log"
	"net/mail"
	"os"
)

func ValidateEmailAddress(email string) (bool, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return false, fmt.Errorf("invalid email format")
	}
	return true, nil
}

func GetEnvVariable(searchValue string) string {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	resolvedValue := os.Getenv(searchValue)

	if resolvedValue == "" {
		log.Fatal(searchValue + " environment variable not set")
	}
	return resolvedValue
}
