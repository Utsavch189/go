package main

import (
	"fmt"
	"time"
)

func producer(id int, ch chan<- string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("Producer %d: %d", id, i)
		time.Sleep(100 * time.Millisecond) // Artificial delay
	}
	close(ch)
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go producer(1, ch1)
	go producer(2, ch2)

	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

/*
select allows waiting on multiple channel operations.

Use Case: Fan-In Pattern
Combine multiple data streams into a single channel.
*/
