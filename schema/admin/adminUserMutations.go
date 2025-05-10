package admin

import (
	"go_server/helper/adminUserReusables"
	"go_server/resolver/mutations"

	"github.com/graphql-go/graphql"
)

func CreateAdminUserMutation(requiredType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        requiredType,
		Description: "Create a new admin user",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{ // Wrap it in a FieldConfigArgument
				Type:        graphql.NewNonNull(adminUserReusables.AdminUserInputDef), // Use the InputObject here
				Description: "Input for creating an admin user",                       // Description for the "input" argument
			},
		},
		Resolve: mutations.CreateAdminUser,
	}
}

func RequestAdminAccess(requiredType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Name:        "AdminUserRequestInput",
		Type:        requiredType,
		Description: "Request admin access",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: mutations.CreateAdminUserRequest,
	}
}
