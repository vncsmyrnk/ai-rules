---
id: lang-go
scope: language
tags:
  - language
  - go
  - golang
references:
  - https://go.dev/doc/effective_go
  - https://github.com/golang/go/wiki/CodeReviewComments
  - https://github.com/uber-go/guide
---

# Go (Golang) Guidelines

## Principle
Write clear, idiomatic, and maintainable Go with explicit error handling and minimal abstraction.

## Source priority
1. Effective Go
2. Go Code Review Comments
3. Uber Go Style Guide
4. Standard library docs and source (`pkg.go.dev`)

## Apply when
- Writing or modifying Go code, tests, packages, or services.

## Do
- Run `gofmt` (or `goimports`) on every change.
- Keep package names short, lowercase, and meaningful.
- Add comments to exported identifiers starting with the identifier name.
- Handle errors explicitly; wrap with `%w` when propagating context.
- Return early to reduce nesting.
- Pass `context.Context` as the first parameter for request-scoped work.
- Define interfaces at the consumer side; keep them small.
- Prefer composition over inheritance-like patterns.
- Write table-driven tests where useful; run `go test` and `go test -race` when concurrency is involved.
- Use standard tools: `go vet`, `staticcheck` (if available).

## Do not
- Ignore errors with `_` without a clear reason.
- Use `panic` for normal control flow.
- Add interfaces or layers with only one trivial implementation.
- Use `init()` for hidden complex setup.
- Create package-level mutable state unless required and synchronized.

## API and style specifics
- Accept interfaces, return concrete types when practical.
- Keep zero values useful.
- Use initialisms consistently (`HTTPServer`, `userID`).
- Keep receiver names short and consistent (`func (s *Service) ...`).

## Example
````go
// bad
func Fetch(id string) (User, error) {
    u, _ := repo.Load(id)
    return u, nil
}

// good
func Fetch(ctx context.Context, id string, repo UserRepo) (User, error) {
    u, err := repo.Load(ctx, id)
    if err != nil {
        return User{}, fmt.Errorf("load user %s: %w", id, err)
    }
    return u, nil
}
````

## Checklist
- Is code formatted and idiomatic?
- Are errors handled and wrapped with context?
- Is `context.Context` used correctly?
- Are abstractions justified by real usage?
