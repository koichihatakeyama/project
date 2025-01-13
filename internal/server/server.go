package server

import (
	"net/http"
	"project/internal/handler"
	"project/pkg/logger"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	logger *logger.Logger
}

func NewServer(logger *logger.Logger) *Server {
	return &Server{
		router: mux.NewRouter(),
		logger: logger,
	}
}

func (s *Server) SetupRoutes(userHandler *handler.UserHandler) {
	s.router.HandleFunc("/users", userHandler.Create).Methods("POST")
	s.router.HandleFunc("/users/{id}", userHandler.GetByID).Methods("GET")
	s.router.Use(s.loggingMiddleware)
}

func (s *Server) Start(addr string) error {
	s.logger.Info("Starting server on %s", addr)
	return http.ListenAndServe(addr, s.router)
}

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.LogRequest(r.Method, r.URL.Path, 0)
		next.ServeHTTP(w, r)
	})
}
