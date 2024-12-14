package main

import (
	"fmt"
	"time"
)

/*
Timers and tickers manage scheduled tasks.

Use Case: Periodic Data Refresh
Refresh a dashboard every second.
*/

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		for t := range ticker.C {
			fmt.Println("Refreshing data at", t)
		}
	}()

	time.Sleep(5 * time.Second) // Run for 5 seconds
	fmt.Println("Done.")
}
