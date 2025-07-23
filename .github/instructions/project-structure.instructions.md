---
appliesTo: '*'
description: These instructions are to ensure a consistent project structure no matter the type of project (single application, microservices, multi-agent, etc.).
---

# Project Structure Instructions

## Overview
These instructions provide a consistent project structure for all types of projects, including single applications, microservices, multi-agent systems, and more. The goal is to ensure that all projects follow a standardized structure that is easy to navigate and maintain.

## Project Structure
All projects should adhere to the following directory structure:

```
project-root/
├── .github/
│   └── workflows/
│   └── chatmodes/
│   └── instructions/
│   └── prompts/
├── src/
│   ├── main/
│   │   └── agents/ (optional)
│   │   |   └── agent-1/
│   │   └── services/ (optional)
│   │   |   └── service-1/
│   │   └── applications/ (optional)
│   │   |   └── application-1/
│   └── test/
│   │   └── agents/ (optional)
│   │   |   └── agent-1/
│   │   └── services/ (optional)
│   │   |   └── service-1/
│   │   └── applications/ (optional)
│   │   |   └── application-1/
├── docs/
├── scripts/
└── README.md
```

### Directory Descriptions
- **.github/**: Contains GitHub-specific files, such as workflows and issue templates.
- **src/**: Contains the source code for the project.
  - **main/**: Contains the main application code.
    - **agents/**: Contains code for individual agents (if applicable).
      - **agent-1/**: Contains code for agent 1.
    - **services/**: Contains code for services (if applicable).
      - **service-1/**: Contains code for service 1.
    - **applications/**: Contains code for applications (if applicable).
      - **application-1/**: Contains code for application 1.
  - **test/**: Contains the test code for the project.
    - **agents/**: Contains tests for individual agents (if applicable).
      - **agent-1/**: Contains tests for agent 1.
    - **services/**: Contains tests for services and APIs (if applicable).
      - **service-1/**: Contains tests for service 1.
    - **applications/**: Contains tests for applications (if applicable).
      - **application-1/**: Contains tests for application 1.
- **docs/**: Contains documentation files for the project.
- **scripts/**: Contains scripts for building, testing, and deploying the project.
- **README.md**: The main documentation file for the project, providing an overview and instructions for use.