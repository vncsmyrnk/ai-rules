---
id: workflow-test-authoring
priority: 3
scope: workflow
tags:
  - workflow
  - testing
  - quality
applies_when: Adding or updating tests for new/fixed behavior.
conflicts_with: []
decision_rule: Prefer deterministic, behavior-focused tests with clear failure signals.
---

# Test Authoring Workflow

## Goal
Create reliable tests that validate behavior and prevent regressions.

## When to use
- New feature added.
- Bug fixed.
- Critical refactor performed.

## Inputs
- Behavior requirements.
- Existing test patterns and tooling.

## Steps
1. Define observable behavior and edge cases.
2. Select proper test level (unit/integration/e2e).
3. Write deterministic tests (control time/network/randomness).
4. Cover at least one failure mode per critical path.
5. Keep tests focused and readable.
6. Run and verify stability.

## Output
- `Behavior covered`
- `Test cases added/updated`
- `Gaps or deferred coverage`

## Done criteria
- Do tests assert behavior (not internals)?
- Are tests deterministic and stable?
- Are critical failure modes covered?
