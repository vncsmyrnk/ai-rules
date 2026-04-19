---
id: error-handling
priority: 2
scope: reliability
tags:
  - principle
  - reliability
  - robustness
  - errors
applies_when: Writing logic that can fail due to input, I/O, network, or state.
conflicts_with:
  - kiss
decision_rule: Handle failures explicitly and fail fast with actionable context.
---

# Error Handling

## Principle
Handle errors explicitly; do not fail silently.

## Intent
Improve reliability, debuggability, and user trust.

## Apply when
- Any operation can fail (I/O, parsing, network, validation).

## Avoid when
- A hard crash is the explicit and documented strategy for unrecoverable startup failure.

## Do
- Validate inputs early.
- Return/raise explicit errors with context.
- Preserve original error cause when wrapping.
- Fail fast on invalid state.

## Do not
- Swallow exceptions/errors.
- Return ambiguous failure values without context.
- Continue after critical invariants break.

## Conflict resolution
- Follow `principles/_priority-model.md`.

## Example (pseudo)
```pseudo
# bad
parse(config)

# good
result = parse(config)
if result.is_error:
  return error("invalid config", cause=result.error)
```

## Checklist
- Are failure paths explicit?
- Is error context actionable?
- Are critical invariants enforced?
