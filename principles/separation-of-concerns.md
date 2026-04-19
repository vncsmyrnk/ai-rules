---
tags:
  - principle
  - architecture
  - decoupling
  - modularity
---

# Separation of Concerns (SoC)

## Principle
Each module/function should have one clear responsibility.

## Intent
Improve readability, testability, and change safety through clear boundaries.

## Apply when
- Writing new code.
- Modifying mixed-responsibility code.

## Do
- Split multi-purpose functions into focused units.
- Separate UI/presentation, business rules, and data access.
- Hide internals behind clear module interfaces.

## Do not
- Mix rendering, persistence, and domain decisions in one unit.
- Create “god” modules handling unrelated concerns.

## Trade-offs
- Better maintainability and easier testing.
- More files/functions to navigate.

## Conflict resolution
- If strict separation adds unjustified complexity, apply Pragmatic Consensus.
- Priority order: Correctness > Existing project conventions > Pragmatic Consensus > SoC > Boy Scout cleanup.

## Example (pseudo)
```pseudo
# bad
handle_request() {
  html = render()
  db.save(html)
  apply_discount_rules()
}

# good
handle_request() {
  result = apply_discount_rules()
  repo.save(result)
  return presenter.render(result)
}
```

## Checklist
- Does each unit have a single primary reason to change?
- Are UI, domain, and data concerns separated?
- Are boundaries explicit and testable?

