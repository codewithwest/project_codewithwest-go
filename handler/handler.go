package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/digvijay17july/golang-projects/go-graphql-rest-api/api"    // Adjust path
	"github.com/digvijay17july/golang-projects/go-graphql-rest-api/schema" // Adjust path
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

var LOGGER zerolog.Logger

func init() {
	LOGGER = zerolog.New(os.Stdout).With().Timestamp().
		Str("app", "codewithwest-go").
		Logger().With().Caller().Logger()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "logger", LOGGER)

	userType := schema.GetUserType()

	schemaObj, err := schema.GetSchema(userType)
	if err != nil {
		LOGGER.Err(err).Msg(err.Error())
		http.Error(w, "Error getting schema", http.StatusInternalServerError)
		return
	}

	userController := &api.UserController{Schema: schemaObj}

	router := mux.NewRouter()
	router.Handle("/graphql", userController.GetUser())

	router.ServeHTTP(w, r) // Use the router to serve the request
}
