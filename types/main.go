package types

import (
	"github.com/graphql-go/graphql"
	"go_server/schema/admin"
	"go_server/schema/user"
)

func GetSchema() (graphql.Schema, error) {

	var adminUserType = GetAdminUserType()
	var userType = GetUserType()
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getUser":       user.GetUserSchema(userType),
				"getUsers":      user.GetUsersSchema(userType),
				"getAdminUser":  admin.GetAdminUserSchema(adminUserType),
				"getAdminUsers": admin.GetAdminUsersSchema(adminUserType),
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"createAdminUser": admin.CreateUserMutation(adminUserType),
			},
		}),
	})

	return schema, err
}
