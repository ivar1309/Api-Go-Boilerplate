package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/ivar1309/Api-Go-Boilerplate/internal/utils"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "username", claims.Username)
		ctx = context.WithValue(ctx, "role", claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
