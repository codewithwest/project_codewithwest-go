package adminUserReusables

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"go_server/helper"
)

func ValidateAdminUserInput(params graphql.ResolveParams) (*AdminUserInputData, error) {
	inputArg, isInput := params.Args["input"].(map[string]interface{})
	if !isInput {
		return nil, fmt.Errorf("invalid input arguments: expected map[string]interface{}")
	}

	// Create a helper function for type assertion to reduce repetition
	getValueUsingKey := func(m map[string]interface{}, key string) (string, bool) {
		val, ok := m[key].(string)
		return val, ok && val != ""
	}

	email, isEmail := getValueUsingKey(inputArg, "email")
	password, isPassword := getValueUsingKey(inputArg, "password")

	var validateErrors []string
	if !isEmail {
		validateErrors = append(validateErrors, "email")
	}

	if !isPassword {
		validateErrors = append(validateErrors, "password")
	}

	if len(validateErrors) > 0 {
		return nil, fmt.Errorf("invalid input arguments: missing %s", validateErrors)
	}

	_, err := helper.ValidateEmailAddress(email)
	if err != nil {
		return nil, err
	}

	username, _ := getValueUsingKey(inputArg, "username")
	if username == "" {
		username = "newRandomUsername"
	}

	return &AdminUserInputData{
		UserName: username,
		Email:    email,
		Password: password,
	}, nil
}
