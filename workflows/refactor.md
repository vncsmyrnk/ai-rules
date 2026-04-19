---
id: workflow-refactor
priority: 5
scope: workflow
tags:
  - workflow
  - refactor
  - maintainability
applies_when: Improving structure without changing behavior.
conflicts_with: []
decision_rule: Preserve behavior; optimize readability and maintainability incrementally.
---

# Refactor Workflow

## Goal
Improve code structure while preserving external behavior.

## When to use
- User requests refactor.
- Local cleanup needed around touched code.

## Inputs
- Current code and tests.
- Behavior constraints/public API expectations.

## Steps
1. Define non-negotiable behavior constraints.
2. Add missing safety tests if coverage is weak.
3. Refactor in small, reversible steps.
4. Keep APIs stable unless requested.
5. Validate after each meaningful step.
6. Summarize structural improvements.

## Output
- `Constraints`
- `Refactor steps`
- `Behavior verification`
- `Risk notes`

## Done criteria
- Is behavior preserved?
- Are changes incremental and reviewable?
- Is readability/maintainability improved?
