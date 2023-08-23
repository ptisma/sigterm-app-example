package server

import (
	"fmt"
	"net/http"
	"sync"
)

// Server represents the HTTP server.
type Server struct {
	Port int
	UnHealtyhResponse bool
	healthResponseLock   sync.Mutex
}

// HealthHandler handles the /healthz endpoint.
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	if s.UnHealtyhResponse {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, "Service is unavailable")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Service is healthy")
	}
}

// Start starts the HTTP server.
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.Port)
	http.HandleFunc("/healthz", s.healthHandler)
	return http.ListenAndServe(addr, nil)
}

// Change the healt-check to unhealthy
func (s *Server) ChangeHealthCheck() {
	s.healthResponseLock.Lock()
	defer s.healthResponseLock.Unlock()
	s.UnHealtyhResponse = true
}

// NewServer creates a new instance of Server.
func NewServer(port int) *Server {
	return &Server{
		Port: port,
	}
}





