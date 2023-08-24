package server

import (
	"fmt"
	"net/http"
	"sync"
	"github.com/ptisma/sigterm-app-example/task"
)

// Server represents the HTTP server.
type Server struct {
	port int
	unHealtyhResponse bool
	healthResponseLock   sync.Mutex
	taskRunner		     *task.TaskRunner
}

// HealthHandler handles the /healthz endpoint.
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	if s.unHealtyhResponse {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, "Service is unavailable")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Service is healthy")
	}
}

// TaskHandler handles the /task endpoint.
// Single GET request starts the task
func (s *Server) taskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	msg := ""
	if s.taskRunner.CheckStatus(){
		w.WriteHeader(http.StatusInternalServerError)
		msg = "Task already started"
	} else {
		w.WriteHeader(http.StatusOK)
		msg = "Starting the task"
		go s.taskRunner.Start()
	}
	
	
	fmt.Fprint(w, msg)
	
}

// Start starts the HTTP server.
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.port)
	http.HandleFunc("/healthz", s.healthHandler)
	http.HandleFunc("/task", s.taskHandler)
	return http.ListenAndServe(addr, nil)
}

// Change the health-check to unhealthy
func (s *Server) ChangeHealthCheck() {
	s.healthResponseLock.Lock()
	defer s.healthResponseLock.Unlock()
	s.unHealtyhResponse = true
}

// NewServer creates a new instance of Server.
func NewServer(port int, taskRunner *task.TaskRunner) *Server {
	return &Server{
		port: port,
		taskRunner: taskRunner,
	}
}





