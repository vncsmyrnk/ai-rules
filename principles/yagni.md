---
id: yagni
priority: 5
scope: global
tags:
  - principle
  - scope
  - simplicity
  - product
applies_when: Deciding whether to build optional or speculative functionality.
conflicts_with:
  - clean-architecture
  - dry
decision_rule: Implement only what is required now; defer speculative needs.
---

# YAGNI (You Aren't Gonna Need It)

## Principle
Do not implement functionality until it is needed.

## Intent
Prevent overengineering and reduce unused code.

## Apply when
- You are tempted to add “future” support without a current requirement.

## Avoid when
- A small addition is required for correctness, safety, or contractual compatibility.

## Do
- Implement current requirements only.
- Leave clear extension points when cheap and obvious.
- Remove dead speculative code.

## Do not
- Add unused configuration flags, hooks, or abstractions.
- Build generalized frameworks for one concrete use case.

## Conflict resolution
- Follow `principles/_priority-model.md`.

## Example (pseudo)
```pseudo
# bad
build_plugin_system_for_single_formatter()

# good
build_single_formatter()
```

## Checklist
- Is each added part required now?
- Is there evidence for new abstraction needs?
- Did we avoid speculative features?
