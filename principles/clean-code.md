---
id: clean-code
priority: 6
scope: code-quality
tags:
  - principle
  - readability
  - maintainability
  - clean-code
applies_when: Writing or refactoring production code.
conflicts_with:
  - pragmatic-consensus
decision_rule: Optimize for readability and maintainability without needless abstraction.
---

# Clean Code

## Principle
Code should be easy to read, reason about, and change.

## Intent
Improve long-term maintainability and reduce defects.

## Apply when
- Writing new code.
- Refactoring touched code.

## Avoid when
- Large style-only rewrites would delay critical fixes.

## Do
- Use clear, intention-revealing names.
- Keep functions small and focused.
- Replace magic values with named constants.
- Prefer explicit over implicit behavior.

## Do not
- Use ambiguous names (`data`, `tmp`, `misc`) without context.
- Mix unrelated responsibilities in one function.
- Hide side effects.

## Conflict resolution
- Follow `principles/_priority-model.md`.

## Example (pseudo)
```pseudo
# bad
fn(x, y) { return x * 0.2 + y }

# good
apply_discount(price, fee) {
  DISCOUNT_RATE = 0.2
  return price * DISCOUNT_RATE + fee
}
```

## Checklist
- Are names intention-revealing?
- Are functions focused and short?
- Are constants explicit and meaningful?
