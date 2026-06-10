# ADR-002 — tenant_id everywhere

Status: Accepted (approved, 05/11/2026)

## Context
Pangea is a multi-tenant SaaS platform. Even with a single tenant today, retro-
fitting tenancy later is expensive and dangerous (data leaks between tenants).

## Decision
Every domain entity, database table, and query carries a `tenant_id` from line
one. `tenant_id` is injected from the request context server-side and is never
set by the client. Tenancy is a primitive, not a later feature.

## Consequences
- Positive: tenant isolation is built in; multi-tenant scaling (R6) and data
  residency (R9) are evolutions, not rewrites.
- Negative: a small constant tax on every model and query now.
- Enforcement: the `Product` entity requires a non-empty `tenant_id` in its
  constructor.