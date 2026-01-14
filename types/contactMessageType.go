package types

import (
	"go_server/helper"

	"github.com/graphql-go/graphql"
)

var ContactMessageType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ContactMessage",
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
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var ContactMessageRequestQueryType = helper.GlobalPaginatedQueryResolver(
	ContactMessageType,
	"ContactMessageRequestQuery",
)
