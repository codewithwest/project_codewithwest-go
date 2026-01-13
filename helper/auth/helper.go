package auth

import (
	"encoding/json"
	"go_server/helper"
	"go_server/helper/mongoDB"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/rs/zerolog"
)

var LOGGER zerolog.Logger

func IsQueryOrMutation(r *http.Request) bool {
	var Body = r.Body
	if Body == nil {
		return false
	}

	var requestBody struct {
		Query string `json:"query"`
	}

	body, err := ioutil.ReadAll(Body)
	if err != nil {
		return false
	}

	if err := json.Unmarshal(body, &requestBody); err != nil {
		return false
	}

	// Reset the body for later use
	r.Body = ioutil.NopCloser(strings.NewReader(string(body)))

	query := strings.ToLower(requestBody.Query)

	// Allow introspection queries
	if strings.Contains(query, "__schema") || strings.Contains(query, "__type") {
		return false
	}

	// Check if the query contains any protected mutations
	for _, mutation := range ProtectedMutationsAndQueries {
		if strings.Contains(query, strings.ToLower(mutation)) {
			return true
		}
	}

	return false
}

func ValidateSession(next http.HandlerFunc, key string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !IsQueryOrMutation(r) {
			next.ServeHTTP(w, r)
			return
		}

		// if token is valid the check with mongo sessions if it is still valid
		_, err := mongoDB.GetSessionFromRequest(r, key)
		if err != nil {
			LOGGER.Error().Msg(err.Error())
			http.Error(w, "Sorry! you do not have a valid session", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func EnableCors(w *http.ResponseWriter) {
	allowedOrigins := helper.GetEnvVariable("ALLOWED_ORIGINS")
	header := (*w).Header()

	header.Add("Access-Control-Allow-Origin", allowedOrigins)
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Signature, user_id") // Add "signature"
	header.Add("Access-Control-Allow-Credentials", "true")
}
