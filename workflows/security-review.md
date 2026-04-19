---
id: workflow-security-review
priority: 2
scope: workflow
tags:
  - workflow
  - security
  - review
applies_when: Reviewing code paths handling input, auth, secrets, or sensitive data.
conflicts_with: []
decision_rule: Prioritize risk reduction and secure defaults over convenience.
---

# Security Review Workflow

## Goal
Identify and reduce security risks before merge or release.

## When to use
- Features touching auth, permissions, input parsing, secrets, or data storage.

## Inputs
- Changed code and threat-relevant context.
- Data flow and trust boundaries.

## Steps
1. Map trust boundaries and untrusted inputs.
2. Check validation/sanitization and output encoding.
3. Check authentication and authorization decisions.
4. Check secret handling and sensitive logging.
5. Check dependency and configuration exposure risks.
6. Classify risks by severity and propose mitigations.

## Output
- `High risk`
- `Medium risk`
- `Low risk`
- `Mitigations`

## Done criteria
- Are high-risk issues explicit and actionable?
- Are trust boundaries and controls documented?
- Are secure defaults enforced?
