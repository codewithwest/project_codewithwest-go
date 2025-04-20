package queries

import (
	"context"
	"errors"
	"fmt"
	"go_server/helper"
	"go_server/helper/adminUserReusables"
	"go_server/helper/mongoDB"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/graphql-go/graphql"
)

func LoginAdminUser(params graphql.ResolveParams) (interface{}, error) {
	adminUserLoginData, validationError := adminUserReusables.ValidateAdminUserInput(params)
	if validationError != nil {
		return nil, validationError
	}

	collection, dbConnError := mongoDB.ConnectMongoDB("admin_users")
	if dbConnError != nil {
		return nil, fmt.Errorf("database connection error: %w", dbConnError)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var adminUser adminUserReusables.AdminUserInputMongo
	findUserError := collection.FindOne(ctx, bson.M{"email": adminUserLoginData.Email}).Decode(&adminUser)
	if findUserError != nil {
		if errors.Is(findUserError, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("invalid email or password combination")
		}
		return nil, fmt.Errorf("internal server error: %w", findUserError)
	}

	if adminUser.Password == nil {
		return nil, fmt.Errorf("invalid email or password combination")
	}

	if !helper.ValidatePassword(adminUserLoginData.Password, *adminUser.Password) {
		return nil, fmt.Errorf("invalid email or password combination")
	}

	session, sessionError := mongoDB.CreateSession(strconv.Itoa(adminUser.ID), adminUserLoginData.Email, false)
	if sessionError != nil {
		return nil, fmt.Errorf("session creation failed: %w", sessionError)
	}

	return map[string]interface{}{
		"token": session.Token,
		"id":    strconv.Itoa(adminUser.ID),
		"email": adminUserLoginData.Email,
	}, nil
}

func GetAdminUsers(params graphql.ResolveParams) (interface{}, error) {
	isAuthorized, err := mongoDB.UserDataAccessIsAuthorized(params)
	if err != nil {
		return nil, fmt.Errorf("not authorized")
	}

	adminUserId, strToIntErr := strconv.Atoi(isAuthorized)
	if strToIntErr != nil {
		return nil, strToIntErr
	}

	filter := bson.M{
		"id": adminUserId,
	}
	limit, ok := params.Args["limit"].(int)
	if !ok {
		return nil, fmt.Errorf("missing limit Argument")
	}

	collection, err := mongoDB.ConnectMongoDB("admin_users")

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var administrator adminUserReusables.AdminUserIsAdministrator
	isAdministratorError := collection.FindOne(context.Background(), filter).Decode(&administrator)
	if isAdministratorError != nil {
		if errors.Is(isAdministratorError, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf(isAdministratorError.Error())
		}

		return nil, fmt.Errorf("oops looks like an error occurred on our side, if the error continues contact support or create new account if you don't already have one please reset your password")
	}

	if administrator.Role != "administrator" {
		return nil, fmt.Errorf("oops! you do not have access to this resource")
	}

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

	var adminUsers []adminUserReusables.AdminUserInputMongo

	for cursor.Next(context.Background()) {
		var doc adminUserReusables.AdminUserInputMongo
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		adminUsers = append(adminUsers, doc)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return adminUsers, nil
}

func GetAdminUserRequests(params graphql.ResolveParams) (interface{}, error) {
	_, err := mongoDB.UserDataAccessIsAuthorized(params)
	if err != nil {
		return nil, fmt.Errorf("not authorized")
	}

	limit, ok := params.Args["limit"].(int)
	if !ok {
		return nil, fmt.Errorf("missing limit Argument")
	}

	collection, err := mongoDB.ConnectMongoDB(
		"admin_user_request")

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

	var adminUserRequests []adminUserReusables.AdminUserRequest
	for cursor.Next(context.Background()) {
		var doc adminUserReusables.AdminUserRequest
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		adminUserRequests = append(adminUserRequests, doc)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return adminUserRequests, nil
}
