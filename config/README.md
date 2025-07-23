# Configuration Directory

This directory contains environment-specific configuration files for the multi-agent microservices architecture.

## Structure

```
config/
├── development/        # Development environment configurations
├── staging/           # Staging environment configurations
└── production/        # Production environment configurations
```

## Overview

Configuration management follows the principle of environment-specific settings while maintaining a consistent structure across all environments. Each environment directory contains:

- Application configurations
- Database connection settings
- External service configurations
- Feature flags and toggles
- Security and authentication settings

## Environment Configurations

### Development (`development/`)
Local development environment settings:
- Local database connections
- Debug logging enabled
- Development-specific API endpoints
- Relaxed security settings for testing
- Mock external services

### Staging (`staging/`)
Pre-production testing environment:
- Staging database connections
- Production-like configurations
- Limited external service access
- Enhanced logging and monitoring
- Security settings similar to production

### Production (`production/`)
Production environment settings:
- Production database connections
- Optimized performance settings
- Full external service integrations
- Comprehensive security configurations
- Production monitoring and alerting

## Configuration Files

### Common Configuration Files

#### `.env`
Environment variables for each environment:
```bash
# Application
NODE_ENV=development
APP_PORT=3000
APP_HOST=localhost

# Database
DATABASE_URL=postgresql://localhost:5432/context_engineering_dev
REDIS_URL=redis://localhost:6379/0

# Authentication
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24h

# External Services
API_GATEWAY_URL=http://localhost:8080
NOTIFICATION_SERVICE_URL=http://localhost:3001

# Monitoring
LOG_LEVEL=debug
METRICS_ENABLED=true
```

#### `database.yml`
Database configuration:
```yaml
development:
  adapter: postgresql
  database: context_engineering_dev
  username: postgres
  password: password
  host: localhost
  port: 5432
  pool: 5

staging:
  adapter: postgresql
  database: context_engineering_staging
  username: <%= ENV['DB_USERNAME'] %>
  password: <%= ENV['DB_PASSWORD'] %>
  host: <%= ENV['DB_HOST'] %>
  port: 5432
  pool: 10

production:
  adapter: postgresql
  database: context_engineering_prod
  username: <%= ENV['DB_USERNAME'] %>
  password: <%= ENV['DB_PASSWORD'] %>
  host: <%= ENV['DB_HOST'] %>
  port: 5432
  pool: 20
  ssl: true
```

#### `services.yml`
Service discovery and endpoints:
```yaml
services:
  api_gateway:
    development: "http://localhost:8080"
    staging: "https://api-staging.context-engineering.com"
    production: "https://api.context-engineering.com"
  
  auth_service:
    development: "http://localhost:3001"
    staging: "https://auth-staging.context-engineering.com"
    production: "https://auth.context-engineering.com"
  
  config_service:
    development: "http://localhost:3002"
    staging: "https://config-staging.context-engineering.com"
    production: "https://config.context-engineering.com"
```

#### `features.yml`
Feature flags and toggles:
```yaml
features:
  new_authentication_flow:
    development: true
    staging: true
    production: false
  
  enhanced_monitoring:
    development: true
    staging: true
    production: true
  
  experimental_ai_features:
    development: true
    staging: false
    production: false
```

## Configuration Management

### Environment Variable Precedence
1. System environment variables (highest priority)
2. Environment-specific `.env` file
3. Default configuration values
4. Application defaults (lowest priority)

### Secrets Management
- **Development**: Local `.env` files (not committed to git)
- **Staging/Production**: External secret management (AWS Secrets Manager, Azure Key Vault, etc.)
- **Kubernetes**: Kubernetes Secrets
- **Docker**: Docker Secrets

### Configuration Validation
All configurations should be validated at application startup:
- Required variables are present
- Data types are correct
- Ranges and formats are valid
- External dependencies are accessible

## Security Considerations

### Sensitive Data
- Never commit secrets to version control
- Use placeholder values in committed files
- Implement proper access controls
- Rotate secrets regularly

### Configuration Security
- Encrypt sensitive configuration files
- Use secure channels for configuration delivery
- Implement configuration drift detection
- Audit configuration changes

## Loading Configurations

### Python Example
```python
import os
from pathlib import Path
import yaml

def load_config(environment='development'):
    config_dir = Path(__file__).parent / 'config' / environment
    
    # Load environment variables
    env_file = config_dir / '.env'
    if env_file.exists():
        from dotenv import load_dotenv
        load_dotenv(env_file)
    
    # Load YAML configurations
    config = {}
    for config_file in config_dir.glob('*.yml'):
        with open(config_file) as f:
            config[config_file.stem] = yaml.safe_load(f)
    
    return config
```

### Node.js Example
```javascript
const path = require('path');
const fs = require('fs');
const yaml = require('js-yaml');
require('dotenv').config({ 
    path: path.join(__dirname, 'config', process.env.NODE_ENV || 'development', '.env')
});

function loadConfig(environment = 'development') {
    const configDir = path.join(__dirname, 'config', environment);
    const config = {};
    
    // Load YAML configurations
    const configFiles = fs.readdirSync(configDir).filter(file => file.endsWith('.yml'));
    configFiles.forEach(file => {
        const content = fs.readFileSync(path.join(configDir, file), 'utf8');
        const key = path.basename(file, '.yml');
        config[key] = yaml.load(content);
    });
    
    return config;
}
```

### Go Example
```go
package config

import (
    "os"
    "path/filepath"
    "gopkg.in/yaml.v2"
)

type Config struct {
    Database DatabaseConfig `yaml:"database"`
    Services ServicesConfig `yaml:"services"`
    Features FeaturesConfig `yaml:"features"`
}

func LoadConfig(environment string) (*Config, error) {
    configDir := filepath.Join("config", environment)
    
    // Load environment variables
    if err := loadEnvFile(filepath.Join(configDir, ".env")); err != nil {
        return nil, err
    }
    
    // Load YAML configurations
    config := &Config{}
    configFiles := []string{"database.yml", "services.yml", "features.yml"}
    
    for _, file := range configFiles {
        if err := loadYAMLFile(filepath.Join(configDir, file), config); err != nil {
            return nil, err
        }
    }
    
    return config, nil
}
```

## Configuration Templates

### Creating New Environment
1. Copy configuration template from `development/`
2. Update environment-specific values
3. Validate configuration syntax
4. Test with target environment
5. Document environment-specific settings

### Configuration Drift Detection
```bash
#!/bin/bash
# Script to detect configuration drift

ENVIRONMENTS=("development" "staging" "production")
REFERENCE_ENV="development"

for env in "${ENVIRONMENTS[@]}"; do
    if [ "$env" != "$REFERENCE_ENV" ]; then
        echo "Comparing $REFERENCE_ENV with $env..."
        diff -u "config/$REFERENCE_ENV/" "config/$env/" || true
    fi
done
```

## Best Practices

### Configuration Design
- Use consistent naming conventions
- Group related configurations
- Provide sensible defaults
- Document all configuration options
- Validate configurations at startup

### Environment Management
- Minimize differences between environments
- Use infrastructure as code for consistency
- Implement automated configuration testing
- Monitor configuration changes

### Documentation
- Document all configuration options
- Provide examples and default values
- Explain environment-specific differences
- Maintain change logs

## Troubleshooting

### Common Issues
- Missing environment variables
- Incorrect file paths
- YAML syntax errors
- Permission denied errors
- Network connectivity issues

### Debugging
- Enable verbose logging for configuration loading
- Validate YAML syntax before deployment
- Check file permissions and accessibility
- Test configuration loading in isolation

## Integration with Deployment

### Kubernetes ConfigMaps
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
data:
  database.yml: |
    production:
      adapter: postgresql
      host: {{ .Values.database.host }}
      port: {{ .Values.database.port }}
      database: {{ .Values.database.name }}
```

### Docker Environment
```dockerfile
# Copy configuration files
COPY config/production/ /app/config/production/

# Set environment
ENV NODE_ENV=production
ENV CONFIG_PATH=/app/config/production
```

### Terraform Variables
```hcl
variable "environment" {
  description = "Environment name"
  type        = string
  default     = "development"
}

variable "database_config" {
  description = "Database configuration"
  type = object({
    host     = string
    port     = number
    database = string
  })
}
```
