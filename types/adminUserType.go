package types

import "github.com/graphql-go/graphql"

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

var AdminUserPaginatedType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AdminUsers",
		Fields: graphql.Fields{
			"data": &graphql.Field{
				Type:        graphql.NewList(AdminUserType),
				Description: "List of admin users for the current page",
			},
			"page": &graphql.Field{
				Type:        graphql.Int,
				Description: "Current page number",
			},
			"totalPages": &graphql.Field{
				Type:        graphql.Int,
				Description: "Total number of pages available",
			},
			"totalItems": &graphql.Field{
				Type:        graphql.Int,
				Description: "Total number of items across all pages",
			},
		},
	},
)
