package mutations

import (
	"context"
	"fmt"
	"go_server/helper"
	"go_server/helper/adminUserReusables"
	"go_server/helper/mongoDB"
	"log"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAdminUser(params graphql.ResolveParams) (interface{}, error) {

	adminUserInputData, validationError := adminUserReusables.ValidateAdminUserInput(
		params)
	if validationError != nil {
		return nil, validationError
	}

	collection, err := mongoDB.ConnectMongoDB("admin_users")
	if err != nil {
		log.Fatal(err)
	}

	hashedPassword, isPassword := helper.HashPassword(adminUserInputData.Password)
	if !isPassword {
		return nil, fmt.Errorf("oops! something went wrong on our side while creating your password! Please contact support")
	}

	emailExist, isEmailExists := mongoDB.EmailExist(
		collection, adminUserInputData.Email)

	if isEmailExists != nil {
		return nil, fmt.Errorf("failed to convert inserted ID to ObjectID %s", isEmailExists)
	}
	if emailExist {
		return nil, fmt.Errorf("email already exists")
	}

	userId, userIdError := mongoDB.GetHighestIdInCollection(collection)
	if userIdError != nil {
		return nil, userIdError
	}

	adminUserInputData.Password = hashedPassword
	user := adminUserReusables.NewAdminUser(userId, adminUserInputData)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	insertedID := result.InsertedID // No type assertion here yet

	objectID, ok := insertedID.(primitive.ObjectID) // Type assertion to primitive.ObjectID
	if !ok {
		return nil, fmt.Errorf("failed to convert inserted ID to ObjectID")
	}

	var createdUser adminUserReusables.AdminUserInputMongo
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&createdUser)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve created user: %w", err)
	}

	return createdUser, nil
}

func CreateAdminUserRequest(params graphql.ResolveParams) (interface{}, error) {
	isAuthorized, err := mongoDB.UserDataAccessIsAuthorized(params)
	if err != nil {
		return nil, fmt.Errorf("not authorized: %w", err)
	}

	_, err = strconv.Atoi(isAuthorized)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	email, isEmail := params.Args["email"].(string)

	if !isEmail {
		return nil, fmt.Errorf("missing required argument(s)")
	}
	_, err = helper.ValidateEmailAddress(email)
	if err != nil {
		return nil, err
	}

	collection, err := mongoDB.ConnectMongoDB("admin_user_request") // Replace placeholders
	if err != nil {
		log.Fatal(err)
	}

	userId, userIdError := mongoDB.GetHighestIdInCollection(collection)
	if userIdError != nil {
		return nil, userIdError
	}

	emailExist, isEmailExists := mongoDB.EmailExist(
		collection, email)

	if isEmailExists != nil {
		return nil, fmt.Errorf("failed to convert inserted id to objectId %s", isEmailExists)
	}
	if emailExist {
		return nil, fmt.Errorf("email already exists")
	}

	newRequestUser := adminUserReusables.AdminUserRequestMongo{
		ID:        userId + 1,
		Email:     email,
		CreatedAt: helper.GetCurrentDateTime(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, newRequestUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return newRequestUser, nil
}
