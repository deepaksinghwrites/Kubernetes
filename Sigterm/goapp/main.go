package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
	"io"        // Import the io package

	"github.com/gin-gonic/gin"
)

var (
	requestCount  int64 // Counter for incoming requests
	responseCount int64 // Counter for responses
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
	log.Printf("Total requests: %d, Total responses: %d\n", atomic.LoadInt64(&requestCount), atomic.LoadInt64(&responseCount))

	log.Println("Server shutdown complete.")
}

func slackHandler(c *gin.Context) {
    // Read the entire payload from the request body
	atomic.AddInt64(&requestCount, 1) // Increment request count
	body, err := io.ReadAll(c.Request.Body)                        
	atomic.AddInt64(&responseCount, 1) // Increment response count after request is handled
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
        return
    }

    // Log the size of the payload
    log.Printf("Received payload of size: %d bytes", len(body))

    // Respond to the client
    c.JSON(http.StatusOK, gin.H{"message": "Slack endpoint hit!"})
}
