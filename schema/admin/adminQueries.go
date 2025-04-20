package admin

import (
	"github.com/graphql-go/graphql"
	"go_server/helper/adminUserReusables"
	"go_server/resolver/queries"
)

func GetLoginAdminUserSchema(adminUserType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        adminUserType,
		Description: "Login admin user by email",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(adminUserReusables.AdminUserInputDef), // Use the InputObject here
				Description: "Input for creating an admin user",                       // Description for the "input" argument
			},
		},
		Resolve: queries.LoginAdminUser,
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
		Resolve: queries.GetAdminUsers,
	}
}

func GetAdminUserRequests(adminUserType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Name: "AdminUserRequest",
		Type: graphql.NewList(adminUserType),
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve:     queries.GetAdminUserRequests,
		Description: "Get all admin user requests",
	}
}
