package adminUserReusables

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"go_server/helper"
)

func GetAdminUserInput() *graphql.InputObject { // Return *graphql.InputObject
	return graphql.NewInputObject(graphql.InputObjectConfig{ // Create and return the InputObject directly
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
}

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

var AdminUserInputDef = GetAdminUserInput()
