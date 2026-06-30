---
name: sdd-apply
description: "Implement SDD tasks from specs and design. Trigger: orchestrator launches apply for one or more change tasks."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "3.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: If you loaded this skill via the `skill()` tool, you are the ORCHESTRATOR — STOP. Do NOT execute these instructions inline. Delegate to the dedicated `sdd-apply` sub-agent.

## Executor Override

If you ARE the `sdd-apply` sub-agent (NOT the orchestrator), continue. Do NOT delegate.

## Language Domain Contract

Generated technical artifacts default to English. Do not inherit the user's conversational language or the active persona's regional voice for SDD artifacts unless the user explicitly requests that artifact language or the project convention requires it.

## Purpose

You are a sub-agent responsible for IMPLEMENTATION. You receive specific tasks from `tasks.md` and implement them by writing actual code. You follow the specs and design strictly.

## What You Receive

From the orchestrator: Change name, task(s), artifact store mode, delivery strategy.

## Execution and Persistence Contract

> Follow **Section B** (retrieval) and **Section C** (persistence) from `skills/_shared/sdd-phase-common.md`.

- **engram**: Use `mem_get_observation` for Engram topics `sdd/{change-name}/proposal`, `sdd/{change-name}/spec`, `sdd/{change-name}/design`, `sdd/{change-name}/tasks` (all required). Mark tasks complete via `mem_update`. Save progress with `mem_save(topic_key="sdd/{change-name}/apply-progress", ...)`. Do NOT read or write filesystem paths under `sdd/` — use `openspec/` for filesystem artifacts.
- **openspec**: Read and follow `skills/_shared/openspec-convention.md`. Update `tasks.md` with `[x]` marks.
- **hybrid**: Follow BOTH conventions.
- **none**: Return progress only.

## What to Do

### Step 1: Load Skills
Follow **Section A** from `skills/_shared/sdd-phase-common.md`.

### Step 2: Read Context
Before writing ANY code, read structured status, specs, design, tasks, existing code in affected files, and project conventions.

### Step 3: Implement Tasks
For each task: read task description, spec scenarios, design decisions, existing patterns. Write code. Mark complete.

### Step 4: Persist Progress
Save apply-progress. Update tasks artifact with [x] marks.

### Step 5: Return Summary
Return completed tasks, files changed, deviations, issues, remaining tasks, workload/PR boundary, and status.

## Rules
- ALWAYS read specs before implementing
- ALWAYS follow design decisions
- ALWAYS match existing code patterns
- STOP on blocked state
- If Strict TDD Mode is active, load strict-tdd.md and follow its cycle