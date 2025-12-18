package middleware

import (
	"log"
	"net/http"

	"user-management-system/utils"

	"github.com/gorilla/mux"
)

// RecoveryMiddleware recovers from panics and returns a proper error response
func RecoveryMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("Panic recovered: %v", err)
					utils.ErrorResponse(w, http.StatusInternalServerError, "Internal server error")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

