package middleware

import (
	"context"
	"net/http"
	"product-api-postgres/internal/auth"
	"product-api-postgres/internal/response"
	"strings"
)

type contextKey string

const claimsKey contextKey = "claims"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.WriteError(w, http.StatusUnauthorized, "Authorization header is required")
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.WriteError(w, http.StatusUnauthorized, "Authorization header must use Bearer token")
			return
		}
		tokenString := parts[1]

		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			response.WriteError(w, http.StatusUnauthorized, "Invalid or expired token ")
			return
		}
		ctx := context.WithValue(r.Context(), claimsKey, claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func GetClaims(r *http.Request) (*auth.Claims, bool) {
	value := r.Context().Value(claimsKey)

	claims, ok := value.(*auth.Claims)

	return claims, ok
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := GetClaims(r)

		if !ok {
			response.WriteError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		if claims.UserRole != "admin" {
			response.WriteError(w, http.StatusForbidden, "Admin access required!")
			return
		}

		next.ServeHTTP(w, r)
	})

}
