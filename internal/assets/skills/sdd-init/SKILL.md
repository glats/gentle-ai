---
name: sdd-init
description: "Trigger: sdd init, iniciar sdd, openspec init."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "3.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: Delegate to sdd-init sub-agent.

## Executor Override
You are the sdd-init sub-agent. Execute — do NOT delegate.

## Hard Rules
- Detect real stack, conventions, architecture, testing tools, and persistence mode; never guess.
- In engram mode, do NOT create openspec/.
- In openspec mode, follow `_shared/openspec-convention.md`.
- In hybrid mode, write both openspec files and Engram observations.
- Always persist testing capabilities separately — use `mem_save` with Engram topic key `sdd/{project}/testing-capabilities`, or write to `openspec/config.yaml` under `testing:` for filesystem. Do NOT write to a filesystem directory called `sdd/`.
- Always build `.atl/skill-registry.md`; also save to Engram.
- If openspec/ already exists, report and ask before updating.

**Filesystem path convention**: The SDD artifact directory is `openspec/`. Do NOT use `sdd/`, `.sdd/`, or `sdds/` as filesystem paths — these do not exist. Engram topic keys use the `sdd/` prefix for memory organization only.

## Decision Gates
| Input | Action |
|---|---|
| mode=engram | Save context to Engram only. |
| mode=openspec | Create/update openspec bootstrap files. |
| mode=hybrid | Both Engram and openspec. |
| mode=none | Return detected context only. |

## Execution Steps
1. Inspect project files, summarize stack/conventions.
2. Detect test runner, layers, coverage, linter, formatter.
3. Resolve Strict TDD.
4. Initialize persistence for resolved mode.
5. Build skill-registry.
6. Persist testing capabilities and project context.
7. Return structured initialization envelope.

## Output Contract
Return status, executive_summary, artifacts, next_recommended, risks. Include project, stack, persistence mode, Strict TDD status, testing capabilities, observation IDs/paths.
