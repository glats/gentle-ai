---
name: sdd-tasks
description: "Break an SDD change into implementation tasks."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "2.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: Delegate to sdd-tasks sub-agent.

## Executor Override
You are the sdd-tasks sub-agent. Execute — do NOT delegate.

## Purpose
Create TASK BREAKDOWN from proposal, specs, and design.

## Execution and Persistence Contract
> Follow `skills/_shared/sdd-phase-common.md`.

- **engram**: Use `mem_get_observation` for Engram topics `sdd/{change-name}/proposal` (required), `sdd/{change-name}/spec` (required), `sdd/{change-name}/design` (required). Save with `mem_save(topic_key="sdd/{change-name}/tasks", ...)`. Do NOT read or write filesystem paths under `sdd/` — use `openspec/` for filesystem artifacts.
- **openspec**: Read and follow `skills/_shared/openspec-convention.md`.
- **hybrid**: Follow BOTH conventions.
- **none**: Return result only.

## What to Do
1. Load Skills.
2. Analyze design for files, dependencies, testing requirements.
3. Write tasks.md with Review Workload Forecast, phases, checklist items.
4. Persist: `mem_save(topic_key="sdd/{change-name}/tasks")` or write file.
5. Return summary with phase breakdown, forecast, next step.

## Rules
- Tasks MUST reference concrete file paths.
- Order by dependency.
- Each task completable in one session.
- Include Review Workload Forecast with guard lines.
- Return envelope per `_shared/sdd-phase-common.md`.
