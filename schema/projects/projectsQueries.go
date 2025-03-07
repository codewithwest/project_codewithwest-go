package projects

import (
	"go_server/resolver/queries"

	"github.com/graphql-go/graphql"
)

func GetProjects(projectsType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(projectsType),
		Description: "Get all projects",
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: queries.GetProjects,
	}
}

func GetProjectCategories(projectsType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(projectsType),
		Description: "Get all project categories",
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: queries.GetProjectCategories,
	}
}
