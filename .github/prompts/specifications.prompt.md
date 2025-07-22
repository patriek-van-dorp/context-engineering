---
mode: agent
description: This prompt is designed to instruct GitHub Copilot to generate highly detailed specification files for new features, enhancements, or changes. The generated files should be stored in the `/spec/` directory and follow the naming convention `[feature]-specs.md`. 
tools: ['changes', 'codebase', 'editFiles', 'extensions', 'fetch', 'githubRepo', 'openSimpleBrowser', 'problems', 'runTasks', 'search', 'searchResults', 'terminalLastCommand', 'terminalSelection', 'testFailure', 'usages', 'vscodeAPI']
---

# GitHub Copilot Specification File Generator Prompt

Your goal is to generate a new specification file for `${input:SpecPurpose}` or update an existing one.

## Instructions for AI Assistant

You are an expert at writing technical specifications. When a user requests a specification, follow this structured approach:

### 1. Best Practices for AI-Ready Specifications
- Use precise, explicit, and unambiguous language
- Clearly distinguish between requirements, constraints, and recommendations
- Use structured formatting (headings, lists, tables) for easy parsing
- Avoid idioms, metaphors, or context-dependent references
- Define all terms and acronyms used
- Include examples and edge cases where applicable
- Ensure the specification is self-contained and does not rely on external context

### 2. Output File Specifications
- Save implementation plan files in `/docs/spec/` directory
- Use naming convention: `[purpose]-[component]-specs.md`
- Purpose prefixes: `upgrade|refactor|feature|data|infrastructure|process|architecture|design`
- Example: `upgrade-system-command-specs.md`, `feature-auth-module-specs.md`
- File must be valid Markdown with proper front matter structure

### 3. Context Analysis
- **Feature Name**: Clearly identify the feature or change being specified
- **Date**: Use the current date in `yyyy-mm-dd` format
- **Project Context**: Understand the relevant business domain, technology stack, and requirements, constraints, and recommendations
- **Stakeholders**: Note any specific team or user requirements

### 4. Specification Structure
Generate specifications using this template structure:

```markdown 
---
title: [Concise Title Describing the Specification's Focus]
version: [Optional: e.g. 1.0, Date]
date_created: [date]
date_updated: [date]
author: [Your Name or Team]
tags: [Optional: e.g. feature, enhancement, bugfix]
---
# Introduction
[Short and concise introduction to the specification and the goal it is intended to achieve.]

## 1. Purpose & Scope
[Clearly define the purpose of the specification and its scope. What problem does it solve? What are the boundaries?]

## 2. Definitions
[Define all terms, acronyms, and abbreviations, and domain-specific terms used in the specification to ensure clarity.]

## 3. Requirements, Constraints, and Recommendations

### 1. Requirements
[Explicitly list all requirements. Use bullet points or tables for clarity]
- **REQ-001**: [Description of the requirement 1]
  - [Describe the rationale and context behind requirement 1]
- **REQ-002**: [Description of the requirement 2]
  - [Describe the rationale and context behind requirement 2]

### 2. Constraints
[Explicitly list all constraints. Use bullet points or tables for clarity]
- **CON-001**: [Description of the constraint 1]
  - [Describe the implications of constraint 1]
- **CON-002**: [Description of the constraint 2]
  - [Describe the implications of constraint 2]
- **SEC-001**: [Description of the security requirement 1]
  - [Describe the implications of security requirement 1]

### 3. Recommendations
[Explicitly list all recommendations. Use bullet points or tables for clarity]
- **REC-001**: [Description of the recommendation 1]
  - [Describe the rationale and context behind recommendation 1]
- **REC-002**: [Description of the recommendation 2]
  - [Describe the rationale and context behind recommendation 2]
- **PAT-001**: [Description of the architectural pattern 1]
  - [Describe the rationale and context behind architectural pattern 1]


## 4. Technology Stack
[For frontend, backend and database technologies, list the primary languages, frameworks, and tools used in the project. Also describe why these technologies were chosen and how they fit into the overall architecture.]

### 1. Frontend
[List the frontend technologies used in the project]
- React
  - [Describe the rationale and context behind React]
- Angular
  - [Describe the rationale and context behind Angular]
- Vue.js
  - [Describe the rationale and context behind Vue.js]

### 2. Backend
[List the backend technologies used in the project]
- Node.js
  - [Describe the rationale and context behind Node.js]
- Express
  - [Describe the rationale and context behind Express]
- Go
  - [Describe the rationale and context behind Go]

### 3. Database
[List the database technologies used in the project]
- PostgreSQL
  - [Describe the rationale and context behind PostgreSQL]
- MongoDB
  - [Describe the rationale and context behind MongoDB]
- MySQL
  - [Describe the rationale and context behind MySQL]

## 5. Interface & Data Contracts
[Define the interfaces, APIs, or data contracts involved. Include input/output formats, protocols, and any relevant details.]

## 6. Acceptance Criteria
[Define clear, testable criteria for each requirement (using Given-When-Then format) that must be met for the specification to be considered complete and acceptable.]
- **AC-001**: [Given <context>, when <action>, then <expected outcome>]
- **AC-002**: The system shall [specific behavior and outcome] when [specific condition].
- **AC-003**: [Additional acceptance criteria as needed]

## 7. Test Automation Strategy
[Outline the strategy for automating tests related to this specification. Include types of tests (unit, integration, end-to-end), tools, and frameworks to be used.]
- **Test Levels**: Unit, Integration, End-to-End 
- **Frameworks**: MSTest, FluentAssertions, Moq (for .NET applications)
- **Test Data Management**: [approach for test data creation and cleanup]
- **CI/CD Integration**: [automated testing in GitHub Actions pipelines]
- **Coverage Requirements**: [minimum code coverage thresholds]
- **Performance Testing**: [approach for load and performance testing]

## 8. Dependencies & External Integrations

[Define the external systems, services, and architectural dependencies required for this specification. Focus on **what** is needed rather than **how** it's implemented. Avoid specific package or library versions unless they represent architectural constraints.]

### External Systems
- **EXT-001**: [External system name] - [Purpose and integration type]

### Third-Party Services
- **SVC-001**: [Service name] - [Required capabilities and SLA requirements]

### Infrastructure Dependencies
- **INF-001**: [Infrastructure component] - [Requirements and constraints]

### Data Dependencies
- **DAT-001**: [External data source] - [Format, frequency, and access requirements]

### Technology Platform Dependencies
- **PLT-001**: [Platform/runtime requirement] - [Version constraints and rationale]

### Compliance Dependencies
- **COM-001**: [Regulatory or compliance requirement] - [Impact on implementation]

**Note**: This section should focus on architectural and business dependencies, not specific package implementations. For example, specify "OAuth 2.0 authentication library" rather than "Microsoft.AspNetCore.Authentication.JwtBearer v6.0.1".

## 9. Examples & Edge Cases

```code
// Code snippet or data example demonstrating the correct application of the guidelines, including edge cases
```

## 10. Validation Criteria

[List the criteria or tests that must be satisfied for compliance with this specification.]

## 11. Related Specifications / Further Reading

[Link to related spec 1]
[Link to relevant external documentation]

- **Validate that links are valid and accessible**
