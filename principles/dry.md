---
id: dry
priority: 6
scope: code-quality
tags:
  - principle
  - maintainability
  - duplication
applies_when: Similar logic appears in multiple places.
conflicts_with:
  - yagni
  - kiss
decision_rule: Remove proven duplication with minimal useful abstraction.
---

# DRY (Don't Repeat Yourself)

## Principle
Avoid duplicated knowledge and logic.

## Intent
Make changes safer by keeping behavior defined in one place.

## Apply when
- The same behavior is duplicated with the same change pattern.

## Avoid when
- Similarity is accidental or likely to diverge soon.

## Do
- Extract shared logic after duplication is clear.
- Keep abstractions small and concrete.
- Centralize domain rules, not just syntax.

## Do not
- Create generic frameworks for one duplicate.
- Force unrelated cases into one abstraction.

## Conflict resolution
- Follow `principles/_priority-model.md`.

## Example (pseudo)
```pseudo
# bad
total_a = price_a + price_a * tax
total_b = price_b + price_b * tax

# good
apply_tax(price, tax_rate) = price + price * tax_rate
total_a = apply_tax(price_a, tax)
total_b = apply_tax(price_b, tax)
```

## Checklist
- Is duplication semantic (knowledge), not just textual?
- Is extraction simple and justified?
- Will this reduce future change risk?
