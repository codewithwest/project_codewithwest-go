package admin

import (
	"github.com/graphql-go/graphql"
	"go_server/repository"
)

func GetAdminUserSchema(adminUserType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        adminUserType,
		Description: "Get admin user by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: repository.GetAdminUser,
	}
}

func GetAdminUsersSchema(adminUserType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(adminUserType),
		Description: "Get all admin users",
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: repository.GetAdminUsers,
	}
}
