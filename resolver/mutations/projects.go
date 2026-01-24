package mutations

import (
	"context"
	"errors"
	"fmt"
	"go_server/helper"
	"go_server/helper/mongoDB"
	"go_server/helper/projectReusables"
	"log"
	"time"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateProject(params graphql.ResolveParams) (interface{}, error) {
	projectData, err := projectReusables.ValidateCreateProjectInput(params)
	if err != nil {
		return nil, fmt.Errorf("input validation failed: %w", err)
	}

	projectCategoryCollection, err := mongoDB.ConnectMongoDB("project_categories")
	if err != nil {
		log.Fatal(err)
	}

	var projectCategoryResult struct {
		ID int `bson:"id" json:"id"`
	}

	if err := projectCategoryCollection.FindOne(context.TODO(),
		bson.M{"id": projectData.ProjectCategoryId}).Decode(&projectCategoryResult); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("project category with ID %d does not exist",
				projectData.ProjectCategoryId)
		}
		return nil, fmt.Errorf("error checking project category: %w", err)
	}

	projectCollection, err := mongoDB.ConnectMongoDB("projects")
	if err != nil {
		log.Fatal(err)
	}

	projectId, userIdError := mongoDB.GetHighestIdInCollection(projectCollection)
	if userIdError != nil {
		return nil, userIdError
	}
	projectData.ID = projectId

	project := projectReusables.NewProject(projectData)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err = projectCollection.InsertOne(ctx, project)
	if err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	return project, nil
}

func UpdateProject(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if !ok {
		return nil, fmt.Errorf("invalid project id")
	}

	projectData, err := projectReusables.ValidateCreateProjectInput(params)
	if err != nil {
		return nil, fmt.Errorf("input validation failed: %w", err)
	}

	projectCategoryCollection, err := mongoDB.ConnectMongoDB("project_categories")
	if err != nil {
		log.Fatal(err)
	}

	var projectCategoryResult struct {
		ID int `bson:"id" json:"id"`
	}

	if err := projectCategoryCollection.FindOne(context.TODO(),
		bson.M{"id": projectData.ProjectCategoryId}).Decode(&projectCategoryResult); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("project category with ID %d does not exist",
				projectData.ProjectCategoryId)
		}
		return nil, fmt.Errorf("error checking project category: %w", err)
	}

	projectCollection, err := mongoDB.ConnectMongoDB("projects")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	now := helper.GetCurrentDateTime()
	update := bson.M{
		"$set": bson.M{
			"name":                projectData.Name,
			"project_category_id": projectData.ProjectCategoryId,
			"description":         projectData.Description,
			"tech_stacks":         projectData.TechStacks,
			"github_link":         projectData.GithubLink,
			"live_link":           projectData.LiveLink,
			"test_link":           projectData.TestLink,
			"updated_at":          now,
		},
	}

	result, err := projectCollection.UpdateOne(ctx, bson.M{"id": id}, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update project: %w", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("project with ID %d not found", id)
	}

	var updatedProject projectReusables.ProjectMongo
	if err := projectCollection.FindOne(ctx, bson.M{"id": id}).Decode(&updatedProject); err != nil {
		return nil, fmt.Errorf("error fetching updated project: %w", err)
	}

	return updatedProject, nil
}
