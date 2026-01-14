package adminUserReusables

import (
	"go_server/helper"

	"github.com/graphql-go/graphql"
)

var AdminUserInputDef = graphql.NewInputObject(graphql.InputObjectConfig{ // Create and return the InputObject directly
	Name:        "AdminUserInput",
	Description: "Input for creating a new admin user",
	Fields: graphql.InputObjectConfigFieldMap{
		"username": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Name of the user",
		},
		"email": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Email of the user",
		},
		"password": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Password of the user",
		},
	},
})

var AdminLoginInputDef = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "AdminLoginInput",
	Description: "Input for admin user login",
	Fields: graphql.InputObjectConfigFieldMap{
		"identifier": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Email or Username of the user",
		},
		"password": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Password of the user",
		},
	},
})

type AdminUserInput struct {
	ID       int     `json:"id"`
	UserName string  `json:"name"`
	Email    string  `json:"email"`
	Password *string `json:"password"`
}

type AdminUserInputMongo struct {
	ID        int     `json:"id" bson:"id"`
	UserName  string  `json:"name" bson:"username"`
	Email     string  `json:"email" bson:"email"`
	Password  *string `json:"password" bson:"password"`
	Role      string  `json:"role" bson:"role"`
	Type      string  `json:"type" bson:"type"`
	Status    string  `json:"status" bson:"status"`
	CreatedAt string  `json:"created_at" bson:"created_at"`
	UpdatedAt string  `json:"updated_at" bson:"updated_at"`
	LastLogin *string `json:"last_login" bson:"last_login"`
}

type AdminUserIsAdministrator struct {
	ID   int    `json:"id" bson:"id"`
	Role string `json:"role" bson:"role"`
}

type AdminUserInputData struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAdminUser(
	userId int,
	adminUserInputData *AdminUserInputData,
) *AdminUserInputMongo {
	return &AdminUserInputMongo{
		ID:        userId + 1,
		UserName:  adminUserInputData.UserName,
		Email:     adminUserInputData.Email,
		Password:  &adminUserInputData.Password,
		Role:      "administrator",
		Type:      "user",
		Status:    "active",
		CreatedAt: helper.GetCurrentDateTime(),
		UpdatedAt: helper.GetCurrentDateTime(),
		LastLogin: nil,
	}
}
