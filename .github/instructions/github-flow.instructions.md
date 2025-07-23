# GitHub Flow Instructions for GitHub Copilot

## Overview
These instructions guide developers and agents on working with Git according to GitHub Flow, ensuring efficient, collaborative, and high-quality software development. Tehse instructions are used whenever a developer requests Copilot to implement a GitHub Issue.

## Core Principles
- Work in small, incremental branches
- Collaborate through pull requests and code reviews
- Maintain a deployable main branch
- Automate testing and CI/CD for every change
- Communicate clearly and document all changes

## Coding Standards
### Language-Specific Guidelines
- Use clear, descriptive branch names (e.g., `feature/login-form`, `fix/typo-header`)
- Write concise, informative commit messages
- Organize commits logically (one purpose per commit)
- Handle merge conflicts promptly and transparently

### Architecture Patterns
- Keep the `main` branch always deployable
- Use feature branches for all work
- Merge via pull requests only after review and tests pass
- Delete branches after merging to keep the repo clean

## Code Generation Guidelines
### Function/Method Generation
- Document code changes in commit messages
- Validate code with automated tests before merging
- Specify related issues in pull request descriptions
- Handle errors and edge cases before merging

### Class/Component Generation
- Reference related branches and issues in documentation
- Organize code changes by feature or fix
- Include tests for new components or features

### API/Endpoint Generation
- Document API changes in pull requests
- Validate endpoints with automated tests
- Ensure backward compatibility when possible

## Quality Requirements
### Testing
- Name test files after the feature or fix (e.g., `login.test.js`)
- Structure tests to cover new and changed code
- Require passing tests before merging
- Use mocks/stubs for external dependencies

### Documentation
- Add inline comments for complex changes
- Update README and API docs for major changes
- Include code examples in pull requests

### Security
- Review code for security vulnerabilities before merging
- Validate input and handle sensitive data securely
- Use protected branches and required reviews for sensitive repos

## Integration Guidelines
### Dependencies
- Document new dependencies in pull requests
- Use semantic versioning for package updates
- Update configuration files as needed

### Database/Storage
- Document migration scripts in pull requests
- Test migrations before merging
- Maintain data model consistency

### External Services
- Document changes to API clients or integrations
- Handle errors and retries for external calls
- Log integration changes in pull requests

## Performance Optimization
- Review code for efficiency before merging
- Test performance-critical changes
- Use async/await for non-blocking operations
- Document caching strategies in pull requests

## Maintenance Guidelines
- Use code review checklists for every pull request
- Refactor legacy code in separate branches
- Track version history in Git
- Regularly prune old branches

## Examples
### Well-Structured Function
```markdown
- Create a branch: `git checkout -b feature/user-auth`
- Make changes and commit: `git commit -m "Add user authentication"`
- Push branch: `git push origin feature/user-auth`
- Open a pull request and request review
```

### Proper Error Handling
```markdown
- Resolve merge conflicts before merging
- Add tests for error cases
- Document fixes in commit messages
```

### Documentation Format
```markdown
- Use clear pull request descriptions
- Reference related issues (e.g., `Closes #42`)
- Include code examples and screenshots if relevant
```

### Test Pattern
```markdown
- Add or update tests for every feature or fix
- Require passing tests before merging
- Use CI/CD to automate test runs
```

## How to Use These Instructions
1. Enable GitHub Copilot in your development environment
2. Reference these instructions when working with Git and GitHub Flow
3. Use the patterns and examples as templates for branching, committing, and reviewing
4. Regularly update instructions based on team practices and project evolution
5. Share with team members for consistency
6. Before working on a new GitHub Issues, ensure you create a new feature branch and follow the naming conventions outlined in these instructions.

## Updating Instructions
- Review quarterly or when major changes occur
- Incorporate feedback from code reviews and team retrospectives
- Update examples with real project scenarios
- Maintain version history in git
