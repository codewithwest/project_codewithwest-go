package types

import (
	"github.com/graphql-go/graphql"
	"go_server/schema/admin"
	"go_server/schema/projects"
	"go_server/schema/user"
)

func GetSchema() (graphql.Schema, error) {

	var adminUserType = GetAdminUserType()
	var adminUserRequestType = GetAdminUserRequestType()
	var userType = GetUserType()
	var projectCategoryType = GetProjectCategoryType()
	var projectType = GetProjectType()

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getUser":        user.GetUserSchema(userType),
				"getUsers":       user.GetUsersSchema(userType),
				"loginAdminUser": admin.GetLoginAdminUserSchema(adminUserType),
				"getAdminUsers":  admin.GetAdminUsersSchema(adminUserType),
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"createAdminUser":       admin.CreateAdminUserMutation(adminUserType),
				"createProjectCategory": projects.CreateProjectCategoryMutation(projectCategoryType),
				"createProject":         projects.CreateProjectMutation(projectType),
				"requestAdminAccess":    admin.RequestAdminAccess(adminUserRequestType),
			},
		}),
	})

	return schema, err
}
