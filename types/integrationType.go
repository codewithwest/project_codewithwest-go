package types

import "github.com/graphql-go/graphql"

var IntegrationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Integration",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"is_revoked": &graphql.Field{
				Type: graphql.Boolean,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
			"updated_at": &graphql.Field{
				Type: graphql.String,
			},
			"last_used_at": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
