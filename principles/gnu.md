---
tags:
  - principle
  - gnu
  - posix
  - standards
  - cli
---

# CLI GNU/POSIX Conventions

## Principle
CLI tools should follow standard UNIX/GNU behavior.

## Intent
Maximize predictability, scriptability, and pipeline compatibility.

## Apply when
- Creating or modifying CLI commands or shell scripts.

## Do
- Provide `-h`/`--help` and `-v`/`--version` when relevant.
- Support short and long options.
- Write data output to `stdout`.
- Write errors/diagnostics to `stderr`.
- Return `0` on success and non-zero on failure.

## Do not
- Use non-standard flags for common behavior.
- Print errors to `stdout`.
- Ignore exit status from invoked commands.

## Trade-offs
- Standard behavior improves interoperability.
- Strict conventions may add small implementation overhead.

## Conflict resolution
- If project CLI conventions exist, follow them first.
- Priority order: Correctness > Existing project conventions > GNU/POSIX conventions > Boy Scout cleanup.

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
- Are help/version flags available where appropriate?
- Are stdout/stderr separated correctly?
- Are exit codes meaningful?

