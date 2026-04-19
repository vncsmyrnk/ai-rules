---
id: lang-nix
scope: language
tags:
  - language
  - nix
references:
  - https://nix.dev/
  - https://nixos.org/manual/nix/stable/
  - https://noogle.dev/
---

# Nix Guidelines

## Principle
Keep Nix expressions pure, reproducible, and explicit.

## Apply when
- Writing flakes, derivations, modules, or Nix development shells.

## Do
- Pin inputs (use `flake.lock`).
- Keep expressions pure and deterministic.
- Prefer explicit attribute access (`pkgs.foo`) over broad `with`.
- Use `let ... in` for local structure and readability.
- Provide hashes for fetchers.
- Use `nix fmt`/`alejandra` and `nix flake check` when applicable.

## Do not
- Use unpinned or impure fetches for reproducible builds.
- Overuse `with` at top level.
- Use `rec` unless intra-attr references are required.
- Hide important behavior in stringly shell code when Nix options exist.

## Module and package basics
- Prefer `stdenv.mkDerivation` conventions.
- Keep derivations minimal and declarative.
- In modules, separate options, config, and logic clearly.

## Example
````nix
# bad
with pkgs; { buildInputs = [ git jq ]; }

# good
{ pkgs, ... }: {
  buildInputs = [ pkgs.git pkgs.jq ];
}
````

## Checklist
- Are dependencies pinned?
- Are fetchers hashed and deterministic?
- Is scope explicit and readable?
- Is module/derivation structure minimal and clear?
