# Package Standards

Version: 1.0

Status: Accepted

---

# Purpose

This document defines mandatory rules for every package in Dataset Factory.

These rules ensure consistency, maintainability, and long-term architectural stability.

---

# Single Responsibility

Each package owns exactly one business capability.

Examples

GOOD

```
knowledge
dataset
pipeline
quality
```

BAD

```
utils
helpers
common2
misc
```

---

# Public Surface

Minimize exported symbols.

Only export APIs intended for use outside the package.

Everything else remains private.

---

# Package Size

Target:

- 5–15 files

Maximum:

- 25 files

If exceeded, split the package.

---

# File Size

Target:

- 150–300 LOC

Maximum:

- 500 LOC

Split large files.

---

# Function Size

Target:

20–40 LOC

Maximum:

80 LOC

Large functions indicate missing abstractions.

---

# Constructor Rules

Every entity has a constructor.

Never expose invalid zero values.

GOOD

NewDocument()

BAD

Document{}

---

# Validation

Validation belongs to the aggregate.

Never validate objects in multiple places.

---

# Immutability

Prefer immutable value objects.

Mutable entities must expose explicit state transitions.

---

# Error Handling

Never panic.

Return typed errors.

Wrap external errors.

Never lose context.

---

# Logging

Domain packages never log.

Infrastructure logs.

---

# Configuration

Domain packages never read configuration.

Configuration is injected.

---

# Context

Only infrastructure and use cases receive context.Context.

Never store Context in structs.

Never pass nil Context.

---

# Concurrency

The domain is single-threaded.

Concurrency belongs to the execution engine.

---

# Dependencies

Dependencies point inward.

```
Infrastructure

↓

Application

↓

Domain
```

Never reverse.

---

# Testing

Every package requires:

- Unit tests
- Validation tests
- Benchmarks (when applicable)

Coverage target:

95%

---

# Documentation

Every package contains:

doc.go

README.md (optional for complex packages)

Every exported symbol has GoDoc.

---

# Performance

Avoid allocations.

Avoid reflection.

Avoid interface{}.

Prefer concrete types.

Benchmark critical paths.

---

# External Dependencies

Every dependency requires justification.

Prefer the Go standard library.

No dependency without a documented reason.

---

# Breaking Changes

Public APIs follow Semantic Versioning.

Internal APIs may evolve.

---

# Summary

Packages are small, cohesive, deterministic and independently testable.

Package boundaries are architectural boundaries.