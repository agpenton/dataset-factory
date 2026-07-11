# ADR-0005 — Document Model

| Field | Value |
|-------|-------|
| Status | Accepted |
| Version | 1.0 |
| Date | 2026-07-11 |
| Project | Dataset Factory |
| Depends On | ADR-0001, ADR-0002, ADR-0003, ADR-0004 |

---

# 1. Purpose

Define the canonical representation of every document entering Dataset Factory.

Every crawler, importer and parser SHALL produce a Document.

Every downstream component SHALL consume only the normalized Document model.

---

# 2. Problem

Engineering knowledge comes from heterogeneous sources:

- HTML
- Markdown
- PDF
- GitHub
- ADRs
- RFCs
- Books
- Blog posts
- Videos
- Conference slides

Each source has a different structure.

Without normalization every parser would require custom logic throughout the pipeline.

---

# 3. Decision

Every input source SHALL be converted into a canonical Document.

The Document is the only object accepted by the Normalization stage.

```
Crawler

↓

Raw Source

↓

Document

↓

Normalizer

↓

Knowledge Assets
```

---

# 4. Philosophy

A Document represents content.

A Knowledge Asset represents engineering knowledge.

The Document exists only to normalize heterogeneous inputs.

Knowledge extraction begins after normalization.

---

# 5. Lifecycle

```
Discover

↓

Fetch

↓

Parse

↓

Normalize

↓

Validate

↓

Persist

↓

Knowledge Extraction
```

---

# 6. Canonical Model

```yaml
id:

version:

type:

source:

metadata:

content:

attachments:

relationships:

checksum:
```

---

# 7. Identity

Every document SHALL have a deterministic identifier.

Example

```
document://github/kubernetes/issues/12345

document://aws/blog/landing-zones

document://terraform/docs/backend-s3
```

IDs shall remain stable.

---

# 8. Document Types

Supported types

```
Documentation

Blog

Article

Tutorial

Reference

ADR

RFC

KEP

Issue

PullRequest

Discussion

Book

Presentation

VideoTranscript

Runbook

FAQ

CodeRepository
```

---

# 9. Source

Every document records its origin.

```yaml
source:

provider:

url:

organization:

repository:

branch:

commit:

author:

license:
```

---

# 10. Metadata

```yaml
metadata:

title:

subtitle:

description:

language:

published:

updated:

author:

estimated_tokens:

estimated_read_time:
```

---

# 11. Content

The canonical representation SHALL be Markdown.

Regardless of source:

```
HTML

↓

Markdown

PDF

↓

Markdown

Word

↓

Markdown

Slides

↓

Markdown
```

Markdown becomes the canonical document format.

---

# 12. Sections

Documents are divided into semantic sections.

Each section contains

```yaml
id:

heading:

level:

content:
```

Section hierarchy must be preserved.

---

# 13. Code Blocks

Code SHALL NOT be merged with prose.

Each code block records

```yaml
language:

filename:

purpose:

content:
```

---

# 14. Images

Images are represented as metadata.

```yaml
image:

id:

caption:

alt_text:

path:
```

Future OCR is supported.

---

# 15. Tables

Tables remain structured.

Never flatten tables into plain text.

Supported formats

- Markdown
- CSV
- JSON

---

# 16. Hyperlinks

Links are normalized.

Each link stores

```yaml
text:

url:

type:
```

Types

```
Internal

External

Reference
```

---

# 17. Attachments

Supported attachments

```
PDF

Images

Archives

Code

Diagrams

Spreadsheets
```

Attachments are stored separately.

Documents reference them.

---

# 18. Relationships

Documents may reference other documents.

Relationships include

```
References

Supersedes

Duplicates

DerivedFrom

RelatedTo
```

---

# 19. Checksum

Every normalized document stores a checksum.

Purpose

- deduplication
- integrity verification
- change detection

Supported algorithms

```
SHA256
```

---

# 20. Versioning

Every document is versioned.

```
v1

v2

v3
```

Document history is immutable.

---

# 21. Validation

Every document must satisfy

Required

```
id

type

title

source

content

checksum
```

Validation failures stop processing.

---

# 22. Immutability

Documents are immutable.

Pipeline stages never modify existing documents.

Transformations always create new versions.

---

# 23. Storage

Documents are stored separately from Knowledge Assets.

```
Raw Sources

↓

Documents

↓

Knowledge Assets

↓

Datasets
```

Each layer has independent persistence.

---

# 24. Serialization

Supported formats

```
JSON

YAML

Parquet
```

JSON is canonical.

---

# 25. Performance

Normalization must

- stream input
- avoid loading entire files
- support files larger than 1 GB
- preserve ordering

---

# 26. Error Handling

Failures are classified

```
Unsupported Format

Malformed Content

Encoding Error

Checksum Failure

Validation Failure
```

Errors are recoverable whenever possible.

---

# 27. Future Extensions

Future support

- OCR
- Speech-to-text
- Diagram parsing
- Embedded notebooks
- CAD documents
- API specifications
- CAD diagrams
- Interactive documentation

Backward compatibility is required.

---

# 28. Decisions

Accepted

- Canonical Document model
- Markdown as canonical representation
- Immutable documents
- Deterministic IDs
- Structured metadata
- Streaming normalization
- Versioned documents

---

# 29. Consequences

Advantages

- One ingestion format
- Simplified parsers
- Easier testing
- Stable contracts
- Independent crawlers

Trade-offs

- Initial normalization cost
- Additional storage
- Conversion complexity

---

# 30. Summary

The Document model is the canonical ingestion format for Dataset Factory.

Every crawler produces Documents.

Every downstream stage relies on a consistent representation independent of the original source.

Documents represent content.

Knowledge Assets represent engineering knowledge.

The separation between these two models is a foundational architectural decision.