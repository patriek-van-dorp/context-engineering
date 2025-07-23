# Agents Directory

This directory contains all agent microservices that are part of the multi-agent system. Each agent is an autonomous microservice responsible for specific tasks or decision-making.

## Structure

```
agents/
├── agent-template/        # Template for creating new agents
├── cognitive-agent/       # Agent responsible for cognitive tasks and reasoning
├── decision-agent/        # Agent responsible for decision-making processes
└── monitoring-agent/      # Agent responsible for system monitoring and health
```

## Agent Guidelines

### Requirements
- Each agent must be implemented as an independent microservice (REQ-001)
- All inter-agent communication must occur via APIs (REQ-002)
- All agents must be containerized for deployment (REQ-004)
- Agents must not share state directly; use APIs or message queues (CON-001)
- All agents must be stateless or manage state externally (CON-002)

### Recommendations
- Use OpenAPI/Swagger for API contract definition (REC-001)
- Follow Domain-Driven Design for service boundaries (PAT-001)
- Implement health checks and monitoring endpoints
- Use structured logging for observability

## Technology Stack

### Supported Languages
- **Python** with FastAPI framework
- **Node.js** with Express framework  
- **Go** with Gin framework

### Database Options
- **PostgreSQL** for relational data and transactional integrity
- **MongoDB** for flexible, document-based storage
- **Redis** for caching and message brokering

## Creating a New Agent

1. Copy the `agent-template/` directory
2. Rename it to your new agent name
3. Update the configuration files
4. Implement the agent logic according to the template
5. Add corresponding test directory in `src/test/agents/`
6. Add Dockerfile in `infrastructure/docker/agents/`
7. Add Kubernetes manifests in `infrastructure/kubernetes/`

## API Standards

- All APIs must follow RESTful conventions
- API endpoints must be versioned (e.g., `/api/v1/`)
- Input/output formats: JSON (default), gRPC for performance-critical scenarios
- All API communication must be authenticated and encrypted (SEC-001)

## Testing

Each agent should have comprehensive tests in the corresponding `src/test/agents/[agent-name]/` directory:
- Unit tests (minimum 80% coverage)
- Integration tests
- End-to-end tests
- Performance tests

## Documentation

Each agent directory should contain:
- `README.md` - Agent-specific documentation
- `API.md` - API documentation and examples
- `DEPLOYMENT.md` - Deployment instructions
- OpenAPI/Swagger specifications in `src/main/shared/contracts/openapi/`
