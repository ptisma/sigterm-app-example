package task

import (
	"fmt"
	"sync"
	"time"
)

// TaskRunner represents the long running task.
type TaskRunner struct {
	taskLength int // in seconds
	taskWG          *sync.WaitGroup
	taskRunning bool
	taskLock   sync.Mutex
}

// Start starts the long running task.
func (t *TaskRunner) Start() {
	defer t.taskWG.Done()
	t.taskLock.Lock()
	t.taskRunning = true
	t.taskLock.Unlock()

	fmt.Println("Long-running task started...")
	time.Sleep(time.Duration(t.taskLength)*time.Second)
	fmt.Println("Long-running task completed.")
}

// CheckStatus returns if the task is running
func (t *TaskRunner) CheckStatus() bool{
	return t.taskRunning
}


// NewServer creates a new instance of Server.
func NewTaskRunner(taskLength int, taskWG *sync.WaitGroup) *TaskRunner {
	return &TaskRunner{
		taskLength: taskLength,
		taskWG: taskWG,
	}
}