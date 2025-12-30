package main

import (
	"fmt"
	"log"

	"go-url-shortener/internal/config"
	"go-url-shortener/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration
	cfg := config.LoadConfig()

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
