---
tags:
  - principle
  - pragmatism
  - teamwork
  - decision-making
  - simplicity
---

# Pragmatic Consensus

## Principle
Pragmatic Consensus is the balance between theoretical architectural purity and the practical realities of delivering software within a team. It means choosing the "good enough" solution that the entire team can understand, maintain, and agree upon, rather than dogmatically enforcing complex "best practices" that slow down development or create cognitive overload.

## Instructions for AI Agents
When suggesting solutions or generating code, prioritize clarity, maintainability, and standard language idioms over over-engineered abstractions.

**Do:**
- Favor simple, procedural, or functional code for straightforward tasks instead of heavy object-oriented hierarchies.
- Choose native standard library features over introducing new third-party dependencies when reasonable.
- Respect the existing conventions and style of the codebase, even if they differ slightly from theoretical ideals.
- Document *why* a particular trade-off was made if choosing a less "pure" but more pragmatic approach.

**Do Not:**
- Introduce complex design patterns (like Abstract Factories or generic metaprogramming) unless the problem's complexity strictly requires it.
- Optimize prematurely for performance or scale before a bottleneck is proven.
- Refactor working, readable code purely for the sake of conforming to an abstract architectural rule.
