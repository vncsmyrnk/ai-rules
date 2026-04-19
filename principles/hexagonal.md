---
tags:
  - principle
  - architecture
  - ports-and-adapters
  - decoupling
  - testing
---

# Hexagonal Architecture (Ports and Adapters)

## Principle
Hexagonal Architecture, or Ports and Adapters, aims to create loosely coupled application components that can be easily connected to their software environment. It makes the application agnostic to the nature of the clients interacting with it (UI, test runners, other apps) and the infrastructure it uses (databases, message queues).

- **Core Application:** Contains the domain logic.
- **Ports:** Interfaces defined by the Core that dictate how it can be controlled (Driving/Primary Ports) and how it can control external dependencies (Driven/Secondary Ports).
- **Adapters:** Implementations that connect the outside world to the Ports.

## Instructions for AI Agents
Structure the codebase so that the domain logic does not depend on technical implementation details.

**Do:**
- Define interfaces (Ports) inside the domain layer for any external service (e.g., Database, Email Service).
- Write adapters outside the domain layer that implement these interfaces.
- Ensure the domain layer only imports domain code, never external libraries or adapters.
- Use this architecture to make the core logic highly testable in isolation.

**Do Not:**
- Hardcode database connection logic within service classes.
- Leak HTTP request objects (like Express `req` or Django `request`) into the domain layer.
