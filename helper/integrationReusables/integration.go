package integrationReusables

import (
	"go_server/helper"
	"go_server/helper/mongoDB"

	"github.com/graphql-go/graphql"
)

type Integration struct {
	ID         string `json:"id" bson:"_id"`
	Name       string `json:"name" bson:"name"`
	Token      string `json:"token" bson:"token"`
	IsRevoked  bool   `json:"is_revoked" bson:"is_revoked"`
	CreatedAt  string `json:"created_at" bson:"created_at"`
	UpdatedAt  string `json:"updated_at" bson:"updated_at"`
	LastUsedAt string `json:"last_used_at" bson:"last_used_at"`
}

var IntegrationInputDef = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "IntegrationInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

func NewIntegration(name string, token string) *Integration {
	return &Integration{
		ID:         mongoDB.GenerateObjectID(),
		Name:       name,
		Token:      token,
		IsRevoked:  false,
		CreatedAt:  helper.GetCurrentDateTime(),
		UpdatedAt:  helper.GetCurrentDateTime(),
		LastUsedAt: "",
	}
}
