package types

import (
	"github.com/graphql-go/graphql"
	"go_server/schema/admin"
	"go_server/schema/projects"
	"go_server/schema/user"
)

func GetSchema() (graphql.Schema, error) {

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getUser":                    user.GetUserSchema(UserType),
				"getUsers":                   user.GetUsersSchema(UserType),
				"loginAdminUser":             admin.GetLoginAdminUserSchema(AdminUserType),
				"getAdminUsers":              admin.GetAdminUsersSchema(AdminUserType),
				"getAdminUserAccessRequests": admin.GetAdminUserRequests(AdminUserRequestType),
				"getProjects":                projects.GetProjects(ProjectType),
				"getProjectCategories":       projects.GetProjectCategories(ProjectCategoryType),
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"createAdminUser":        admin.CreateAdminUserMutation(AdminUserType),
				"createProjectCategory":  projects.CreateProjectCategoryMutation(ProjectCategoryType),
				"createProject":          projects.CreateProjectMutation(ProjectType),
				"adminUserAccessRequest": admin.RequestAdminAccess(AdminUserRequestType),
			},
		}),
	})

	return schema, err
}
