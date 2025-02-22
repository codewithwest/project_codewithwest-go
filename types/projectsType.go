package types

import "github.com/graphql-go/graphql"

func GetProjectCategoryType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "ProjectCategory",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
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
}

func GetProjectType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Project",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"project_category_id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
				"tech_stacks": &graphql.Field{
					Type: graphql.NewList(graphql.String),
				},
				"github_link": &graphql.Field{
					Type: graphql.String,
				},
				"live_link": &graphql.Field{
					Type: graphql.String,
				},
				"test_link": &graphql.Field{
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
}
