package handler

import (
	"context"
	"go_server/api"
	"go_server/helper"
	"go_server/types"
	"net/http"
	"os"
	"strings"
	
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
	enableCors(&w)

	ctx := r.Context()
	ctx = context.WithValue(ctx, "logger", LOGGER)

	w.Header().Set("Content-Type", "application/json")

	schemaObj, err := types.GetSchema()
	if err != nil {
		LOGGER.Err(err).Msg(err.Error())
		http.Error(w, "Error getting schema", http.StatusInternalServerError)
		return
	}

	mainController := &api.MainController{Schema: schemaObj}

	router := mux.NewRouter()
	router.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		mainController.GetData().ServeHTTP(w, r)
	}).Methods("GET", "POST", "OPTIONS") // Explicitly allow OPTIONS
	router.ServeHTTP(w, r)
}

func enableCors(w *http.ResponseWriter) {
	allowedOrigins := helper.GetEnvVariable("ALLOWED_ORIGINS")
        origins := strings.Split(allowedOrigins, ",")
	header := (*w).Header()
	
	header.Add("Access-Control-Allow-Origin", origins)
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, signature, user_id") // Add "signature"
	header.Add("Access-Control-Allow-Credentials", "true")
}
