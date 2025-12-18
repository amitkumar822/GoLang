package routes

import (
	"net/http"

	"user-management-system/config"
	"user-management-system/handlers"
	"user-management-system/middleware"
	"user-management-system/services"

	"github.com/gorilla/mux"
)

// SetupRoutes configures all application routes
func SetupRoutes(
	userService *services.UserService,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	cfg *config.Config,
) *mux.Router {
	router := mux.NewRouter()

	// Apply global middleware
	router.Use(middleware.RecoveryMiddleware())
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.CORSMiddleware())

	// Root endpoint (welcome message)
	homeHandler := handlers.NewHomeHandler()
	router.HandleFunc("/", homeHandler.Welcome).Methods("GET")

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// Auth routes (public)
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", authHandler.Register).Methods("POST")
	auth.HandleFunc("/login", authHandler.Login).Methods("POST")

	// User routes (protected)
	users := api.PathPrefix("/users").Subrouter()
	users.Use(middleware.JWTMiddleware(cfg))
	users.Use(middleware.RequireAuth())

	// Get all users (with pagination)
	users.HandleFunc("", userHandler.GetAllUsers).Methods("GET")

	// Get single user
	users.HandleFunc("/{id}", userHandler.GetUser).Methods("GET")

	// Update user (user can update own account, admin can update any)
	users.HandleFunc("/{id}", applyMiddleware(
		userHandler.UpdateUser,
		middleware.CanModifyUser(),
	)).Methods("PUT")

	// Delete user (admin only)
	users.HandleFunc("/{id}", applyMiddleware(
		userHandler.DeleteUser,
		middleware.RequireAdmin(),
	)).Methods("DELETE")

	return router
}

// applyMiddleware applies middleware to an http.HandlerFunc
func applyMiddleware(handler http.HandlerFunc, mw mux.MiddlewareFunc) http.HandlerFunc {
	return mw(http.HandlerFunc(handler)).ServeHTTP
}
