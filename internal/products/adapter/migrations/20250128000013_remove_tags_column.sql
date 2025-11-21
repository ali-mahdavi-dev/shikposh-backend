-- migrate:up
-- Remove the tags JSONB column from products table since we're using a separate tags table now
DROP INDEX IF EXISTS idx_products_tags;
ALTER TABLE products DROP COLUMN IF EXISTS tags;

-- migrate:down
-- Restore the tags column (as empty array for existing records)
ALTER TABLE products ADD COLUMN tags JSONB DEFAULT '[]'::jsonb;
CREATE INDEX idx_products_tags ON products USING GIN(tags);

