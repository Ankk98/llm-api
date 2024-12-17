// internal/handler/llm_handler.go
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yourusername/llm-api/internal/service"
)

type LLMHandler struct {
	llmService service.LLMService
}

func NewLLMHandler(service service.LLMService) *LLMHandler {
	return &LLMHandler{
		llmService: service,
	}
}

func (h *LLMHandler) HandleCompletion(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Prompt string `json:"prompt"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	completion, err := h.llmService.GenerateCompletion(ctx, request.Prompt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"completion": completion,
	})
}
