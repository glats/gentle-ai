---
name: sdd-propose
description: "Create an SDD change proposal with intent, scope, and approach. Trigger: orchestrator launches proposal work for a change."
disable-model-invocation: true
user-invocable: false
license: MIT
metadata:
  author: gentleman-programming
  version: "2.0"
  delegate_only: true
---

> **ORCHESTRATOR GATE**: If you loaded this skill via the `skill()` tool, you are
> the ORCHESTRATOR — STOP. Do NOT execute these instructions inline. Delegate to
> the dedicated `sdd-propose` sub-agent.

## Executor Override
If you ARE the `sdd-propose` sub-agent (NOT the orchestrator), continue. Do NOT delegate.

## Language Domain Contract
Generated technical artifacts default to English.

## Purpose
You are a sub-agent responsible for creating PROPOSALS from exploration analysis or direct user input.

## What You Receive
From the orchestrator: Change name, exploration analysis, artifact store mode.

## Execution and Persistence Contract
> Follow `skills/_shared/sdd-phase-common.md`.

- **engram**: Use `mem_get_observation` for Engram topic `sdd/{change-name}/explore` (optional) and `sdd-init/{project}` (optional). Save with `mem_save(topic_key="sdd/{change-name}/proposal", ...)`. Do NOT read or write filesystem paths under `sdd/` — use `openspec/` for filesystem artifacts.
- **openspec**: Read and follow `skills/_shared/openspec-convention.md`.
- **hybrid**: Follow BOTH conventions.
- **none**: Return result only.

## What to Do

### Step 1: Load Skills
Follow Section A from `_shared/sdd-phase-common.md`.

### Step 2: Create Change Directory
In openspec/hybrid mode, create `openspec/changes/{change-name}/proposal.md`.

### Step 3: Read Existing Specs
If openspec/specs/ has relevant specs, read them.

### Step 4: Write proposal.md
Write a structured proposal with: Intent, Scope, Capabilities, Approach, Affected Areas, Risks, Rollback Plan, Dependencies, Success Criteria.

### Step 5: Persist Artifact
Save with `mem_save(topic_key="sdd/{change-name}/proposal")` or write file per mode.

### Step 6: Return Summary
Return change name, location, summary, next step.

## Rules
- Every proposal MUST have rollback plan and success criteria.
- Fill Capabilities section as contract with sdd-spec.
- Size budget: under 450 words.
- Return envelope per `_shared/sdd-phase-common.md`.
