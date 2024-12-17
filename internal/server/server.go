// internal/server/server.go
package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/llm-api/internal/config"
	"github.com/yourusername/llm-api/internal/handler"
)

type Server struct {
	router     *mux.Router
	llmHandler *handler.LLMHandler
	config     *config.Configuration
}

func NewServer(cfg *config.Configuration, llmHandler *handler.LLMHandler) *Server {
	router := mux.NewRouter()
	server := &Server{
		router:     router,
		llmHandler: llmHandler,
		config:     cfg,
	}
	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	s.router.HandleFunc("/v1/completions", s.llmHandler.HandleCompletion).Methods("POST")
}

func (s *Server) Start() error {
	log.Printf("Starting server on port %s", s.config.ServerPort)
	return http.ListenAndServe(":" + s.config.ServerPort, s.router)
}
