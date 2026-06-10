# Product — Data Dictionary & Validation (v1)

## Fields
| Field        | Type            | Required | Source         | Rule |
|--------------|-----------------|----------|----------------|------|
| id           | UUID (string)   | yes      | system         | assigned on persist (later sprint) |
| tenant_id    | UUID (string)   | yes      | system         | injected from request context; never client-set |
| name         | string          | yes      | merchant       | 1–140 chars after trimming whitespace |
| price_cents  | integer (int64) | yes      | merchant       | >= 0 (money stored in cents, never floats) |
| currency     | string (ISO-4217)| no      | merchant/tenant| defaults to "USD" if blank |
| category     | string          | yes      | merchant       | non-empty after trimming |
| in_stock     | boolean         | no       | merchant       | defaults to false |
| created_at   | timestamp (UTC) | yes      | system         | set on creation |
| updated_at   | timestamp (UTC) | yes      | system         | set on creation and every update |

## Validation / decision table (create product)
| Condition                         | Result | Error message            |
|-----------------------------------|--------|--------------------------|
| tenant_id missing/blank           | reject | "tenant_id is required"  |
| name blank after trim             | reject | "name is required"       |
| name longer than 140 chars        | reject | "name must be at most 140 characters" |
| price_cents < 0                   | reject | "price_cents must be >= 0" |
| category blank after trim         | reject | "category is required"   |
| all rules pass                    | accept | (product created)        |

## Creation workflow
merchant submits -> validate (table above) -> inject tenant_id ->
persist -> return created product

## Decisions logged today
- Money is stored as integer cents (`price_cents`), never as a float.
- Duplicate names within a tenant are allowed for now (revisit later).