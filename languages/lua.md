---
tags:
  - language
  - lua
---

# Lua Guidelines

## Introduction
Lua is a powerful, efficient, lightweight, embeddable scripting language. It is commonly used in game development, embedded systems, and for scripting existing applications.

**Crucial Instruction for AI Agents:** Always search the [Lua 5.4 Reference Manual](https://www.lua.org/manual/5.4/) (or the specific version in use) and [LuaRocks](https://luarocks.org/) for packages. Validate syntax and idiomatic usage, especially concerning metamethods and environments.

## Best Practices
- **Local Variables:** **Always** declare variables as `local`. Global variables are the default in Lua if `local` is omitted, which pollutes the global namespace and significantly hurts performance.
- **Tables for Everything:** Tables are Lua's only data structuring mechanism. Use them for arrays (1-based), dictionaries, objects, and modules.
- **1-Based Indexing:** Remember that standard Lua arrays (sequences) start at index `1`, not `0`. The length operator `#` only works reliably on sequences without "holes" (nil values).
- **Modules:** Return a local table from a file to create a module. Avoid the deprecated `module()` function.

## Code Smells to Avoid
- **Unintended Globals:** Forgetting the `local` keyword. Use linters like `luacheck` to catch these.
- **Metatable Abuse:** While metatables (using `__index`, `__newindex`, etc.) are powerful for OOP and operator overloading, overusing them makes code hard to follow and debug.
- **String Concatenation in Loops:** Use `table.concat()` when building large strings instead of the `..` operator inside a loop to avoid excessive memory allocation.

## Battle-Tested Frameworks and Patterns
- **Game Development:** [LÖVE (Love2D)](https://love2d.org/) is the standard framework for 2D games.
- **Web and API:** [OpenResty](https://openresty.org/) embeds Lua into Nginx for ultra-high performance web applications.
- **Testing:** [Busted](https://olivinelabs.com/busted/) for unit testing.

## Quoting Official Docs
> "Local variables have their scope limited to the block where they are declared... It is good programming style to use local variables whenever possible." — [Programming in Lua](https://www.lua.org/pil/4.2.html)
