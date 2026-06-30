---
name: sdd-propose
description: "Create an SDD change proposal with intent, scope, and approach."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "2.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: If you loaded this skill via the `skill()` tool, you are the ORCHESTRATOR — STOP.

## Executor Override
If you ARE the `sdd-propose` sub-agent, continue. Do NOT delegate.

## Execution and Persistence Contract
> Follow `skills/_shared/sdd-phase-common.md`.

- **engram**: Use `mem_get_observation` for Engram topic `sdd/{change-name}/explore` (optional) and `sdd-init/{project}` (optional). Save with `mem_save(topic_key="sdd/{change-name}/proposal", ...)`. Do NOT read or write filesystem paths under `sdd/` — use `openspec/` for filesystem artifacts.
- **openspec**: Read and follow `skills/_shared/openspec-convention.md`.
- **hybrid**: Follow BOTH conventions.
- **none**: Return result only.

## Purpose
You create PROPOSALS from exploration analysis.

## Rules
- Every proposal MUST have rollback plan and success criteria.
- Return envelope per `_shared/sdd-phase-common.md`.
