package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Middleware to log each incoming request
	r.Use(requestLogger())

	// Routes
	r.POST("/slack", slackHandler)

	// Retrieve the port from the environment variable or default to 8080
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

// Middleware to log details of each incoming request
func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Incoming request: %s %s from %s", c.Request.Method, c.Request.URL.Path, c.ClientIP())
		c.Next()
	}
}

// Handler for the /slack endpoint
func slackHandler(c *gin.Context) {
	// Your handler logic here
	c.JSON(http.StatusOK, gin.H{"message": "Slack endpoint hit!"})
}
