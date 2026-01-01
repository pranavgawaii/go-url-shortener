package main

import (
	"fmt"
	"log"

	"go-url-shortener/internal/config"
	"go-url-shortener/internal/handler"
	"go-url-shortener/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration
	cfg := config.LoadConfig()

	// Initialize Database
	db, err := repository.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize Repositories
	// urlRepo := repository.NewURLRepository(db) // Will be used in service layer later

	// Initialize Gin engine
	r := gin.Default()

	// Initialize handlers
	healthHandler := handler.NewHealthHandler()

	// Register routes
	r.GET("/health", healthHandler.GetHealth)

	// Start server
	serverAddr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Starting server on port %s", cfg.Port)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
