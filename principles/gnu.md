---
id: cli-gnu-posix
priority: 3
scope: cli
tags:
  - principle
  - gnu
  - posix
  - standards
  - cli
applies_when: Building or editing command-line tools and scripts.
conflicts_with: []
decision_rule: Follow project CLI conventions first, then GNU/POSIX defaults.
---

# CLI GNU/POSIX Conventions

## Principle
CLI tools should follow standard UNIX/GNU behavior.

## Intent
Maximize predictability, scriptability, and pipeline compatibility.

## Apply when
- Creating or modifying CLI commands or shell scripts.

## Avoid when
- The project is not CLI-oriented.

## Do
- Provide `-h`/`--help` and `-v`/`--version` when relevant.
- Support short and long options.
- Write data output to `stdout`.
- Write errors and diagnostics to `stderr`.
- Return `0` on success and non-zero on failure.

## Do not
- Use non-standard flags for common behavior.
- Print errors to `stdout`.
- Ignore exit status from invoked commands.

## Conflict resolution
- Follow `principles/_priority-model.md`.

## Example (pseudo)
```pseudo
# bad
if error then print_stdout("failed")
exit(0)

# good
if error then print_stderr("failed")
exit(1)
```

## Checklist
- Are help/version flags available when appropriate?
- Are stdout/stderr separated correctly?
- Are exit codes meaningful?

