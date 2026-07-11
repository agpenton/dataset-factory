# Value Objects

Value Objects are immutable business primitives.

They have:

- no identity
- immutable state
- deterministic equality

Every Value Object is comparable by value.

---

## Rules

Value Objects SHALL

- be immutable
- validate themselves
- expose no setters
- never perform I/O

---

## Types

ID

Version

Checksum

Language

Difficulty

License

Reference

URI

Hash

Duration

Timestamp

TokenCount

ByteSize