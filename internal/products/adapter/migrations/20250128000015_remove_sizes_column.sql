-- migrate:up
-- Remove the sizes JSONB column from products table since we're using a separate sizes table now
ALTER TABLE products DROP COLUMN IF EXISTS sizes;

-- migrate:down
-- Restore the sizes column (as empty array for existing records)
ALTER TABLE products ADD COLUMN sizes JSONB DEFAULT '[]'::jsonb;

