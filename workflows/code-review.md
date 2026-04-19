---
id: workflow-code-review
priority: 3
scope: workflow
tags:
  - workflow
  - code-review
  - quality
applies_when: Reviewing a patch, PR, or file changes.
conflicts_with: []
decision_rule: Prioritize correctness, security, and clear actionable feedback.
---

# Code Review Workflow

## Goal
Produce a clear, prioritized review that improves correctness, security, and maintainability.

## When to use
- User asks for a code review.
- Reviewing changed files before merge.

## Inputs
- Changed files or diff.
- Task intent (feature, bug fix, refactor).
- Relevant language/principle rules.

## Steps
1. Gather context: intent, scope, language, architecture constraints.
2. Check high-risk issues first: correctness, security, data loss, race conditions.
3. Check design fit: SoC, simplicity, existing conventions.
4. Check code quality: readability, error handling, test coverage, edge cases.
5. Classify findings by severity.
6. Provide concrete fixes for critical issues.

## Severity model
- **Critical**: correctness/security/data-loss issues; must fix.
- **Major**: reliability/maintainability risks; should fix.
- **Minor**: style/readability improvements.

## Output
Use this structure:
- `Summary`
- `Critical`
- `Major`
- `Minor`
- `Positives`

Each item should include: `location`, `issue`, `recommendation`.

## Done criteria
- Are critical risks identified first?
- Is feedback specific and actionable?
- Is output concise and structured?

