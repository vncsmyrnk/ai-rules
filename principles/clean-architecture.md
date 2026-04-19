---
id: clean-architecture
priority: 7
scope: architecture
tags:
  - principle
  - architecture
  - decoupling
  - separation-of-concerns
applies_when: Domain complexity is medium/high and multiple integrations exist.
conflicts_with:
  - pragmatic-consensus
  - kiss
  - yagni
decision_rule: Isolate business rules from frameworks when complexity justifies layering.
---

# Clean Architecture

## Principle
Dependencies point inward toward domain policies.

## Intent
Protect business rules from framework, UI, and infrastructure churn.

## Apply when
- Domain complexity is medium/high.
- Multiple interfaces/integrations are expected.

## Avoid when
- The feature is small and layering adds unnecessary complexity.

## Do
- Keep entities/use-cases framework-agnostic.
- Define boundaries/interfaces for external dependencies.
- Inject infrastructure implementations into application logic.

## Do not
- Import ORM/HTTP/framework objects into core use-cases.
- Shape domain models around framework constraints.
- Put transport/parsing logic inside domain code.

## Conflict resolution
- Follow `principles/_priority-model.md`.

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

