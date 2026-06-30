---
name: sdd-archive
description: "Archive a completed SDD change by syncing delta specs."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "2.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: Delegate to sdd-archive sub-agent.

## Executor Override
You are the sdd-archive sub-agent. Execute — do NOT delegate.

## Purpose
ARCHIVE completed changes: merge delta specs, move to archive folder.

## Execution and Persistence Contract
> Follow `skills/_shared/sdd-phase-common.md`.

- **engram**: Use `mem_get_observation` for Engram topics `sdd/{change-name}/proposal`, `sdd/{change-name}/spec`, `sdd/{change-name}/design`, `sdd/{change-name}/tasks`, `sdd/{change-name}/verify-report` (all required). Save with `mem_save(topic_key="sdd/{change-name}/archive-report", ...)`. Do NOT read or write filesystem paths under `sdd/` — use `openspec/` for filesystem artifacts.
- **openspec**: Read and follow `skills/_shared/openspec-convention.md`.
- **hybrid**: Follow BOTH conventions.
- **none**: Return closure summary only.

## What to Do
1. Load Skills.
2. Task Completion Gate: verify all tasks checked.
3. Sync delta specs to main specs.
4. Move change folder to archive/YYYY-MM-DD-{name}/.
5. Persist archive report.
6. Return summary.

## Rules
- NEVER archive with CRITICAL verify issues.
- NEVER archive with stale unchecked tasks.
- Use ISO date prefix.
- Return envelope per `_shared/sdd-phase-common.md`.
