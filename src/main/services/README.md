# Services Directory

This directory contains all supporting microservices that provide shared functionality for the multi-agent system. These services are not agents themselves but provide essential capabilities that agents and other services can use.

## Structure

```
services/
├── service-template/      # Template for creating new services
├── api-gateway/          # API Gateway for routing and security
├── auth-service/         # Authentication and authorization service
├── config-service/       # Configuration management service
└── notification-service/ # Notification and messaging service
```

## Service Guidelines

### Requirements
- Supporting services must be implemented as microservices with APIs (REQ-003)
- All inter-service communication must occur via APIs (REQ-002)
- All services must be containerized for deployment (REQ-004)
- Services must not share state directly; use APIs or message queues (CON-001)
- All services must be stateless or manage state externally (CON-002)

### Recommendations
- Use OpenAPI/Swagger for API contract definition (REC-001)
- Follow Domain-Driven Design for service boundaries (PAT-001)
- Implement health checks and monitoring endpoints
- Use structured logging for observability
- Employ service mesh for observability and traffic management (REC-002)

## Technology Stack

### Supported Languages
- **Python** with FastAPI framework
- **Node.js** with Express framework  
- **Go** with Gin framework

### Database Options
- **PostgreSQL** for relational data and transactional integrity
- **MongoDB** for flexible, document-based storage
- **Redis** for caching and message brokering

## Service Descriptions

### API Gateway
- **Purpose**: Central entry point for all API requests
- **Responsibilities**: Request routing, load balancing, authentication, rate limiting
- **Technology**: Kong, NGINX, or cloud-native solutions

### Auth Service
- **Purpose**: Centralized authentication and authorization
- **Responsibilities**: User authentication, JWT token management, role-based access control
- **Technology**: OAuth2, OpenID Connect

### Config Service
- **Purpose**: Centralized configuration management
- **Responsibilities**: Environment-specific configurations, feature flags, secrets management
- **Technology**: Consul, etcd, or cloud-native configuration services

### Notification Service
- **Purpose**: Centralized notification and messaging
- **Responsibilities**: Email, SMS, push notifications, webhook delivery
- **Technology**: Message queues (RabbitMQ, Apache Kafka)

## Creating a New Service

1. Copy the `service-template/` directory
2. Rename it to your new service name
3. Update the configuration files
4. Implement the service logic according to the template
5. Add corresponding test directory in `src/test/services/`
6. Add Dockerfile in `infrastructure/docker/services/`
7. Add Kubernetes manifests in `infrastructure/kubernetes/`

## API Standards

- All APIs must follow RESTful conventions
- API endpoints must be versioned (e.g., `/api/v1/`)
- Input/output formats: JSON (default), gRPC for performance-critical scenarios
- All API communication must be authenticated and encrypted (SEC-001)
- Services must be discoverable via service discovery mechanism

## Testing

Each service should have comprehensive tests in the corresponding `src/test/services/[service-name]/` directory:
- Unit tests (minimum 80% coverage)
- Integration tests
- End-to-end tests
- Performance tests
- Contract tests

## Documentation

Each service directory should contain:
- `README.md` - Service-specific documentation
- `API.md` - API documentation and examples
- `DEPLOYMENT.md` - Deployment instructions
- OpenAPI/Swagger specifications in `src/main/shared/contracts/openapi/`
