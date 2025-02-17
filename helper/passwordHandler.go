package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, bool) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", false // Return empty string and false on error
	}
	return string(bytes), true // Return hashed password and true on success
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
