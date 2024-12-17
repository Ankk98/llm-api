# Open Router LLM API Server

## Overview
This is a Go-based microservice that provides a simple API for generating text completions using Open Router's LLM services.

## Features
- OpenAI-compatible API endpoint
- Support for Open Router LLM providers
- Configurable via environment variables
- Modular design with separation of concerns

## Prerequisites
- Go 1.21+
- Open Router API Key

## Installation

### 1. Clone the Repository
```bash
git clone https://github.com/ankk98/llm-api.git
cd llm-api
```

### 2. Install Dependencies
```bash
go mod tidy
```

## Configuration

### Environment Variables
- `SERVER_PORT`: HTTP server port (default: 8080)
- `OPEN_ROUTER_API_KEY`: Your Open Router API key (required)
- `DEFAULT_MODEL`: LLM model to use (default: anthropic/claude-3-haiku)

## Running the Server

### Development Mode
```bash
# Set your Open Router API key
export OPEN_ROUTER_API_KEY=your_open_router_api_key

# Run the server
go run main.go
```

### Building and Running
```bash
# Build the application
go build -o llm-api

# Run the executable
./llm-api
```

## API Endpoints

### Text Completion
- **Endpoint**: `/v1/completions`
- **Method**: POST
- **Content-Type**: `application/json`

#### Request Body
```json
{
  "prompt": "Your text prompt here"
}
```

## Testing the API

### Curl Command Examples

#### Simple Prompt
```bash
curl -X POST http://localhost:8080/v1/completions \
     -H "Content-Type: application/json" \
     -d '{"prompt": "Write a haiku about technology"}'
```

#### Creative Writing Prompt
```bash
curl -X POST http://localhost:8080/v1/completions \
     -H "Content-Type: application/json" \
     -d '{"prompt": "Create a short story about an AI discovering its purpose"}'
```

#### Code Generation Prompt
```bash
curl -X POST http://localhost:8080/v1/completions \
     -H "Content-Type: application/json" \
     -d '{"prompt": "Write a Python function to calculate fibonacci sequence"}'
```

## Development

### Running Tests
```bash
go test ./...
```

### Project Structure
```
llm-api/
├── go.mod
├── main.go
└── internal/
    ├── config/      # Configuration management
    ├── handler/     # HTTP request handlers
    ├── service/     # LLM service implementations
    └── server/      # HTTP server setup
```

## Error Handling
- Invalid JSON: Returns 400 Bad Request
- LLM Service Errors: Returns 500 Internal Server Error

## Contributing
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License
Distributed under the MIT License. See `LICENSE` for more information.

## Contact
Ankit Khandelwal - ankk98@gmail.com

Project Link: [https://github.com/ankk98/llm-api](https://github.com/ankk98/llm-api)
