-- migrate:up
CREATE TABLE product_specs (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    product_id BIGINT NOT NULL,
    key VARCHAR(255) NOT NULL,
    value TEXT NOT NULL,
    "order" INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_product_specs_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    CONSTRAINT uk_product_specs_product_key UNIQUE (product_id, key)
);

CREATE INDEX idx_product_specs_deleted_at ON product_specs(deleted_at);
CREATE INDEX idx_product_specs_product_id ON product_specs(product_id);
CREATE INDEX idx_product_specs_order ON product_specs(product_id, "order");

-- migrate:down
DROP INDEX IF EXISTS idx_product_specs_order;
DROP INDEX IF EXISTS idx_product_specs_product_id;
DROP INDEX IF EXISTS idx_product_specs_deleted_at;
DROP TABLE IF EXISTS product_specs;

