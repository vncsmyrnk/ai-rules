---
id: lang-zig
scope: language
tags:
  - language
  - zig
references:
  - https://ziglang.org/documentation/master/
  - https://ziglang.org/learn/
---

# Zig Guidelines

## Principle
Write explicit Zig code with clear memory ownership and explicit error flow.

## Apply when
- Writing Zig applications, libraries, or build scripts.

## Do
- Pass allocator explicitly when allocation is required.
- Use error unions and `try`/`catch` intentionally.
- Use `defer`/`errdefer` immediately after acquiring resources.
- Keep pointer/slice ownership and lifetimes explicit.
- Use `comptime` only when it simplifies or removes duplication.
- Run `zig fmt`, `zig test`, and build checks.

## Do not
- Hide allocations or control flow.
- Use `catch unreachable` unless impossibility is proven.
- Return references to invalid stack data.
- Overuse `comptime` for simple runtime logic.

## Reliability basics
- Prefer explicit bounds-safe operations.
- Test with allocators that detect leaks in tests.
- Keep APIs small and predictable.

## Example
````zig
// bad
const data = try allocator.alloc(u8, n);
process(data);

// good
const data = try allocator.alloc(u8, n);
defer allocator.free(data);
process(data);
````

## Checklist
- Is memory ownership explicit?
- Are all errors handled or propagated?
- Are resource releases guaranteed?
- Is `comptime` usage justified?
