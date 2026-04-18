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

## Example

**Bad (Coupled):**
```python
# The business logic is tied directly to the database layer
from app.database import get_db_connection

def create_user(username, email):
    db = get_db_connection()
    db.execute("INSERT INTO users (username, email) VALUES (?, ?)", (username, email))
```

**Good (Clean):**
```python
# The core logic relies on an interface, not an implementation
class UserRepository:
    def save(self, user):
        pass

def create_user(username: str, email: str, repo: UserRepository):
    user = User(username, email)
    repo.save(user)

# The database adapter is implemented elsewhere and injected at runtime
```