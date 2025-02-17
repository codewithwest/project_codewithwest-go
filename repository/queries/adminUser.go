package queries

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go_server/helper"
	"go_server/helper/adminUserReusables"
	"go_server/helper/mongoDB"
	"log"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
)

func LoginAdminUser(params graphql.ResolveParams) (interface{}, error) {
	// In a real application, you would typically fetch the user data from a database or other data source
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

	err := mongoDB.ConnectMongoDB(
		helper.GetEnvVariable("MONGO_DB_URL"),
		"codewithwest",
		"admin_users")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var adminUser adminUserReusables.AdminUserInputMongo

	findOneError := mongoDB.RetrievedCollection.FindOne(
		ctx, bson.M{"email": email}).Decode(&adminUser)
	fmt.Println("password compared with adminUser" + password + "second: " + *adminUser.Password)
	passwordInvalid := helper.CheckPasswordHash(password, *adminUser.Password)
	if !passwordInvalid {
		return nil, fmt.Errorf("invalid email or password combination")
	}
	if findOneError != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to retrieve user: %w", findOneError)
	}

	return adminUser, nil
}

func GetAdminUsers(params graphql.ResolveParams) (interface{}, error) {
	limit, ok := params.Args["limit"].(int)
	if !ok {
		return nil, fmt.Errorf("missing limit Argument")
	}

	var users []adminUserReusables.AdminUser

	for userId := 0; userId < limit; userId++ {

		users = append(users, NewRandomAdminUser(strconv.Itoa(userId)))
	}
	return users, nil

}

func NewRandomAdminUser(id string) adminUserReusables.AdminUser {
	return adminUserReusables.AdminUser{
		ID:        id,
		Username:  "user" + id,
		Email:     "user" + id + "@example.com",
		Role:      "user",
		Type:      "user",
		Status:    "active",
		CreatedAt: helper.GetCurrentDateTime(),
		UpdatedAt: helper.GetCurrentDateTime(),
		LastLogin: helper.GetCurrentDateTime(),
	}
}
