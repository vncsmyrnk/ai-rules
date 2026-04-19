---
tags:
  - principle
  - pragmatism
  - teamwork
  - decision-making
  - simplicity
---

# Pragmatic Consensus

## Principle
Choose the simplest solution the team can understand, maintain, and ship.

## Intent
Balance ideal architecture with delivery speed and team clarity.

## Apply when
- Multiple valid implementations exist.
- A “pure” approach increases complexity without clear value.

## Do
- Prefer simple, idiomatic solutions.
- Reuse standard library/platform features before adding dependencies.
- Follow established project conventions.
- Document key trade-offs briefly when needed.

## Do not
- Introduce complex patterns without clear need.
- Optimize prematurely.
- Refactor readable working code only for theoretical purity.

## Trade-offs
- Faster delivery and lower cognitive load.
- May defer “ideal” architecture until justified.

## Conflict resolution
- Tie-breaker rule for architecture disagreements.
- Priority order: Correctness > Existing project conventions > Pragmatic Consensus > Other design principles > Boy Scout cleanup.

## Example (pseudo)
```pseudo
# bad
add_factory_builder_strategy_for_simple_parse()

# good
parse_config_with_standard_library()
```

## Checklist
- Is the solution understandable by the current team?
- Is added abstraction justified by current complexity?
- Did we avoid unnecessary dependencies/patterns?

