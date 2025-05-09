package types

import (
	"github.com/graphql-go/graphql"
	"go_server/helper"
)

var AdminUserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AdminUser",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"role": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
			"updated_at": &graphql.Field{
				Type: graphql.String,
			},
			"last_login": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var LoginAdminUserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "LoginAdminUser",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var AdminUserRequestType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AdminUserRequest",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var AdminUsersQueryType = helper.GlobalPaginatedQueryResolver(AdminUserType, "AdminUsersType")
var AdminUserRequestQueryType = helper.GlobalPaginatedQueryResolver(AdminUserRequestType, "AdminUserRequestType")
