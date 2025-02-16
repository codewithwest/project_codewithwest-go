package schema

import (
	"github.com/graphql-go/graphql"
	"go_server/schema/admin"
	"go_server/schema/user"
)

func GetSchema() (graphql.Schema, error) {
	userType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	adminType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Admin",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"username": &graphql.Field{
					Type: graphql.String,
				},
				"password": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getUser":       user.GetUserSchema(userType),
				"getUsers":      user.GetUsersSchema(userType),
				"getAdminUser":  admin.GetAdminUserSchema(adminType),
				"getAdminUsers": admin.GetAdminUsersSchema(adminType),
			},
		}),
	})

	return schema, err
}
