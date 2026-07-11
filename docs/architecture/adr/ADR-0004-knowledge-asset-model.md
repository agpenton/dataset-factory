# ADR-0004 — Knowledge Asset Model

| Field | Value |
|-------|-------|
| Status | Accepted |
| Version | 1.0 |
| Date | 2026-07-11 |
| Project | Dataset Factory |
| Depends On | ADR-0001, ADR-0002, ADR-0003 |

---

# 1. Purpose

Define the canonical data model used throughout Dataset Factory.

Every component in the pipeline SHALL consume and produce a `KnowledgeAsset`.

This ADR establishes the project's canonical representation of engineering knowledge.

---

# 2. Problem

Engineering knowledge exists in many forms:

- Documentation
- GitHub Issues
- Pull Requests
- RFCs
- ADRs
- Blog Posts
- Books
- Videos
- Conference Talks

Every source uses different formats.

Without a canonical model, every pipeline stage must understand every source.

This creates tight coupling and unnecessary complexity.

---

# 3. Decision

Dataset Factory SHALL normalize every source into a single canonical object called:

```
KnowledgeAsset
```

Every stage after extraction SHALL operate exclusively on Knowledge Assets.

Raw documents SHALL NOT be processed directly after normalization.

---

# 4. Philosophy

Documents are temporary.

Knowledge is permanent.

Knowledge Assets represent reusable engineering knowledge independent of its original source.

---

# 5. Processing Flow

```
Raw Document

↓

Normalized Document

↓

Semantic Sections

↓

Knowledge Assets

↓

Prompt Generation

↓

Answer Generation

↓

Review

↓

Export
```

Knowledge Assets are the contract between extraction and generation.

---

# 6. Granularity

One Knowledge Asset SHALL represent exactly one coherent engineering concept.

Examples:

Good

```
Terraform Remote State

IRSA

ApplicationSets

Blue/Green Deployments

Canary Releases

OpenTelemetry Collector
```

Bad

```
Terraform

AWS

Kubernetes
```

Large documents become multiple Knowledge Assets.

---

# 7. Canonical Model

```yaml
id:

version:

type:

title:

summary:

description:

source:

metadata:

taxonomy:

content:

relationships:

quality:

embeddings:
```

---

# 8. Identity

Every Knowledge Asset SHALL have a globally unique identifier.

```
knowledge://aws/iam/irsa

knowledge://terraform/backend/s3

knowledge://argocd/applicationsets
```

IDs must remain stable across pipeline executions.

---

# 9. Asset Types

Supported asset types:

```
Concept

Procedure

Architecture

Pattern

AntiPattern

BestPractice

Decision

Tradeoff

Incident

Runbook

Tutorial

FAQ

Code

Diagram

InterviewQuestion

Evaluation
```

Future types may be added.

---

# 10. Source

Every asset SHALL preserve provenance.

```yaml
source:

provider:

url:

title:

author:

published:

license:

document_id:

section:

chunk:
```

Every generated dataset entry must be traceable.

---

# 11. Metadata

```yaml
metadata:

language:

domain:

subcategory:

difficulty:

reading_time:

estimated_tokens:

updated:

created:
```

Metadata supports filtering and balancing.

---

# 12. Taxonomy

Every asset SHALL be categorized.

Example:

```yaml
taxonomy:

platform:

cloud:

technology:

topic:

subtopic:

tags:
```

Example

```
Platform Engineering

↓

GitOps

↓

ArgoCD

↓

ApplicationSets
```

---

# 13. Concepts

Concepts define the primary engineering subjects.

Example

```yaml
concepts:

- GitOps

- Continuous Delivery

- Kubernetes
```

---

# 14. Technologies

Explicit technology references.

Example

```yaml
technologies:

- Kubernetes

- Terraform

- AWS

- Helm
```

---

# 15. Patterns

Architectural patterns extracted from the document.

Examples

```
Hub and Spoke

Cell Architecture

GitOps

Sidecar

Controller Pattern

Event Sourcing
```

---

# 16. Anti-Patterns

Examples

```
Shared Production Account

Manual Infrastructure Changes

Long-Lived Credentials

Snowflake Servers
```

Anti-patterns improve troubleshooting datasets.

---

# 17. Best Practices

Example

```yaml
best_practices:

- Enable Versioning

- Least Privilege

- Immutable Infrastructure
```

---

# 18. Trade-offs

Trade-offs distinguish expert knowledge.

Example

```yaml
tradeoffs:

option_a:

option_b:

advantages:

disadvantages:

recommendation:
```

Example

```
FluxCD

vs

ArgoCD
```

---

# 19. Decisions

Engineering decisions extracted from source material.

Example

```yaml
decision:

problem:

selected_option:

reasoning:

consequences:
```

This field is essential for architecture datasets.

---

# 20. Procedures

Step-by-step operational knowledge.

Example

```yaml
procedure:

preconditions:

steps:

validation:

rollback:
```

Useful for operational runbooks.

---

# 21. Code

Code snippets are stored separately.

```yaml
code:

language:

purpose:

snippet:

references:
```

Large code samples are referenced, not embedded.

---

# 22. Diagrams

Architectural diagrams.

```yaml
diagram:

type:

description:

source:
```

Diagram extraction is optional.

---

# 23. Relationships

Knowledge Assets form a graph.

Relationships include:

```
DependsOn

Implements

Uses

References

Replaces

Extends

RelatedTo
```

Example

```
IRSA

↓

DependsOn

↓

OIDC Provider
```

---

# 24. Quality

Every asset carries quality metadata.

```yaml
quality:

confidence:

completeness:

reviewed:

review_score:

hallucination_risk:
```

Assets below minimum quality thresholds may be excluded.

---

# 25. Embeddings

Embeddings are optional.

Purpose:

- semantic search
- deduplication
- clustering
- balancing

The pipeline SHALL NOT depend on embeddings.

---

# 26. Serialization

Supported formats:

```
JSON

YAML

JSONL

Parquet
```

JSON is the canonical format.

---

# 27. Immutability

Knowledge Assets SHALL be immutable.

Pipeline stages produce new versions.

Original assets remain unchanged.

---

# 28. Versioning

Assets are versioned.

```
v1

v2

v3
```

Version history enables reproducibility.

---

# 29. Validation

Every asset must pass validation.

Required fields:

```
id

type

title

summary

source

concepts

quality
```

Validation failures stop the pipeline.

---

# 30. Knowledge Graph

Knowledge Assets collectively form a directed graph.

```
IRSA

↓

OIDC

↓

IAM

↓

AWS Organizations

↓

Landing Zone
```

Future capabilities:

- graph traversal
- recommendation
- gap analysis
- automatic benchmark generation

---

# 31. Pipeline Contract

All pipeline stages SHALL consume and produce:

```
KnowledgeAsset
```

No stage shall depend on raw documents after extraction.

---

# 32. Future Extensions

Future fields may include:

```
Images

Audio

Video

Benchmarks

Synthetic Variants

Evaluation Results

Human Feedback

Preference Labels
```

Backward compatibility must be preserved.

---

# 33. Decisions

Accepted decisions:

- Canonical Knowledge Asset model
- Immutable assets
- Stable identifiers
- Provenance required
- Explicit relationships
- Quality metadata
- Versioned assets
- JSON as canonical format

---

# 34. Consequences

Advantages

- Single canonical model
- Loose coupling
- Easier testing
- Better deduplication
- Explainability
- Provenance
- Future knowledge graph support

Trade-offs

- Larger metadata footprint
- More normalization effort
- Additional validation logic

---

# 35. Summary

Knowledge Assets are the fundamental unit of Dataset Factory.

Every pipeline stage, plugin, exporter, reviewer and benchmark operates on this canonical model.

By separating engineering knowledge from document structure, Dataset Factory becomes a knowledge engineering platform rather than a document processing system.