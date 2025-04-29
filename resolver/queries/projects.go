package queries

import (
	"context"
	"fmt"
	"go_server/helper"
	"go_server/helper/mongoDB"
	"log"
	"time"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProjects(params graphql.ResolveParams) (interface{}, error) {

	limit, ok := params.Args["limit"].(int)
	if !ok {
		return nil, fmt.Errorf("missing limit Argument")
	}

	collection, err := mongoDB.ConnectMongoDB("projects")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	findOptions := options.Find().SetLimit(int64(limit))
	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.Background())

	var projects []helper.ProjectMongo
	for cursor.Next(context.Background()) {
		var doc helper.ProjectMongo
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		projects = append(projects, doc)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return projects, nil
}
