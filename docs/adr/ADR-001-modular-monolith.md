# ADR-001 — Modular monolith over microservices for the MVP

Status: Accepted (CTO approved, Day 1)

## Context
Tiny team, unknown domain boundaries, need to ship a vertical slice fast and learn.

## Decision
Build Pangea as ONE deployable Go service composed of internal modules
(catalog, later orders, payments, etc.) with clean boundaries between them.
We do NOT split into microservices yet.

## Consequences
- Positive: fast to build, easy to refactor, one deploy, no distributed-systems
  tax before we have proven seams.
- Negative: we must keep module boundaries disciplined so future extraction is cheap.
- Later: when a seam is proven and a scaling reason exists, extract it into a
  service. This baseline is referenced by every later release.