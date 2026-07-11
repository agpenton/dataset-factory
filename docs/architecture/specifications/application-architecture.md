# Application Architecture

| Field | Value |
|-------|-------|
| Status | Accepted |
| Version | 1.0 |
| Date | 2026-07-11 |
| Project | Dataset Factory |

---

# Purpose

Define the application architecture and dependency rules.

This specification governs every package in the repository.

Violating these rules requires an ADR.

---

# Architectural Style

Dataset Factory follows a layered architecture inspired by:

- Domain-Driven Design (DDD)
- Clean Architecture
- Hexagonal Architecture (Ports & Adapters)

---

# Architecture

```
                    +----------------------+
                    |         CLI          |
                    +----------------------+
                               |
                               ▼
                    +----------------------+
                    |     Application      |
                    |      Use Cases       |
                    +----------------------+
                               |
                               ▼
                    +----------------------+
                    |        Domain        |
                    +----------------------+
                               ▲
                               |
          +--------------------------------------------+
          |                                            |
          ▼                                            ▼
+----------------------+                +----------------------+
|     Infrastructure   |                |       Plugins        |
+----------------------+                +----------------------+
```

---

# Layers

## Presentation

Responsible for user interaction.

Contains:

```
CLI

REST API

gRPC

Web UI (future)
```

No business logic.

---

## Application

Responsible for orchestration.

Contains:

```
Use Cases

Commands

Queries

Pipeline Coordination

Transactions

Authorization
```

Application coordinates.

It does not implement business rules.

---

## Domain

Contains business rules.

Includes:

```
Document

Knowledge Asset

Dataset

Pipeline

Recipe

Benchmark
```

No infrastructure dependencies.

---

## Infrastructure

Responsible for implementation.

Examples

```
SQLite

PostgreSQL

Filesystem

OpenAI

Ollama

GitHub

AWS

Logging

Metrics
```

---

## Plugins

Feature implementations.

Examples

```
GitHub Crawler

Markdown Parser

PDF Parser

OpenAI Generator

JSONL Exporter
```

Plugins depend on interfaces.

Never the reverse.

---

# Dependency Rules

Allowed

```
Presentation

↓

Application

↓

Domain
```

Infrastructure

↓

Domain

Plugins

↓

Domain

Plugins

↓

SDK

---

Forbidden

```
Domain

↓

Infrastructure
```

```
Domain

↓

Plugins
```

```
Application

↓

Concrete Plugin
```

```
Plugin

↓

Plugin
```

---

# Dependency Inversion

Application depends on interfaces.

Infrastructure implements interfaces.

Example

```
Application

↓

Crawler Interface

↓

GitHub Crawler
```

---

# Use Cases

Every business operation becomes a Use Case.

Examples

```
RunPipeline

ImportDocument

ExtractKnowledge

GenerateDataset

ReviewDataset

ExportDataset

RunBenchmark
```

Each Use Case performs exactly one task.

---

# Commands

Commands mutate state.

Examples

```
CreatePipeline

DeleteDataset

ImportDocuments

GeneratePrompts
```

---

# Queries

Queries never modify state.

Examples

```
ListDatasets

ShowStatistics

PipelineStatus

BenchmarkHistory
```

---

# Ports

Ports define application contracts.

Examples

```
Crawler

Parser

Generator

Storage

AIProvider

Exporter
```

---

# Adapters

Adapters implement Ports.

Examples

```
GitHubCrawler

FilesystemCrawler

OpenAIProvider

SQLiteStorage

JSONLExporter
```

---

# Dependency Injection

Every dependency is injected.

Never instantiate infrastructure inside business logic.

Allowed

```
func NewRunner(storage Storage)
```

Forbidden

```
storage := sqlite.New(...)
```

inside business logic.

---

# Transactions

Transactions belong to the Application layer.

The Domain remains transaction agnostic.

---

# Events

Business events originate from the Domain.

Infrastructure publishes them.

Examples

```
DocumentImported

KnowledgeExtracted

DatasetGenerated

DatasetReleased
```

---

# Error Handling

Domain defines errors.

Infrastructure wraps errors.

Presentation formats errors.

---

# Logging

Infrastructure concern.

Domain never logs.

---

# Configuration

Configuration is loaded before application startup.

Domain never reads configuration files.

---

# Testing Strategy

Presentation

↓

Integration Tests

Application

↓

Use Case Tests

Domain

↓

Unit Tests

Infrastructure

↓

Integration Tests

Plugins

↓

Contract Tests

---

# Summary

Dataset Factory follows Clean Architecture.

Business rules remain independent.

Infrastructure remains replaceable.

Plugins remain isolated.

Every dependency points toward the Domain.