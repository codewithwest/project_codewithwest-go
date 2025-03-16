package queries

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"go_server/helper"
	"go_server/helper/clientReusables"
	"go_server/helper/mongoDB"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthenticateClient(params graphql.ResolveParams) (interface{}, error) {
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

	collection, err := mongoDB.ConnectMongoDB("clients")

	if err != nil {
		return nil, fmt.Errorf(" ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var client clientReusables.ClientInputMongo

	findOneError := collection.FindOne(
		ctx, bson.M{"email": email}).Decode(&client)

	if client.Password == nil {
		return nil, fmt.Errorf("invalid email or password combination")
	}

	passwordInvalid := helper.ValidatePassword(password, *client.Password)

	if !passwordInvalid {
		return nil, fmt.Errorf("invalid email or password combination Error")
	}

	if findOneError != nil {
		if errors.Is(findOneError, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("incorrect Email and Password combination")
		}

		return nil, fmt.Errorf("oops looks like an error occurred on our side, if the error continues contact support or create new account if you don't already have one please reset your password")
	}

	session, err := mongoDB.CreateSession(strconv.Itoa(client.ID), email, true)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token":        session.Token,
		"id":           client.ID,
		"email":        client.Email,
		"username":     client.UserName,
		"company_name": client.CompanyName,
	}, nil

}
