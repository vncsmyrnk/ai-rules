---
id: priority-model
priority: 0
scope: global
tags:
  - principle
  - governance
  - conflict-resolution
applies_when: Multiple principles suggest different actions.
conflicts_with: []
decision_rule: Apply the highest-priority principle first; de-scope lower-priority rules if needed.
---

# Principle Priority Model

## Principle
When principles conflict, follow a single global priority order.

## Intent
Make agent decisions deterministic and reduce contradictory outputs.

## Apply when
- Two or more principles lead to different implementation choices.

## Avoid when
- No conflict exists.

## Do
- Resolve decisions using this order (highest first):
  1. Correctness and Safety
  2. Security
  3. Existing project conventions and constraints
  4. KISS and Pragmatic Consensus
  5. YAGNI
  6. Separation of Concerns
  7. Clean/Hexagonal Architecture
  8. Boy Scout cleanup
- Record trade-off briefly when skipping a lower-priority rule.

## Do not
- Apply lower-priority principles if they reduce correctness, safety, or security.
- Mix conflicting choices in the same change.

## Conflict resolution
- This file is the source of truth for rule precedence.

## Example (pseudo)
```pseudo
if architecture_rule_conflicts_with_kiss_for_small_feature:
  choose(kiss)
```

## Checklist
- Was priority order used explicitly?
- Was lower-priority scope reduced when needed?
- Is the final choice consistent and explainable?
