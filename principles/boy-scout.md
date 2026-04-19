---
tags:
  - principle
  - maintenance
  - refactoring
  - clean-code
---

# Boy Scout Rule

## Principle
Leave touched code slightly better than you found it.

## Intent
Improve local code quality continuously without expanding scope.

## Apply when
- You are already editing a file for a requested task.

## Do
- Fix nearby low-risk issues (typos, naming, dead imports/vars, minor formatting).
- Keep cleanup in the same area as the primary change.
- Keep cleanup small (target: <= 20% of changed lines).

## Do not
- Start broad refactors unrelated to the requested task.
- Change public behavior/API unless requested.
- Hide the main purpose of the change with cleanup noise.

## Trade-offs
- Small local improvements increase maintainability.
- Excess cleanup increases review risk and delivery time.

## Conflict resolution
- Priority order: Correctness > Existing project conventions > Pragmatic Consensus > Architecture rules > Boy Scout cleanup.
- If cleanup conflicts with delivery scope, reduce cleanup.

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
- Is every cleanup change local to edited code?
- Is behavior unchanged (unless requested)?
- Is cleanup volume small vs primary change?

