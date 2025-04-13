package clientReusables

import "github.com/graphql-go/graphql"

var AuthenticateClientInputDef = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name:        "AuthenticateClientInput",
		Description: "Input for creating a new client",
		Fields: graphql.InputObjectConfigFieldMap{
			"email": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Email of the user",
			},
			"password": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Password of the user",
			},
		},
	},
)

type AuthenticateClientInputData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
