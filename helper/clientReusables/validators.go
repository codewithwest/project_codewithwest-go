package clientReusables

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"go_server/helper"
	"strings"
)

func ValidateClientInput(params graphql.ResolveParams) (*ClientInputData, error) {
	// Use a type assertion with direct error checking
	inputArg, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input arguments: expected map[string]interface{}")
	}

	// Create a helper function for type assertion to reduce repetition
	getString := func(m map[string]interface{}, key string) (string, bool) {
		val, ok := m[key].(string)
		return val, ok && val != ""
	}

	// Perform all validations at once
	email, isEmail := getString(inputArg, "email")
	password, isPassword := getString(inputArg, "password")
	companyName, isCompanyName := getString(inputArg, "company_name")

	// Collect validation errors
	var validationErrors []string
	if !isEmail {
		validationErrors = append(validationErrors, "email")
	}

	if !isPassword {
		validationErrors = append(validationErrors, "password")
	}

	if !isCompanyName {
		validationErrors = append(validationErrors, "company_name")
	}

	if len(validationErrors) > 0 {
		return nil, fmt.Errorf("invalid input: missing or empty %s", strings.Join(validationErrors, ", "))
	}

	_, err := helper.ValidateEmailAddress(email)
	if err != nil {
		return nil, err
	}

	// Get username with default value
	username, _ := getString(inputArg, "username")
	if username == "" {
		username = "newRandomUsername"
	}

	return &ClientInputData{
		UserName:    username,
		Email:       email,
		Password:    password,
		CompanyName: companyName,
	}, nil
}
