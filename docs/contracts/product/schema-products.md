# Schema Design — products (v1)

| Column      | Type        | Constraints                              | Why |
|-------------|-------------|------------------------------------------|-----|
| id          | UUID        | PK, default gen_random_uuid()            | global, non-guessable IDs; no cross-tenant sequence leakage |
| tenant_id   | UUID        | NOT NULL                                 | ADR-002; every row is tenant-scoped |
| name        | TEXT        | NOT NULL, CHECK length 1..140            | mirrors domain invariant |
| price_cents | BIGINT      | NOT NULL, CHECK >= 0                     | money as integer cents; never float |
| currency    | CHAR(3)     | NOT NULL, default 'USD'                  | ISO-4217 code |
| category    | TEXT        | NOT NULL, CHECK length >= 1              | mirrors domain invariant |
| in_stock    | BOOLEAN     | NOT NULL, default FALSE                  | new products start out of stock |
| created_at  | TIMESTAMPTZ | NOT NULL, default now()                  | UTC; audit |
| updated_at  | TIMESTAMPTZ | NOT NULL, default now()                  | UTC; bumped on update |

## Indexes
- PK on id (automatic).
- idx_products_tenant_id on (tenant_id): the list endpoint always filters by
  tenant, so this is the hot read path.

## Defense in depth
The CHECK constraints duplicate the Go domain invariants on purpose: the
application validates for good UX (clear errors), and the database validates as
the last line of defense (Requirement — a direct write must still be safe).