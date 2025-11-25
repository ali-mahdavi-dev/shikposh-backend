-- migrate:up
-- Insert default categories (only if they don't exist)
INSERT INTO categories (name, slug, description, created_at, updated_at)
SELECT * FROM (VALUES
  ('پیراهن و لباس مجلسی', 'dresses', 'انواع پیراهن و لباس مجلسی زنانه', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('بلوز و تاپ', 'tops', 'انواع بلوز و تاپ زنانه', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('دامن', 'skirts', 'انواع دامن زنانه', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('شلوار', 'pants', 'انواع شلوار زنانه', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('اکسسوری', 'accessories', 'انواع اکسسوری و لوازم جانبی', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
) AS v(name, slug, description, created_at, updated_at)
WHERE NOT EXISTS (SELECT 1 FROM categories WHERE categories.slug = v.slug);

-- migrate:down
DELETE FROM categories WHERE slug IN ('dresses', 'tops', 'skirts', 'pants', 'accessories');

