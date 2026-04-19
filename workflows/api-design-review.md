---
id: workflow-api-design-review
priority: 3
scope: workflow
tags:
  - workflow
  - api
  - rest
  - design-review
applies_when: Designing or reviewing HTTP API contracts.
conflicts_with: []
decision_rule: Enforce correct HTTP semantics, compatibility, and clear error contracts.
---

# API Design Review Workflow

## Goal
Validate API design for correctness, consistency, and client usability.

## When to use
- New API endpoint design.
- API contract changes.

## Inputs
- Endpoint specs.
- Request/response examples.
- Versioning and compatibility constraints.

## Steps
1. Validate resource naming and URI structure.
2. Validate HTTP method semantics and idempotency.
3. Validate status codes and error model consistency.
4. Validate pagination/filter/sort conventions.
5. Validate auth, rate-limit, and caching expectations.
6. Validate backward compatibility and deprecation path.

## Output
- `Contract issues`
- `Compatibility risks`
- `Recommended changes`

## Done criteria
- Are semantics and status codes correct?
- Are errors standardized and actionable?
- Are compatibility and versioning clear?
