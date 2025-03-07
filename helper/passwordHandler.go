package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, bool) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", false
	}

	return string(bytes), true
}

func ValidatePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
