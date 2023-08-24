package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"github.com/ptisma/sigterm-app-example/server"
	"github.com/ptisma/sigterm-app-example/task"
)

func main() {
	// Set up a waiting group
	taskWG := &sync.WaitGroup{}
	taskWG.Add(1)
	// Set up a Task Runner
	tr := task.NewTaskRunner(60, taskWG)
	// Start a long-running task in the background
	
	// Set up an HTTP server
	srv := server.NewServer(8080, tr)
	go func() {
		if err := srv.Start(); err != nil {
			fmt.Println("HTTP server error:", err)
		}
	}()

	// Set up a signal handler to catch SIGTERM
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)
	go func() {
		// Wait for either the signal or the WaitGroup to complete
		select {
		case <-sigChan:
			fmt.Println("Received SIGTERM. Changing response code to 503.")
			srv.ChangeHealthCheck()
			if !tr.CheckStatus() {
				taskWG.Done()
			} 
		}
	}()
	// Wait for the long-running task to finish
	taskWG.Wait()
	fmt.Println("Shutting down gracefully...")
}










