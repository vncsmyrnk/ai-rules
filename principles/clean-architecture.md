---
tags:
  - principle
  - architecture
  - decoupling
  - separation-of-concerns
---

# Clean Architecture

## Principle
Clean Architecture emphasizes the separation of concerns by dividing software into layers. The fundamental rule is the **Dependency Rule**: source code dependencies must point only inward, toward higher-level policies. 

The inner layers contain the business rules (Entities, Use Cases) and must not know anything about the outer layers (UI, Database, Frameworks).

## Instructions for AI Agents
When designing or modifying systems, ensure that business logic is completely isolated from delivery mechanisms and infrastructure.

**Do:**
- Define core business logic in pure, framework-agnostic classes/functions.
- Use Dependency Injection to provide infrastructure implementations to the core domain.
- Create explicit interfaces (boundaries) for how the core communicates with the outside world.

**Do Not:**
- Import database ORM models directly into business logic (Use Cases).
- Put HTTP request parsing logic inside domain entities.
- Let framework-specific constraints dictate the shape of the domain model.
