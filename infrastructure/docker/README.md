# Docker Infrastructure

This directory contains optimized Dockerfiles for the multi-agent microservices architecture. Each Dockerfile implements security best practices, multi-stage builds, and production-ready configurations.

## Directory Structure

```
infrastructure/docker/
├── agents/           # Agent-specific Dockerfiles
│   ├── Dockerfile.python   # Python/FastAPI agents
│   ├── Dockerfile.nodejs   # Node.js/Express agents
│   └── Dockerfile.go       # Go/Gin agents
└── services/         # Service-specific Dockerfiles
    ├── Dockerfile.python   # Python/FastAPI services
    ├── Dockerfile.nodejs   # Node.js/Express services
    └── Dockerfile.go       # Go/Gin services
```

## Dockerfile Features

### Security
- **Non-root users**: All containers run as non-privileged users (uid/gid 1000)
- **Minimal base images**: Alpine and slim variants for reduced attack surface
- **Dependency verification**: Package integrity checks where applicable
- **Build-time security**: Separate build and runtime environments

### Performance
- **Multi-stage builds**: Separate build and production stages for smaller images
- **Layer optimization**: Strategic copying for better Docker layer caching
- **Dependency management**: Production-only dependencies in final stage
- **Static compilation**: Go binaries compiled statically for maximum portability

### Production Readiness
- **Health checks**: HTTP health endpoints on all services
- **Signal handling**: Proper init systems (dumb-init for Node.js)
- **Environment configuration**: Configurable via environment variables
- **Logging**: Structured logging with configurable levels

## Usage

### Building Images

#### Python Services/Agents
```bash
# Build Python service
docker build -f infrastructure/docker/services/Dockerfile.python -t my-python-service .

# Build Python agent (requires AGENT_TYPE build arg)
docker build -f infrastructure/docker/agents/Dockerfile.python --build-arg AGENT_TYPE=coordinator -t my-coordinator-agent .
```

#### Node.js Services/Agents
```bash
# Build Node.js service
docker build -f infrastructure/docker/services/Dockerfile.nodejs -t my-nodejs-service .

# Build Node.js agent
docker build -f infrastructure/docker/agents/Dockerfile.nodejs --build-arg AGENT_TYPE=processor -t my-processor-agent .
```

#### Go Services/Agents
```bash
# Build Go service
docker build -f infrastructure/docker/services/Dockerfile.go -t my-go-service .

# Build Go agent
docker build -f infrastructure/docker/agents/Dockerfile.go --build-arg AGENT_TYPE=orchestrator -t my-orchestrator-agent .
```

### Running Containers

#### Basic Usage
```bash
# Run Python service
docker run -p 8000:8000 --env LOG_LEVEL=debug my-python-service

# Run Node.js agent
docker run -p 3000:3000 --env AGENT_TYPE=processor my-processor-agent

# Run Go service with custom port
docker run -p 8080:8080 --env PORT=8080 my-go-service
```

#### Production Deployment
```bash
# Run with resource limits and restart policy
docker run -d \
  --name my-service \
  --restart unless-stopped \
  --memory 512m \
  --cpus 0.5 \
  -p 8000:8000 \
  --env LOG_LEVEL=info \
  --env NODE_ENV=production \
  my-python-service
```

## Environment Variables

### Common Variables
- `PORT`: Application port (default varies by framework)
- `LOG_LEVEL`: Logging verbosity (debug, info, warn, error)
- `NODE_ENV`: Environment mode (development, production)

### Agent-Specific Variables
- `AGENT_TYPE`: Type of agent being deployed (coordinator, processor, orchestrator, etc.)

### Framework-Specific Variables

#### Python/FastAPI
- `PYTHONDONTWRITEBYTECODE=1`: Prevent .pyc file creation
- `PYTHONUNBUFFERED=1`: Force stdout/stderr to be unbuffered

#### Node.js/Express
- `NODE_ENV`: Environment mode
- `NODE_OPTIONS`: Node.js runtime options

#### Go/Gin
- `GIN_MODE`: Gin framework mode (debug, release)
- `CGO_ENABLED`: CGO compilation setting

## Health Checks

All Dockerfiles include health checks that verify service availability:

- **Endpoint**: `/health`
- **Interval**: 30 seconds
- **Timeout**: 10 seconds
- **Start Period**: 5 seconds
- **Retries**: 3

Example health check response:
```json
{
  "status": "healthy",
  "timestamp": "2024-01-15T10:30:00Z",
  "version": "1.0.0",
  "uptime": 3600
}
```

## Best Practices

### Image Optimization
1. **Use specific tags**: Avoid `latest` tags in production
2. **Multi-stage builds**: Separate build dependencies from runtime
3. **Layer caching**: Order COPY commands by change frequency
4. **Dependency pinning**: Lock dependency versions for reproducibility

### Security Hardening
1. **Non-root users**: Never run containers as root
2. **Minimal images**: Use Alpine or distroless base images
3. **Dependency scanning**: Regularly scan for vulnerabilities
4. **Secret management**: Use external secret stores, not environment variables

### Production Deployment
1. **Resource limits**: Set memory and CPU constraints
2. **Health monitoring**: Implement comprehensive health checks
3. **Graceful shutdown**: Handle SIGTERM signals properly
4. **Logging**: Use structured logging with correlation IDs

## Troubleshooting

### Common Issues

#### Build Failures
```bash
# Check build context size
docker system df

# Build with no cache to debug layer issues
docker build --no-cache -f infrastructure/docker/services/Dockerfile.python .

# Build with verbose output
docker build --progress=plain -f infrastructure/docker/services/Dockerfile.python .
```

#### Runtime Issues
```bash
# Check container logs
docker logs <container-id>

# Execute into running container for debugging
docker exec -it <container-id> /bin/sh

# Check health status
docker inspect <container-id> | grep Health
```

#### Performance Issues
```bash
# Monitor resource usage
docker stats <container-id>

# Profile container startup
time docker run --rm my-service

# Check image layers
docker history my-service
```

## Development Workflow

### Local Development
1. Build images locally for testing
2. Use bind mounts for code changes during development
3. Override health check intervals for faster feedback

### CI/CD Integration
1. Build images in CI pipeline
2. Run security scans and tests
3. Push to container registry
4. Deploy with orchestration tools (Kubernetes, Docker Compose)

## Related Documentation

- [Architecture Specifications](../../docs/spec/architecture-multi-agent-microservices-specs.md)
- [Project Structure](../../src/main/README.md)
- [Deployment Guidelines](../kubernetes/README.md)
- [Security Policies](../../docs/security/README.md)
