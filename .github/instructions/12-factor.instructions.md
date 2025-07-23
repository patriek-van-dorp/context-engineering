---
appliesTo: '*'
description: This instruction is to ensure that all applications and services designed and build follow the 12-factor methodology.
---

# 12-Factor Methodology Instructions

## Overview
The 12-factor methodology is a set of best practices for building software-as-a-service applications that are portable, scalable, and maintainable. It is designed to help developers create applications that can be easily deployed and managed in a cloud environment.

## 12-Factor Principles
1. **Codebase**: One codebase tracked in revision control, many deploys.
   - Use a single codebase for the application, which can be deployed to multiple environments (development, staging, production).
2. **Dependencies**: Explicitly declare and isolate dependencies.
   - Use a dependency management tool to declare and manage application dependencies.
3. **Config**: Store configuration in the environment.
   - Use environment variables to store configuration settings, allowing for different configurations in different environments without changing the codebase.
4. **Backing Services**: Treat backing services as attached resources.
   - Backing services (databases, message queues, etc.) should be treated as resources that can be attached and detached from the application.
5. **Build, Release, Run**: Strictly separate build and run stages.
   - The build stage is responsible for compiling the application, the release stage is responsible for packaging the application, and the run stage is responsible for executing the application.
6. **Processes**: Execute the app as one or more stateless processes.
   - Applications should be designed to run as stateless processes, allowing for easy scaling and recovery.
7. **Port Binding**: Export services via port binding.
   - Applications should expose their services on a port, allowing them to be accessed by other services or clients.
8. **Concurrency**: Scale out via the process model.
   - Applications should be designed to scale out by running multiple instances of the application processes.
   - This can be achieved by using a process manager or orchestration tool to manage the application instances.
9. **Disposability**: Maximize robustness with fast startup and graceful shutdown.
   - Applications should be designed to start quickly and shut down gracefully, allowing for easy deployment and recovery.
10. **Dev/Prod Parity**: Keep development, staging, and production as similar as possible.
    - Ensure that the development, staging, and production environments are as similar as possible to avoid issues when deploying code.
11. **Logs**: Treat logs as event streams.
    - Applications should generate logs as event streams, which can be collected and processed by a log management system.
12. **Admin Processes**: Run admin/management tasks as one-off processes.
    - Administrative tasks (database migrations, backups, etc.) should be run as one-off processes that can be executed independently of the application.

## Implementation Guidelines
- Ensure that all applications and services follow the 12-factor principles. 
- Use tools and frameworks that support the 12-factor methodology.
- Regularly review and refactor code to ensure compliance with the 12-factor principles.
- Document any deviations from the 12-factor principles and provide justifications for them.

## Reference
For more information on the 12-factor methodology, refer to the [12-Factor App](https://12factor.net/) website.