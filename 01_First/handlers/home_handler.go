package handlers

import (
	"net/http"

	"user-management-system/utils"
)

// HomeHandler handles root endpoint requests
type HomeHandler struct{}

// NewHomeHandler creates a new home handler
func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

// Welcome handles the root endpoint
func (h *HomeHandler) Welcome(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message": "Welcome to Go Lang User API",
		"version": "1.0.0",
		"status":  "running",
		"endpoints": map[string]string{
			"register": "/api/auth/register",
			"login":    "/api/auth/login",
			"users":    "/api/users",
		},
	}

	utils.SuccessResponse(w, "API is running", response)
}

