---
name: sdd-explore
description: "Explore SDD ideas before committing to a change."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "2.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: Delegate to sdd-explore sub-agent.

## Executor Override
You are the sdd-explore sub-agent. Execute — do NOT delegate.

## Purpose
EXPLORE the codebase, think through problems, compare approaches, return structured analysis.

## Execution and Persistence Contract
> Follow `skills/_shared/sdd-phase-common.md`.

- **engram**: Use `mem_get_observation` for Engram topic `sdd-init/{project}` for project context (optional). Save with `mem_save(topic_key="sdd/{change-name}/explore", ...)`. Do NOT read or write filesystem paths under `sdd/` — use `openspec/` for filesystem artifacts.
- **openspec**: Read and follow `skills/_shared/openspec-convention.md`.
- **hybrid**: Follow BOTH conventions.
- **none**: Return result only.

### Retrieving Context
> Follow Section B from `_shared/sdd-phase-common.md`.

- **engram**: Use `mem_search` for Engram topic key `sdd-init/{project}` and optionally any `sdd/*` topic key. Do NOT grep or glob the filesystem for `sdd/` — use `openspec/` for filesystem artifact searches.
- **openspec**: Read `openspec/config.yaml` and `openspec/specs/`.
- **none**: Use orchestrator-passed context.

**Filesystem path convention**: The SDD artifact directory is `openspec/`. Do NOT use `sdd/`, `.sdd/`, or `sdds/` as filesystem paths — these do not exist. Engram topic keys use the `sdd/` prefix for memory organization only.

## What to Do
1. Load Skills.
2. Understand the request.
3. Investigate the codebase.
4. Analyze options with pros/cons/complexity table.
5. Persist artifact: `mem_save(topic_key="sdd/{change-name}/explore")`.
6. Return structured analysis with Current State, Affected Areas, Approaches, Recommendation, Risks.

## Rules
- ONLY create exploration.md inside change folder.
- DO NOT modify existing code.
- ALWAYS read real code, never guess.
- Return envelope per `_shared/sdd-phase-common.md`.
