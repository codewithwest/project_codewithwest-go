package mutations

import (
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go_server/helper/mongoDB"
	"go_server/helper/projectCategoryReusables"
	"time"
)

func CreateProjectCategory(params graphql.ResolveParams) (interface{}, error) {
	name, isName := params.Args["name"].(string)
	if !isName || name == "" {
		return nil, fmt.Errorf("missing or invalid name argument")
	}

	collection, err := mongoDB.ConnectMongoDB("project_categories")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	exists, err := mongoDB.NameExists(collection, name)
	if err != nil {
		return nil, fmt.Errorf("error checking category existence: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("project category '%s' already exists", name)
	}

	categoryID, err := mongoDB.GetHighestIdInCollection(collection)
	if err != nil {
		return nil, fmt.Errorf("error getting highest ID: %w", err)
	}

	newCategory := projectCategoryReusables.NewProjectCategory(categoryID, name)

	result, err := collection.InsertOne(ctx, newCategory)
	if err != nil {
		return nil, fmt.Errorf("failed to create project category: %w", err)
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("invalid inserted ID type")
	}

	var createdCategory projectCategoryReusables.ProjectCategoryMongo
	if err := collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&createdCategory); err != nil {
		return nil, fmt.Errorf("failed to retrieve created category: %w", err)
	}

	return createdCategory, nil
}
