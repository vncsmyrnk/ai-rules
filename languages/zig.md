---
tags:
  - language
  - zig
---

# Zig Guidelines

## Introduction
Zig is a general-purpose programming language and toolchain for maintaining robust, optimal, and reusable software. It is a modern alternative to C, focusing on manual memory management, explicit control flow, and compile-time execution (`comptime`).

**Crucial Instruction for AI Agents:** Zig is an evolving language. Always check the [Official Zig Documentation](https://ziglang.org/documentation/master/) and standard library source code to validate syntax, functions, and best practices.

## Best Practices
- **Explicit Memory Allocation:** Zig has no hidden memory allocation. Functions that allocate memory must explicitly accept an `Allocator` as an argument (usually `std.mem.Allocator`).
- **Error Handling:** Use Zig's native error sets and the `!` operator. Use `try` to propagate errors and `catch` to handle them. Never silently ignore an error.
- **Defer:** Use `defer` (and `errdefer`) to execute code (like freeing memory or closing files) at the exit of the current scope, regardless of how the scope is exited.
- **Comptime:** Leverage `comptime` for generic programming and compile-time evaluation, but keep it simple. It replaces macros and templates.

## Code Smells to Avoid
- **Hidden Control Flow:** Zig's philosophy is "no hidden control flow." Do not attempt to overload operators or create macros that obscure what the CPU is actually doing.
- **Ignoring Errors:** Using `catch unreachable` when an error is actually possible in a production environment. Only use it when you can mathematically prove the error will never occur.
- **Memory Leaks:** Forgetting to `defer allocator.free(memory)` immediately after allocation. Always test with `std.testing.allocator` which catches memory leaks during tests.

## Battle-Tested Frameworks and Patterns
- **Build System:** Zig includes its own build system (`build.zig`). Use it instead of Make or CMake to define executables, libraries, and tests.
- **C Interoperability:** Zig can natively import C headers (`@cImport`) and compile C code. Use this feature to easily leverage existing C libraries without writing bindings.
- **Standard Library:** Rely heavily on `std`, but be aware that it is subject to change.

## Quoting Official Docs
> "Focus on debugging your application rather than debugging your programming language knowledge." — Zig Philosophy
>
> "No hidden control flow. No hidden memory allocation. No preprocessor, no macros." — [Zig Language Reference](https://ziglang.org/learn/why_zig_rust_d_cpp/)
