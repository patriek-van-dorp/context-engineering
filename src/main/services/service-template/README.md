# Service Template

This is a template for creating new services in the multi-agent microservices architecture.

## Usage

To create a new service based on this template:

1. Copy this directory and rename it to your service name
2. Replace all instances of `service-template` with your service name
3. Update the configuration files with service-specific settings
4. Implement your service logic in the appropriate files
5. Create corresponding test files in `src/test/services/[service-name]/`

## Structure

```
service-template/
├── README.md           # This file
├── Dockerfile         # Container configuration (to be created)
├── requirements.txt   # Python dependencies (for Python services)
├── package.json       # Node.js dependencies (for Node.js services)
├── go.mod            # Go dependencies (for Go services)
├── src/              # Source code
│   ├── main.py       # Main application file (Python)
│   ├── index.js      # Main application file (Node.js)
│   ├── main.go       # Main application file (Go)
│   ├── api/          # API route definitions
│   ├── models/       # Data models
│   ├── services/     # Business logic
│   ├── middleware/   # Custom middleware
│   └── utils/        # Utility functions
├── tests/            # Unit tests
├── docs/             # Service-specific documentation
└── config/           # Service-specific configuration
```

## Implementation Checklist

### Basic Setup
- [ ] Rename directory from `service-template` to `[your-service-name]`
- [ ] Update README.md with service-specific information
- [ ] Choose programming language (Python/Node.js/Go)
- [ ] Set up dependencies file (requirements.txt/package.json/go.mod)
- [ ] Create Dockerfile for containerization

### API Implementation
- [ ] Define API endpoints in OpenAPI specification
- [ ] Implement health check endpoint (`/health`)
- [ ] Implement metrics endpoint (`/metrics`)
- [ ] Implement main service functionality endpoints
- [ ] Add proper error handling and validation
- [ ] Implement authentication middleware

### Data Layer
- [ ] Set up database models and schemas
- [ ] Implement data access layer
- [ ] Set up database migrations
- [ ] Configure connection pooling
- [ ] Add data validation

### Configuration
- [ ] Define environment-specific configurations
- [ ] Set up logging configuration
- [ ] Configure database connections
- [ ] Set up external service integrations
- [ ] Configure security settings

### Testing
- [ ] Create unit tests with minimum 80% coverage
- [ ] Implement integration tests
- [ ] Add performance tests
- [ ] Set up test data and fixtures
- [ ] Add contract tests for APIs

### Documentation
- [ ] Write comprehensive API documentation
- [ ] Document deployment procedures
- [ ] Create usage examples
- [ ] Document configuration options
- [ ] Add troubleshooting guide

### Deployment
- [ ] Create Kubernetes manifests
- [ ] Set up monitoring and health checks
- [ ] Configure logging and metrics
- [ ] Test deployment in development environment
- [ ] Set up automated deployments

## Service Requirements

Based on the architecture specifications, all services must:

1. **Microservice Architecture** (REQ-003): Implement as microservice with APIs
2. **API Communication** (REQ-002): Use APIs for all communication
3. **Containerization** (REQ-004): Be containerized for deployment
4. **Stateless Design** (CON-002): Be stateless or manage state externally
5. **Security** (SEC-001): Implement authenticated and encrypted communication

## Technology Guidelines

### Python with FastAPI
```python
from fastapi import FastAPI, Depends, HTTPException, status
from fastapi.security import HTTPBearer
from fastapi.responses import JSONResponse
import logging

# Initialize FastAPI app
app = FastAPI(
    title="Service Template",
    version="1.0.0",
    description="Template for creating new services"
)

# Security
security = HTTPBearer()
logger = logging.getLogger(__name__)

@app.get("/health")
async def health_check():
    """Health check endpoint for monitoring"""
    return {
        "status": "healthy",
        "service": "service-template",
        "timestamp": "2025-07-23T09:00:00Z"
    }

@app.get("/metrics")
async def get_metrics():
    """Metrics endpoint for monitoring"""
    return {
        "service": "service-template",
        "version": "1.0.0",
        "uptime": "1h 30m",
        "requests_total": 1000,
        "requests_per_second": 10.5
    }

@app.get("/api/v1/service-template")
async def get_service_data(token: str = Depends(security)):
    """Main service endpoint"""
    try:
        # Service logic here
        return {"message": "Service template working correctly"}
    except Exception as e:
        logger.error(f"Error in service endpoint: {e}")
        raise HTTPException(
            status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
            detail="Internal server error"
        )

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
```

### Node.js with Express
```javascript
const express = require('express');
const jwt = require('jsonwebtoken');
const helmet = require('helmet');
const cors = require('cors');

const app = express();
const PORT = process.env.PORT || 3000;

// Middleware
app.use(helmet());
app.use(cors());
app.use(express.json());

// Authentication middleware
const authenticateToken = (req, res, next) => {
    const authHeader = req.headers['authorization'];
    const token = authHeader && authHeader.split(' ')[1];

    if (!token) {
        return res.sendStatus(401);
    }

    jwt.verify(token, process.env.JWT_SECRET, (err, user) => {
        if (err) return res.sendStatus(403);
        req.user = user;
        next();
    });
};

// Health check endpoint
app.get('/health', (req, res) => {
    res.json({
        status: 'healthy',
        service: 'service-template',
        timestamp: new Date().toISOString()
    });
});

// Metrics endpoint
app.get('/metrics', (req, res) => {
    res.json({
        service: 'service-template',
        version: '1.0.0',
        uptime: process.uptime(),
        memory: process.memoryUsage(),
        pid: process.pid
    });
});

// Main service endpoint
app.get('/api/v1/service-template', authenticateToken, (req, res) => {
    try {
        // Service logic here
        res.json({ message: 'Service template working correctly' });
    } catch (error) {
        console.error('Error in service endpoint:', error);
        res.status(500).json({ error: 'Internal server error' });
    }
});

// Error handling middleware
app.use((err, req, res, next) => {
    console.error(err.stack);
    res.status(500).json({ error: 'Something went wrong!' });
});

app.listen(PORT, () => {
    console.log(`Service Template running on port ${PORT}`);
});
```

### Go with Gin
```go
package main

import (
    "net/http"
    "time"
    
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

type HealthResponse struct {
    Status    string `json:"status"`
    Service   string `json:"service"`
    Timestamp string `json:"timestamp"`
}

type MetricsResponse struct {
    Service string  `json:"service"`
    Version string  `json:"version"`
    Uptime  string  `json:"uptime"`
}

func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
            c.Abort()
            return
        }
        
        // Validate token here
        c.Next()
    }
}

func main() {
    r := gin.Default()
    
    // Middleware
    r.Use(cors.Default())
    
    // Health check endpoint
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, HealthResponse{
            Status:    "healthy",
            Service:   "service-template",
            Timestamp: time.Now().Format(time.RFC3339),
        })
    })
    
    // Metrics endpoint
    r.GET("/metrics", func(c *gin.Context) {
        c.JSON(http.StatusOK, MetricsResponse{
            Service: "service-template",
            Version: "1.0.0",
            Uptime:  "1h 30m",
        })
    })
    
    // Protected API routes
    api := r.Group("/api/v1")
    api.Use(authMiddleware())
    {
        api.GET("/service-template", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{
                "message": "Service template working correctly",
            })
        })
    }
    
    r.Run(":8080")
}
```

## Configuration Template

### Environment Variables
```bash
# Service Configuration
SERVICE_NAME=service-template
SERVICE_VERSION=1.0.0
SERVICE_PORT=8000

# Database Configuration
DATABASE_URL=postgresql://localhost:5432/service_template
REDIS_URL=redis://localhost:6379

# Authentication
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24h

# External Services
API_GATEWAY_URL=http://localhost:8080
AUTH_SERVICE_URL=http://localhost:3001

# Logging
LOG_LEVEL=info
LOG_FORMAT=json

# Monitoring
METRICS_ENABLED=true
HEALTH_CHECK_INTERVAL=30s

# Performance
MAX_CONNECTIONS=100
REQUEST_TIMEOUT=30s
```

### Docker Template
```dockerfile
FROM python:3.9-slim

WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Copy dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy source code
COPY src/ ./src/

# Create non-root user
RUN useradd -m -u 1000 appuser && chown -R appuser:appuser /app
USER appuser

# Expose port
EXPOSE 8000

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8000/health || exit 1

# Run application
CMD ["python", "src/main.py"]
```

## Database Integration

### PostgreSQL with SQLAlchemy (Python)
```python
from sqlalchemy import create_engine, Column, Integer, String
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

Base = declarative_base()

class ServiceData(Base):
    __tablename__ = "service_data"
    
    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, index=True)
    value = Column(String)

# Database setup
DATABASE_URL = "postgresql://user:password@localhost/dbname"
engine = create_engine(DATABASE_URL)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()
```

## API Documentation Template

### OpenAPI Specification
```yaml
openapi: 3.0.0
info:
  title: Service Template API
  version: 1.0.0
  description: Template for creating new services

servers:
  - url: http://localhost:8000
    description: Development server

paths:
  /health:
    get:
      summary: Health check endpoint
      responses:
        '200':
          description: Service is healthy
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: string
                  service:
                    type: string
                  timestamp:
                    type: string

  /api/v1/service-template:
    get:
      summary: Main service endpoint
      security:
        - bearerAuth: []
      responses:
        '200':
          description: Successful response
        '401':
          description: Unauthorized
        '500':
          description: Internal server error

components:
  securitySchemes:
    bearerAuth:
      type: http
      scheme: bearer
      bearerFormat: JWT
```

## Testing Template

### Unit Test Example (Python)
```python
import pytest
from fastapi.testclient import TestClient
from src.main import app

client = TestClient(app)

def test_health_check():
    response = client.get("/health")
    assert response.status_code == 200
    assert response.json()["status"] == "healthy"

def test_metrics_endpoint():
    response = client.get("/metrics")
    assert response.status_code == 200
    assert "service" in response.json()

@pytest.fixture
def auth_headers():
    return {"Authorization": "Bearer test-token"}

def test_protected_endpoint(auth_headers):
    response = client.get("/api/v1/service-template", headers=auth_headers)
    assert response.status_code == 200
```

## Next Steps

After customizing this template:

1. Run the code generation tool: `./tools/generators/generate-service.sh --from-template`
2. Implement your specific service logic
3. Set up database models and migrations
4. Add comprehensive tests
5. Create API documentation
6. Test locally with Docker
7. Deploy to development environment
8. Add monitoring and alerting
9. Document usage and maintenance procedures
10. Set up CI/CD pipeline
