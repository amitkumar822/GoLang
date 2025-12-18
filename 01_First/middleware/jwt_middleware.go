package middleware

import (
	"context"
	"net/http"
	"strings"

	"user-management-system/config"
	"user-management-system/utils"

	"github.com/gorilla/mux"
)

// ContextKey is a custom type for context keys
type ContextKey string

const (
	UserIDKey ContextKey = "userId"
	EmailKey  ContextKey = "email"
	RoleKey   ContextKey = "role"
)

// JWTMiddleware validates JWT tokens and extracts user information
func JWTMiddleware(cfg *config.Config) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get token from Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.ErrorResponse(w, http.StatusUnauthorized, "Authorization header required")
				return
			}

			// Extract token from "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid authorization header format")
				return
			}

			tokenString := parts[1]

			// Validate token
			claims, err := utils.ValidateToken(tokenString, cfg)
			if err != nil {
				utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid or expired token")
				return
			}

			// Add user information to context
			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, EmailKey, claims.Email)
			ctx = context.WithValue(ctx, RoleKey, claims.Role)

			// Call next handler with updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserID extracts user ID from context
func GetUserID(ctx context.Context) string {
	if userID, ok := ctx.Value(UserIDKey).(string); ok {
		return userID
	}
	return ""
}

// GetEmail extracts email from context
func GetEmail(ctx context.Context) string {
	if email, ok := ctx.Value(EmailKey).(string); ok {
		return email
	}
	return ""
}

// GetRole extracts role from context
func GetRole(ctx context.Context) string {
	if role, ok := ctx.Value(RoleKey).(string); ok {
		return role
	}
	return ""
}

