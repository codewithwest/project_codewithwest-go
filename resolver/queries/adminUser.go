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
	inputArg, isInput := params.Args["input"].(map[string]interface{})
	if !isInput {
		return nil, fmt.Errorf("invalid Input arguments")
	}

	email, isEmail := inputArg["email"].(string)
	username, isUsername := inputArg["username"].(string)
	password, isPassword := inputArg["password"].(string)

	if !isEmail || !isUsername || !isPassword {
		return nil, fmt.Errorf("invalid email or password" + username + password)
	}

	collection, err := mongoDB.ConnectMongoDB("admin_users")

	if err != nil {
		return nil, fmt.Errorf(" ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var adminUser adminUserReusables.AdminUserInputMongo

	findOneError := collection.FindOne(
		ctx, bson.M{"email": email}).Decode(&adminUser)

	if adminUser.Password == nil {
		return nil, fmt.Errorf("invalid email or password combination")
	}

	passwordInvalid := helper.ValidatePassword(password, *adminUser.Password)
	if !passwordInvalid {
		return nil, fmt.Errorf("invalid email or password combination")
	}

	if findOneError != nil {
		if errors.Is(findOneError, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("incorrect Email and Password combination")
		}

		return nil, fmt.Errorf("oops looks like an error occurred on our side, if the error continues contact support or create new account if you don't already have one please reset your password")
	}

	session, err := mongoDB.CreateSession(strconv.Itoa(adminUser.ID), email, false)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token": session.Token,
	}, nil

}

func GetAdminUsers(params graphql.ResolveParams) (interface{}, error) {
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
