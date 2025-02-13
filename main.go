package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/digvijay17july/golang-projects/go-graphql-rest-api/api"
	"github.com/digvijay17july/golang-projects/go-graphql-rest-api/schema"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

var LOGGER zerolog.Logger

func init() { // Initialize logger at package level
	LOGGER = zerolog.New(os.Stdout).With().Timestamp().
		Str("app", "codewithwest-go").
		Logger().With().Caller().Logger()
}

func Handler(w http.ResponseWriter, r *http.Request) { // Correct signature
	ctx := r.Context()
	ctx = context.WithValue(ctx, "logger", LOGGER) // Add logger to context

	userType := schema.GetUserType()

	schemaObj, err := schema.GetSchema(userType)
	if err != nil {
		LOGGER.Err(err).Msg(err.Error())
		http.Error(w, "Error getting schema", http.StatusInternalServerError) // Handle error
		return                                                                // Important: Return after an error
	}

	userController := &api.UserController{Schema: schemaObj}

	router := mux.NewRouter()
	router.Handle("/graphql", userController.GetUser())

	// Serve the request using the router
	router.ServeHTTP(w, r)
}

func main() { // This is for local testing only
	LOGGER = zerolog.New(os.Stdout).With().Timestamp().
		Str("app", "codewithwest-go").
		Logger().With().Caller().Logger()
	http.HandleFunc("/graphql", Handler)
	log.Fatal(http.ListenAndServe(":3072", nil))
}
