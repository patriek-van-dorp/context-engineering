# Context Engineering

A multi-agent microservices architecture for scalable, containerized agent systems.

## 🚀 Quick Start

This project is currently under development. The base directory structure has been implemented in [PR #187](https://github.com/patriek-van-dorp/context-engineering/pull/187).

## 📁 Project Structure

Following the project structure instructions, all source code is organized under the `src/` directory:

```
context-engineering/
├── .github/                     # GitHub workflows and templates
├── docs/                        # Documentation
├── src/                         # All source code
│   ├── main/
│   │   ├── agents/              # Agent microservices
│   │   ├── services/            # Supporting microservices
│   │   └── shared/              # Shared components
│   └── test/                    # Test structure mirroring main
├── infrastructure/              # Infrastructure as Code
├── scripts/                     # Automation scripts
├── config/                      # Environment configurations
└── tools/                       # Development tools
```

## 🛠️ Development Workflow

This project follows GitHub Flow as specified in the [GitHub Flow Instructions](.github/instructions/github-flow.instructions.md):

1. **Create a feature branch**: `git checkout -b feature/your-feature-name`
2. **Make your changes** and commit with clear messages
3. **Push the branch**: `git push origin feature/your-feature-name`
4. **Open a Pull Request** for review
5. **Merge after review** and tests pass
6. **Delete the feature branch** to keep the repo clean

## 📋 Current Status

- ✅ **Base Directory Structure** (Issue #165) - [PR #187](https://github.com/patriek-van-dorp/context-engineering/pull/187)
- 🔄 **Technology Stack Initialization** (Issue #166) - Pending
- 🔄 **API Contract & Documentation** (Issue #167) - Pending
- 🔄 **Containerization & Orchestration** (Issue #168) - Pending

## 🏗️ Architecture

This project implements a multi-agent microservices architecture with:

- **Independent Agents**: Each agent is a separate microservice
- **API Communication**: All inter-service communication via REST/gRPC APIs
- **Containerization**: All services containerized with Docker
- **Orchestration**: Kubernetes for container management
- **Multi-language Support**: Python, Node.js, and Go frameworks

## 📖 Documentation

- [Architecture Specifications](docs/spec/architecture-multi-agent-microservices-specs.md)
- [Implementation Plan](docs/plan/architecture-multi-agent-microservices-plan.md)
- [Project Structure Instructions](.github/instructions/project-structure.instructions.md)
- [GitHub Flow Instructions](.github/instructions/github-flow.instructions.md)

## 🤝 Contributing

Please follow the GitHub Flow instructions and ensure:

1. All work is done in feature branches
2. Pull requests are created for all changes
3. Code follows the project structure guidelines
4. Tests are included for new functionality
5. Documentation is updated as needed

## 📄 License

[Add your license information here]