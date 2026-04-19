---
id: separation-of-concerns
priority: 6
scope: architecture
tags:
  - principle
  - architecture
  - decoupling
  - modularity
applies_when: Code mixes unrelated responsibilities.
conflicts_with:
  - pragmatic-consensus
decision_rule: Separate concerns proportionally; avoid over-fragmenting trivial logic.
---

# Separation of Concerns (SoC)

## Principle
Each module or function should have one clear responsibility.

## Intent
Improve readability, testability, and change safety through clear boundaries.

## Apply when
- Writing new code.
- Modifying mixed-responsibility code.

## Avoid when
- Splitting trivial logic would reduce readability.

## Do
- Split multi-purpose functions into focused units.
- Separate presentation, domain, and data access.
- Hide internals behind clear module interfaces.

## Do not
- Mix rendering, persistence, and domain decisions in one unit.
- Create “god” modules handling unrelated concerns.

## Conflict resolution
- Follow `principles/_priority-model.md`.

## Example (pseudo)
```pseudo
# bad
handle_request() {
  view = render()
  db.save(view)
  apply_rules()
}

# good
handle_request() {
  result = apply_rules()
  repo.save(result)
  return presenter.render(result)
}
```

## Checklist
- Does each unit have one primary reason to change?
- Are presentation, domain, and data concerns separated?
- Are boundaries explicit and testable?

