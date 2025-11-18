-- PostgreSQL schema for product_codes table

-- Drop table if it already exists (optional, use with care in non-production environments)
-- DROP TABLE IF EXISTS product_codes;

CREATE TABLE IF NOT EXISTS product_codes (
    id           BIGSERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    code         VARCHAR(19) NOT NULL UNIQUE
);

-- Optional: add a CHECK constraint to enforce format XXXX-XXXX-XXXX-XXXX (A-Z, 0-9)
ALTER TABLE product_codes
    ADD CONSTRAINT product_codes_code_format_chk
    CHECK (code ~ '^[A-Z0-9]{4}(-[A-Z0-9]{4}){3}$');

-- Indexes

-- Unique index on code (redundant with UNIQUE in table definition but explicit for clarity)
CREATE UNIQUE INDEX IF NOT EXISTS idx_product_codes_code
    ON product_codes (code);

-- Non-unique index on product_name to speed up search/filtering by name (if needed)
CREATE INDEX IF NOT EXISTS idx_product_codes_product_name
    ON product_codes (product_name);

-- Example: index on id is automatically created by PRIMARY KEY (no need to add manually)


