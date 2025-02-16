package repository

import (
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUser(params graphql.ResolveParams) (interface{}, error) {
	// In a real application, you would typically fetch the user data from a database or other data source
	id, ok := params.Args["id"].(string)
	if !ok {
		return nil, fmt.Errorf("missing id Argument")
	}
	return NewRandomUser(id), nil
}

func GetUsers(params graphql.ResolveParams) (interface{}, error) {
	limit, ok := params.Args["limit"].(int)
	if !ok {
		return nil, fmt.Errorf("missing limit Argument")
	}

	var users []User

	for userId := 0; userId < limit; userId++ {

		users = append(users, NewRandomUser(strconv.Itoa(userId)))
	}
	return users, nil

}

func NewRandomUser(id string) User {
	return User{
		ID:    id,
		Name:  "Alice",
		Email: "alice@example.com",
	}
}
