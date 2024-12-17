// main.go
package main

import (
	"log"
	"os"

	"github.com/yourusername/llm-api/internal/config"
	"github.com/yourusername/llm-api/internal/handler"
	"github.com/yourusername/llm-api/internal/service"
	"github.com/yourusername/llm-api/internal/server"
)

func main() {
	cfg := config.NewConfiguration(
		os.Getenv("SERVER_PORT"),
		os.Getenv("OPEN_ROUTER_API_KEY"),
		os.Getenv("DEFAULT_MODEL"),
	)

	llmService := service.NewOpenRouterLLMService(cfg)
	llmHandler := handler.NewLLMHandler(llmService)
	srv := server.NewServer(cfg, llmHandler)

	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
