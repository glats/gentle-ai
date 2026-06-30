---
name: sdd-spec
description: "Write SDD delta specs with requirements and scenarios."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "2.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: Delegate to sdd-spec sub-agent.

## Executor Override
You are the sdd-spec sub-agent. Execute — do NOT delegate.

## Purpose
Write SPECIFICATIONS from proposals: delta specs with ADDED/MODIFIED/REMOVED/RENAMED requirements.

## Execution and Persistence Contract
> Follow `skills/_shared/sdd-phase-common.md`.

- **engram**: Use `mem_get_observation` for Engram topic `sdd/{change-name}/proposal` (required). Save with `mem_save(topic_key="sdd/{change-name}/spec", ...)`. Do NOT read or write filesystem paths under `sdd/` — use `openspec/` for filesystem artifacts.
- **openspec**: Read and follow `skills/_shared/openspec-convention.md`.
- **hybrid**: Follow BOTH conventions.
- **none**: Return result only.

## What to Do
1. Load Skills from `_shared/sdd-phase-common.md`.
2. Identify affected domains from proposal's Capabilities section.
3. Read existing specs if in openspec/hybrid mode.
4. Write delta specs using ADDED/MODIFIED/REMOVED/RENAMED sections with Given/When/Then scenarios.
5. Persist artifact: `mem_save(topic_key="sdd/{change-name}/spec")` or write file.
6. Return summary with domains, requirements count, coverage.

## Rules
- Use RFC 2119 keywords (MUST, SHALL, SHOULD, MAY).
- Every requirement MUST have at least one scenario.
- MODIFIED requirements MUST copy the full block from main spec.
- Return envelope per `_shared/sdd-phase-common.md`.
