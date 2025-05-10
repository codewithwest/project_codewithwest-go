package projects

import (
	"go_server/helper"
	"go_server/resolver/queries"

	"github.com/graphql-go/graphql"
)

func GetProjects(projectsType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        projectsType,
		Description: "Get all projects",
		Args:        helper.GlobalPaginatedQueriesInput,
		Resolve:     queries.GetProjects,
	}
}

func GetProjectCategories(projectsType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        projectsType,
		Description: "Get all project categories",
		Args:        helper.GlobalPaginatedQueriesInput,
		Resolve:     queries.GetProjectCategories,
	}
}
