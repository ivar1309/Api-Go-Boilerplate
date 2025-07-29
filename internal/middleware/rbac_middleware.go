package middleware

import (
	"net/http"

	"github.com/ivar1309/Api-Go-Boilerplate/internal/utils"
)

func RequirePermission(permission utils.Permission) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			roleVal := r.Context().Value("role")
			role, ok := roleVal.(string)
			if !ok || !utils.HasPermission(role, permission) {
				http.Error(w, "Forbidden: insufficient permissions", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
