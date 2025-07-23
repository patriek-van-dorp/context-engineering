# Scripts Directory

This directory contains automation scripts for building, testing, deploying, and monitoring the multi-agent microservices architecture.

## Structure

```
scripts/
├── setup/              # Environment setup and initialization scripts
├── build/              # Build automation scripts
├── deploy/             # Deployment automation scripts
└── monitoring/         # Monitoring and maintenance scripts
```

## Overview

These scripts provide automation for common development and operations tasks:
- **Setup**: Environment initialization and dependency installation
- **Build**: Code compilation, testing, and packaging
- **Deploy**: Application deployment and configuration
- **Monitoring**: System monitoring, health checks, and maintenance

## Setup Scripts (`setup/`)

### Environment Initialization
- `setup-dev-environment.sh`: Complete development environment setup
- `install-dependencies.sh`: Install required tools and dependencies
- `init-database.sh`: Initialize databases with schemas and seed data
- `setup-kubernetes.sh`: Set up local Kubernetes environment

### Configuration
- `generate-configs.sh`: Generate environment-specific configurations
- `setup-secrets.sh`: Create and configure secrets and certificates
- `init-monitoring.sh`: Set up monitoring and logging infrastructure

### Usage Examples
```bash
# Complete development setup
./scripts/setup/setup-dev-environment.sh

# Install dependencies only
./scripts/setup/install-dependencies.sh

# Initialize databases
./scripts/setup/init-database.sh --env development
```

## Build Scripts (`build/`)

### Compilation and Packaging
- `build-all.sh`: Build all agents and services
- `build-agent.sh`: Build specific agent
- `build-service.sh`: Build specific service
- `build-docker-images.sh`: Build all Docker images

### Testing
- `run-tests.sh`: Execute complete test suite
- `run-unit-tests.sh`: Execute unit tests only
- `run-integration-tests.sh`: Execute integration tests
- `run-performance-tests.sh`: Execute performance tests

### Quality Assurance
- `lint-code.sh`: Run code linting and formatting
- `security-scan.sh`: Run security vulnerability scans
- `generate-docs.sh`: Generate API documentation

### Usage Examples
```bash
# Build everything
./scripts/build/build-all.sh

# Build specific agent
./scripts/build/build-agent.sh cognitive-agent

# Run full test suite
./scripts/build/run-tests.sh --coverage

# Build Docker images
./scripts/build/build-docker-images.sh --push
```

## Deploy Scripts (`deploy/`)

### Environment Deployment
- `deploy-dev.sh`: Deploy to development environment
- `deploy-staging.sh`: Deploy to staging environment
- `deploy-prod.sh`: Deploy to production environment
- `rollback.sh`: Rollback to previous version

### Infrastructure Management
- `provision-infrastructure.sh`: Provision cloud infrastructure
- `update-infrastructure.sh`: Update infrastructure configurations
- `destroy-infrastructure.sh`: Clean up infrastructure resources

### Database Management
- `migrate-database.sh`: Run database migrations
- `backup-database.sh`: Create database backups
- `restore-database.sh`: Restore database from backup

### Usage Examples
```bash
# Deploy to development
./scripts/deploy/deploy-dev.sh --version v1.2.3

# Deploy specific service to staging
./scripts/deploy/deploy-staging.sh --service auth-service

# Run database migrations
./scripts/deploy/migrate-database.sh --env production

# Rollback deployment
./scripts/deploy/rollback.sh --env production --version v1.2.2
```

## Monitoring Scripts (`monitoring/`)

### Health Checks
- `health-check.sh`: Comprehensive system health check
- `service-status.sh`: Check status of specific services
- `database-health.sh`: Check database connectivity and performance

### Metrics and Alerting
- `collect-metrics.sh`: Collect system and application metrics
- `generate-reports.sh`: Generate performance and health reports
- `alert-check.sh`: Check and trigger alerts

### Maintenance
- `cleanup-logs.sh`: Clean up old log files
- `cleanup-images.sh`: Clean up unused Docker images
- `backup-system.sh`: Create system backups
- `update-certificates.sh`: Update SSL/TLS certificates

### Usage Examples
```bash
# Full health check
./scripts/monitoring/health-check.sh --detailed

# Check specific service
./scripts/monitoring/service-status.sh auth-service

# Generate weekly report
./scripts/monitoring/generate-reports.sh --period weekly

# Clean up old resources
./scripts/monitoring/cleanup-logs.sh --days 30
```

## Script Standards

### Shell Script Guidelines
- Use bash for portability
- Include proper error handling
- Provide help/usage information
- Log all operations with timestamps
- Support dry-run mode where applicable

### Script Template
```bash
#!/bin/bash
set -euo pipefail  # Exit on error, undefined vars, pipe failures

# Script metadata
SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"

# Default values
DRY_RUN=false
VERBOSE=false

# Functions
usage() {
    cat << EOF
Usage: $SCRIPT_NAME [OPTIONS]

Description of what this script does.

OPTIONS:
    -h, --help      Show this help message
    -v, --verbose   Enable verbose output
    -n, --dry-run   Show what would be done without executing
    
EXAMPLES:
    $SCRIPT_NAME --verbose
    $SCRIPT_NAME --dry-run

EOF
}

log() {
    echo "[$(date +'%Y-%m-%d %H:%M:%S')] $*"
}

# Main script logic here
```

### Error Handling
- Always use `set -euo pipefail`
- Implement proper cleanup on exit
- Provide meaningful error messages
- Log errors with context

### Configuration
- Use environment variables for configuration
- Provide sensible defaults
- Support configuration files
- Document all configuration options

## Environment Variables

### Common Variables
```bash
# Environment
export ENVIRONMENT=development
export PROJECT_NAME=context-engineering

# Kubernetes
export KUBECONFIG=~/.kube/config
export KUBE_NAMESPACE=default

# Docker
export DOCKER_REGISTRY=your-registry.com
export DOCKER_TAG=latest

# Database
export DATABASE_URL=postgresql://localhost:5432/context_engineering
export REDIS_URL=redis://localhost:6379

# Monitoring
export MONITORING_ENABLED=true
export LOG_LEVEL=info
```

### Loading Configuration
```bash
# Load environment-specific config
if [ -f "$PROJECT_ROOT/config/$ENVIRONMENT/.env" ]; then
    source "$PROJECT_ROOT/config/$ENVIRONMENT/.env"
fi
```

## CI/CD Integration

### GitHub Actions
Scripts are integrated with GitHub Actions workflows:
- `build-all.sh` in build workflow
- `run-tests.sh` in test workflow
- `deploy-dev.sh` in deployment workflow

### Hooks
- Pre-commit hooks for code quality
- Pre-push hooks for testing
- Post-deploy hooks for verification

## Security Considerations

### Secrets Management
- Never hardcode secrets in scripts
- Use environment variables or secret managers
- Implement proper access controls
- Audit script access and usage

### Script Security
- Validate all inputs
- Use absolute paths
- Avoid shell injection vulnerabilities
- Implement proper file permissions

## Cross-Platform Support

### Platform Detection
```bash
# Detect operating system
case "$(uname -s)" in
    Linux*)     MACHINE=Linux;;
    Darwin*)    MACHINE=Mac;;
    CYGWIN*)    MACHINE=Cygwin;;
    MINGW*)     MACHINE=MinGw;;
    *)          MACHINE="UNKNOWN:${unameOut}"
esac
```

### Tool Compatibility
- Use portable commands when possible
- Provide platform-specific alternatives
- Test on multiple platforms
- Document platform requirements

## Troubleshooting

### Common Issues
- Permission denied errors
- Missing dependencies
- Configuration file not found
- Network connectivity issues

### Debugging
- Use `bash -x script.sh` for debugging
- Add verbose logging options
- Implement debug mode
- Provide troubleshooting guides

## Documentation

Each script should include:
- Purpose and functionality description
- Usage examples and options
- Prerequisites and dependencies
- Troubleshooting information
- Change log and version history
