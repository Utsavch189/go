package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Utsavch189/api_mysql/internal/models/responses"
	"github.com/Utsavch189/api_mysql/internal/utils"
)

type contextKey string

const ClaimsContextKey contextKey = "claims"

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jwt_token := r.Header.Get("Authorization")

		if jwt_token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(responses.ErrorResponse(nil, "Please provide a JWT token as a Bearer token in the Authorization header!"))
			return
		}

		jwt_tokens := strings.Split(jwt_token, " ")

		if len(jwt_tokens) != 2 || jwt_tokens[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(responses.ErrorResponse(nil, "Invalid Authorization header format. Expected 'Bearer <token>'!"))
			return
		}

		jwt_token = jwt_tokens[1]

		claims, err := utils.ValidateJWT(jwt_token)
		if err != nil || claims == nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(responses.ErrorResponse(err, "JWT token validation failed!"))
			return
		}

		ctx := context.WithValue(r.Context(), ClaimsContextKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
