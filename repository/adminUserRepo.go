package repository

import (
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
)

type AdminUser struct {
	ID        string  `json:"id"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Password  *string `json:"password"`
	Role      string  `json:"role"`
	Type      string  `json:"type"`
	Status    string  `json:"status"`
	CreatedAt int     `json:"created_at"`
	UpdatedAt int     `json:"updated_at"`
	LastLogin string  `json:"last_login"`
}

func GetAdminUser(params graphql.ResolveParams) (interface{}, error) {
	// In a real application, you would typically fetch the user data from a database or other data source
	id, ok := params.Args["id"].(string)
	if !ok {
		return nil, fmt.Errorf("missing id Argument")
	}
	return NewRandomAdminUser(id), nil
}

func GetAdminUsers(params graphql.ResolveParams) (interface{}, error) {
	limit, ok := params.Args["limit"].(int)
	if !ok {
		return nil, fmt.Errorf("missing limit Argument")
	}

	var users []AdminUser

	for userId := 0; userId < limit; userId++ {

		users = append(users, NewRandomAdminUser(strconv.Itoa(userId)))
	}
	return users, nil

}

func NewRandomAdminUser(id string) AdminUser {
	return AdminUser{
		ID:        id,
		Username:  "user" + id,
		Email:     "user" + id + "@example.com",
		Role:      "user",
		Type:      "user",
		Status:    "active",
		CreatedAt: 1234567890,
		UpdatedAt: 1234567890,
		LastLogin: "2021-01-01",
	}
}
