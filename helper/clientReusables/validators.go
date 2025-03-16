package clientReusables

import (
	"fmt"
	"go_server/helper"

	"github.com/graphql-go/graphql"
)

func ValidateClientInput(params graphql.ResolveParams) ([]string, error) {
	inputArg, isInput := params.Args["input"].(map[string]interface{})
	if !isInput {
		return nil, fmt.Errorf("invalid Input arguments")
	}

	email, isEmail := inputArg["email"].(string)
	password, isPassword := inputArg["password"].(string)
	companyName, isCompanyName := inputArg["company_name"].(string)

	if !isEmail || !isPassword || !isCompanyName {
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
	return []string{email, password, username, companyName}, nil
}

func NewClient(userId int, username string, email string, companyName string, password string) *ClientInputMongo {
	return &ClientInputMongo{
		ID:          userId + 1,
		UserName:    username,
		Email:       email,
		Password:    &password,
		CompanyName: companyName,
		Type:        "client",
		Status:      "active",
		CreatedAt:   helper.GetCurrentDateTime(),
		UpdatedAt:   helper.GetCurrentDateTime(),
		LastLogin:   nil,
	}
}
