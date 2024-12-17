// internal/service/llm_service.go
package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sashabaranov/go-openai"
	"github.com/yourusername/llm-api/internal/config"
)

type LLMService interface {
	GenerateCompletion(ctx context.Context, prompt string) (string, error)
}

type OpenRouterLLMService struct {
	client *openai.Client
	config *config.Configuration
}

func NewOpenRouterLLMService(cfg *config.Configuration) *OpenRouterLLMService {
	// Create a custom OpenAI client configuration
	clientConfig := openai.DefaultConfig(cfg.OpenRouterAPIKey)
	
	// Set Open Router's API base URL
	clientConfig.BaseURL = "https://openrouter.ai/api/v1"
	
	// Create a custom HTTP client with additional headers
	customClient := &http.Client{}
	clientConfig.HTTPClient = customClient

	client := openai.NewClientWithConfig(clientConfig)

	return &OpenRouterLLMService{
		client: client,
		config: cfg,
	}
}

func (s *OpenRouterLLMService) GenerateCompletion(ctx context.Context, prompt string) (string, error) {
	// Create a chat completion request with Open Router specific configurations
	req := openai.ChatCompletionRequest{
		Model: s.config.DefaultModel,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		// Optional: Add Open Router specific parameters
		MaxTokens: 150,  // Example token limit
	}

	resp, err := s.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("open router completion error: %v", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no completion choices returned")
	}

	return resp.Choices[0].Message.Content, nil
}
