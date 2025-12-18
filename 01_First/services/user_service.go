package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"user-management-system/models"
	"user-management-system/repositories"

	"golang.org/x/crypto/bcrypt"
)

// UserService handles business logic for users
type UserService struct {
	userRepo *repositories.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// Register creates a new user account
func (s *UserService) Register(ctx context.Context, req *models.RegisterRequest) (*models.UserResponse, error) {
	// Validate input
	if err := s.validateRegisterRequest(req); err != nil {
		return nil, err
	}

	// Convert email to lowercase
	email := strings.ToLower(strings.TrimSpace(req.Email))

	// Check if user already exists
	existingUser, _ := s.userRepo.FindByEmail(ctx, email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create user
	user := &models.User{
		Name:      req.Name,
		Email:     email,
		Password:  string(hashedPassword),
		Role:      "user", // Default role
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, errors.New("failed to create user")
	}

	return user.ToUserResponse(), nil
}

// Login authenticates a user and returns user info
func (s *UserService) Login(ctx context.Context, req *models.LoginRequest) (*models.User, error) {
	// Validate input
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}

	// Convert email to lowercase
	email := strings.ToLower(strings.TrimSpace(req.Email))

	// Find user by email
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user.ToUserResponse(), nil
}

// UpdateUser updates user information
func (s *UserService) UpdateUser(ctx context.Context, id string, req *models.UpdateUserRequest) (*models.UserResponse, error) {
	// Build update data
	updateData := make(map[string]interface{})

	if req.Name != "" {
		updateData["name"] = req.Name
	}

	if req.Email != "" {
		// Convert email to lowercase
		email := strings.ToLower(strings.TrimSpace(req.Email))

		// Check if email is already taken by another user
		existingUser, _ := s.userRepo.FindByEmail(ctx, email)
		if existingUser != nil && existingUser.ID.Hex() != id {
			return nil, errors.New("email already in use")
		}
		updateData["email"] = email
	}

	if req.Password != "" {
		// Hash new password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.New("failed to hash password")
		}
		updateData["password"] = string(hashedPassword)
	}

	if req.Role != "" {
		updateData["role"] = req.Role
	}

	if req.IsActive != nil {
		updateData["isActive"] = *req.IsActive
	}

	// Update user
	if err := s.userRepo.Update(ctx, id, updateData); err != nil {
		return nil, err
	}

	// Fetch updated user
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user.ToUserResponse(), nil
}

// DeleteUser removes a user from the system
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.userRepo.Delete(ctx, id)
}

// GetAllUsers retrieves all users with pagination
func (s *UserService) GetAllUsers(ctx context.Context, page, limit int) ([]*models.UserResponse, int, int64, error) {
	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100 // Max limit
	}

	users, total, err := s.userRepo.FindAll(ctx, page, limit)
	if err != nil {
		return nil, 0, 0, err
	}

	// Convert to response format
	userResponses := make([]*models.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = user.ToUserResponse()
	}

	// Calculate total pages
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return userResponses, totalPages, total, nil
}

// validateRegisterRequest validates registration input
func (s *UserService) validateRegisterRequest(req *models.RegisterRequest) error {
	if req.Name == "" || len(req.Name) < 2 {
		return errors.New("name must be at least 2 characters")
	}

	if req.Email == "" {
		return errors.New("email is required")
	}

	if req.Password == "" || len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	return nil
}
