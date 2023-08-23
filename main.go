package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"github.com/ptisma/sigterm-app-example/server"
)

var (
	taskWG          sync.WaitGroup
)

func main() {
	// Start a long-running task in the background
	taskWG.Add(1)
	go longRunningTask()

	// Set up an HTTP server
	srv := server.NewServer(8080)
	go func() {
		if err := srv.Start(); err != nil {
			fmt.Println("HTTP server error:", err)
		}
	}()

	// Set up a signal handler to catch SIGTERM
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)
	<-sigChan

	fmt.Println("Received SIGTERM. Changing response code to 503.")
	srv.ChangeHealthCheck()

	// Wait for the long-running task to finish
	taskWG.Wait()

	fmt.Println("Shutting down gracefully...")
}

func longRunningTask() {
	defer taskWG.Done() // Decrement the WaitGroup counter when the task is done
	fmt.Println("Long-running task started...")
	// Simulate a long-running task
	time.Sleep(time.Minute * 60)
	fmt.Println("Long-running task completed.")
}








