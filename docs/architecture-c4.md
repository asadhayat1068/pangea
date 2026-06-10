# C4 — Pangea Marketplace MVP (v1)

## Level 1 — System Context
[FreshCart Merchant] --creates/lists products--> (Pangea Marketplace System)
[FreshCart Customer] --browses (later)----------> (Pangea Marketplace System)
(Pangea Marketplace System) --reads/writes------> [PostgreSQL]

## Level 2 — Containers
- Web App (Next.js + TypeScript)  -- HTTP/JSON --> Marketplace Service
- Marketplace Service (Go, modular monolith, hexagonal) -- SQL --> PostgreSQL
- PostgreSQL (single database, tenant_id-scoped rows)

## Internal modules of the Marketplace Service (Level 3, brief)
- catalog/domain   : entities + invariants (no I/O)
- catalog/ports    : interfaces (e.g. ProductRepository)
- catalog/adapters : http (inbound), postgres (outbound) — built later
- platform/tenant  : tenant_id context helper