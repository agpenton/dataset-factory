# Dataset Factory Roadmap

This roadmap defines the implementation order of Dataset Factory.

The order is intentional.

Do not skip milestones.

Each milestone must:

- compile
- be tested
- be documented
- be merged before the next begins

---

# Phase 0

Repository Bootstrap

Goal

Build a compiling repository.

Deliverables

- go.mod
- Makefile
- CI
- lint
- formatting
- release automation

---

# Phase 1

Artifact System

Goal

Create the universal data model.

Deliverables

- Artifact interface
- Artifact IDs
- Artifact metadata
- Validation
- Serialization

Artifacts

- Document
- Chunk
- Knowledge
- Prompt
- Conversation
- Dataset

Acceptance

Artifacts compile.

100% unit tests.

---

# Phase 2

Operator System

Goal

Create the execution abstraction.

Deliverables

- Operator interface
- Metadata
- Capabilities
- Validation

Acceptance

Dummy operator executes.

---

# Phase 3

Workflow Model

Goal

Represent DAGs.

Deliverables

- Workflow
- Edge
- Graph
- Validation

Acceptance

Workflow compiler validates DAG.

---

# Phase 4

Planner

Goal

Compile workflows into execution plans.

Deliverables

- Validation
- Dependency resolution
- Type checking
- Execution plan

Acceptance

Workflow → Execution Plan

---

# Phase 5

Runtime

Goal

Execute plans.

Deliverables

- Scheduler
- Workers
- Streams
- Retry
- Checkpoints

Acceptance

Workflow executes locally.

---

# Phase 6

Registry

Goal

Dynamic operator discovery.

Deliverables

- Registration
- Lookup
- Metadata

Acceptance

Operators discoverable.

---

# Phase 7

Filesystem Operator

Goal

Import files.

Acceptance

Markdown files become Document artifacts.

---

# Phase 8

Markdown Operator

Goal

Normalize Markdown.

Acceptance

Markdown → normalized Document.

---

# Phase 9

Chunk Operator

Goal

Split documents.

Acceptance

Document → Chunks.

---

# Phase 10

Knowledge Operator

Goal

Extract engineering knowledge.

Acceptance

Chunks → Knowledge.

---

# Phase 11

Prompt Operator

Goal

Generate prompts.

Acceptance

Knowledge → Prompt.

---

# Phase 12

LLM Operator

Goal

Generate answers.

Acceptance

Prompt → Conversation.

---

# Phase 13

Review Operator

Goal

Score generated samples.

Acceptance

Conversation → Reviewed Conversation.

---

# Phase 14

Exporter

Goal

Export datasets.

Acceptance

JSONL
ShareGPT
ChatML

---

# Phase 15

Benchmarks

Goal

Evaluate datasets.

Acceptance

Quality report generated.

---

# Phase 16

Production

Goal

Production-ready release.

Acceptance

CI/CD
Coverage >95%
Benchmarks
Documentation
Release