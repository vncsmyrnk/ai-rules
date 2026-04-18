---
tags:
  - language
  - go
  - golang
---

# Go (Golang) Guidelines

## Introduction
Go is designed for simplicity, concurrency, and performance. Its design philosophy favors straightforward, readable code over clever tricks or complex abstractions.

**Crucial Instruction for AI Agents:** Always search for the latest packages and documentation on [pkg.go.dev](https://pkg.go.dev) and [go.dev/doc](https://go.dev/doc). Validate your assumptions by checking official standard library implementations when in doubt.

## Best Practices
- **Effective Go:** Always follow the conventions outlined in [Effective Go](https://go.dev/doc/effective_go) and [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).
- **Error Handling:** Errors in Go are values. Always handle them explicitly. Never use `_` to ignore an error unless you have a rigorously documented reason. 
- **Concurrency:** Use goroutines and channels for concurrent state, but prefer simpler synchronization (like `sync.Mutex`) if the concurrency is just protecting shared state. "Do not communicate by sharing memory; instead, share memory by communicating."
- **Interfaces:** Define interfaces where they are used (consumer side), not where they are implemented (producer side). Keep them small (e.g., `io.Reader`, `io.Writer`).
- **Formatting:** Always run `gofmt` or `goimports` on generated code.

## Code Smells to Avoid
- **Deep Nesting:** Return early to avoid deeply nested `if` statements (the "line of sight" rule).
- **Package-Level State:** Avoid global variables and `init()` functions that set up complex hidden state.
- **Overusing Interfaces:** Don't create an interface for a struct unless there is more than one implementation or it is required for mocking in tests.
- **Ignoring Contexts:** Always pass `context.Context` as the first argument to functions that do I/O or RPC to handle timeouts and cancellation.

## Battle-Tested Frameworks and Patterns
- **Standard Library:** Prefer `net/http` for routing and HTTP servers. The standard library is extremely powerful.
- **Web Frameworks:** [Gin](https://gin-gonic.com/) or [Echo](https://echo.labstack.com/) for high-performance REST APIs.
- **Database:** `database/sql` combined with [sqlx](https://github.com/jmoiron/sqlx) or [pgx](https://github.com/jackc/pgx) for PostgreSQL. Avoid heavy ORMs unless specifically requested; prefer raw SQL or query builders like [Squirrel](https://github.com/Masterminds/squirrel).

## Quoting Official Docs
> "Clear is better than clever." — Go Proverbs
>
> "Errors are just values... The most important lesson to learn about error handling in Go is that it is just programming." — [Error handling and Go](https://go.dev/blog/error-handling-and-go)
