package middleware

import (
	"net/http"

	"user-management-system/utils"

	"github.com/gorilla/mux"
)

// RequireRole checks if the user has the required role
func RequireRole(allowedRoles ...string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := GetRole(r.Context())

			if userRole == "" {
				utils.ErrorResponse(w, http.StatusUnauthorized, "User role not found")
				return
			}

			// Check if user role is in allowed roles
			allowed := false
			for _, role := range allowedRoles {
				if userRole == role {
					allowed = true
					break
				}
			}

			if !allowed {
				utils.ErrorResponse(w, http.StatusForbidden, "Insufficient permissions")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// RequireAdmin is a convenience middleware for admin-only routes
func RequireAdmin() mux.MiddlewareFunc {
	return RequireRole("admin")
}

// RequireAuth ensures user is authenticated (can be used after JWT middleware)
func RequireAuth() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID := GetUserID(r.Context())
			if userID == "" {
				utils.ErrorResponse(w, http.StatusUnauthorized, "Authentication required")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// CanModifyUser checks if user can modify another user (own account or admin)
func CanModifyUser() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			userID := GetUserID(ctx)
			userRole := GetRole(ctx)

			// Get target user ID from URL
			vars := mux.Vars(r)
			targetUserID := vars["id"]

			// Allow if user is admin or modifying their own account
			if userRole == "admin" || userID == targetUserID {
				next.ServeHTTP(w, r)
				return
			}

			utils.ErrorResponse(w, http.StatusForbidden, "You can only modify your own account")
		})
	}
}

