package handlers

import (
	"encoding/json"
	"net/http"

	"user-management-system/config"
	"user-management-system/models"
	"user-management-system/services"
	"user-management-system/utils"
)

// AuthHandler handles authentication-related requests
type AuthHandler struct {
	authService *services.UserService
	config      *config.Config
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService *services.UserService, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		config:      cfg,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, err := h.authService.Register(r.Context(), &req)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(w, "User registered successfully", user)
}

// Login handles user login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, err := h.authService.Login(r.Context(), &req)
	if err != nil {
		utils.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID.Hex(), user.Email, user.Role, h.config)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Return user info and token
	response := map[string]interface{}{
		"user":  user.ToUserResponse(),
		"token": token,
	}

	utils.SuccessResponse(w, "Login successful", response)
}

