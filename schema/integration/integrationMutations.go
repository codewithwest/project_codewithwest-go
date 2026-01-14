package integration

import (
	"context"
	"go_server/helper/integrationReusables"
	"go_server/helper/mongoDB"
	"go_server/types"
	"time"

	"github.com/graphql-go/graphql"
)

var CreateIntegrationMutation = &graphql.Field{
	Type: types.IntegrationType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name := p.Args["name"].(string)

		token, err := mongoDB.GenerateSecureToken()
		if err != nil {
			return nil, err
		}

		newIntegration := integrationReusables.NewIntegration(name, token)

		collection, err := mongoDB.ConnectMongoDB("integrations")
		if err != nil {
			return nil, err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err = collection.InsertOne(ctx, newIntegration)
		if err != nil {
			return nil, err
		}

		return newIntegration, nil
	},
}

var RevokeIntegrationMutation = &graphql.Field{
	Type: types.IntegrationType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := p.Args["id"].(string)

		collection, err := mongoDB.ConnectMongoDB("integrations")
		if err != nil {
			return nil, err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		filter := mongoDB.CreateBSONID(id)
		update := mongoDB.CreateBSONUpdate("is_revoked", true)

		var updatedIntegration integrationReusables.Integration
		err = collection.FindOneAndUpdate(ctx, filter, update).Decode(&updatedIntegration)
		if err != nil {
			return nil, err
		}

		updatedIntegration.IsRevoked = true
		return updatedIntegration, nil
	},
}
