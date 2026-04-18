---
tags:
  - workflow
  - task
  - code-review
  - quality-assurance
---

# Code Review Workflow

## Purpose
This workflow provides instructions for an AI agent to perform a comprehensive code review on a set of changes, a Pull Request, or a specific file.

## Instructions for AI Agents
When tasked with performing a code review, follow these steps systematically. Your goal is not just to find bugs, but to improve the overall quality, maintainability, and readability of the code while respecting the project's established conventions.

### 1. Context Gathering
- Identify the language and frameworks being used.
- Search for and read any relevant `languages/` or `principles/` guidelines provided in the context.
- Understand the intent of the changes (e.g., bug fix, feature addition, refactoring). If the intent is unclear, ask the user for clarification before proceeding.

### 2. Structural & Architectural Review
- **Separation of Concerns:** Does the code mix different layers of abstraction (e.g., UI logic in database queries)?
- **Clean Architecture:** Do dependencies point in the correct direction?
- **Pragmatism:** Is the solution overly complex for the problem it solves? 
- Check if the changes align with the established architectural patterns of the project.

### 3. Code-Level Review (The Boy Scout Rule)
- **Readability:** Are variables and functions named descriptively? Can the code be easily understood by a newcomer?
- **Simplicity:** Are there deep nested conditionals? Can they be refactored using early returns?
- **Error Handling:** Are errors caught and handled gracefully? Are error messages informative?
- **Performance:** Are there obvious bottlenecks (e.g., N+1 queries, unnecessary loops, unbounded memory allocations)?
- **Security:** Are inputs validated and sanitized? Are secrets hardcoded?

### 4. Testing
- Do the changes include appropriate unit or integration tests?
- Do the tests cover edge cases and failure modes, not just the happy path?

### 5. Formatting the Review Output
Present your findings clearly and constructively. 
- Use a polite, collaborative tone.
- Group your feedback into categories: **Critical Issues** (must fix), **Suggestions** (nice to have), and **Praise** (point out good practices).
- Provide specific code snippets demonstrating how to improve the code when making suggestions.

## Execution Prompt Example
*If a user asks "Review this code", apply the instructions above and generate a response structured like this:*

```markdown
### Code Review Summary
[Brief summary of the changes and overall impression]

### 🔴 Critical Issues (Must Fix)
- [File/Line]: [Explanation of the issue, e.g., security flaw, major bug, architectural violation]
  - *Recommendation:* [Code snippet showing the fix]

### 🟡 Suggestions (Nice to Have)
- [File/Line]: [Explanation of potential improvement, e.g., readability, minor performance tweak]
  - *Recommendation:* [Code snippet showing the improvement]

### 🟢 Praise
- [File/Line]: [Acknowledge well-written code, good use of patterns, or clever solutions]
```