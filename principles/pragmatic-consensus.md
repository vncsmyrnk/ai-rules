---
id: pragmatic-consensus
priority: 4
scope: global
tags:
  - principle
  - pragmatism
  - teamwork
  - decision-making
  - simplicity
applies_when: Multiple valid solutions exist and complexity trade-offs are unclear.
conflicts_with:
  - clean-architecture
  - hexagonal
  - separation-of-concerns
decision_rule: Prefer the simplest maintainable option aligned with project conventions.
---

# Pragmatic Consensus

## Principle
Choose the simplest solution the team can understand, maintain, and ship.

## Intent
Balance ideal design with delivery speed and team clarity.

## Apply when
- Multiple valid implementations exist.
- A “pure” approach adds complexity without clear value.

## Avoid when
- Simplicity would reduce correctness or security.

## Do
- Prefer simple, idiomatic solutions.
- Reuse standard library/platform features before adding dependencies.
- Follow established project conventions.
- Document key trade-offs briefly.

## Do not
- Introduce complex patterns without clear need.
- Optimize prematurely.
- Refactor readable working code only for theoretical purity.

## Conflict resolution
- Follow `principles/_priority-model.md`.

## Example (pseudo)
```pseudo
# bad
add_factory_builder_strategy_for_simple_parse()

# good
parse_config_with_standard_library()
```

## Checklist
- Is the solution easy for the team to maintain?
- Is extra abstraction clearly justified?
- Did we avoid unnecessary dependencies?

