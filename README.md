# Dataset Factory

> Transform engineering knowledge into production-grade datasets for LLM fine-tuning.

---

## Vision

Dataset Factory is a production-grade knowledge engineering platform that transforms heterogeneous engineering knowledge into high-quality supervised datasets suitable for modern Large Language Model (LLM) training.

Unlike traditional crawlers or document chunkers, Dataset Factory extracts engineering knowledge, architecture decisions, best practices, trade-offs, troubleshooting procedures, and operational experience, then converts them into reproducible instruction datasets.

The long-term objective is to become the reference open-source platform for dataset engineering.

---

# Goals

Dataset Factory SHALL:

- Discover engineering knowledge from multiple sources.
- Normalize heterogeneous content.
- Extract structured knowledge.
- Generate high-quality instruction datasets.
- Review generated examples using AI.
- Score dataset quality.
- Remove semantic duplicates.
- Balance datasets.
- Export multiple training formats.
- Benchmark resulting datasets and models.

---

# Supported Domains

- Platform Engineering
- Kubernetes
- GitOps
- Terraform
- Helm
- FluxCD
- Argo CD
- Argo Rollouts
- Flagger
- AWS
- Azure
- Google Cloud
- Linux
- Networking
- Security
- Observability
- CI/CD
- Go
- Python
- TypeScript
- Agentic AI
- AI Infrastructure
- System Design

---

# High-Level Architecture

```
Knowledge Sources
        │
        ▼
Crawler
        │
        ▼
Parser
        │
        ▼
Normalizer
        │
        ▼
Semantic Segmenter
        │
        ▼
Knowledge Extractor
        │
        ▼
Knowledge Assets
        │
        ▼
Instruction Generator
        │
        ▼
Answer Generator
        │
        ▼
AI Reviewer
        │
        ▼
Quality Scorer
        │
        ▼
Deduplicator
        │
        ▼
Dataset Balancer
        │
        ▼
Exporter
        │
        ▼
Training Dataset
```

---

# Project Status

Current Stage:

> Architecture & Design

---

# Documentation

Architecture documentation is available under:

```
docs/architecture/
```

---

# License

Apache License 2.0