-- migrate:up
CREATE TABLE product_features (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    product_id BIGINT NOT NULL,
    feature VARCHAR(500) NOT NULL,
    "order" INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_product_features_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE INDEX idx_product_features_deleted_at ON product_features(deleted_at);
CREATE INDEX idx_product_features_product_id ON product_features(product_id);
CREATE INDEX idx_product_features_order ON product_features(product_id, "order");

-- migrate:down
DROP INDEX IF EXISTS idx_product_features_order;
DROP INDEX IF EXISTS idx_product_features_product_id;
DROP INDEX IF EXISTS idx_product_features_deleted_at;
DROP TABLE IF EXISTS product_features;

