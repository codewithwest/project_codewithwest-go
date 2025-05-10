package admin

import (
	"github.com/graphql-go/graphql"
	"go_server/helper"
	"go_server/helper/adminUserReusables"
	"go_server/resolver/queries"
)

func GetLoginAdminUserSchema(adminUserType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        adminUserType,
		Description: "Login admin user by email",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(adminUserReusables.AdminUserInputDef),
				Description: "Input for creating an admin user",
			},
		},
		Resolve: queries.LoginAdminUser,
	}
}

func GetAdminUsersSchema(adminUserType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        adminUserType,
		Description: "Get all admin users",
		Args:        adminUserReusables.AdminUserQueriesInput,
		Resolve:     queries.GetAdminUsers,
	}
}

func GetAdminUserRequests(adminUserType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Name:        "AdminUserRequest",
		Type:        adminUserType,
		Args:        helper.GlobalPaginatedQueriesInput,
		Resolve:     queries.GetAdminUserRequests,
		Description: "Get all admin user requests",
	}
}
