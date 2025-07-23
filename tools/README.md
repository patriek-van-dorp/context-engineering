# Tools Directory

This directory contains development tools and utilities for the multi-agent microservices architecture.

## Structure

```
tools/
├── generators/         # Code generation tools
└── validators/         # Validation and verification tools
```

## Overview

The tools directory provides automation and utilities to improve developer productivity and maintain code quality:

- **Generators**: Automated code generation for common patterns
- **Validators**: Code validation, linting, and verification tools

## Generators (`generators/`)

### Code Generation Tools
- `generate-agent.sh`: Generate new agent from template
- `generate-service.sh`: Generate new service from template
- `generate-api-contract.sh`: Generate API contracts and documentation
- `generate-tests.sh`: Generate test scaffolding

### Template-Based Generation
- Agent templates with boilerplate code
- Service templates with standard patterns
- API contract templates
- Test suite templates
- Docker configuration templates

### Usage Examples
```bash
# Generate new agent
./tools/generators/generate-agent.sh --name analytics-agent --language python

# Generate new service
./tools/generators/generate-service.sh --name metrics-service --language nodejs

# Generate API contract
./tools/generators/generate-api-contract.sh --service auth-service --version v1
```

### Agent Generator (`generate-agent.sh`)
Creates a new agent with:
- Basic project structure
- Language-specific templates (Python/FastAPI, Node.js/Express, Go/Gin)
- Dockerfile and Kubernetes manifests
- Test scaffolding
- API contract stub
- Documentation templates

Example usage:
```bash
./tools/generators/generate-agent.sh \
  --name recommendation-agent \
  --language python \
  --description "Agent for generating recommendations" \
  --port 3003
```

Generated structure:
```
src/main/agents/recommendation-agent/
├── README.md
├── Dockerfile
├── requirements.txt (Python) / package.json (Node.js) / go.mod (Go)
├── src/
│   ├── main.py / index.js / main.go
│   ├── api/
│   ├── models/
│   └── utils/
├── tests/
└── docs/
```

### Service Generator (`generate-service.sh`)
Creates a new service with:
- Service-specific templates
- Database integration templates
- Authentication middleware
- Health check endpoints
- Monitoring and logging setup

### API Contract Generator (`generate-api-contract.sh`)
Generates:
- OpenAPI/Swagger specifications
- Protocol Buffer definitions
- Client SDK stubs
- Documentation

## Validators (`validators/`)

### Code Quality Tools
- `validate-code-quality.sh`: Run all code quality checks
- `validate-api-contracts.sh`: Validate API contract consistency
- `validate-configs.sh`: Validate configuration files
- `validate-security.sh`: Security vulnerability scanning

### Linting and Formatting
- `lint-python.sh`: Python code linting (flake8, black, isort)
- `lint-javascript.sh`: JavaScript/TypeScript linting (ESLint, Prettier)
- `lint-go.sh`: Go code linting (golangci-lint, gofmt)
- `lint-docker.sh`: Dockerfile linting (hadolint)
- `lint-yaml.sh`: YAML file validation

### Security Validation
- `security-scan.sh`: Comprehensive security scanning
- `dependency-check.sh`: Check for vulnerable dependencies
- `secrets-scan.sh`: Scan for accidentally committed secrets
- `container-scan.sh`: Container security scanning

### Usage Examples
```bash
# Run all code quality checks
./tools/validators/validate-code-quality.sh

# Validate specific language
./tools/validators/lint-python.sh src/main/agents/cognitive-agent/

# Security scan
./tools/validators/security-scan.sh --full

# Validate configurations
./tools/validators/validate-configs.sh --env production
```

## Tool Configuration

### Linting Configurations

#### Python (`.flake8`, `pyproject.toml`)
```ini
# .flake8
[flake8]
max-line-length = 88
extend-ignore = E203, W503
exclude = .git,__pycache__,dist,build,.venv

# pyproject.toml
[tool.black]
line-length = 88
target-version = ['py39']

[tool.isort]
profile = "black"
line_length = 88
```

#### JavaScript/TypeScript (`.eslintrc.js`, `.prettierrc`)
```javascript
// .eslintrc.js
module.exports = {
  extends: [
    'eslint:recommended',
    '@typescript-eslint/recommended',
    'prettier'
  ],
  parser: '@typescript-eslint/parser',
  plugins: ['@typescript-eslint'],
  rules: {
    'no-console': 'warn',
    '@typescript-eslint/no-unused-vars': 'error'
  }
};

// .prettierrc
{
  "semi": true,
  "trailingComma": "es5",
  "singleQuote": true,
  "printWidth": 80,
  "tabWidth": 2
}
```

#### Go (`.golangci.yml`)
```yaml
run:
  timeout: 5m
  
linters:
  enable:
    - gofmt
    - golint
    - govet
    - ineffassign
    - misspell
    - goconst
    - gocyclo
    
linters-settings:
  gocyclo:
    min-complexity: 10
```

### Security Scanning Configuration

#### Container Scanning (`.hadolint.yaml`)
```yaml
ignored:
  - DL3008
  - DL3009
  
failure-threshold: error

override:
  error:
    - DL3001
    - DL3002
  warning:
    - DL3003
```

## Development Workflow Integration

### Pre-commit Hooks
```bash
#!/bin/sh
# .git/hooks/pre-commit

echo "Running code quality checks..."
./tools/validators/validate-code-quality.sh --fast

if [ $? -ne 0 ]; then
    echo "Code quality checks failed. Please fix issues before committing."
    exit 1
fi
```

### CI/CD Integration
Tools are integrated with GitHub Actions:
```yaml
# .github/workflows/code-quality.yml
name: Code Quality

on: [push, pull_request]

jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Run code quality checks
        run: ./tools/validators/validate-code-quality.sh
      
      - name: Run security scan
        run: ./tools/validators/security-scan.sh
```

## Custom Tool Development

### Tool Template
```bash
#!/bin/bash
# Template for new tools

set -euo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"

# Tool metadata
TOOL_VERSION="1.0.0"
TOOL_DESCRIPTION="Description of what this tool does"

usage() {
    cat << EOF
Usage: $SCRIPT_NAME [OPTIONS] [ARGUMENTS]

$TOOL_DESCRIPTION

OPTIONS:
    -h, --help      Show this help message
    -v, --version   Show tool version
    --verbose       Enable verbose output
    --dry-run       Show what would be done

EXAMPLES:
    $SCRIPT_NAME --verbose
    $SCRIPT_NAME --dry-run

EOF
}

main() {
    # Tool logic here
    echo "Tool execution completed successfully"
}

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help)
            usage
            exit 0
            ;;
        -v|--version)
            echo "$SCRIPT_NAME version $TOOL_VERSION"
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            usage
            exit 1
            ;;
    esac
    shift
done

main "$@"
```

### Tool Development Guidelines
1. Follow consistent naming conventions
2. Provide comprehensive help documentation
3. Implement proper error handling
4. Support dry-run mode
5. Include version information
6. Write tests for complex tools
7. Document tool dependencies

## IDE Integration

### VS Code Settings
```json
{
    "python.linting.enabled": true,
    "python.linting.flake8Enabled": true,
    "python.formatting.provider": "black",
    "eslint.enable": true,
    "prettier.enable": true,
    "go.useLanguageServer": true,
    "go.lintTool": "golangci-lint"
}
```

### Extensions Recommendations
- Python: Python, Pylance, Black Formatter
- JavaScript/TypeScript: ESLint, Prettier, TypeScript
- Go: Go, Go Outliner
- Docker: Docker, hadolint
- YAML: YAML, Red Hat YAML

## Performance Optimization

### Parallel Execution
```bash
# Run multiple validators in parallel
./tools/validators/lint-python.sh &
./tools/validators/lint-javascript.sh &
./tools/validators/lint-go.sh &
wait

echo "All linting completed"
```

### Caching
- Cache dependency installations
- Cache linting results for unchanged files
- Use incremental validation where possible

## Troubleshooting

### Common Issues
- Missing tool dependencies
- Configuration file syntax errors
- Permission denied on script execution
- Path resolution issues

### Debugging Tools
- Enable verbose output with `--verbose`
- Use dry-run mode to test without changes
- Check tool logs for detailed error information
- Validate configuration files independently

## Tool Maintenance

### Updates
- Regular tool version updates
- Configuration file maintenance
- Performance optimization
- Security updates

### Documentation
- Keep tool documentation current
- Update usage examples
- Document new features and options
- Maintain troubleshooting guides

## Contributing

### Adding New Tools
1. Create tool in appropriate subdirectory
2. Follow naming conventions
3. Implement standard options (help, version, verbose)
4. Add comprehensive documentation
5. Write tests if applicable
6. Update this README

### Tool Standards
- Use bash for shell scripts
- Implement proper error handling
- Provide meaningful error messages
- Support common command-line options
- Include usage examples
