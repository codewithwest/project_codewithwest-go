package handler

import (
	"context"
	"go_server/api"
	"go_server/types"
	"net/http"
	"os"

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

	//types := schema.GetAllTypes()

	schemaObj, err := types.GetSchema()
	if err != nil {
		LOGGER.Err(err).Msg(err.Error())
		http.Error(w, "Error getting schema", http.StatusInternalServerError)
		return
	}

	mainController := &api.MainController{Schema: schemaObj}

	router := mux.NewRouter()
	router.Handle("/graphql", mainController.GetData())

	router.ServeHTTP(w, r) // Use the router to serve the request
}
