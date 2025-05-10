package helper

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	"log"
	"net/mail"
	"os"
	"time"
)

type Pagination struct {
	CurrentPage int `json:"currentPage"`
	PerPage     int `json:"perPage"`
	Count       int `json:"count"`
	Offset      int `json:"offset"`
	TotalPages  int `json:"totalPages"`
	TotalItems  int `json:"totalItems"`
}

func ValidateEmailAddress(email string) (bool, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return false, fmt.Errorf("invalid email format")
	}
	return true, nil
}

func GetEnvVariable(searchValue string) string {
	if os.Getenv("VERCEL") == "" {
		// The "VERCEL" env variable is set by Vercel
		if err := godotenv.Load(); err != nil {
			log.Println("Error loading .env file (development only):", err)
			// You can choose to continue or exit if the .env file is missing
		}
	}
	resolvedValue := os.Getenv(searchValue)

	if resolvedValue == "" {
		log.Fatal(searchValue + " environment variable not set")
	}
	return resolvedValue
}

func GetCurrentDateTime() string {
	resolvedTime := time.Now()
	local, err := time.LoadLocation("Local")

	if err != nil {
		fmt.Println("Error loading location:", err)

		return resolvedTime.Format("02-01-2006 15:04:05")
	} else {
		currentTime := resolvedTime.In(local)

		return currentTime.Format("02-01-2006 15:04:05")
	}
}

var genericPagination = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Pagination",
		Fields: graphql.Fields{
			"currentPage": &graphql.Field{
				Type:        graphql.Int,
				Description: "Current page number",
			},
			"perPage": &graphql.Field{
				Type:        graphql.Int,
				Description: "Number of items per page",
			},
			"count": &graphql.Field{
				Type:        graphql.Int,
				Description: "Number of items in the current page",
			},
			"offset": &graphql.Field{
				Type:        graphql.Int,
				Description: "Offset for the current page",
			},
			"totalPages": &graphql.Field{
				Type:        graphql.Int,
				Description: "Total number of pages available",
			},
			"totalItems": &graphql.Field{
				Type:        graphql.Int,
				Description: "Total number of items across all pages",
			},
		},
	},
)

func GlobalPaginatedQueryResolver(data graphql.Type, name string) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: name,
			Fields: graphql.Fields{
				"data": &graphql.Field{
					Type:        graphql.NewList(data),
					Description: "List of data for the current page",
				},
				"pagination": &graphql.Field{
					Type:        genericPagination,
					Description: "Pagination information",
				},
			},
		},
	)
}

var GlobalPaginatedQueriesInput = graphql.FieldConfigArgument{
	"limit": &graphql.ArgumentConfig{
		Type:         graphql.Int,
		DefaultValue: 10,
		Description:  "Number of items per page",
	},
	"page": &graphql.ArgumentConfig{
		Type:         graphql.Int,
		DefaultValue: 1,
		Description:  "Page number",
	},
}
