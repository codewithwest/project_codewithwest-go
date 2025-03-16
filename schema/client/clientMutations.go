package client

import (
	"go_server/helper/clientReusables"
	"go_server/resolver/mutations"

	"github.com/graphql-go/graphql"
)

func CreateClientMutation(requiredType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        requiredType,
		Description: "Create a new client",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(clientReusables.ClientInputDef),
				Description: "Input for creating a new client",
			},
		},
		Resolve: mutations.CreateClient,
	}
}
