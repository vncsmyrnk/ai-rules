---
id: testing-strategy
priority: 3
scope: quality
tags:
  - principle
  - testing
  - quality
  - reliability
applies_when: Adding or changing behavior in production code.
conflicts_with:
  - boy-scout
decision_rule: Test behavior at the right level with deterministic, maintainable tests.
---

# Testing Strategy

## Principle
Test behavior, not implementation details.

## Intent
Catch regressions with stable, high-signal tests.

## Apply when
- Adding new behavior.
- Fixing defects.
- Refactoring critical logic.

## Avoid when
- Temporary exploratory scripts not part of production code.

## Do
- Prefer deterministic tests.
- Test core logic close to domain boundaries.
- Add regression tests for bug fixes.
- Keep tests readable and focused.

## Do not
- Over-mock internal details.
- Assert private implementation steps.
- Add flaky timing/network dependencies without control.

## Conflict resolution
- Follow `principles/_priority-model.md`.

## Example (pseudo)
```pseudo
# bad
test_calls_internal_method_three_times()

# good
test_discount_rule_returns_expected_total()
```

## Checklist
- Does each test verify observable behavior?
- Are tests deterministic and isolated?
- Is there a regression test for each fixed bug?
