---
tags:
  - language
  - nix
---

# Nix Guidelines

## Introduction
Nix is a purely functional package manager and build system, and also the name of the expression language used to configure it. It is designed for reproducible builds and declarative system configuration.

**Crucial Instruction for AI Agents:** The Nix language and ecosystem can be complex. Always search and validate against [nix.dev](https://nix.dev/), the [NixOS Manual](https://nixos.org/manual/nixos/stable/), and [Noogle](https://noogle.dev/) for standard library functions. 

## Best Practices
- **Purity:** Functions should be pure. They should return the same output given the same input, without relying on external state (unless explicitly using impurities like `builtins.currentTime` or `fetchurl`).
- **Reproducibility:** Always pin dependencies. Use Nix Flakes (`flake.nix` and `flake.lock`) to ensure that your inputs are strictly pinned to specific git commits.
- **Let Bindings:** Use `let ... in` for local variables to keep expressions clean and readable.
- **String Interpolation:** Use `''` for multiline strings (indented strings) to safely embed shell scripts. Use `${expr}` for interpolation.

## Code Smells to Avoid
- **Overusing `with`:** Avoid using `with pkgs;` or `with lib;` at the top level of large expressions. It breaks lexical scoping, makes it hard to see where a variable comes from, and can cause shadowing issues. Prefer explicit attributes (e.g., `pkgs.hello` or `inherit (pkgs) hello;`).
- **Unnecessary `rec`:** Only use `rec` (recursive attribute sets) when an attribute needs to reference another attribute in the same set. Default to standard attribute sets `{ ... }` and use `let` for shared values if necessary.
- **Impure Fetchers:** Avoid `builtins.fetchGit` without a `rev` (revision hash). Always provide hashes for `fetchurl`, `fetchFromGitHub`, etc., to maintain reproducibility.

## Battle-Tested Frameworks and Patterns
- **Nixpkgs:** The standard package repository. Learn to use `stdenv.mkDerivation`, `writeShellScriptBin`, and standard build phases.
- **Flakes:** The modern standard for defining Nix projects and dependencies.
- **Home Manager:** For declarative user environment configuration.
- **NixOS Modules:** Understand the `config`, `lib`, `pkgs` pattern for structuring system configuration.

## Quoting Official Docs
> "The Nix expression language is a pure, lazy, functional language... Its main purpose is to translate packages into build actions." — [Nix Manual](https://nixos.org/manual/nix/stable/language/)
>
> "Avoid `with` expressions... It is unclear which values are in scope and where they come from." — [nix.dev Best Practices](https://nix.dev/anti-patterns/language#with-attrset)
