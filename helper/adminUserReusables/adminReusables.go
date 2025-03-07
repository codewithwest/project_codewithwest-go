package adminUserReusables

import (
	"fmt"
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

var AdminUserRequestInputDef = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "AdminUserRequestInput",
	Description: "Input for creating a new admin user request",
	Fields: graphql.InputObjectConfigFieldMap{
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

func ValidateAdminUserInput(params graphql.ResolveParams) ([]string, error) {
	inputArg, isInput := params.Args["input"].(map[string]interface{})
	if !isInput {
		return nil, fmt.Errorf("invalid Input arguments")
	}

	email, isEmail := inputArg["email"].(string)
	password, isPassword := inputArg["password"].(string)

	if !isEmail || !isPassword {
		return nil, fmt.Errorf("invalid email or password input")
	}

	_, err := helper.ValidateEmailAddress(email)
	if err != nil {
		return nil, err
	}

	username, isUsername := inputArg["username"].(string)
	if !isUsername {
		username = "newRandomUsername"
	}
	return []string{email, password, username}, nil
}

func NewAdminUser(userId int, username string, email string, password string) *AdminUserInputMongo {
	return &AdminUserInputMongo{
		ID:        userId + 1,
		UserName:  username,
		Email:     email,
		Password:  &password,
		Role:      "administrator",
		Type:      "user",
		Status:    "active",
		CreatedAt: helper.GetCurrentDateTime(),
		UpdatedAt: helper.GetCurrentDateTime(),
		LastLogin: nil,
	}
}
