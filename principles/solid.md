---
id: solid
priority: 6
scope: architecture
tags:
  - principle
  - design
  - solid
  - maintainability
applies_when: Code has growing complexity, frequent changes, or multiple implementation variants.
conflicts_with:
  - kiss
  - yagni
decision_rule: Apply SOLID only where it reduces coupling and change risk with clear net benefit.
---

# SOLID (Pragmatic Applicability)

## Principle
Use SOLID to reduce coupling and improve maintainability when complexity justifies abstraction.

## Intent
Keep code change-friendly without adding unnecessary indirection.

## Apply when
- A module has multiple reasons to change.
- New variants/features repeatedly require editing stable code.
- Clients depend on methods they do not use.
- Business logic is tightly coupled to infrastructure.

## Avoid when
- Code is small, stable, and easy to understand.
- Abstractions would add indirection without real change pressure.

## Do
- **S (Single Responsibility):** one primary reason to change per unit.
- **O (Open/Closed):** extend behavior via composition/polymorphism, not repeated edits in stable flows.
- **L (Liskov):** subtypes preserve base contract behavior.
- **I (Interface Segregation):** expose small, role-focused interfaces.
- **D (Dependency Inversion):** depend on abstractions at boundaries.

## Do not
- Create interface-per-class by default.
- Introduce deep inheritance hierarchies.
- Generalize for hypothetical future needs.

## Conflict resolution
- Follow `principles/_priority-model.md`.
- If SOLID increases complexity for a simple case, prefer KISS/YAGNI.

## Example (pseudo)
```pseudo
# bad (SRP violation)
class ReportService {
  generate_report()
  save_to_database()
  send_email()
}

# good (SRP + DIP)
class ReportGenerator { generate_report() }
class ReportRepositoryPort { save(report) }
class ReportNotifierPort { notify(report) }
class ReportUseCase(repo_port, notifier_port) {
  run() { report = generator.generate_report(); repo_port.save(report); notifier_port.notify(report) }
}

# bad (ISP violation)
interface Worker {
  code()
  deploy()
  approve_budget()
}

# good (ISP)
interface Coder { code() }
interface Deployer { deploy() }
```

## Checklist
- Does each unit have one primary responsibility?
- Are extensions possible without editing stable core logic repeatedly?
- Do clients depend only on methods they use?
- Are infrastructure details inverted behind boundaries where needed?
