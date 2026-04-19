---
tags:
  - principle
  - maintenance
  - refactoring
  - clean-code
---

# The Boy Scout Rule

## Principle
"Always leave the campground cleaner than you found it." 

In software development, this means that whenever you touch a piece of code, you should strive to leave it in a slightly better state than it was before. 

## Instructions for AI Agents
When modifying a file to implement a feature or fix a bug, actively look for minor, low-risk improvements you can make in the immediate vicinity of your changes. 

**Do:**
- Fix obvious typos in comments or variable names.
- Remove unused variables or imports.
- Improve the naming of a poorly named local variable.
- Reformat misaligned code to match the surrounding style.
- Extract a complex conditional into a well-named boolean variable if it improves readability.

**Do Not:**
- Embark on massive refactoring missions outside the scope of your primary task.
- Change the public API or behavior of classes/functions unless explicitly requested.
- Let the "cleanup" obscure the primary intent of the commit/change.
