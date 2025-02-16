package admin

import (
	"github.com/graphql-go/graphql"
	"go_server/helper/adminUserReusables"
	"go_server/repository"
)

func GetAdminUserSchema(adminUserType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        adminUserType,
		Description: "Get admin user by ID",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{ // Wrap it in a FieldConfigArgument
				Type:        graphql.NewNonNull(adminUserReusables.AdminUserInputDef), // Use the InputObject here
				Description: "Input for creating an admin user",                       // Description for the "input" argument
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
