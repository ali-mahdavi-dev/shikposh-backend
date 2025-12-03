-- migrate:up
-- product_categories (M:N - یک محصول میتونه چند کتگوری داشته باشه)
CREATE TABLE IF NOT EXISTS product_categories (
    product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    category_id BIGINT NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    PRIMARY KEY (product_id, category_id)
);
CREATE INDEX IF NOT EXISTS idx_product_categories_product ON product_categories(product_id);
CREATE INDEX IF NOT EXISTS idx_product_categories_category ON product_categories(category_id);

-- Migrate existing category_id to junction table
INSERT INTO product_categories (product_id, category_id)
SELECT id, category_id FROM products WHERE category_id IS NOT NULL AND deleted_at IS NULL
ON CONFLICT DO NOTHING;

-- migrate:down
DROP INDEX IF EXISTS idx_product_categories_category;
DROP INDEX IF EXISTS idx_product_categories_product;
DROP TABLE IF EXISTS product_categories;

