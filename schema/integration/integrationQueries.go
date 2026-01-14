package integration

import (
	"context"
	"fmt"
	"go_server/helper"
	"go_server/helper/integrationReusables"
	"go_server/helper/mongoDB"
	"go_server/types"
	"math"
	"time"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var GetIntegrationsSchema = &graphql.Field{
	Type: types.IntegrationRequestQueryType,
	Args: helper.GlobalPaginatedQueriesInput,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// Get pagination parameters
		limit := int32(10) // Default limit
		if limitArg, ok := p.Args["limit"].(int); ok && limitArg > 0 {
			limit = int32(limitArg)
		}

		page := int32(1) // Default page
		if pageArg, ok := p.Args["page"].(int); ok && pageArg > 0 {
			page = int32(pageArg)
		}

		skip := (page - 1) * limit

		collection, err := mongoDB.ConnectMongoDB("integrations")
		if err != nil {
			return nil, err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Get total count for pagination
		totalItems, err := collection.CountDocuments(ctx, bson.D{})
		if err != nil {
			return nil, fmt.Errorf("error counting documents: %w", err)
		}

		findOptions := options.Find().
			SetLimit(int64(limit)).
			SetSkip(int64(skip)).
			SetSort(bson.D{{Key: "created_at", Value: -1}})

		cursor, err := collection.Find(ctx, bson.M{}, findOptions)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)

		var integrations []integrationReusables.Integration
		if err = cursor.All(ctx, &integrations); err != nil {
			return nil, err
		}

		totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

		return map[string]interface{}{
			"data": integrations,
			"pagination": helper.Pagination{
				CurrentPage: int(page),
				PerPage:     int(limit),
				Count:       len(integrations),
				Offset:      int(skip),
				TotalPages:  totalPages,
				TotalItems:  int(totalItems),
			},
		}, nil
	},
}
