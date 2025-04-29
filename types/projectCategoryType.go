package types

import (
	"github.com/graphql-go/graphql"
	"go_server/helper"
)

var ProjectCategoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ProjectCategory",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
			"updated_at": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var ProjectCategoryRequestQueryType = helper.GlobalPaginatedQueryResolver(
	ProjectCategoryType,
	"ProjectCategoriesType",
)
