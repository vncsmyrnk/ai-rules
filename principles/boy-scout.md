---
id: boy-scout
priority: 8
scope: code-quality
tags:
  - principle
  - maintenance
  - refactoring
  - clean-code
applies_when: A file is already being modified for another task.
conflicts_with:
  - yagni
  - pragmatic-consensus
decision_rule: Apply only small local cleanups that do not expand scope.
---

# Boy Scout Rule

## Principle
Leave touched code slightly better than you found it.

## Intent
Improve local code quality continuously without expanding scope.

## Apply when
- You are already editing a file for a requested task.

## Avoid when
- Cleanup would materially expand task scope or risk.

## Do
- Fix nearby low-risk issues (typos, naming, dead imports/vars, minor formatting).
- Keep cleanup in the same area as the primary change.
- Keep cleanup small (target: <= 20% of changed lines).

## Do not
- Start broad refactors unrelated to the requested task.
- Change public behavior/API unless requested.
- Hide the main purpose of the change with cleanup noise.

## Conflict resolution
- Follow `principles/_priority-model.md`.

## Example (pseudo)
```pseudo
# bad
change(feature)
rename_many_unrelated_modules()

# good
change(feature)
fix_local_typo()
remove_local_unused_import()
```

## Checklist
- Is cleanup local to edited code?
- Is behavior unchanged unless requested?
- Is cleanup small vs primary change?

