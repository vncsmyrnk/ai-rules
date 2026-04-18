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

## Example

**Bad (Mixed Concerns):**
```javascript
async function handleUserRegistration(req, res) {
  // 1. Parsing request
  const { username, password, email } = req.body;
  
  // 2. Business logic & Validation
  if (!email.includes('@')) {
    return res.status(400).send('Invalid email');
  }
  
  // 3. Data Access
  const db = await connectToDatabase();
  const existingUser = await db.query(`SELECT * FROM users WHERE email = '${email}'`);
  if (existingUser) {
    return res.status(400).send('User exists');
  }
  
  // 4. More Business logic & Data Access
  const hashedPassword = hashPassword(password);
  await db.query(`INSERT INTO users (username, password, email) VALUES ('${username}', '${hashedPassword}', '${email}')`);
  
  // 5. Response
  res.status(201).send('User registered');
}
```

**Good (Separated Concerns):**
```javascript
// Controller (HTTP concern)
async function handleUserRegistration(req, res) {
  try {
    const userData = extractUserData(req.body); // Parsing concern
    await userService.registerUser(userData);
    res.status(201).send('User registered');
  } catch (error) {
    res.status(400).send(error.message);
  }
}

// Service (Business logic concern)
class UserService {
  async registerUser(userData) {
    if (!isValidEmail(userData.email)) {
      throw new Error('Invalid email');
    }
    if (await userRepository.findByEmail(userData.email)) {
      throw new Error('User exists');
    }
    const hashedPassword = hashPassword(userData.password);
    await userRepository.save({ ...userData, password: hashedPassword });
  }
}

// Repository (Data access concern)
class UserRepository {
  async findByEmail(email) {
    // DB query...
  }
  async save(user) {
    // DB insert...
  }
}
```
