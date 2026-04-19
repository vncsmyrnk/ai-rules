---
tags:
  - principle
  - architecture
  - decoupling
  - separation-of-concerns
---

# Clean Architecture

## Principle
Dependencies point inward toward domain policies.

## Intent
Protect business rules from framework, UI, and infrastructure churn.

## Apply when
- The system has medium/high domain complexity.
- Multiple interfaces/integrations are expected.

## Do
- Keep entities/use-cases framework-agnostic.
- Define boundaries/interfaces for external dependencies.
- Inject infrastructure implementations into application logic.

## Do not
- Import ORM/HTTP/framework objects into core use-cases.
- Shape domain models around framework constraints.
- Put transport/parsing logic inside domain code.

## Trade-offs
- Better testability and longevity.
- More upfront structure for simple apps.

## Conflict resolution
- For simple tasks, prefer Pragmatic Consensus over heavy layering.
- Priority order: Correctness > Existing project conventions > Pragmatic Consensus > Clean Architecture > Boy Scout cleanup.

## Example (pseudo)
```pseudo
# bad
use_case() {
  row = orm.query("...")
  return http_response(row)
}

# good
use_case(repo_port) {
  data = repo_port.load()
  return domain_result(data)
}
```

## Checklist
- Does core code avoid framework imports?
- Are external dependencies behind interfaces?
- Can use-cases run in isolation tests?

