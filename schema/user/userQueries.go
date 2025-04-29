package user

import (
	"github.com/graphql-go/graphql"
	"go_server/helper"
	"go_server/resolver/queries"
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
		Resolve: queries.GetUser,
	}
}

func GetUsersSchema(userType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(userType),
		Description: "Get all users",
		Args:        helper.GlobalPaginatedQueriesInput,
		Resolve:     queries.GetUsers,
	}
}
