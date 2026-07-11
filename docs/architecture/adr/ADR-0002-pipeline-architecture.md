# ADR-0002 — Pipeline Architecture

| Field | Value |
|-------|-------|
| Status | Accepted |
| Version | 1.0 |
| Date | 2026-07-11 |
| Project | Dataset Factory |
| Depends On | ADR-0001 |

---

# 1. Purpose

Define the execution model of Dataset Factory.

This ADR specifies:

- pipeline architecture
- stage lifecycle
- execution model
- concurrency
- streaming
- checkpointing
- retries
- failure handling
- observability

This document is the blueprint for the pipeline engine.

---

# 2. Goals

The pipeline SHALL:

- be deterministic
- support streaming
- support millions of documents
- support resumability
- support checkpointing
- support distributed execution (future)
- support plugins
- support incremental execution

---

# 3. Non-Goals

The pipeline SHALL NOT:

- contain business logic
- implement AI providers
- understand document formats
- perform crawling
- perform exporting

The pipeline only orchestrates execution.

---

# 4. Pipeline Philosophy

The pipeline is a Directed Acyclic Graph (DAG).

Each node performs one responsibility.

```

```
Input

↓

Stage A

↓

Stage B

↓

Stage C

↓

Output
```

No cycles are allowed.

---

# 5. Pipeline Components

```
Pipeline

├── Runner

├── Scheduler

├── Executor

├── Checkpointer

├── Retry Manager

├── Event Bus

├── Metrics

├── Logger

└── State Store
```

Each component owns exactly one responsibility.

---

# 6. Stage Model

Every processing unit is a Stage.

Stages are isolated.

Stages communicate only through contracts.

```
Input

↓

Stage

↓

Output
```

Stages never call each other directly.

---

# 7. Pipeline Stages

Initial implementation:

```
Discovery

↓

Fetch

↓

Normalize

↓

Segment

↓

Extract

↓

Knowledge Assets

↓

Generate Prompts

↓

Generate Answers

↓

Review

↓

Score

↓

Deduplicate

↓

Balance

↓

Export
```

Each stage becomes an independent package.

---

# 8. Stage Lifecycle

Every stage follows the same lifecycle.

```
Initialize

↓

Validate Configuration

↓

Start

↓

Process

↓

Checkpoint

↓

Complete

↓

Shutdown
```

---

# 9. Stage Contract

Every stage SHALL:

- receive one input type
- produce one output type
- never mutate previous stages
- support cancellation
- emit metrics
- emit structured logs

---

# 10. Streaming Model

The pipeline processes streams.

Never collections.

```
Document

↓

Document

↓

Document

↓

Document
```

Streaming minimizes memory usage.

---

# 11. Concurrency Model

Stages process independent items concurrently.

Concurrency is configurable.

Example:

```
Workers

Worker 1

Worker 2

Worker 3

Worker N
```

Ordering is preserved when deterministic mode is enabled.

---

# 12. Determinism

Deterministic mode guarantees:

Same Input

+

Same Configuration

↓

Same Output

Requirements:

- fixed ordering
- deterministic hashing
- deterministic IDs
- deterministic scheduling

---

# 13. Checkpointing

Each stage periodically creates checkpoints.

Checkpoint contains:

- stage
- progress
- offsets
- statistics
- hashes

Checkpoint format:

```yaml
stage:

processed:

remaining:

last_id:

timestamp:
```

---

# 14. Resume

Pipeline execution must resume from the latest checkpoint.

No stage should repeat completed work.

---

# 15. Retry Policy

Failures are classified.

Recoverable

↓

Retry

Non Recoverable

↓

Stop

Default retry policy:

- exponential backoff
- configurable attempts
- jitter disabled by default

---

# 16. Error Handling

Errors are categorized.

```
Configuration

Validation

Network

Parser

AI

Storage

Internal
```

Every error includes:

- stage
- source
- document
- severity

---

# 17. Event Bus

Every stage emits events.

Examples:

```
PipelineStarted

StageStarted

DocumentProcessed

CheckpointCreated

Retry

Warning

Failure

Completed
```

Future integrations may consume these events.

---

# 18. Metrics

Every stage exports:

Counters

Timers

Histograms

Gauges

Examples:

```
documents_processed

processing_latency

retry_count

errors_total

memory_usage

queue_size
```

---

# 19. Logging

Structured logging only.

Required fields:

```
timestamp

pipeline

stage

worker

document

duration

status
```

No printf logging.

---

# 20. State Store

Pipeline state is external.

Supported implementations:

- SQLite
- PostgreSQL (future)
- Cloud Storage (future)

The pipeline never stores execution state in memory only.

---

# 21. Configuration

Every pipeline is configured declaratively.

Example:

```yaml
pipeline:

workers: 8

checkpoint_interval: 100

retry_attempts: 3

deterministic: true
```

---

# 22. Plugin Integration

Pipeline discovers plugins.

Pipeline does not instantiate implementations directly.

Every stage depends only on interfaces.

---

# 23. Observability

Future integrations:

- OpenTelemetry
- Prometheus
- Grafana
- Tempo
- Loki

Pipeline must expose instrumentation hooks.

---

# 24. Scalability

Pipeline should support:

- millions of documents
- hundreds of workers
- resumable execution
- distributed execution

No architecture changes should be required.

---

# 25. Testing Strategy

Every stage must support:

- unit tests
- integration tests
- benchmark tests
- deterministic replay tests

---

# 26. Future Evolution

Future capabilities:

- distributed scheduler
- Kubernetes execution
- remote workers
- cloud execution
- DAG optimization
- pipeline visualization

These must not require breaking changes.

---

# 27. Decisions

Accepted decisions:

- Streaming execution
- DAG pipeline
- Stateless stages
- External state store
- Structured events
- Plugin architecture
- Deterministic execution
- Configurable concurrency
- Automatic checkpointing

---

# 28. Consequences

Advantages

- reproducible
- scalable
- observable
- fault tolerant
- resumable
- modular

Trade-offs

- higher implementation complexity
- more interfaces
- more metadata
- additional state management

---

# 29. Summary

The Dataset Factory pipeline is a deterministic streaming execution engine.

Business logic belongs inside plugins.

The pipeline only orchestrates execution.

Every stage remains independent, observable, resumable, and replaceable.