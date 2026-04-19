---
id: hexagonal-architecture
priority: 7
scope: architecture
tags:
  - principle
  - architecture
  - ports-and-adapters
  - decoupling
  - testing
applies_when: Core logic must remain independent from replaceable external systems.
conflicts_with:
  - pragmatic-consensus
  - kiss
  - yagni
decision_rule: Use ports/adapters when integration boundaries are central to the problem.
---

# Hexagonal Architecture (Ports and Adapters)

## Principle
Domain logic communicates through ports; adapters handle external systems.

## Intent
Keep the core independent from transport, storage, and vendor details.

## Apply when
- The app has multiple external dependencies (DB, queue, API, UI).
- You need high testability and replaceable integrations.

## Avoid when
- A simple internal feature does not need replaceable adapters.

## Do
- Define ports in the application/domain boundary.
- Implement adapters outside the core.
- Keep domain imports limited to domain/application code.

## Do not
- Embed database/network/framework calls in domain services.
- Pass transport objects directly into domain logic.

## Conflict resolution
- Follow `principles/_priority-model.md`.

## Example (pseudo)
```pseudo
# bad
domain_service(request, sql_client) {
  row = sql_client.query(request.id)
  return row
}

# good
domain_service(load_user_port, user_id) {
  user = load_user_port.load(user_id)
  return user
}
```

## Checklist
- Are external calls isolated in adapters?
- Are ports defined at core boundaries?
- Can core logic be tested without real infrastructure?

