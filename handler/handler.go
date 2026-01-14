package handler

import (
	"context"
	"go_server/api"
	"go_server/helper/auth"
	"go_server/schema"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func init() {
	auth.LOGGER = zerolog.New(os.Stdout).With().Timestamp().
		Str("app", "codewithwest-go").
		Logger().With().Caller().Logger()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	auth.EnableCors(&w)

	ctx := r.Context()
	_ = context.WithValue(ctx, "logger", auth.LOGGER)

	w.Header().Set("Content-Type", "application/json")

	schemaObj, err := schema.GetSchema()
	if err != nil {
		auth.LOGGER.Err(err).Msg(err.Error())
		http.Error(w, "Error getting schema", http.StatusInternalServerError)
		return
	}

	mainController := &api.MainController{Schema: schemaObj}

	router := mux.NewRouter()
	router.HandleFunc("/graphql", auth.ValidateIntegrationToken(auth.ValidateSession(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "OPTIONS" {
				auth.LOGGER.Info().Msg("OPTIONS request received")
				w.WriteHeader(http.StatusOK)
				return
			}
			ctx := context.WithValue(r.Context(), "http.Request", r)
			r = r.WithContext(ctx)

			mainController.GetData().ServeHTTP(w, r)
		}, "Authorization"))).Methods("GET", "POST", "OPTIONS")

	router.ServeHTTP(w, r)
}
