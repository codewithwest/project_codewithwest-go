package queries

import (
	"context"
	"fmt"
	"go_server/helper"
	"go_server/helper/mongoDB"
	"go_server/helper/projectReusables"
	"log"
	"math"
	"sync"
	"time"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProjects(params graphql.ResolveParams) (interface{}, error) {
	limit, page := 10, 1
	if resolvedLimit, ok := params.Args["limit"].(int); ok {
		limit = resolvedLimit
	} else {
		return nil, fmt.Errorf("invalid or missing limit argument")
	}

	if resolvedPage, ok := params.Args["page"].(int); ok && resolvedPage > 0 {
		page = resolvedPage
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection, err := mongoDB.ConnectMongoDB("projects")
	if err != nil {
		return nil, fmt.Errorf("database connection error: %v", err)
	}

	var (
		totalCount int64
		cursor     *mongo.Cursor
		errCount   error
		errFind    error
		waitGroup  sync.WaitGroup
	)

	waitGroup.Add(2)
	go func() {
		defer waitGroup.Done()
		totalCount, errCount = collection.CountDocuments(ctx, bson.D{})
	}()

	go func() {
		defer waitGroup.Done()
		skip := int64((page - 1) * limit)
		findOptions := options.Find().
			SetSkip(skip).
			SetLimit(int64(limit))
		cursor, errFind = collection.Find(ctx, bson.D{}, findOptions)
	}()
	waitGroup.Wait()

	if errCount != nil {
		return nil, fmt.Errorf("error counting documents: %v", errCount)
	}

	if errFind != nil {
		return nil, fmt.Errorf("error finding documents: %v", errFind)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		if err := cursor.Close(ctx); err != nil {
			log.Printf("error closing cursor: %v", err)
		}
	}(cursor, ctx)

	projects := make([]projectReusables.ProjectMongo, 0, limit)

	if err := cursor.All(ctx, &projects); err != nil {
		return nil, fmt.Errorf("error decoding documents: %v", err)
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return &projectReusables.ProjectResponse{
		Data: projects,
		Pagination: helper.Pagination{
			CurrentPage: page,
			PerPage:     limit,
			Count:       int(totalCount),
			Offset:      (page - 1) * limit,
			TotalPages:  totalPages,
			TotalItems:  int(totalCount),
		},
	}, nil
}
