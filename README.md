# AI rules

This project provides a structured set of guidelines, principles, workflows, and language rules to build high-quality prompts for AI coding agents.

## Structure

- **`principles/`**: Software engineering principles and decision rules.
- **`workflows/`**: Task-specific execution workflows.
- **`languages/`**: Language-specific coding guidelines.
- **`bin/`**: CLI tools for indexing and selecting rules.
- **`index/`**: Generated deterministic rule index.

## Rule selection MVP

The project now includes a deterministic minimum viable retrieval pipeline:

1. Build a stable JSON index from markdown files (`principles/`, `languages/`, `workflows/`).
2. Apply deterministic hard filtering (language/topic inference).
3. Apply deterministic scoring and ranking.
4. Output selected rules as JSON or assembled prompt text.


### Build index

- `./bin/rules-index`

This generates `index/rules-index.json` with stable ordering.

### Select rules from a prompt

- JSON output:
  - `./bin/rules-select "Build a REST API in Go with tests" --top-k 8 --format json`
- Prompt block output:
  - `./bin/rules-select "Build a REST API in Go with tests" --top-k 8 --format prompt`
- Restrict categories:
  - `./bin/rules-select "Review this API" --category workflows,principles --top-k 8 --format json`
  - `./bin/rules-select "Write Go code" --category languages --top-k 5 --format json`
