# Repository Layout Specification

| Field | Value |
|-------|-------|
| Status | Accepted |
| Version | 1.0 |
| Date | 2026-07-11 |
| Project | Dataset Factory |
| Depends On | ADR-0001 ... ADR-0006 |

---

# 1. Purpose

Define the repository structure of Dataset Factory.

This document establishes ownership boundaries between packages.

Every directory has exactly one responsibility.

No package may violate these boundaries.

---

# 2. Design Principles

The repository SHALL be:

- Modular
- Layered
- Interface-first
- Plugin-oriented
- Testable
- Scalable

---

# 3. Repository Layout

```
dataset-factory/

├── cmd/
│
├── internal/
│
├── pkg/
│
├── plugins/
│
├── sdk/
│
├── api/
│
├── configs/
│
├── recipes/
│
├── datasets/
│
├── benchmarks/
│
├── docs/
│
├── scripts/
│
├── testdata/
│
├── examples/
│
├── deployments/
│
├── tools/
│
├── .github/
│
├── Makefile
├── go.mod
├── go.sum
└── README.md
```

---

# 4. cmd/

Contains executable entrypoints.

```
cmd/

server/

cli/

worker/

benchmark/
```

No business logic.

---

# 5. internal/

Private implementation.

Cannot be imported externally.

```
internal/

pipeline/

engine/

scheduler/

runner/

workers/

queue/

checkpoint/

events/

metrics/

logging/

config/

validation/

storage/

telemetry/

registry/
```

---

# 6. pkg/

Public reusable libraries.

```
pkg/

document/

knowledge/

pipeline/

plugin/

crawler/

parser/

segmenter/

generator/

reviewer/

scorer/

exporter/

benchmark/

errors/

types/

utils/
```

Everything inside pkg is considered public API.

---

# 7. plugins/

Concrete implementations.

```
plugins/

crawler/

parser/

generator/

reviewer/

exporter/

storage/

benchmark/

provider/
```

---

# 8. Crawlers

```
plugins/crawler/

aws/

azure/

gcp/

github/

gitlab/

cncf/

terraform/

helm/

argocd/

fluxcd/

rollouts/

flagger/

ansible/

golang/

python/

typescript/

agentic-ai/

openai/

anthropic/
```

Each crawler is an independent Go module.

---

# 9. Parsers

```
plugins/parser/

markdown/

html/

pdf/

json/

yaml/

asciidoc/

rst/
```

---

# 10. AI Providers

```
plugins/provider/

openai/

anthropic/

gemini/

ollama/

llamacpp/

mlx/

vllm/

sglang/
```

---

# 11. Exporters

```
plugins/exporter/

chatml/

sharegpt/

openai/

alpaca/

hf/

parquet/

csv/

jsonl/
```

---

# 12. SDK

Public SDK for plugin developers.

```
sdk/

plugin/

config/

logging/

metrics/

testing/

events/

validation/
```

No business logic.

---

# 13. API

REST and future gRPC APIs.

```
api/

http/

grpc/

middleware/

handlers/

dto/
```

---

# 14. Configurations

```
configs/

pipeline/

crawler/

generator/

reviewer/

storage/

examples/
```

---

# 15. Recipes

Pipeline definitions.

```
recipes/

aws/

terraform/

kubernetes/

gitops/

platform/

linux/

golang/

python/

typescript/

agentic-ai/

custom/
```

Recipes are declarative YAML.

---

# 16. Datasets

Generated datasets.

```
datasets/

raw/

normalized/

knowledge/

generated/

reviewed/

balanced/

released/
```

Datasets are never committed.

Only metadata may be committed.

---

# 17. Benchmarks

```
benchmarks/

platform-engineering/

aws/

terraform/

kubernetes/

gitops/

linux/

coding/

agentic-ai/
```

Benchmark prompts are versioned.

---

# 18. Documentation

```
docs/

architecture/

guides/

tutorials/

api/

examples/

development/
```

---

# 19. Scripts

```
scripts/

bootstrap/

release/

lint/

test/

benchmark/
```

Scripts must be idempotent.

---

# 20. Test Data

```
testdata/

documents/

knowledge/

datasets/

configs/

fixtures/
```

No production data.

---

# 21. Examples

```
examples/

simple/

advanced/

plugins/

recipes/

datasets/
```

Examples are runnable.

---

# 22. Deployments

```
deployments/

docker/

kubernetes/

helm/

compose/
```

Deployment manifests only.

---

# 23. Tools

Developer tooling.

```
tools/

codegen/

schema/

lint/

migration/
```

---

# 24. Package Ownership

Every package has one owner.

No package may contain unrelated functionality.

---

# 25. Import Rules

Allowed

```
cmd

↓

internal

↓

pkg

↓

sdk
```

Plugins

```
plugins

↓

sdk

↓

pkg
```

Forbidden

```
pkg

↓

internal
```

```
sdk

↓

internal
```

```
plugin A

↓

plugin B
```

---

# 26. Dependency Rules

```
cmd

↓

internal

↓

pkg

↓

sdk
```

Plugins depend only on:

- sdk
- pkg

Never on:

- internal
- cmd
- other plugins

---

# 27. Package Naming

Rules

- lowercase
- singular
- no abbreviations
- one responsibility

Good

```
pipeline

knowledge

crawler

benchmark
```

Bad

```
utils2

misc

common

helpers

stuff
```

---

# 28. Internal Package Layout

Example

```
internal/pipeline/

builder.go

engine.go

executor.go

runner.go

scheduler.go

state.go

worker.go
```

One file per responsibility.

---

# 29. Public Package Layout

Example

```
pkg/knowledge/

asset.go

builder.go

validator.go

serializer.go

types.go
```

---

# 30. Plugin Layout

Example

```
plugins/crawler/github/

plugin.go

config.go

crawler.go

parser.go

client.go

mapper.go

validator.go

errors.go

README.md

plugin_test.go
```

---

# 31. Testing Layout

```
pkg/

foo.go

foo_test.go

integration/

benchmark/
```

Unit tests remain beside implementation.

---

# 32. Generated Code

Generated code belongs in

```
internal/generated/

or

pkg/generated/
```

Never mixed with handwritten code.

---

# 33. Configuration Files

Supported

```
yaml

json

toml
```

YAML is preferred.

---

# 34. Assets

Static assets

```
assets/

schemas/

templates/

prompts/

icons/
```

---

# 35. Future Expansion

Reserved directories

```
operators/

web/

desktop/

mobile/

cloud/
```

No implementation until required.

---

# 36. Decisions

Accepted

- Layered architecture
- Plugin isolation
- Public/private separation
- Interface-first design
- Modular repository
- One responsibility per package

---

# 37. Consequences

Advantages

- Easy navigation
- Clear ownership
- Scalable architecture
- Independent plugins
- Stable APIs

Trade-offs

- Larger repository
- More packages
- More interfaces
- Additional documentation

---

# 38. Summary

The repository structure reflects the architecture of Dataset Factory.

Directory boundaries define architectural boundaries.

Violating these boundaries requires a new ADR.