package mutations

import (
	"context"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_server/helper"
	"go_server/helper/mongoDB"
	"log"
	"time"
)

func CreateProjectCategory(params graphql.ResolveParams) (interface{}, error) {
	// check if params exists
	name, isname := params.Args["name"].(string)

	// check if params are valid
	if !isname {
		return nil, fmt.Errorf("missing required argument(s)")
	}

	collection, err := mongoDB.ConnectMongoDB(
		helper.GetEnvVariable("MONGO_DB_URL"),
		"codewithwest",
		"project_categories") // Replace placeholders
	if err != nil {
		log.Fatal(err)
	}

	projectCategoryExist, isProjectCategoryExists := mongoDB.NameExists(
		collection, name)

	if isProjectCategoryExists != nil {
		return nil, fmt.Errorf(
			"failed to convert inserted ID to ObjectID",
			isProjectCategoryExists)
	}
	if projectCategoryExist {
		return nil, fmt.Errorf("project category already exists")
	}

	projectCategoryId, userIdError := mongoDB.GetHighestIdInCollection(collection)
	if userIdError != nil {
		return nil, userIdError
	}
	projectCategory := helper.NewProjectCategory(projectCategoryId, name)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, projectCategory)
	if err != nil {
		return nil, fmt.Errorf("failed to project category: %w", err)
	}

	insertedID := result.InsertedID // No type assertion here yet
	// Convert ObjectID to string
	objectID, ok := insertedID.(primitive.ObjectID) // Type assertion to primitive.ObjectID
	if !ok {
		return nil, fmt.Errorf("failed to convert inserted ID to ObjectID")
	}

	// Adjust type assertion if needed
	var createProjectCategory helper.ProjectCategoryMongo
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&createProjectCategory)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve created project Category: %w", err)
	}

	return createProjectCategory, nil
}

func CreateProject(params graphql.ResolveParams) (interface{}, error) {
	inputArg, isInput := params.Args["input"].(map[string]interface{})
	if !isInput {
		return nil, fmt.Errorf("invalid Input arguments")
	}

	// check if params exists
	name, isname := inputArg["name"].(string)
	projectCategoryId, isProjectCategoryId := inputArg["project_category_id"].(int)
	description, isDescription := inputArg["description"].(string)
	techStacksInterface, isTechStacks := inputArg["tech_stacks"].([]interface{})
	// check if params are valid
	if !isname || !isProjectCategoryId || !isDescription || !isTechStacks {
		return nil, fmt.Errorf("missing required argument(s)")
	}

	techStacksList := make([]string, len(techStacksInterface))
	for position, value := range techStacksInterface {
		str, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("tech_stacks element at index %d is not a string, but a %T", position, value) // Include index and type!
		}
		techStacksList[position] = str
	}

	projectCategoryCollection, err := mongoDB.ConnectMongoDB(
		helper.GetEnvVariable("MONGO_DB_URL"),
		"codewithwest",
		"project_categories") // Replace placeholders
	if err != nil {
		log.Fatal(err)
	}

	var projectCategoryResult struct {
		ID interface{} `bson:"id"` // _id can be of any type
	}

	err = projectCategoryCollection.FindOne(context.TODO(), bson.M{"id": projectCategoryId}, options.FindOne()).Decode(&projectCategoryResult)
	if err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return nil, fmt.Errorf("project category does not exist")
		}
	}

	projectCollection, err := mongoDB.ConnectMongoDB(
		helper.GetEnvVariable("MONGO_DB_URL"),
		"codewithwest",
		"projects") // Replace placeholders
	if err != nil {
		log.Fatal(err)
	}
	projectId, userIdError := mongoDB.GetHighestIdInCollection(projectCollection)
	if userIdError != nil {
		return nil, userIdError
	}
	// create project using data sent from client
	project := helper.ProjectMongo{
		ID:                projectId + 1,
		ProjectCategoryId: projectCategoryId,
		Name:              name,
		Description:       description,
		TechStacks:        techStacksList,
		GithubLink:        inputArg["github_link"].(string),
		LiveLink:          inputArg["live_link"].(string),
		TestLink:          inputArg["test_link"].(string),
		CreatedAt:         helper.GetCurrentDateTime(),
		UpdatedAt:         helper.GetCurrentDateTime(),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err = projectCollection.InsertOne(ctx, project)
	if err != nil {
		return nil, fmt.Errorf("failed to project category: %w", err)
	}

	return project, nil
}
