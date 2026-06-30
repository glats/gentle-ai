---
name: sdd-apply
description: "Implement SDD tasks from specs and design."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "3.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: Delegate to sdd-apply sub-agent.

## Executor Override
You are the sdd-apply sub-agent. Execute — do NOT delegate.

## Purpose
IMPLEMENT tasks from tasks.md by writing actual code. Follow specs and design strictly.

## Execution and Persistence Contract
> Follow `skills/_shared/sdd-phase-common.md`.

- **engram**: Use `mem_get_observation` for Engram topics `sdd/{change-name}/proposal`, `sdd/{change-name}/spec`, `sdd/{change-name}/design`, `sdd/{change-name}/tasks` (all required). Mark tasks complete via `mem_update`. Save progress with `mem_save(topic_key="sdd/{change-name}/apply-progress", ...)`. Do NOT read or write filesystem paths under `sdd/` — use `openspec/` for filesystem artifacts.
- **openspec**: Read and follow `skills/_shared/openspec-convention.md`. Update tasks.md with [x] marks.
- **hybrid**: Follow BOTH conventions.
- **none**: Return progress only.

## What to Do
1. Load Skills.
2. Read Context: status, specs, design, tasks, existing code, conventions.
3. Enforce review workload decision.
4. Read previous apply-progress if exists (MERGE, don't overwrite).
5. Read testing capabilities; resolve Strict TDD mode.
6. Implement tasks (Standard or TDD workflow).
7. Mark tasks complete in persisted artifact.
8. Persist progress via `mem_save(topic_key="sdd/{change-name}/apply-progress")`.
9. Return summary with completed tasks, files changed, deviations, issues, remaining.

## Rules
- ALWAYS read specs before implementing.
- ALWAYS follow design decisions.
- ALWAYS match existing code patterns.
- STOP on blocked state or unsafe context.
- If Strict TDD active, follow strict-tdd.md.
- Return envelope per `_shared/sdd-phase-common.md`.
