package repository

import (
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
)

type AdminUser struct {
	ID        string `json:"id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetAdminUser(params graphql.ResolveParams) (interface{}, error) {
	// In a real application, you would typically fetch the user data from a database or other data source
	id, ok := params.Args["id"].(string)
	if !ok {
		return nil, fmt.Errorf("missing id Argument")
	}
	return NewRandomAdminUser(id), nil
}

func NewRandomAdminUser(id string) AdminUser {
	return AdminUser{
		ID:        id,
		UserName:  "Alice",
		Email:     "alice@example.com",
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
}
