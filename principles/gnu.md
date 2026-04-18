---
tags:
  - principle
  - gnu
  - posix
  - standards
  - cli
---

# GNU Coding Standards & Philosophy

## Principle
The GNU project advocates for software freedom, portability, and standardized user interfaces. When writing command-line tools or system software, adhering to GNU/POSIX standards ensures that tools behave predictably and integrate seamlessly into the Unix-like ecosystem.

## Instructions for AI Agents
When generating or modifying command-line interfaces (CLI) and scripts, enforce standard UNIX/GNU conventions.

**Do:**
- Support standard flags: `--help` (or `-h`) for usage and `--version` (or `-v`) for version info.
- Accept long-form options (`--option`) and short-form options (`-o`).
- Send standard output (data meant for pipelines) to `stdout`.
- Send diagnostics, prompts, debug info, and error messages to `stderr`.
- Return `0` for success and non-zero (e.g., `1`) for failure.
- Allow arguments to be provided via standard input (`stdin`) if appropriate.

**Do Not:**
- Invent custom, non-standard flags for common operations (e.g., using `-?` instead of `-h`).
- Print error messages to `stdout`, which can break chained commands in a pipeline.
- Ignore the exit status of commands run within a script.

## Example

**Bash script adhering to standard output streams and flags:**
```bash
#!/usr/bin/env bash

# Print errors to stderr
error() {
  echo "Error: $1" >&2
  exit 1
}

# Standard help flag
if [[ "$1" == "--help" || "$1" == "-h" ]]; then
  echo "Usage: mytool [OPTIONS] FILE"
  echo "Process FILE and output results."
  echo ""
  echo "Options:"
  echo "  -h, --help     Show this help message and exit"
  exit 0
fi

if [[ -z "$1" ]]; then
  error "Missing expected argument."
fi

# Actual output to stdout
echo "Processing complete."
```