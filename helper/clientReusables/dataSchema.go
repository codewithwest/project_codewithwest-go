package clientReusables

import "github.com/graphql-go/graphql"

var ClientInputDef = graphql.NewInputObject(graphql.InputObjectConfig{ // Create and return the InputObject directly
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
	}})

type ClientInputMongo struct {
	ID          int     `json:"id" bson:"id"`
	UserName    string  `json:"name" bson:"username"`
	Email       string  `json:"email" bson:"email"`
	Password    *string `json:"password" bson:"password"`
	CompanyName string  `json:"company_name" bson:"company_name"`
	Type        string  `json:"type" bson:"type"`
	Status      string  `json:"status" bson:"status"`
	CreatedAt   string  `json:"created_at" bson:"created_at"`
	UpdatedAt   string  `json:"updated_at" bson:"updated_at"`
	LastLogin   *string `json:"last_login" bson:"last_login"`
}
