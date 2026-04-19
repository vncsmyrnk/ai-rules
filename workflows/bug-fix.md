---
id: workflow-bug-fix
priority: 3
scope: workflow
tags:
  - workflow
  - bug-fix
  - regression
applies_when: Fixing incorrect or failing behavior.
conflicts_with: []
decision_rule: Reproduce first, fix root cause, add regression protection.
---

# Bug Fix Workflow

## Goal
Resolve a defect safely and prevent recurrence.

## When to use
- User reports a bug or failing behavior.

## Inputs
- Expected vs actual behavior.
- Reproduction steps or failing test.
- Logs/errors/context.

## Steps
1. Reproduce issue (or create failing test).
2. Isolate root cause in minimal code path.
3. Implement smallest correct fix.
4. Add regression test for the bug.
5. Run tests and impacted validations.
6. Document root cause and fix scope.

## Output
- `Reproduction`
- `Root cause`
- `Fix`
- `Regression test`
- `Validation`

## Done criteria
- Is bug reproducible before and fixed after?
- Is there a regression test?
- Did fix avoid unrelated refactors?
