-- migrate:up
CREATE TABLE product_details (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    product_id BIGINT NOT NULL,
    color_key VARCHAR(100),
    color_name VARCHAR(255),
    size_key VARCHAR(100),
    price DECIMAL(10,2) NOT NULL,
    original_price DECIMAL(10,2),
    stock INTEGER DEFAULT 0,
    discount INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_product_details_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE INDEX idx_product_details_deleted_at ON product_details(deleted_at);
CREATE INDEX idx_product_details_product_id ON product_details(product_id);
CREATE INDEX idx_product_details_color_key ON product_details(color_key) WHERE color_key IS NOT NULL;
CREATE INDEX idx_product_details_size_key ON product_details(size_key) WHERE size_key IS NOT NULL;
CREATE UNIQUE INDEX uk_product_details_unique ON product_details(product_id, COALESCE(color_key, ''), COALESCE(size_key, ''));

-- migrate:down
DROP INDEX IF EXISTS uk_product_details_unique;
DROP INDEX IF EXISTS idx_product_details_size_key;
DROP INDEX IF EXISTS idx_product_details_color_key;
DROP INDEX IF EXISTS idx_product_details_product_id;
DROP INDEX IF EXISTS idx_product_details_deleted_at;
DROP TABLE IF EXISTS product_details;

