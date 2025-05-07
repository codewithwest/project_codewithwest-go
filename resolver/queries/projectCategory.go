package queries

import (
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_server/helper"
	"go_server/helper/mongoDB"
	"go_server/helper/projectCategoryReusables"
	"math"
	"sync"
	"time"
)

func GetProjectCategories(params graphql.ResolveParams) (interface{}, error) {
	limit, page := 10, 1
	if l, ok := params.Args["limit"].(int); ok {
		limit = l
	} else {
		return nil, fmt.Errorf("invalid or missing limit argument")
	}

	if p, ok := params.Args["page"].(int); ok && p > 0 {
		page = p
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection, err := mongoDB.ConnectMongoDB("project_categories")
	if err != nil {
		return nil, fmt.Errorf("database connection error: %v", err)
	}

	var (
		totalCount int64
		cursor     *mongo.Cursor
		errCount   error
		errFind    error
		wg         sync.WaitGroup
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		totalCount, errCount = collection.CountDocuments(ctx, bson.D{})
	}()

	go func() {
		defer wg.Done()
		skip := int64((page - 1) * limit)
		findOptions := options.Find().
			SetSkip(skip).
			SetLimit(int64(limit))
		cursor, errFind = collection.Find(ctx, bson.D{}, findOptions)
	}()

	wg.Wait()

	if errCount != nil {
		return nil, fmt.Errorf("error counting documents: %v", errCount)
	}

	if errFind != nil {
		return nil, fmt.Errorf("error finding documents: %v", errFind)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
		}
	}(cursor, ctx)

	projects := make([]projectCategoryReusables.ProjectCategoryMongo, 0, limit)

	if err := cursor.All(ctx, &projects); err != nil {
		return nil, fmt.Errorf("error decoding documents: %v", err)
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return &projectCategoryReusables.ProjectCategoryResponse{
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
