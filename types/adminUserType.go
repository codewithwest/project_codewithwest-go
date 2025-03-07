package types

import "github.com/graphql-go/graphql"

func GetAdminUserType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "AdminUser",
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
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"role": &graphql.Field{
					Type: graphql.String,
				},
				"type": &graphql.Field{
					Type: graphql.String,
				},
				"created_at": &graphql.Field{
					Type: graphql.String,
				},
				"updated_at": &graphql.Field{
					Type: graphql.String,
				},
				"last_login": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}

func GetAdminUserRequestType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "AdminUserRequest",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"created_at": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
}
