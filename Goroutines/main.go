package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetchURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Fetched %s: %d\n", url, resp.StatusCode)
}

func firstWayGoroutine() {
	urls := []string{"https://example.com", "https://golang.org", "https://github.com"}
	for _, url := range urls {
		go fetchURL(url) // Concurrently fetch URLs
	}
	time.Sleep(3 * time.Second) // Prevent main goroutine from exiting early
	/*
		This a bad way beacuse here we are hold up main func for 3 seconds now all three fetchUrl's might be executed less or greater than 3 seconds.
	*/
}

func fetchURLNew(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wg.Done() // signal that I'm done
	defer resp.Body.Close()
	fmt.Printf("Fetched %s: %d\n", url, resp.StatusCode)
}

func secondwayGoroutine() {
	// We will sync all async fetchUrl() with main
	waitGroup := &sync.WaitGroup{}
	urls := []string{"https://example.com", "https://golang.org", "https://github.com"}
	for _, url := range urls {
		waitGroup.Add(1)               // Ensure this is called before the goroutine starts
		go fetchURLNew(url, waitGroup) // Concurrently fetch URLs with wait group
	}
	waitGroup.Wait() // join point for go routines
}

func main() {
	// firstWayGoroutine()
	secondwayGoroutine()
}
