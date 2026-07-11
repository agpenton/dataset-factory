# ADR-0001 — Project Vision & Architecture

| Field | Value |
|-------|-------|
| Status | Accepted |
| Version | 1.0 |
| Date | 2026-07-11 |
| Project | Dataset Factory |
| Authors | Asdrubal Gonzalez Penton |
| Decision Type | Foundational |

---

# 1. Abstract

Dataset Factory is a production-grade knowledge engineering platform that transforms heterogeneous engineering knowledge into high-quality supervised datasets for Large Language Model (LLM) fine-tuning.

The project is designed to be deterministic, reproducible, extensible, source-agnostic and model-agnostic.

This document defines the architectural principles governing every future implementation decision.

---

# 2. Problem Statement

Current dataset generation tools focus primarily on crawling documentation and splitting text into chunks.

They generally do not:

- extract engineering knowledge
- capture architecture decisions
- preserve trade-offs
- generate diverse instruction datasets
- verify factual correctness
- remove semantic duplicates
- benchmark generated datasets
- provide deterministic outputs

Consequently, many fine-tuned models behave like documentation search engines rather than expert assistants.

Dataset Factory addresses these limitations.

---

# 3. Vision

Dataset Factory shall become the reference open-source platform for engineering dataset generation.

Instead of treating documents as the primary asset, Dataset Factory treats **knowledge** as the primary asset.

The system transforms engineering knowledge into reusable, benchmarked, explainable instruction datasets.

---

# 4. Mission

Enable organizations to continuously transform engineering knowledge into production-quality datasets through deterministic and reproducible pipelines.

---

# 5. Scope

Dataset Factory is responsible for:

- Knowledge discovery
- Content acquisition
- Normalization
- Semantic segmentation
- Knowledge extraction
- Instruction generation
- Answer generation
- AI review
- Quality scoring
- Deduplication
- Dataset balancing
- Dataset exporting
- Dataset benchmarking

Dataset Factory is **not** responsible for:

- Model training
- Model serving
- Vector databases
- RAG systems
- Inference APIs

---

# 6. Core Philosophy

The project follows one fundamental rule.

> Knowledge is the primary artifact.

Documents are temporary.

Knowledge is permanent.

---

# 7. Engineering Principles

## 7.1 Deterministic by Default

Running the same pipeline with identical configuration shall produce identical output.

Randomness must always be optional.

---

## 7.2 Reproducibility

Every generated example must be reproducible.

Every dataset must be rebuildable.

---

## 7.3 Streaming

Every pipeline stage should process streams whenever possible.

The system must support millions of documents.

---

## 7.4 Modular

Every stage must expose interfaces.

Implementations must be replaceable.

---

## 7.5 Explainability

Every generated example must contain provenance.

Example:

```
Source

↓

Section

↓

Knowledge Asset

↓

Prompt

↓

Answer

↓

Reviewer

↓

Dataset
```

---

## 7.6 AI-Assisted

AI improves quality.

AI never replaces deterministic processing.

---

## 7.7 Open Standards

Avoid proprietary formats.

Prefer:

- Markdown
- YAML
- JSON
- JSONL
- Parquet
- SQLite

---

# 8. Knowledge-Centric Architecture

Traditional systems process:

```
Document

↓

Chunks

↓

LLM
```

Dataset Factory processes:

```
Document

↓

Knowledge Assets

↓

Instructions

↓

Answers

↓

Datasets
```

Knowledge Assets are the core abstraction of the system.

---

# 9. Knowledge Asset

Every document is converted into one or more Knowledge Assets.

A Knowledge Asset represents one coherent engineering concept.

Example:

```yaml
id:

title:

summary:

concepts:

technologies:

patterns:

anti_patterns:

tradeoffs:

best_practices:

procedures:

references:

difficulty:

tags:
```

Every downstream stage consumes Knowledge Assets instead of raw documents.

---

# 10. Processing Pipeline

```
Discovery

↓

Fetch

↓

Normalize

↓

Segment

↓

Knowledge Extraction

↓

Knowledge Assets

↓

Instruction Generation

↓

Answer Generation

↓

AI Review

↓

Quality Scoring

↓

Deduplication

↓

Balancing

↓

Export
```

Every stage has one responsibility.

---

# 11. Supported Knowledge Sources

Documentation

Markdown

HTML

PDF

GitHub

GitHub Issues

GitHub Discussions

GitHub Pull Requests

Architecture Decision Records

RFCs

KEPs

Blogs

Books (metadata only)

Videos (future)

Audio (future)

---

# 12. Supported Domains

Platform Engineering

Cloud Engineering

DevOps

Kubernetes

GitOps

Terraform

OpenTofu

Helm

FluxCD

Argo CD

Argo Rollouts

Flagger

AWS

Azure

Google Cloud

Linux

Networking

Security

Observability

Go

Python

TypeScript

Rust

Agentic AI

LLM Infrastructure

System Design

Software Architecture

---

# 13. Dataset Types

Instruction Following

Question Answer

Architecture Review

Troubleshooting

Incident Response

Design

Code Generation

Code Review

Migration

Optimization

Security Review

Interview Questions

Evaluation Datasets

Preference Datasets

---

# 14. Quality Model

Every generated example shall be scored.

Minimum quality dimensions:

- Correctness
- Completeness
- Consistency
- Originality
- Groundedness
- Hallucination Risk
- Difficulty
- Readability

---

# 15. Plugin Architecture

Everything is a plugin.

Including:

- Crawlers
- Parsers
- Reviewers
- Exporters
- Benchmarks
- AI Providers

No implementation should be hardcoded.

---

# 16. Technology Decisions

Implementation Language

Go

Reason:

- static binaries
- concurrency
- streaming
- performance
- portability

Python may be used only when integration with AI tooling is required.

---

# 17. Configuration Philosophy

Configuration must be declarative.

Prefer YAML.

Configuration must be versionable.

Configuration must be deterministic.

---

# 18. Observability

Every pipeline stage shall emit:

- logs
- metrics
- tracing

Future integrations:

- OpenTelemetry
- Prometheus
- Grafana

---

# 19. Security

No secrets stored in configuration.

Support:

- environment variables
- secret managers
- encrypted configuration

---

# 20. Performance Goals

The architecture should support:

- millions of documents
- hundreds of concurrent crawlers
- streaming processing
- resumable pipelines
- distributed execution (future)

---

# 21. Success Metrics

Dataset Factory succeeds when it can automatically produce datasets that:

- are reproducible
- are explainable
- outperform raw documentation datasets
- require minimal manual review
- are benchmarked before release

---

# 22. Risks

Large implementation scope.

High maintenance cost.

AI model quality dependency.

Rapid evolution of LLM ecosystem.

---

# 23. Future ADRs

ADR-0002 Pipeline Architecture

ADR-0003 Plugin Architecture

ADR-0004 Knowledge Asset Model

ADR-0005 Document Model

ADR-0006 Pipeline Orchestration

ADR-0007 Storage

ADR-0008 AI Review

ADR-0009 Quality Scoring

ADR-0010 Deduplication

ADR-0011 Exporters

ADR-0012 CLI

ADR-0013 REST API

ADR-0014 Security

ADR-0015 Benchmark Framework

ADR-0016 Release Strategy

---

# 24. Conclusion

Dataset Factory is not a crawler.

It is not a document processor.

It is not a training framework.

Dataset Factory is a deterministic knowledge engineering platform whose objective is to continuously transform engineering knowledge into production-grade datasets for LLM fine-tuning.