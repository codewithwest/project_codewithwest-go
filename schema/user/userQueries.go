package user

import (
	"github.com/graphql-go/graphql"
	"go_server/repository"
)

func GetUserSchema(userType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        userType,
		Description: "Get a user by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: repository.GetUser,
	}
}

func GetUsersSchema(userType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(userType),
		Description: "Get all users",
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: repository.GetUsers,
	}
}
