package schema

import (
	"github.com/digvijay17july/golang-projects/go-graphql-rest-api/repository"

	"github.com/graphql-go/graphql"
)

func GetAdminUserType() *graphql.Object {
	adminUserType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "AdminUser",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"username": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"password": &graphql.Field{
					Type: graphql.String,
				},
				"created_at": &graphql.Field{
					Type: graphql.Int,
				},
				"last_login": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
	return adminUserType
}

func GetAdminUserSchema(adminUserType *graphql.Object) (graphql.Schema, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"adminUser": &graphql.Field{
					Type:        adminUserType,
					Description: "Get admin user by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.ID),
						},
					},
					Resolve: repository.GetUser,
				},
			},
		}),
	})

	return schema, err
}
