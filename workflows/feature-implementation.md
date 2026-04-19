---
id: workflow-feature-implementation
priority: 4
scope: workflow
tags:
  - workflow
  - implementation
  - feature
applies_when: Implementing a new feature or endpoint.
conflicts_with: []
decision_rule: Deliver minimal correct scope first, then polish within limits.
---

# Feature Implementation Workflow

## Goal
Implement requested behavior with minimal scope, clear design, and validation.

## When to use
- User requests a new feature.

## Inputs
- Feature request and constraints.
- Existing code conventions.
- Relevant principles/language rules.

## Steps
1. Restate requirements and acceptance criteria.
2. Locate target modules and integration points.
3. Implement smallest viable solution (KISS + YAGNI).
4. Add/adjust tests for expected behavior.
5. Run validations (tests/lint/format).
6. Summarize changes and trade-offs.

## Output
- `Plan`
- `Implemented changes`
- `Validation results`
- `Follow-ups (if any)`

## Done criteria
- Does implementation satisfy explicit requirements?
- Are tests updated for new behavior?
- Is scope limited to requested functionality?
