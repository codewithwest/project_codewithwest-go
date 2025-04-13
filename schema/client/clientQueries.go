package client

import (
	"go_server/helper/clientReusables"
	"go_server/resolver/queries"

	"github.com/graphql-go/graphql"
)

func AuthenticateClient(projectsType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        projectsType,
		Description: "Get all project categories",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{ // Wrap it in a FieldConfigArgument
				Type:        graphql.NewNonNull(clientReusables.AuthenticateClientInputDef), // Use the InputObject here
				Description: "Input for creating a new client",                              // Description for the "input" argument
			},
		},
		Resolve: queries.AuthenticateClient,
	}
}
