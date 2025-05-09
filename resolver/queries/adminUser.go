package queries

import (
	"context"
	"errors"
	"fmt"
	"go_server/helper"
	"go_server/helper/adminUserReusables"
	"go_server/helper/mongoDB"
	"log"
	"math"
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
	// Authorization check
	isAuthorized, err := mongoDB.UserDataAccessIsAuthorized(params)
	if err != nil {
		return nil, fmt.Errorf("not authorized: %w", err)
	}

	adminUserId, err := strconv.Atoi(isAuthorized)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	// Get pagination parameters
	limit := int32(10) // Default limit
	if limitArg, ok := params.Args["limit"].(int); ok && limitArg > 0 {
		limit = int32(limitArg)
	}

	page := int32(1) // Default page
	if pageArg, ok := params.Args["page"].(int); ok && pageArg > 0 {
		page = int32(pageArg)
	}

	skip := (page - 1) * limit

	// Connect to database
	collection, err := mongoDB.ConnectMongoDB("admin_users")
	if err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check administrator privileges
	var administrator adminUserReusables.AdminUserIsAdministrator
	if err := collection.FindOne(ctx, bson.M{"id": adminUserId}).Decode(&administrator); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("internal server error: %w", err)
	}

	if administrator.Role != "administrator" {
		return nil, fmt.Errorf("access denied: administrator privileges required")
	}

	// Get total count for pagination
	totalItems, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error counting documents: %w", err)
	}

	// Configure find options with pagination
	findOptions := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(skip)).
		SetSort(bson.D{{Key: "id", Value: 1}})

	// Execute query
	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("error fetching users: %w", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Printf("Error closing cursor: %v", err)
		}
	}(cursor, ctx)

	// Check if there are no documents
	if !cursor.Next(ctx) {
		return nil, fmt.Errorf("no users found")
	}

	// Decode results
	var adminUsers []adminUserReusables.AdminUserInputMongo
	if err := cursor.All(ctx, &adminUsers); err != nil {
		return nil, fmt.Errorf("error decoding users: %w", err)
	}

	totalPages := int32(math.Ceil(float64(totalItems) / float64(limit)))

	return &adminUserReusables.AdminUsersResponse{
		Data: adminUsers,
		Pagination: helper.Pagination{
			CurrentPage: int(page),
			PerPage:     int(limit),
			Count:       int(totalItems),
			Offset:      0,
			TotalPages:  int(totalPages),
			TotalItems:  int(totalItems),
		},
	}, nil
}

func GetAdminUserRequests(params graphql.ResolveParams) (interface{}, error) {
	// Authorization check
	isAuthorized, err := mongoDB.UserDataAccessIsAuthorized(params)
	if err != nil {
		return nil, fmt.Errorf("not authorized: %w", err)
	}

	adminUserId, err := strconv.Atoi(isAuthorized)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	// Get pagination parameters
	limit := int32(10) // Default limit
	if limitArg, ok := params.Args["limit"].(int); ok && limitArg > 0 {
		limit = int32(limitArg)
	}

	page := int32(1) // Default page
	if pageArg, ok := params.Args["page"].(int); ok && pageArg > 0 {
		page = int32(pageArg)
	}

	skip := (page - 1) * limit

	// Connect to database
	collection, err := mongoDB.ConnectMongoDB("admin_user_request")
	if err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check administrator privileges
	adminCollection, err := mongoDB.ConnectMongoDB("admin_users")
	if err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}

	var administrator adminUserReusables.AdminUserIsAdministrator
	if err := adminCollection.FindOne(ctx, bson.M{"id": adminUserId}).Decode(&administrator); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("internal server error: %w", err)
	}

	if administrator.Role != "administrator" {
		return nil, fmt.Errorf("access denied: administrator privileges required")
	}

	// Get total count for pagination
	totalItems, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error counting documents: %w", err)
	}

	// Configure find options with pagination
	findOptions := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(skip)).
		SetSort(bson.D{{Key: "created_at", Value: -1}}) // Sort by creation date, newest first

	_ = fmt.Errorf("total pages: %v", "dwdes")

	// Execute query
	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("error fetching requests: %w", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Printf("Error closing cursor: %v", err)
		}
	}(cursor, ctx)

	// Check if there are no documents
	if !cursor.Next(ctx) {
		return nil, fmt.Errorf("no requests found")
	}

	// Decode results
	var adminUserRequests []adminUserReusables.AdminUserRequestMongo
	if err := cursor.All(ctx, &adminUserRequests); err != nil {
		return nil, fmt.Errorf("error decoding requests: %w", err)
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	return &adminUserReusables.AdminUsersRequestResponse{
		Data: adminUserRequests,
		Pagination: helper.Pagination{
			CurrentPage: int(page),
			PerPage:     int(limit),
			Count:       int(totalItems),
			Offset:      0,
			TotalPages:  totalPages,
			TotalItems:  int(totalItems),
		},
	}, nil
}
