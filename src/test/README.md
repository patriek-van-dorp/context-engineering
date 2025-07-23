# Test Directory

This directory contains all test suites for the multi-agent microservices architecture, organized to mirror the main source code structure.

## Structure

```
test/
├── agents/              # Tests for agent microservices
│   ├── cognitive-agent/     # Tests for cognitive agent
│   ├── decision-agent/      # Tests for decision agent
│   └── monitoring-agent/    # Tests for monitoring agent
└── services/            # Tests for supporting services
    ├── api-gateway/         # Tests for API gateway
    ├── auth-service/        # Tests for authentication service
    └── config-service/      # Tests for configuration service
```

## Test Strategy

The testing strategy follows a multi-level approach as defined in the architecture specifications:

### Test Levels
1. **Unit Tests**: Test individual functions, methods, and classes
2. **Integration Tests**: Test interactions between components
3. **End-to-End Tests**: Test complete workflows across services
4. **Performance Tests**: Test system performance and scalability

### Coverage Requirements
- **Minimum 80% code coverage** for all services (as per architecture specs)
- **90% coverage** for shared components (higher standard for reusable code)
- **100% coverage** for critical security and authentication functions

## Testing Frameworks

### Python (FastAPI)
- **Framework**: Pytest
- **Coverage**: pytest-cov
- **Mocking**: pytest-mock, unittest.mock
- **API Testing**: httpx (async), requests-mock

```python
# Example test structure
tests/
├── conftest.py          # Shared fixtures
├── unit/
│   ├── test_models.py
│   └── test_utils.py
├── integration/
│   ├── test_api.py
│   └── test_database.py
└── e2e/
    └── test_workflows.py
```

### Node.js (Express)
- **Framework**: Jest
- **Coverage**: Built-in Jest coverage
- **Mocking**: Jest mocks, sinon
- **API Testing**: supertest

```javascript
// Example test structure
tests/
├── setup.js             // Test setup
├── unit/
│   ├── models.test.js
│   └── utils.test.js
├── integration/
│   ├── api.test.js
│   └── database.test.js
└── e2e/
    └── workflows.test.js
```

### Go (Gin)
- **Framework**: Go test (built-in)
- **Coverage**: go test -cover
- **Mocking**: testify/mock, GoMock
- **API Testing**: httptest

```go
// Example test structure
tests/
├── main_test.go         // Test setup
├── unit/
│   ├── models_test.go
│   └── utils_test.go
├── integration/
│   ├── api_test.go
│   └── database_test.go
└── e2e/
    └── workflows_test.go
```

## Test Data Management

### Containerized Databases
- Use containerized databases for test isolation
- Docker Compose for local test environments
- Test containers for integration tests

### Test Data Strategy
- **Unit Tests**: Mock data, in-memory databases
- **Integration Tests**: Dedicated test databases
- **E2E Tests**: Full test environment with seed data

### Data Cleanup
- Clean up test data after each test
- Use transactions that can be rolled back
- Implement test database reset scripts

## CI/CD Integration

### GitHub Actions Pipeline
Tests are automatically run in the CI/CD pipeline:

1. **Pre-commit**: Linting and basic syntax checks
2. **Pull Request**: Full test suite execution
3. **Merge to Main**: Integration and E2E tests
4. **Release**: Performance and load tests

### Test Environments
- **Development**: Local testing with Docker Compose
- **CI**: GitHub Actions with containerized services
- **Staging**: Full environment testing before production

## Performance Testing

### Tools
- **Locust** (Python-based): Load testing and user simulation
- **k6** (JavaScript): Performance testing and monitoring

### Test Types
- **Load Testing**: Normal expected load
- **Stress Testing**: Beyond normal capacity
- **Spike Testing**: Sudden load increases
- **Volume Testing**: Large amounts of data

### Metrics
- Response time percentiles (P50, P95, P99)
- Throughput (requests per second)
- Error rates
- Resource utilization

## Test Organization Guidelines

### File Naming
- Test files should end with `_test.py`, `.test.js`, or `_test.go`
- Mirror the source code directory structure
- Use descriptive test names

### Test Structure
```python
# Example Python test structure
class TestCognitiveAgent:
    def test_should_process_valid_request(self):
        # Arrange
        # Act
        # Assert
        
    def test_should_handle_invalid_input(self):
        # Arrange
        # Act
        # Assert
```

### Fixtures and Mocks
- Create reusable test fixtures
- Mock external dependencies
- Use factory patterns for test data

## Running Tests

### Local Development
```bash
# Python
pytest src/test/agents/cognitive-agent/ -v --cov

# Node.js
npm test -- --coverage

# Go
go test ./src/test/agents/cognitive-agent/... -v -cover
```

### Docker Environment
```bash
# Run all tests in containers
docker-compose -f docker-compose.test.yml up --abort-on-container-exit

# Run specific test suite
docker-compose run test-runner pytest src/test/agents/
```

### CI/CD
Tests are automatically triggered by:
- Pull request creation/updates
- Commits to main branch
- Release tag creation

## Test Documentation

Each test directory should contain:
- `README.md` - Test-specific documentation
- `TEST_PLAN.md` - Detailed test scenarios and cases
- Test data documentation
- Performance test results and baselines

## Best Practices

1. **Write tests first** (Test-Driven Development)
2. **Keep tests independent** and isolated
3. **Use descriptive test names** that explain the scenario
4. **Mock external dependencies** to ensure test reliability
5. **Test both happy path and error scenarios**
6. **Maintain test data integrity** between test runs
7. **Regular test maintenance** and updates
8. **Document complex test scenarios**
