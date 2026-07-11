# Core Interfaces Specification

| Field | Value |
|-------|-------|
| Status | Accepted |
| Version | 1.0 |
| Date | 2026-07-11 |
| Project | Dataset Factory |
| Depends On | ADR-0001 ... ADR-0006 |

---

# 1. Purpose

This document defines the public interfaces that make up Dataset Factory.

These interfaces are considered stable APIs.

All implementations must conform to these contracts.

No implementation-specific behavior shall appear in these interfaces.

---

# 2. Design Principles

Interfaces SHALL:

- represent one responsibility
- remain implementation independent
- be deterministic
- be easy to mock
- support dependency injection
- avoid circular dependencies

---

# 3. Architecture

```
Pipeline

↓

Stage

↓

Plugin

↓

Implementation
```

Business logic never depends on concrete implementations.

---

# 4. Pipeline

Responsible for orchestrating execution.

```go
type Pipeline interface {

    ID() string

    Name() string

    Execute(ctx context.Context) error

    Resume(ctx context.Context) error

    Cancel(ctx context.Context) error

    Status() PipelineStatus
}
```

---

# 5. Stage

Represents one processing step.

```go
type Stage interface {

    Name() string

    Description() string

    Configure(cfg any) error

    Validate() error

    Execute(ctx context.Context, stream Stream) (Stream, error)
}
```

---

# 6. Plugin

Base interface implemented by every extension.

```go
type Plugin interface {

    Metadata() PluginMetadata

    Configure(any) error

    Validate() error

    Start(context.Context) error

    Stop(context.Context) error
}
```

---

# 7. Crawler

Produces Documents.

```go
type Crawler interface {

    Plugin

    Discover(ctx context.Context) ([]DocumentRef, error)

    Fetch(ctx context.Context, ref DocumentRef) (*Document, error)
}
```

---

# 8. Parser

Transforms raw bytes into Documents.

```go
type Parser interface {

    Plugin

    Parse(ctx context.Context, input io.Reader) (*Document, error)
}
```

---

# 9. Normalizer

Produces normalized Documents.

```go
type Normalizer interface {

    Plugin

    Normalize(ctx context.Context, doc *Document) (*Document, error)
}
```

---

# 10. Segmenter

Splits documents into semantic sections.

```go
type Segmenter interface {

    Plugin

    Segment(ctx context.Context, doc *Document) ([]Section, error)
}
```

---

# 11. Knowledge Extractor

Produces Knowledge Assets.

```go
type KnowledgeExtractor interface {

    Plugin

    Extract(
        ctx context.Context,
        section Section,
    ) ([]KnowledgeAsset, error)
}
```

---

# 12. Prompt Generator

Creates training prompts.

```go
type PromptGenerator interface {

    Plugin

    Generate(
        ctx context.Context,
        asset KnowledgeAsset,
    ) ([]Prompt, error)
}
```

---

# 13. Answer Generator

Creates assistant responses.

```go
type AnswerGenerator interface {

    Plugin

    Generate(
        ctx context.Context,
        prompt Prompt,
    ) (*Answer, error)
}
```

---

# 14. Reviewer

Reviews generated examples.

```go
type Reviewer interface {

    Plugin

    Review(
        ctx context.Context,
        sample DatasetSample,
    ) (*Review, error)
}
```

---

# 15. Scorer

Assigns quality scores.

```go
type Scorer interface {

    Plugin

    Score(
        ctx context.Context,
        sample DatasetSample,
    ) (*Score, error)
}
```

---

# 16. Deduplicator

Removes duplicate samples.

```go
type Deduplicator interface {

    Plugin

    Deduplicate(
        ctx context.Context,
        samples []DatasetSample,
    ) ([]DatasetSample, error)
}
```

---

# 17. Balancer

Balances datasets.

```go
type Balancer interface {

    Plugin

    Balance(
        ctx context.Context,
        samples []DatasetSample,
    ) ([]DatasetSample, error)
}
```

---

# 18. Exporter

Exports datasets.

```go
type Exporter interface {

    Plugin

    Export(
        ctx context.Context,
        dataset Dataset,
        writer io.Writer,
    ) error
}
```

---

# 19. Storage

Persists platform state.

```go
type Storage interface {

    Plugin

    Save(context.Context, any) error

    Load(context.Context, string, any) error

    Delete(context.Context, string) error
}
```

---

# 20. AI Provider

Provides LLM capabilities.

```go
type AIProvider interface {

    Plugin

    Complete(
        ctx context.Context,
        request CompletionRequest,
    ) (*CompletionResponse, error)

    Embed(
        ctx context.Context,
        texts []string,
    ) ([]Embedding, error)
}
```

---

# 21. Benchmark

Evaluates generated datasets and models.

```go
type Benchmark interface {

    Plugin

    Run(
        ctx context.Context,
        dataset Dataset,
    ) (*BenchmarkResult, error)
}
```

---

# 22. Event Publisher

Publishes pipeline events.

```go
type EventPublisher interface {

    Publish(
        ctx context.Context,
        event Event,
    ) error
}
```

---

# 23. Metrics

Records metrics.

```go
type Metrics interface {

    Counter(name string)

    Gauge(name string)

    Histogram(name string)

    Timer(name string)
}
```

---

# 24. Logger

Structured logging.

```go
type Logger interface {

    Debug(...any)

    Info(...any)

    Warn(...any)

    Error(...any)
}
```

---

# 25. Configuration

Loads configuration.

```go
type ConfigProvider interface {

    Load(path string) error

    Get(key string) any
}
```

---

# 26. Design Rules

Every interface SHALL:

- have one responsibility
- avoid optional methods
- avoid implementation details
- be mockable
- be deterministic

---

# 27. Dependency Rules

Allowed

```
Pipeline

↓

Interfaces

↓

Implementations
```

Forbidden

```
Pipeline

↓

Concrete Plugins
```

---

# 28. Interface Stability

Public interfaces follow Semantic Versioning.

Breaking changes require:

- major version
- migration guide
- compatibility documentation

---

# 29. Testing

Every interface SHALL have:

- mock implementation
- unit tests
- integration tests
- deterministic replay tests

---

# 30. Summary

These interfaces define the public contracts of Dataset Factory.

Every implementation, plugin and pipeline stage is built on these contracts.

No implementation may bypass or extend these interfaces without a new ADR.