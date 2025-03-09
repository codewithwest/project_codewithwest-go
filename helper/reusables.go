package helper

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/mail"
	"os"
	"time"
)

func ValidateEmailAddress(email string) (bool, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return false, fmt.Errorf("invalid email format")
	}
	return true, nil
}

func GetEnvVariable(searchValue string) string {
	if os.Getenv("VERCEL") == "" { // The "VERCEL" env variable is set by Vercel
		if err := godotenv.Load(); err != nil {
			log.Println("Error loading .env file (development only):", err)
			// You can choose to continue or exit if the .env file is missing
		}
	}
	resolvedValue := os.Getenv(searchValue)

	if resolvedValue == "" {
		log.Fatal(searchValue + " environment variable not set")
	}
	return resolvedValue
}

func GetCurrentDateTime() string {
	resolvedTime := time.Now()
	local, err := time.LoadLocation("Local") // Or specify a specific location like "Africa/Johannesburg"
	if err != nil {
		fmt.Println("Error loading location:", err)
		return resolvedTime.Format("02-01-2006 15:04:05")

	} else {
		currentTime := resolvedTime.In(local)
		return currentTime.Format("02-01-2006 15:04:05")
	}
}
