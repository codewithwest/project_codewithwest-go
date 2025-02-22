package helper

import "github.com/graphql-go/graphql"

func GetProjectInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name:        "ProjectInput",
		Description: "Input for creating a new project",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": (*graphql.InputObjectFieldConfig)(&graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			}),
			"project_category_id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"tech_stacks": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
			},
			"github_link": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"live_link": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"test_link": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})
}

func NewProjectCategory(id int, name string) ProjectCategoryMongo {
	return ProjectCategoryMongo{
		ID:        id + 1,
		Name:      name,
		CreatedAt: GetCurrentDateTime(),
		UpdatedAt: GetCurrentDateTime(),
	}
}

var ProjectInput = GetProjectInput()
