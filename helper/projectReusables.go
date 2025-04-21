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

func GetOptionalString(input map[string]interface{}, key string) string {
	if val, ok := input[key].(string); ok {
		return val
	}
	return ""
}

func NewProject(
	id int,
	projectCategoryId int,
	name string,
	description string,
	techStacks []string,
	inputArg map[string]interface{},
) ProjectMongo {
	return ProjectMongo{
		ID:                id + 1,
		ProjectCategoryId: projectCategoryId,
		Name:              name,
		Description:       description,
		TechStacks:        techStacks,
		GithubLink:        GetOptionalString(inputArg, "github_link"),
		LiveLink:          GetOptionalString(inputArg, "live_link"),
		TestLink:          GetOptionalString(inputArg, "test_link"),
		CreatedAt:         GetCurrentDateTime(),
		UpdatedAt:         GetCurrentDateTime(),
	}
}

var ProjectInput = GetProjectInput()
