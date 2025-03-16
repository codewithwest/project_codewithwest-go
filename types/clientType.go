package types

import "github.com/graphql-go/graphql"

var ClientType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Client",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"company_name": &graphql.Field{
				Type: graphql.String,
			},
			"last_login": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
