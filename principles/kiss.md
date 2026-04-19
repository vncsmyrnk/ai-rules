---
id: kiss
priority: 4
scope: global
tags:
  - principle
  - simplicity
  - design
applies_when: Choosing among multiple valid implementations.
conflicts_with:
  - clean-architecture
  - hexagonal
decision_rule: Prefer the simplest solution that meets current requirements safely.
---

# KISS (Keep It Simple, Stupid)

## Principle
Choose the simplest correct solution.

## Intent
Reduce cognitive load, defects, and maintenance cost.

## Apply when
- Multiple approaches can solve the same requirement.

## Avoid when
- Simplicity would compromise correctness, safety, or security.

## Do
- Prefer straightforward control flow and naming.
- Minimize abstraction layers.
- Choose readability over cleverness.

## Do not
- Add complexity to look “future-proof.”
- Use advanced patterns for simple problems.

## Conflict resolution
- Follow `principles/_priority-model.md`.

## Example (pseudo)
```pseudo
# bad
result = transform(map(filter(data, rule_a), rule_b), rule_c)

# good
for item in data:
  if rule_a(item):
    out = rule_b(item)
    result.add(rule_c(out))
```

## Checklist
- Is the solution understandable on first read?
- Is each abstraction necessary?
- Can complexity be reduced without losing correctness?
