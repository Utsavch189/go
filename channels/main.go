package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d started task %d\n", id, task)
		time.Sleep(time.Second) // Simulate processing
		fmt.Printf("Worker %d finished task %d\n", id, task)
		results <- task * 2 // Return result
	}
}

func logger(logs <-chan string) {
	for log := range logs {
		fmt.Println("Log:", log)
	}
}

func main() {
	// channel()
	BufferedChannel()
}

func BufferedChannel() {
	logs := make(chan string, 3) // Buffered channel

	go logger(logs)

	logs <- "Start Process"
	logs <- "Processing Data"
	logs <- "End Process"

	time.Sleep(time.Second) // Allow logger to finish
	close(logs)

	/*
		Buffered channels allow asynchronous communication.

		Use Case: Logging System
		Buffer logs to avoid blocking the producer.
	*/
}

func channel() {
	tasks := make(chan int, 5)
	results := make(chan int, 5)

	// Start workers
	for i := 1; i <= 3; i++ {
		go worker(i, tasks, results)
	}

	// Send tasks
	for j := 1; j <= 5; j++ {
		tasks <- j
	}
	close(tasks)

	// Collect results
	for k := 1; k <= 5; k++ {
		fmt.Printf("Result: %d\n", <-results)
	}
	/*
	    OUTPUT:
	   	Worker 3 started task 1
	   	Worker 1 started task 2
	   	Worker 2 started task 3
	   	Worker 2 finished task 3
	   	Worker 2 started task 4
	   	Result: 6
	   	Worker 3 finished task 1
	   	Worker 3 started task 5
	   	Result: 2
	   	Worker 1 finished task 2
	   	Result: 4
	   	Worker 3 finished task 5
	   	Worker 2 finished task 4
	   	Result: 10
	   	Result: 8
	*/

	/*
	   This program demonstrates concurrent task processing using worker goroutines in Go.
	   It uses channels for communication between the main function and the workers,
	   employing a worker pool pattern. Here's a detailed breakdown:

	   	Inputs:
	   		id: Identifies the worker.
	   		tasks: A channel from which the worker receives tasks.
	   		results: A channel to send back the results of the task.

	   	Logic:
	   		Workers continuously listen for incoming tasks from the tasks channel.
	   		For each task, they "process" it (simulated by time.Sleep) and then send the result to the results channel.

	   	Task and Results Channels:
	   		tasks: Holds the tasks for workers.
	   		results: Receives the processed results from workers.

	   	Worker Pool:
	   		Three workers (worker(1), worker(2), and worker(3)) are started as goroutines.
	   		These workers process tasks concurrently, each fetching from the shared tasks channel.

	   	Task Dispatch:
	   		Five tasks (integers 1 through 5) are sent to the tasks channel.
	   		The channel is then closed to indicate no more tasks will be sent.

	   	Result Collection:
	   		The program collects and prints the results of the tasks from the results channel.

	   	How It Works
	   		The main function launches 3 workers that listen on the tasks channel.
	   		The main function sends 5 tasks into the tasks channel. Since there are more tasks than workers, workers will handle the tasks in sequence.
	   		Workers process tasks concurrently, printing their progress.
	   		As workers finish, they send their results to the results channel.
	   		The main function waits for all results to be collected before exiting.

	   	Channel Semantics:
	   		Channels in Go ensure that each value sent to the channel is received by exactly one receiver.
	   		Once a worker retrieves a task from the tasks channel, that task is no longer available to other
	   		workers.

	   	Why No WaitGroup Is Needed (In This Case):
	   		sync.WaitGroup is typically used to explicitly manage concurrency and synchronization,
	   		particularly when you want to wait for all goroutines to finish before continuing in the main
	   		function.

	   		However, in this code, the blocking receive from the results channel is acting as a natural
	   		synchronization mechanism. The program will not proceed until all the results are received,
	   		implying that all worker goroutines have completed their tasks and sent their results back.
	*/
}
