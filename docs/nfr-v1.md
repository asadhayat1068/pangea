# Non-Functional Requirements (v1)

- Availability: 99.5% for the MVP.
- Latency: p95 of GET /products < 200 ms; p95 of POST /products < 300 ms.
- Correctness: tenant isolation is non-negotiable — a tenant can never read or
  write another tenant's data.
- Observability: every request emits a structured log line; a request counter
  metric exists (added later this sprint).
- These targets become the seed SLOs the SRE formalises later in Sprint 1.