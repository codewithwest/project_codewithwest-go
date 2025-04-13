package clientReusables

import (
	"github.com/graphql-go/graphql"
	"go_server/helper"
	"go_server/helper/mongoDB"
)

var ClientInputDef = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name:        "ClientInput",
		Description: "Input for creating a new client",
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
			"company_name": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Name of the company",
			},
		},
	},
)

type ClientInputMongo struct {
	ID          int     `json:"id" bson:"id"`
	UserName    string  `json:"username" bson:"username"`
	Email       string  `json:"email" bson:"email"`
	Password    *string `json:"password" bson:"password"`
	CompanyName string  `json:"company_name" bson:"company_name"`
	Type        string  `json:"type" bson:"type"`
	Status      string  `json:"status" bson:"status"`
	CreatedAt   string  `json:"created_at" bson:"created_at"`
	UpdatedAt   string  `json:"updated_at" bson:"updated_at"`
	LastLogin   *string `json:"last_login" bson:"last_login"`
}

type ClientInputData struct {
	UserName    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	CompanyName string `json:"company_name"`
}

type ClientReturnSchema struct {
	ID          int     `json:"id"`
	UserName    string  `json:"username"`
	Email       string  `json:"email"`
	CompanyName string  `json:"company_name"`
	Type        string  `json:"type"`
	Status      string  `json:"status"`
	Token       string  `json:"token"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	LastLogin   *string `json:"last_login"`
}

func CreateNewClientInput(highestClientId int, data *ClientInputData) *ClientInputMongo {
	return &ClientInputMongo{
		ID:          highestClientId + 1,
		UserName:    data.UserName,
		Email:       data.Email,
		Password:    &data.Password,
		CompanyName: data.CompanyName,
		Type:        "client",
		Status:      "active",
		CreatedAt:   helper.GetCurrentDateTime(),
		UpdatedAt:   helper.GetCurrentDateTime(),
		LastLogin:   nil,
	}
}

func ClientInputReturnData(session *mongoDB.Session, createdUserData ClientInputMongo) *ClientReturnSchema {
	return &ClientReturnSchema{
		ID:          createdUserData.ID,
		UserName:    createdUserData.UserName,
		Email:       createdUserData.Email,
		CompanyName: createdUserData.CompanyName,
		Type:        createdUserData.Type,
		Status:      createdUserData.Status,
		Token:       session.Token,
		CreatedAt:   createdUserData.CreatedAt,
		UpdatedAt:   createdUserData.UpdatedAt,
		LastLogin:   createdUserData.LastLogin,
	}
}
