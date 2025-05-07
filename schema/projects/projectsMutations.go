package projects

import (
	"github.com/graphql-go/graphql"
	"go_server/helper/projectReusables"
	"go_server/resolver/mutations"
)

func CreateProjectCategoryMutation(requiredType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        requiredType,
		Description: "Create a new project category",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: mutations.CreateProjectCategory,
	}
}

func CreateProjectMutation(requiredType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        requiredType,
		Description: "Create a new project",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(projectReusables.ProjectInput),
				Description: "Project input",
			},
		},

		Resolve: mutations.CreateProject,
	}
}
