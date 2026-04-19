---
id: lang-lua
scope: language
tags:
  - language
  - lua
references:
  - https://www.lua.org/manual/5.4/
  - https://luarocks.org/
---

# Lua Guidelines

## Principle
Prefer explicit, local, and simple Lua code with predictable table and module patterns.

## Apply when
- Writing Lua scripts, modules, plugins, or embedded runtime logic.

## Do
- Declare variables as `local` by default.
- Use tables intentionally: sequence vs map semantics.
- Remember arrays are 1-based.
- Build modules by returning a local table.
- Use `pcall`/`xpcall` for recoverable boundaries.
- Use `table.concat` for large string assembly.
- Run `stylua` and `luacheck` when available.

## Do not
- Create unintended globals.
- Depend on `#` for sparse arrays with holes.
- Overuse metatables when plain functions/tables are enough.
- Mutate shared module state without clear ownership.

## Performance and safety basics
- Cache hot global lookups locally in tight loops.
- Validate untrusted input and table shapes at boundaries.
- Keep metamethod behavior minimal and documented.

## Example
````lua
-- bad
count = count + 1

-- good
local count = 0
count = count + 1
````

## Checklist
- Are variables local unless explicitly global?
- Are table semantics (array/map) clear?
- Is error handling explicit at boundary calls?
- Is module state controlled?
