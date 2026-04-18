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
Hexagonal Architecture, or Ports and Adapters, aims to create loosely coupled application components that can be easily connected to their software environment. It makes the application agnostic to the nature of the clients interacting with it (UI, test runners, other apps) and the infrastructure it uses (databases, message queues).

- **Core Application:** Contains the domain logic.
- **Ports:** Interfaces defined by the Core that dictate how it can be controlled (Driving/Primary Ports) and how it can control external dependencies (Driven/Secondary Ports).
- **Adapters:** Implementations that connect the outside world to the Ports.

## Instructions for AI Agents
Structure the codebase so that the domain logic does not depend on technical implementation details.

**Do:**
- Define interfaces (Ports) inside the domain layer for any external service (e.g., Database, Email Service).
- Write adapters outside the domain layer that implement these interfaces.
- Ensure the domain layer only imports domain code, never external libraries or adapters.
- Use this architecture to make the core logic highly testable in isolation.

**Do Not:**
- Hardcode database connection logic within service classes.
- Leak HTTP request objects (like Express `req` or Django `request`) into the domain layer.

## Example

**Domain Layer (Port):**
```typescript
// The domain defines what it needs
export interface NotificationPort {
  send(userId: string, message: string): Promise<void>;
}

export class OrderService {
  constructor(private notifier: NotificationPort) {}

  async completeOrder(userId: string) {
    // ... domain logic ...
    await this.notifier.send(userId, "Your order is complete!");
  }
}
```

**Infrastructure Layer (Adapter):**
```typescript
// The infrastructure implements the requirement using a specific technology
import { EmailClient } from 'some-email-library';

export class SendGridNotificationAdapter implements NotificationPort {
  async send(userId: string, message: string): Promise<void> {
    // Translate domain request to SendGrid API call
    const client = new EmailClient('API_KEY');
    await client.sendEmail(userId, message);
  }
}
```