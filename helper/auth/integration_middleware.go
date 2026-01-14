package auth

import (
	"context"
	"go_server/helper/mongoDB"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func ValidateIntegrationToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			next.ServeHTTP(w, r)
			return
		}

		// Mandatory API Key check
		token := r.Header.Get("X-API-Key")
		if token == "" {
			http.Error(w, "Unauthorized: No API key provided", http.StatusUnauthorized)
			return
		}

		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		// Check if token exists in integrations collection
		collection, err := mongoDB.ConnectMongoDB("integrations")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var integration bson.M
		err = collection.FindOne(ctx, bson.M{
			"token":      token,
			"is_revoked": false,
		}).Decode(&integration)

		if err != nil {
			// If not an integration token, it might be a valid session token (for admin ops)
			// BUT the user said "all requests will need the api token set as Authorization"
			// This implies the client itself must have a token.

			// Let's re-verify the request: "From now on all requests will need the api token set as Authorization"
			// If we allow either Integration token OR Session token, we might be safe.
			// But usually, the client token is for the APP, and session token is for the USER.

			// For now, let's treat the Integration tokens as the primary gatekeeper.
			http.Error(w, "Unauthorized: Invalid or revoked integration token", http.StatusUnauthorized)
			return
		}

		// Update last used
		go func() {
			collection.UpdateOne(context.Background(), bson.M{"token": token}, bson.M{
				"$set": bson.M{"last_used_at": time.Now().Format("02-01-2006 15:04:05")},
			})
		}()

		next.ServeHTTP(w, r)
	}
}
