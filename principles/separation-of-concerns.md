---
tags:
  - principle
  - architecture
  - decoupling
  - modularity
---

# Separation of Concerns (SoC)

## Principle
Separation of Concerns is a design principle for separating a computer program into distinct sections such that each section addresses a separate concern. A concern is a set of information that affects the code of a computer program. Good SoC results in modular code that is easier to maintain, test, and understand.

## Instructions for AI Agents
When writing or modifying code, ensure that different responsibilities are handled by different modules, classes, or functions.

**Do:**
- Split large functions that do multiple things into smaller, focused functions.
- Keep UI rendering logic separate from business logic and data access logic.
- Group related features together and hide internal implementation details behind well-defined interfaces.

**Do Not:**
- Mix HTML/UI generation, database queries, and business rules in the same file or function.
- Create "god objects" or massive utility files that handle unrelated tasks.
