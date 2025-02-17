package admin

import (
	"go_server/helper/adminUserReusables"
	"go_server/repository/mutations"

	"github.com/graphql-go/graphql"
)

func CreateUserMutation(requiredType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        requiredType, // Return the user type after creation
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
