# Agent Template

This is a template for creating new agents in the multi-agent microservices architecture.

## Usage

To create a new agent based on this template:

1. Copy this directory and rename it to your agent name
2. Replace all instances of `agent-template` with your agent name
3. Update the configuration files with agent-specific settings
4. Implement your agent logic in the appropriate files
5. Create corresponding test files in `src/test/agents/[agent-name]/`

## Structure

```
agent-template/
├── README.md           # This file
├── Dockerfile         # Container configuration (to be created)
├── requirements.txt   # Python dependencies (for Python agents)
├── package.json       # Node.js dependencies (for Node.js agents)
├── go.mod            # Go dependencies (for Go agents)
├── src/              # Source code
│   ├── main.py       # Main application file (Python)
│   ├── index.js      # Main application file (Node.js)
│   ├── main.go       # Main application file (Go)
│   ├── api/          # API route definitions
│   ├── models/       # Data models
│   ├── services/     # Business logic
│   └── utils/        # Utility functions
├── tests/            # Unit tests
├── docs/             # Agent-specific documentation
└── config/           # Agent-specific configuration
```

## Implementation Checklist

### Basic Setup
- [ ] Rename directory from `agent-template` to `[your-agent-name]`
- [ ] Update README.md with agent-specific information
- [ ] Choose programming language (Python/Node.js/Go)
- [ ] Set up dependencies file (requirements.txt/package.json/go.mod)
- [ ] Create Dockerfile for containerization

### API Implementation
- [ ] Define API endpoints in OpenAPI specification
- [ ] Implement health check endpoint (`/health`)
- [ ] Implement status endpoint (`/status`)
- [ ] Implement main agent functionality endpoints
- [ ] Add proper error handling and validation

### Configuration
- [ ] Define environment-specific configurations
- [ ] Set up logging configuration
- [ ] Configure database connections (if needed)
- [ ] Set up external service integrations

### Testing
- [ ] Create unit tests with minimum 80% coverage
- [ ] Implement integration tests
- [ ] Add performance tests
- [ ] Set up test data and fixtures

### Documentation
- [ ] Write comprehensive API documentation
- [ ] Document deployment procedures
- [ ] Create usage examples
- [ ] Document configuration options

### Deployment
- [ ] Create Kubernetes manifests
- [ ] Set up monitoring and health checks
- [ ] Configure logging and metrics
- [ ] Test deployment in development environment

## Agent Requirements

Based on the architecture specifications, all agents must:

1. **Independence** (REQ-001): Implement as independent microservice
2. **API Communication** (REQ-002): Use APIs for all communication
3. **Containerization** (REQ-004): Be containerized for deployment
4. **Stateless Design** (CON-002): Be stateless or manage state externally
5. **Security** (SEC-001): Implement authenticated and encrypted communication

## Technology Guidelines

### Python with FastAPI
```python
from fastapi import FastAPI
from fastapi.responses import JSONResponse

app = FastAPI(title="Agent Template", version="1.0.0")

@app.get("/health")
async def health_check():
    return {"status": "healthy", "timestamp": "2025-07-23T09:00:00Z"}

@app.get("/status")
async def get_status():
    return {
        "agent": "agent-template",
        "version": "1.0.0",
        "status": "active"
    }

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
```

### Node.js with Express
```javascript
const express = require('express');
const app = express();
const PORT = process.env.PORT || 3000;

app.use(express.json());

// Health check endpoint
app.get('/health', (req, res) => {
    res.json({
        status: 'healthy',
        timestamp: new Date().toISOString()
    });
});

// Status endpoint
app.get('/status', (req, res) => {
    res.json({
        agent: 'agent-template',
        version: '1.0.0',
        status: 'active'
    });
});

app.listen(PORT, () => {
    console.log(`Agent Template running on port ${PORT}`);
});
```

### Go with Gin
```go
package main

import (
    "net/http"
    "time"
    
    "github.com/gin-gonic/gin"
)

type HealthResponse struct {
    Status    string `json:"status"`
    Timestamp string `json:"timestamp"`
}

type StatusResponse struct {
    Agent   string `json:"agent"`
    Version string `json:"version"`
    Status  string `json:"status"`
}

func main() {
    r := gin.Default()
    
    // Health check endpoint
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, HealthResponse{
            Status:    "healthy",
            Timestamp: time.Now().Format(time.RFC3339),
        })
    })
    
    // Status endpoint
    r.GET("/status", func(c *gin.Context) {
        c.JSON(http.StatusOK, StatusResponse{
            Agent:   "agent-template",
            Version: "1.0.0",
            Status:  "active",
        })
    })
    
    r.Run(":8080")
}
```

## Configuration Template

### Environment Variables
```bash
# Agent Configuration
AGENT_NAME=agent-template
AGENT_VERSION=1.0.0
AGENT_PORT=8000

# Database Configuration
DATABASE_URL=postgresql://localhost:5432/agent_template
REDIS_URL=redis://localhost:6379

# External Services
API_GATEWAY_URL=http://localhost:8080
AUTH_SERVICE_URL=http://localhost:3001

# Logging
LOG_LEVEL=info
LOG_FORMAT=json

# Monitoring
METRICS_ENABLED=true
HEALTH_CHECK_INTERVAL=30s
```

### Docker Template
```dockerfile
FROM python:3.9-slim

WORKDIR /app

# Copy dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy source code
COPY src/ ./src/

# Expose port
EXPOSE 8000

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8000/health || exit 1

# Run application
CMD ["python", "src/main.py"]
```

## Next Steps

After customizing this template:

1. Run the code generation tool: `./tools/generators/generate-agent.sh --from-template`
2. Implement your specific agent logic
3. Add comprehensive tests
4. Create API documentation
5. Test locally with Docker
6. Deploy to development environment
7. Add monitoring and alerting
8. Document usage and maintenance procedures
