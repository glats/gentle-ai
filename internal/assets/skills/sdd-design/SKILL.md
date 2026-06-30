---
name: sdd-design
description: "Create the SDD technical design and architecture approach. Trigger: orchestrator launches design for a change."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "2.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: If you loaded this skill via the `skill()` tool, you are the ORCHESTRATOR — STOP. Do NOT execute these instructions inline. Delegate to the dedicated `sdd-design` sub-agent.

## Executor Override
If you ARE the `sdd-design` sub-agent, continue. Do NOT delegate.

## Language Domain Contract
Generated technical artifacts default to English.

## Purpose
You create TECHNICAL DESIGN from proposals and specs.

## Execution and Persistence Contract
> Follow `skills/_shared/sdd-phase-common.md`.

- **engram**: Use `mem_get_observation` for Engram topic `sdd/{change-name}/proposal` (required) and `sdd/{change-name}/spec` (optional). Save with `mem_save(topic_key="sdd/{change-name}/design", ...)`. Do NOT read or write filesystem paths under `sdd/` — use `openspec/` for filesystem artifacts.
- **openspec**: Read and follow `skills/_shared/openspec-convention.md`.
- **hybrid**: Follow BOTH conventions.
- **none**: Return result only.

## What to Do

### Step 1: Load Skills
Follow Section A from `_shared/sdd-phase-common.md`.

### Step 2: Read the Codebase
Read entry points, module structure, existing patterns, dependencies.

### Step 3: Write design.md
Write: Technical Approach, Architecture Decisions, Data Flow, File Changes, Interfaces, Testing Strategy, Migration, Open Questions.

### Step 4: Persist Artifact
Save with `mem_save(topic_key="sdd/{change-name}/design")` or write file per mode.

### Step 5: Return Summary
Return change name, location, summary, key decisions, files affected.

## Rules
- ALWAYS read actual codebase before designing.
- Every decision MUST have rationale.
- Use project's ACTUAL patterns.
- Return envelope per `_shared/sdd-phase-common.md`.
