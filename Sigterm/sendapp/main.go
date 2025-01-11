package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func sendRequest(url string, payloadSize int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when the goroutine completes

	// Create a new HTTP client with the Keep-Alive connection
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false, // Ensure that connections are kept alive
		},
	}

	// Infinite loop to simulate continuous connections
	for {
		// Create the payload (large byte slice) for each request
		largePayload := make([]byte, payloadSize)
		for j := range largePayload {
			largePayload[j] = 'A' // Fill with 'A' for simplicity
		}

		// Create a new POST request with the correct payload
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(largePayload))
		if err != nil {
			log.Printf("Failed to create request: %v", err)
			return
		}

		// Set headers and content length
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Connection", "keep-alive") // Keep the connection open
		req.ContentLength = int64(len(largePayload)) // Set the correct Content-Length

		// Send the request
		startTime := time.Now()
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Request failed: %v\n", err)
			return
		}
		defer resp.Body.Close()

		// Log response and timing
		elapsedTime := time.Since(startTime)
		fmt.Printf("Response status: %s, Time taken: %v\n", resp.Status, elapsedTime)

		// Sleep for a short duration to simulate delay before the next request
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// URL of the server
	url := "http://localhost:10808/slack"

	// Use WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Simulate infinite active connections by starting multiple goroutines
	for i := 0; i < 5; i++ {
		// Set payload size (e.g., 100MB for each request)
		payloadSize := 50 * 1024 * 1024 // 100MB payload for each request

		// Add 1 to the WaitGroup counter for each new goroutine
		wg.Add(1)

		// Run the sendRequest function concurrently in a goroutine
		go sendRequest(url, payloadSize, &wg)
	}

	// Wait for all goroutines to finish before exiting (not really needed here since it's infinite)
	wg.Wait()
}
