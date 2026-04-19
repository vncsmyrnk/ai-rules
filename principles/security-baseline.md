---
id: security-baseline
priority: 2
scope: security
tags:
  - principle
  - security
  - safety
  - secure-by-default
applies_when: Handling input, auth, permissions, secrets, or external data.
conflicts_with:
  - kiss
  - yagni
decision_rule: Favor secure defaults even when they add moderate complexity.
---

# Security Baseline

## Principle
Default to secure behavior in design and implementation.

## Intent
Reduce exploitable risk and protect data/users.

## Apply when
- Processing external input.
- Accessing sensitive data.
- Managing authentication/authorization.

## Avoid when
- Never avoid; reduce scope instead.

## Do
- Validate and sanitize untrusted input.
- Enforce least privilege.
- Keep secrets out of source code and logs.
- Use safe defaults and explicit allowlists.

## Do not
- Trust client-side validation alone.
- Expose sensitive internal error details to users.
- Use broad permissions by default.

## Conflict resolution
- Follow `principles/_priority-model.md`.

## Example (pseudo)
```pseudo
# bad
if user.is_logged_in: allow_admin_action()

# good
if user.has_permission("admin:write"): allow_admin_action()
else: deny()
```

## Checklist
- Is untrusted input validated?
- Are permissions least-privilege?
- Are secrets and sensitive logs protected?
