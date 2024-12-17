// internal/config/config.go
package config

type Configuration struct {
	ServerPort       string
	OpenRouterAPIKey string
	DefaultModel     string
}

func NewConfiguration(port, apiKey, model string) *Configuration {
	cfg := &Configuration{
		ServerPort:       port,
		OpenRouterAPIKey: apiKey,
		DefaultModel:     model,
	}

	// Set default values if not provided
	if cfg.ServerPort == "" {
		cfg.ServerPort = "8080"
	}
	if cfg.DefaultModel == "" {
		cfg.DefaultModel = "anthropic/claude-3-haiku"
	}

	return cfg
}
