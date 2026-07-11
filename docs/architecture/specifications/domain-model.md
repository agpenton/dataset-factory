# Domain Model Specification

| Field | Value |
|-------|-------|
| Status | Accepted |
| Version | 2.0 |
| Date | 2026-07-11 |
| Project | Dataset Factory |

---

# Purpose

This document defines the business domain of Dataset Factory.

The domain model represents business concepts rather than technical implementation details.

The domain is independent of:

- storage
- REST
- CLI
- AI providers
- plugins
- serialization
- databases

---

# Principles

The domain model SHALL:

- model business capabilities
- remain infrastructure independent
- expose immutable value objects whenever possible
- contain deterministic behavior
- define business invariants
- never perform I/O

---

# Domain Overview

```
Content

↓

Knowledge

↓

Dataset

↓

Quality

↓

Benchmark
```

Execution, pipelines and plugins orchestrate the business flow.

---

# Aggregates

The project contains ten aggregates.

| Aggregate | Responsibility |
|------------|----------------|
| Common | Shared value objects |
| Content | Imported content |
| Knowledge | Engineering knowledge |
| Dataset | Training datasets |
| Pipeline | Pipeline definition |
| Execution | Runtime execution |
| Recipe | Declarative workflows |
| Quality | Dataset evaluation |
| Benchmark | Model evaluation |
| Plugin | Extension metadata |

---

# Aggregate Dependencies

```
Common

↑

Content

↓

Knowledge

↓

Dataset

↓

Quality

↓

Benchmark
```

Execution, Pipeline and Recipe coordinate the process.

Plugins implement behavior.

---

# Package Layout

internal/domain/

common/

content/

knowledge/

dataset/

pipeline/

execution/

recipe/

quality/

benchmark/

plugin/

events/

---

# Aggregate Rules

Each aggregate:

- owns its invariants
- validates itself
- exposes deterministic behavior
- hides implementation details

No aggregate performs infrastructure work.

---

# Dependency Rules

Allowed

Pipeline

↓

Knowledge

↓

Dataset

Forbidden

Dataset

↓

Pipeline

Knowledge

↓

Execution

Content

↓

Plugin

Business dependencies must always point toward lower-level concepts.

---

# Domain Events

Business events belong to the domain.

Examples

DocumentImported

KnowledgeExtracted

DatasetGenerated

ReviewCompleted

BenchmarkExecuted

Events contain business meaning only.

---

# Identity

Every aggregate uses a deterministic identifier.

No random UUIDs.

Identifiers are generated from canonical business data.

---

# Validation

Every aggregate validates itself.

No invalid object may exist.

Validation failures are business errors.

---

# Immutability

Value Objects are immutable.

Entities expose controlled state transitions.

---

# Summary

The Dataset Factory domain models business capabilities rather than implementation details.

This structure minimizes coupling and provides long-term architectural stability.