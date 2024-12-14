package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

/*
Contexts handle timeouts, deadlines, and cancellations.

Use Case: API Timeout Handling
Terminate API requests that exceed a timeout.
*/

func fetchWithTimeout(ctx context.Context, url string, done chan bool) {
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// Check if the error is due to context cancellation or timeout
		if ctx.Err() == context.DeadlineExceeded {
			done <- false // Notify that the request timed out
			return
		}
		fmt.Println("Error:", err)
		done <- true // Notify that the request completed but failed for another reason
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.StatusCode)
	done <- true // Notify that the request completed successfully
}

func main() {
	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	done := make(chan bool)

	// Launch the HTTP request in a goroutine
	go fetchWithTimeout(ctx, "https://httpbin.org/delay/3", done) // Simulated delay

	select {
	case success := <-done:
		if !success {
			fmt.Println("Request timed out.")
		} else {
			fmt.Println("Request completed successfully.")
		}
	case <-ctx.Done():
		// Fallback, but this case should rarely happen now
		fmt.Println("Request timed out.")
	}
}

/*
How It Works
	The fetchWithTimeout function sends true to the done channel if the HTTP request completes (even with an error unrelated to timeout).
	If the request times out, it sends false to indicate the timeout.
	The main function listens on the done channel and prints "Request timed out." only when the timeout occurs.
*/
