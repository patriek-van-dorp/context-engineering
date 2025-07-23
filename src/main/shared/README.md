# Shared Components Directory

This directory contains shared libraries, contracts, models, and utilities that are used across multiple agents and services in the multi-agent microservices architecture.

## Structure

```
shared/
├── libraries/           # Shared libraries by programming language
│   ├── python/         # Python shared libraries
│   ├── nodejs/         # Node.js shared libraries
│   └── go/             # Go shared libraries
├── contracts/          # API contracts and schemas
│   ├── openapi/        # OpenAPI/Swagger specifications
│   └── protobuf/       # Protocol Buffer definitions
├── models/             # Shared data models and schemas
└── utils/              # Common utilities and helper functions
```

## Purpose

The shared directory promotes code reuse, consistency, and maintainability across the multi-agent system by providing:

- **Common Libraries**: Shared functionality that multiple services need
- **API Contracts**: Standardized interface definitions
- **Data Models**: Consistent data structures across services
- **Utilities**: Helper functions and common tools

## Guidelines

### Code Organization
- Group shared code by programming language to avoid dependencies
- Keep shared components focused and minimal
- Version shared libraries to manage breaking changes
- Document all shared APIs and utilities

### Dependencies
- Shared libraries should have minimal external dependencies
- Use dependency injection for configuration
- Avoid tight coupling between shared components

### Versioning
- Use semantic versioning for shared libraries
- Maintain backward compatibility when possible
- Provide migration guides for breaking changes

## Libraries

### Python Libraries (`libraries/python/`)
Common Python packages that can be shared across Python-based agents and services:
- Authentication utilities
- Database connection helpers
- Logging configurations
- Common middleware
- Validation utilities

### Node.js Libraries (`libraries/nodejs/`)
Common Node.js modules that can be shared across Node.js-based agents and services:
- Express middleware
- Database ORM configurations
- Authentication utilities
- Logging and monitoring
- Error handling

### Go Libraries (`libraries/go/`)
Common Go packages that can be shared across Go-based agents and services:
- HTTP middleware
- Database utilities
- Authentication helpers
- Logging packages
- Configuration management

## Contracts

### OpenAPI Specifications (`contracts/openapi/`)
- Agent API contracts
- Service API contracts
- Shared API models
- Authentication schemas
- Error response formats

### Protocol Buffers (`contracts/protobuf/`)
- gRPC service definitions
- Message schemas for high-performance communication
- Data serialization formats

## Models

### Shared Data Models (`models/`)
- User models
- Authentication models
- Configuration models
- Common business entities
- Error models

## Utilities

### Common Utilities (`utils/`)
- Configuration parsers
- Health check utilities
- Metrics collection
- Tracing utilities
- Test helpers

## Usage Guidelines

### Including Shared Components

#### Python
```python
from shared.libraries.python.auth import AuthUtils
from shared.libraries.python.db import DatabaseHelper
```

#### Node.js
```javascript
const { AuthUtils } = require('shared/libraries/nodejs/auth');
const { DatabaseHelper } = require('shared/libraries/nodejs/db');
```

#### Go
```go
import (
    "context-engineering/src/main/shared/libraries/go/auth"
    "context-engineering/src/main/shared/libraries/go/db"
)
```

### API Contract Usage
- Reference OpenAPI specs when implementing APIs
- Use protobuf definitions for gRPC services
- Validate requests/responses against contracts

### Contributing to Shared Components

1. Ensure the component is truly reusable across multiple services
2. Write comprehensive tests
3. Document the API and usage examples
4. Follow the established coding standards
5. Version the component appropriately
6. Update this README with new components

## Testing

Shared components should have:
- Unit tests with high coverage (minimum 90%)
- Integration tests where applicable
- Documentation tests/examples
- Performance tests for critical utilities

## Documentation

Each shared component should include:
- Clear API documentation
- Usage examples
- Performance characteristics
- Dependencies and requirements
- Version history and migration guides
