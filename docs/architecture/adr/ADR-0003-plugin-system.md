# ADR-0003 — Plugin System

| Field | Value |
|-------|-------|
| Status | Accepted |
| Version | 1.0 |
| Date | 2026-07-11 |
| Project | Dataset Factory |
| Depends On | ADR-0001, ADR-0002 |

---

# 1. Purpose

Define the plugin architecture of Dataset Factory.

The plugin system enables the platform to be extended without modifying the core engine.

Every source, parser, extractor, AI provider, exporter and benchmark must be implemented as a plugin.

---

# 2. Goals

The plugin system SHALL:

- isolate implementations
- support dependency injection
- support automatic discovery
- support version compatibility
- support capability negotiation
- support configuration validation
- support lifecycle management
- remain deterministic

---

# 3. Non-Goals

The plugin system SHALL NOT:

- implement business logic
- manage pipeline execution
- own configuration storage
- expose plugin internals

---

# 4. Philosophy

The core application knows **interfaces**.

Plugins provide **implementations**.

The core never imports implementation packages directly.

```
Pipeline

↓

Interface

↓

Plugin

↓

Implementation
```

---

# 5. Plugin Categories

Dataset Factory supports the following plugin types.

```
Crawler

Parser

Normalizer

Segmenter

Knowledge Extractor

Prompt Generator

Answer Generator

Reviewer

Scorer

Deduplicator

Balancer

Exporter

Benchmark

AI Provider

Storage

Authentication

Telemetry
```

Each plugin belongs to exactly one category.

---

# 6. Plugin Architecture

```
Dataset Factory

├── Registry

├── Loader

├── Resolver

├── Validator

├── Lifecycle Manager

└── Plugin SDK

            │

            ▼

Plugins
```

---

# 7. Plugin Lifecycle

Every plugin follows the same lifecycle.

```
Register

↓

Configure

↓

Validate

↓

Initialize

↓

Start

↓

Execute

↓

Shutdown
```

No plugin may execute before successful validation.

---

# 8. Registration

Plugins register themselves during application startup.

Example:

```
registry.Register(
    crawler.New(),
)
```

Registration order must not affect execution.

---

# 9. Discovery

The Registry discovers plugins automatically.

The pipeline requests plugins by capability.

Example:

```
Need:

Crawler

↓

Registry

↓

GitHub Crawler
```

The pipeline never requests concrete implementations.

---

# 10. Capability Model

Every plugin declares capabilities.

Example

```yaml
plugin:

name: github

type: crawler

capabilities:

- issues

- pull_requests

- discussions
```

Capabilities are immutable after registration.

---

# 11. Versioning

Every plugin declares:

```
Name

Version

API Version

Capabilities
```

Example

```yaml
name: github

version: 1.2.0

api: v1
```

Plugins with incompatible API versions shall not load.

---

# 12. Configuration

Each plugin owns its configuration.

Example

```yaml
crawler:

github:

token:

organizations:

repositories:
```

Configuration validation occurs before initialization.

---

# 13. Dependency Injection

Plugins receive dependencies through constructors.

Never through globals.

Example

```
Logger

Metrics

Storage

Configuration

AI Client

HTTP Client
```

The plugin must not instantiate shared infrastructure.

---

# 14. Isolation

Plugins shall never communicate directly.

Communication occurs only through:

- interfaces
- events
- pipeline contracts

---

# 15. State

Plugins should be stateless whenever possible.

If state is required, it must be persisted externally.

---

# 16. Error Handling

Plugins return typed errors.

Errors include:

```
Plugin

Stage

Severity

Cause

Recoverable
```

Plugins never terminate the application directly.

---

# 17. Observability

Every plugin exports:

Metrics

Logs

Tracing

Plugins shall use the platform instrumentation.

---

# 18. Security

Plugins never access secrets directly.

Secrets are injected.

Plugins never write credentials to logs.

---

# 19. Determinism

Plugins shall produce identical output when given:

- identical input
- identical configuration
- identical dependencies

Plugins must not depend on:

- random ordering
- wall clock
- global mutable state

unless explicitly configured.

---

# 20. Plugin SDK

Dataset Factory provides an SDK.

Responsibilities:

- registration
- configuration
- validation
- logging
- metrics
- lifecycle helpers
- testing utilities

Plugin authors should never reimplement these features.

---

# 21. Internal Plugins

Core plugins maintained by Dataset Factory.

Examples:

```
Markdown Parser

JSON Exporter

OpenAI Generator

GitHub Crawler

SQLite Storage
```

---

# 22. External Plugins

Third-party plugins are supported.

Examples:

```
Confluence

Jira

Notion

Google Drive

GitLab

Azure DevOps

Bitbucket
```

The core application should not distinguish between internal and external plugins.

---

# 23. Future Plugin Types

Future releases may support:

- WASM plugins
- Remote plugins
- gRPC plugins
- OCI plugin distribution

These additions must remain backward compatible.

---

# 24. Plugin Dependencies

Plugins may depend only on:

- SDK
- public interfaces

Plugins must never depend on:

- another plugin implementation
- pipeline internals
- storage implementation details

---

# 25. Compatibility

Breaking changes require:

- API version increment
- migration guide
- compatibility matrix

Minor releases must preserve plugin compatibility.

---

# 26. Testing

Every plugin shall include:

- unit tests
- integration tests
- configuration validation tests
- deterministic replay tests

The SDK will provide common test helpers.

---

# 27. Packaging

Each plugin is distributed as a Go module.

Repository layout:

```
plugins/

crawler/

github/

gitlab/

aws/

parser/

markdown/

pdf/

html/

generator/

openai/

anthropic/

reviewer/

llm/

rulebased/

exporter/

chatml/

openai/

sharegpt/
```

Each plugin is independently versioned.

---

# 28. Decisions

Accepted decisions:

- Interface-first architecture
- Constructor dependency injection
- Automatic registration
- Capability discovery
- Version compatibility checks
- Stateless plugins
- Shared SDK
- Deterministic execution

---

# 29. Consequences

Advantages

- Highly extensible
- Easy to test
- Independent releases
- Clear ownership
- Loose coupling
- Stable public APIs

Trade-offs

- More interfaces
- Higher initial complexity
- Registry management
- API compatibility requirements

---

# 30. Summary

Dataset Factory is built around a plugin-first architecture.

The core platform owns orchestration.

Plugins own implementation.

Every extension point is defined by stable interfaces, enabling independent evolution of crawlers, parsers, AI providers, exporters, reviewers and future integrations without modifying the pipeline engine.