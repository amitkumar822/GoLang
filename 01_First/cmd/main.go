package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"user-management-system/config"
	"user-management-system/database"
	"user-management-system/handlers"
	"user-management-system/repositories"
	"user-management-system/routes"
	"user-management-system/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to MongoDB
	if err := database.Connect(cfg.MongoURI, cfg.MongoDB); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer database.Disconnect()

	// Initialize repositories
	userCollection := database.GetCollection("users")
	userRepo := repositories.NewUserRepository(userCollection)

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(userService, cfg)
	userHandler := handlers.NewUserHandler(userService)

	// Setup routes
	router := routes.SetupRoutes(userService, authHandler, userHandler, cfg)

	// Create HTTP server
	server := &http.Server{
		Addr:         ":" + cfg.AppPort,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("🚀 Server starting on port %s", cfg.AppPort)
		log.Printf("🌐 Root endpoint: http://localhost:%s/", cfg.AppPort)
		log.Printf("📝 API endpoints available at http://localhost:%s/api", cfg.AppPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("✅ Server exited")
}

