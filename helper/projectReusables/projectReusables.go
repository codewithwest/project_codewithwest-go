package projectReusables

import (
	"fmt"
	"go_server/helper"

	"github.com/graphql-go/graphql"
)

var ProjectInput = graphql.NewInputObject(graphql.InputObjectConfig{
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

func GetOptionalString(input map[string]interface{}, key string) string {
	if val, ok := input[key].(string); ok {
		return val
	}
	return ""
}

func NewProject(
	opts *ProjectMongo,
) ProjectMongo {
	now := helper.GetCurrentDateTime()

	return ProjectMongo{
		ID:                opts.ID + 1,
		ProjectCategoryId: opts.ProjectCategoryId,
		Name:              opts.Name,
		Description:       opts.Description,
		TechStacks:        opts.TechStacks,
		GithubLink:        opts.GithubLink,
		LiveLink:          opts.LiveLink,
		TestLink:          opts.TestLink,
		CreatedAt:         now,
		UpdatedAt:         now,
	}
}

type ProjectResponse struct {
	Data       []ProjectMongo    `json:"data"`
	Pagination helper.Pagination `json:"pagination"`
}

func ValidateCreateProjectInput(params graphql.ResolveParams) (*ProjectMongo, error) {
	inputArg, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input arguments")
	}

	required := struct {
		name              string
		projectCategoryId int
		description       string
		techStacks        []string
	}{}

	var err error
	if required.name, ok = inputArg["name"].(string); !ok {
		err = fmt.Errorf("name is required and must be a string")
	}
	if required.projectCategoryId, ok = inputArg["project_category_id"].(int); !ok {
		err = fmt.Errorf("project_category_id is required and must be an integer")
	}
	if required.description, ok = inputArg["description"].(string); !ok {
		err = fmt.Errorf("description is required and must be a string")
	}

	if techStacksInterface, ok := inputArg["tech_stacks"].([]interface{}); ok {
		required.techStacks = make([]string, 0, len(techStacksInterface))
		for i, v := range techStacksInterface {
			if str, ok := v.(string); ok {
				required.techStacks = append(required.techStacks, str)
			} else {
				err = fmt.Errorf("tech_stacks element at index %d is not a string", i)
				break
			}
		}
	} else {
		err = fmt.Errorf("tech_stacks is required and must be an array of strings")
	}

	if err != nil {
		return nil, err
	}

	return &ProjectMongo{
		Name:              required.name,
		ProjectCategoryId: required.projectCategoryId,
		Description:       required.description,
		TechStacks:        required.techStacks,
		GithubLink:        GetOptionalString(inputArg, "github_link"),
		LiveLink:          GetOptionalString(inputArg, "live_link"),
		TestLink:          GetOptionalString(inputArg, "test_link"),
	}, nil
}
