---
id: rest
priority: 3
scope: api
tags:
  - principle
  - api
  - rest
  - http
  - standards
applies_when: Designing or modifying HTTP APIs intended to be RESTful.
conflicts_with:
  - pragmatic-consensus
decision_rule: Use correct HTTP semantics and resource-oriented design; prioritize correctness and security over stylistic purity.
---

# REST API Standards

## Principle
Design resource-oriented HTTP APIs with correct semantics, predictable behavior, and interoperable error/caching patterns.

## Intent
Improve API consistency, client reliability, and long-term evolvability.

## Apply when
- Building or changing HTTP/JSON APIs.

## Avoid when
- The system is intentionally RPC/event-based and not REST-oriented.

## Do
- Model URIs as resources (nouns), typically plural collections.
- Use methods by semantics (RFC 9110):
  - `GET` read (safe)
  - `POST` create/action on collection
  - `PUT` full replace (idempotent)
  - `PATCH` partial update (RFC 5789)
  - `DELETE` remove (idempotent)
- Use accurate status codes (RFC 9110):
  - `200`, `201` (+ `Location`), `202`, `204`
  - `400`, `401`, `403`, `404`, `409`, `412`, `415`, `422`, `429`
  - `5xx` for server faults
- Return structured errors using Problem Details (RFC 9457).
- Support caching and validators where relevant (RFC 9111, ETag/If-None-Match, Last-Modified).
- Use conditional updates for concurrency (`If-Match` + `ETag`, RFC 7232).
- Use explicit content negotiation (`Accept`, `Content-Type`).

## Do not
- Use verb-style CRUD endpoints (`/createUser`, `/deleteOrder`).
- Return `200` for all outcomes.
- Return unstructured or inconsistent error payloads.
- Break clients silently; version and deprecate intentionally.

## Conflict resolution
- Follow `principles/_priority-model.md`.
- If legacy constraints exist, keep behavior backward-compatible and improve incrementally.

## Example (pseudo)
```pseudo
# create resource
request: POST /users
body: { "name": "Ada" }
response: 201
headers: Location: /users/123
body: { "id": 123, "name": "Ada" }

# conditional partial update
request: PATCH /users/123
headers: If-Match: "etag-v1"
body: { "name": "Ada Lovelace" }
response: 200
headers: ETag: "etag-v2"

# standardized error (Problem Details)
response: 409
body: {
  "type": "https://api.example/problems/conflict",
  "title": "Conflict",
  "status": 409,
  "detail": "Email already exists"
}
```

## Checklist
- Are URIs resource-oriented and method semantics correct?
- Are status codes and error bodies consistent with standards?
- Are caching and concurrency controls used where needed?
- Are compatibility/versioning expectations explicit?
