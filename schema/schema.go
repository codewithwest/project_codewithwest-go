package schema

import (
	"go_server/schema/admin"
	"go_server/schema/client"
	"go_server/schema/projects"
	"go_server/schema/user"
	"go_server/types"

	"github.com/graphql-go/graphql"
)

func GetSchema() (graphql.Schema, error) {

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getUser":                    user.GetUserSchema(types.UserType),
				"getUsers":                   user.GetUsersSchema(types.UserType),
				"loginAdminUser":             admin.GetLoginAdminUserSchema(types.LoginAdminUserType),
				"getAdminUsers":              admin.GetAdminUsersSchema(types.AdminUserType),
				"getAdminUserAccessRequests": admin.GetAdminUserRequests(types.AdminUserRequestType),
				"getProjects":                projects.GetProjects(types.ProjectType),
				"getProjectCategories":       projects.GetProjectCategories(types.ProjectCategoryType),
				"authenticateClient":         client.AuthenticateClient(types.ClientType),
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"createAdminUser":        admin.CreateAdminUserMutation(types.AdminUserType),
				"createProjectCategory":  projects.CreateProjectCategoryMutation(types.ProjectCategoryType),
				"createProject":          projects.CreateProjectMutation(types.ProjectType),
				"adminUserAccessRequest": admin.RequestAdminAccess(types.AdminUserRequestType),
				"createClient":           client.CreateClientMutation(types.ClientType),
			},
		}),
	})

	return schema, err
}
