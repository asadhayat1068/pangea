CREATE TABLE IF NOT EXISTS products (
    id  UUID    PRIMARY KEY     DEFAULT gen_random_uuid(),
    tenant_id   UUID    NOT NULL,
    name        TEXT    NOT NULL    CHECK (char_length(name) BETWEEN 1 and 140),
    price_cents BIGINT  NOT NULL    CHECK (price_cents >= 0),
    currency    CHAR(3) NOT NULL    DEFAULT 'USD',
    category    TEXT    NOT NULL    CHECK(char_length(category) >= 1),
    in_stock    BOOLEAN NOT NULL    DEFAULT FALSE,
    created_at  TIMESTAMPTZ  NOT NULL    DEFAULT now(),
    updated_at  TIMESTAMPTZ  NOT NULL    DEFAULT now()
);


-- Tenant-scoped listing is the core read path (ADR-002). Index it.
CREATE INDEX IF NOT EXISTS idx_products_tenant_id ON products (tenant_id);