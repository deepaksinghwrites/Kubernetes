package main

import (
	"log"
	"net/http"
	"os"
	"io"        // Import the io package
	"os/signal"
	"sync/atomic"
	"syscall"
	
	"github.com/gin-gonic/gin"
)

var (
	requestCount  int64 // Counter for incoming requests
	responseCount int64 // Counter for responses
)

func main() {
	// Create a channel to listen for system signals
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Block until a signal is received
		<-signalChannel
		// Log the total requests and responses
		log.Printf("Total requests: %d, Total responses: %d\n", atomic.LoadInt64(&requestCount), atomic.LoadInt64(&responseCount))
		// Exit the program immediately
		os.Exit(0)
	}()

	r := gin.Default()

	// Routes
	r.POST("/slack", slackHandler)

	// Retrieve the port from the environment variable or default to 10808
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "10808"
	}

	// Start the server
	log.Printf("Starting server on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Handler for the /slack endpoint
func slackHandler(c *gin.Context) {
    // Read the entire payload from the request body
	atomic.AddInt64(&requestCount, 1) // Increment request count
    body, err := io.ReadAll(c.Request.Body)
	atomic.AddInt64(&responseCount, 1) // Increment response count after the request is processed
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
        return
    }

    // Log the size of the payload
    log.Printf("Received payload of size: %d bytes", len(body))

    // Respond to the client
    c.JSON(http.StatusOK, gin.H{"message": "Slack endpoint hit!"})
}

