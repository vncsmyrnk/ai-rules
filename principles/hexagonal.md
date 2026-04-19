---
tags:
  - principle
  - architecture
  - ports-and-adapters
  - decoupling
  - testing
---

# Hexagonal Architecture (Ports and Adapters)

## Principle
Domain logic communicates through ports; adapters handle external systems.

## Intent
Keep the core independent from transport, storage, and vendor details.

## Apply when
- The app has multiple external dependencies (DB, queue, API, UI).
- You need high testability and replaceable integrations.

## Do
- Define ports in the application/domain boundary.
- Implement adapters outside the core.
- Keep domain imports limited to domain/application code.

## Do not
- Embed database/network/framework calls in domain services.
- Pass transport objects directly into domain logic.

## Trade-offs
- Strong isolation and testability.
- Extra abstractions for small/simple applications.

## Conflict resolution
- Prefer simpler layering for low-complexity features (Pragmatic Consensus).
- Priority order: Correctness > Existing project conventions > Pragmatic Consensus > Hexagonal Architecture > Boy Scout cleanup.

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

