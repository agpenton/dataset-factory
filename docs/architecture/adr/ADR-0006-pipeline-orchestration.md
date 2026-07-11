# ADR-0006 — Pipeline Orchestration

| Field | Value |
|-------|-------|
| Status | Accepted |
| Version | 1.0 |
| Date | 2026-07-11 |
| Project | Dataset Factory |
| Depends On | ADR-0001, ADR-0002, ADR-0003, ADR-0004, ADR-0005 |

---

# 1. Purpose

Define how Dataset Factory executes pipelines.

This ADR specifies:

- Pipeline execution
- Scheduling
- Worker pools
- Backpressure
- Cancellation
- Context propagation
- Checkpointing
- Retry orchestration
- Failure recovery
- Pipeline state

This ADR does not define pipeline stages or business logic.

---

# 2. Goals

The orchestration engine SHALL:

- execute deterministic pipelines
- maximize throughput
- minimize memory usage
- support streaming
- support checkpointing
- support resumability
- isolate failures
- expose complete observability
- support future distributed execution

---

# 3. Non-Goals

The orchestration engine SHALL NOT:

- understand documents
- understand knowledge assets
- implement crawlers
- implement AI providers
- contain business logic

---

# 4. Execution Model

The execution model is event-driven.

```
Pipeline

↓

Scheduler

↓

Worker Pool

↓

Stage

↓

Output Queue

↓

Next Stage
```

Each stage consumes one stream and produces another.

---

# 5. Pipeline Definition

A pipeline is a Directed Acyclic Graph (DAG).

Example

```
Discover

↓

Fetch

↓

Normalize

↓

Extract

↓

Generate

↓

Review

↓

Export
```

Cycles are forbidden.

---

# 6. Pipeline State Machine

Every pipeline follows the same lifecycle.

```
Created

↓

Validated

↓

Initialized

↓

Running

↓

Paused

↓

Resuming

↓

Completed

↓

Archived
```

Failure transitions

```
Running

↓

Failed

↓

Retrying

↓

Running

or

↓

Cancelled
```

---

# 7. Stage State Machine

Each stage has an independent lifecycle.

```
Created

↓

Waiting

↓

Running

↓

Checkpoint

↓

Completed
```

Failure

```
Running

↓

Failed

↓

Retry

↓

Running
```

---

# 8. Scheduler

Responsibilities

- start stages
- monitor dependencies
- allocate workers
- enforce execution order
- coordinate retries

The scheduler never performs business logic.

---

# 9. Worker Pools

Each stage owns its own worker pool.

Example

```
Fetch

8 Workers

Normalize

16 Workers

Generate

4 Workers

Export

2 Workers
```

Worker count is configurable.

---

# 10. Context Propagation

Every pipeline owns one Context.

Every stage receives a derived Context.

```
Pipeline Context

↓

Stage Context

↓

Worker Context
```

Cancellation propagates downward.

---

# 11. Cancellation

Cancellation must be cooperative.

Stages must periodically check Context.

No stage may ignore cancellation.

---

# 12. Queues

Communication occurs through bounded queues.

```
Stage A

↓

Queue

↓

Stage B
```

Queues prevent unbounded memory growth.

---

# 13. Backpressure

When a queue becomes full

```
Producer

↓

Wait

↓

Consumer

↓

Continue
```

No stage may continue producing indefinitely.

---

# 14. Checkpointing

Every stage periodically creates checkpoints.

Checkpoint contains

```yaml
pipeline:

stage:

worker:

last_processed:

processed:

remaining:

timestamp:
```

Checkpoint interval is configurable.

---

# 15. Resume

Resume begins at the last successful checkpoint.

Already completed work SHALL NOT execute again.

---

# 16. Retry Policy

Retry is configurable.

Default

```
Attempts

3

Backoff

Exponential

Jitter

Disabled
```

Retries apply only to recoverable failures.

---

# 17. Failure Classification

Recoverable

- timeout
- temporary network
- rate limiting
- transient AI failure

Non-Recoverable

- invalid configuration
- unsupported format
- schema validation
- programming error

---

# 18. Dead Letter Queue

Documents exceeding retry limits move to a Dead Letter Queue.

Example

```
Pipeline

↓

Failure

↓

Retry

↓

Retry

↓

Retry

↓

Dead Letter Queue
```

Dead Letter Queue entries remain inspectable.

---

# 19. Event Bus

Every pipeline event is published.

Events include

```
PipelineStarted

PipelineCompleted

StageStarted

StageCompleted

WorkerStarted

CheckpointCreated

Retry

Failure

Warning

Cancelled
```

Future subscribers may consume events.

---

# 20. Metrics

Pipeline metrics

```
Active Pipelines

Pipeline Duration

Stage Duration

Worker Utilization

Queue Depth

Retries

Failures

Throughput

Memory Usage
```

---

# 21. Logging

Structured logging only.

Required fields

```
Pipeline

Stage

Worker

Event

Duration

Status

DocumentID

Timestamp
```

---

# 22. Pipeline Storage

Execution state is external.

Supported implementations

- SQLite
- PostgreSQL (future)
- Cloud SQL (future)

The orchestrator is stateless.

---

# 23. Configuration

Example

```yaml
pipeline:

workers:

fetch: 8

normalize: 8

extract: 4

generate: 2

review: 2

checkpoint:

interval: 100

retry:

attempts: 3

backoff: exponential

deterministic: true
```

---

# 24. Parallelism

Only independent work executes in parallel.

Example

```
Document A

↓

Worker 1

Document B

↓

Worker 2

Document C

↓

Worker 3
```

Ordering is restored before export.

---

# 25. Determinism

Deterministic mode guarantees

- stable ordering
- stable IDs
- stable scheduling
- stable checkpoints

Thread scheduling must never affect output.

---

# 26. Resource Limits

Every stage may define

```
CPU

Memory

Concurrency

Timeout

Retries
```

Limits prevent resource exhaustion.

---

# 27. Observability

Native support

- OpenTelemetry
- Prometheus
- Grafana
- Tempo
- Loki

Instrumentation is mandatory.

---

# 28. Kubernetes Execution (Future)

Future releases may execute pipelines as Kubernetes Jobs.

```
Scheduler

↓

Job

↓

Worker Pods

↓

Shared State
```

The orchestration model must remain unchanged.

---

# 29. Distributed Execution (Future)

Future architecture

```
Coordinator

↓

Queue

↓

Worker Nodes

↓

Shared Storage
```

Current interfaces must support this evolution.

---

# 30. Decisions

Accepted

- Event-driven orchestration
- Bounded queues
- Worker pools
- Stage isolation
- Context propagation
- Automatic checkpointing
- Cooperative cancellation
- Dead Letter Queue
- External state storage
- Deterministic scheduling

---

# 31. Consequences

Advantages

- Highly scalable
- Resumable
- Fault tolerant
- Observable
- Testable
- Deterministic

Trade-offs

- Increased orchestration complexity
- Additional metadata
- More persistent state
- Scheduler implementation effort

---

# 32. Summary

The orchestration engine is responsible solely for executing pipelines.

Business logic remains inside plugins.

Execution is deterministic, event-driven, observable, resumable, and designed to evolve from a single-process runtime to a distributed execution engine without changing the public pipeline model.