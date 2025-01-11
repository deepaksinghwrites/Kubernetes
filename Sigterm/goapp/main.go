package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/slack", slackHandler)

	// Retrieve the port from the environment variable or default to 8080
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "10808"
	}

	// Create the HTTP server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Graceful shutdown logic
	go func() {
		// Start the server in a goroutine to avoid blocking
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Listen for OS signals for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutdown signal received. Shutting down gracefully...")

	// Create a context with timeout for shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server shutdown complete.")
}

func slackHandler(c *gin.Context) {
	// Your handler logic here
	c.JSON(http.StatusOK, gin.H{"message": "Slack endpoint hit!"})
}