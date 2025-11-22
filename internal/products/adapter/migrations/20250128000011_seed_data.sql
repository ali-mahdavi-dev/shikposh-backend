-- migrate:up
-- Seed data for categories
INSERT INTO categories (name, slug, description, image) VALUES
('لباس زنانه', 'women-clothing', 'لباس‌های زنانه شامل انواع پوشاک روزمره و رسمی', 'https://example.com/images/women-clothing.jpg'),
('کیف و کفش', 'bags-shoes', 'کیف و کفش زنانه و مردانه', 'https://example.com/images/bags-shoes.jpg'),
('زیورآلات', 'jewelry', 'انواع زیورآلات طلا و نقره', 'https://example.com/images/jewelry.jpg'),
('پوشاک مردانه', 'men-clothing', 'لباس‌های مردانه شامل کت و شلوار و پوشاک روزمره', 'https://example.com/images/men-clothing.jpg'),
('اکسسوری', 'accessories', 'انواع اکسسوری و لوازم جانبی', 'https://example.com/images/accessories.jpg')
ON CONFLICT (slug) DO NOTHING;

-- Seed data for featured products
-- Note: We need to get category IDs first, so we'll use subqueries
INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, tags, image, is_new, is_featured, sizes)
SELECT 
    'کت و شلوار زنانه کلاسیک',
    'women-classic-suit',
    'شیک‌پوش',
    5,
    120,
    'کت و شلوار زنانه کلاسیک با طراحی مدرن و کیفیت عالی. مناسب برای مجالس و محیط‌های کاری.',
    (SELECT id FROM categories WHERE slug = 'women-clothing' LIMIT 1),
    '["کلاسیک", "رسمی", "کت و شلوار"]'::jsonb,
    '/images/Women-Formal.avif',
    false,
    true,
    '["S", "M", "L", "XL"]'::jsonb
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'women-classic-suit');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, tags, image, is_new, is_featured, sizes)
SELECT 
    'کیف دستی چرمی زنانه',
    'women-leather-handbag',
    'شیک‌پوش',
    2.8,
    95,
    'کیف دستی چرمی با کیفیت بالا و طراحی شیک. مناسب برای استفاده روزمره و مجالس.',
    (SELECT id FROM categories WHERE slug = 'bags-shoes' LIMIT 1),
    '["چرم", "کیف دستی", "کلاسیک"]'::jsonb,
    '/images/handbag.jpg',
    false,
    false,
    '[]'::jsonb
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'women-leather-handbag');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, tags, image, is_new, is_featured, sizes)
SELECT 
    'گردنبند طلا با نگین',
    'gold-necklace-with-gem',
    'شیک‌پوش',
    3.7,
    150,
    'گردنبند طلای 18 عیار با نگین الماس. طراحی مدرن و لوکس.',
    (SELECT id FROM categories WHERE slug = 'jewelry' LIMIT 1),
    '["طلا", "الماس", "لوکس"]'::jsonb,
    '/images/jewelry.jpg',
    true,
    false,
    '[]'::jsonb
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'gold-necklace-with-gem');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, tags, image, is_new, is_featured, sizes)
SELECT 
    'کت و شلوار مردانه رسمی',
    'men-formal-suit',
    'شیک‌پوش',
    4.6,
    200,
    'کت و شلوار مردانه رسمی با برش عالی و پارچه مرغوب. مناسب برای مجالس و محیط کار.',
    (SELECT id FROM categories WHERE slug = 'men-clothing' LIMIT 1),
    '["رسمی", "کت و شلوار", "کلاسیک"]'::jsonb,
    '/images/suit-Top.jpg',
    true,
    true,
    '["44", "46", "48", "50", "52"]'::jsonb
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'men-formal-suit');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, tags, image, is_new, is_featured, sizes)
SELECT 
    'کفش چرمی مردانه',
    'men-leather-shoes',
    'شیک‌پوش',
    4.9,
    180,
    'کفش چرمی مردانه با کیفیت عالی و طراحی کلاسیک. مناسب برای مجالس و محیط کار.',
    (SELECT id FROM categories WHERE slug = 'bags-shoes' LIMIT 1),
    '["چرم", "رسمی", "کلاسیک"]'::jsonb,
    '/images/shoes.jpg',
    true,
    true,
    '["40", "41", "42", "43", "44", "45"]'::jsonb
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'men-leather-shoes');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, tags, image, is_new, is_featured, sizes)
SELECT 
    'دستبند نقره دست‌ساز',
    'handmade-silver-bracelet',
    'شیک‌پوش',
    4.4,
    75,
    'دستبند نقره دست‌ساز با طراحی سنتی و مدرن. مناسب برای هدیه.',
    (SELECT id FROM categories WHERE slug = 'jewelry' LIMIT 1),
    '["نقره", "دست‌ساز", "سنتی"]'::jsonb,
    '/images/jewelry.jpg',
    false,
    true,
    '[]'::jsonb
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'handmade-silver-bracelet');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, tags, image, is_new, is_featured, sizes)
SELECT 
    'شال حریر ابریشمی',
    'silk-scarf',
    'شیک‌پوش',
    4.3,
    60,
    'شال حریر ابریشمی با طرح‌های زیبا و رنگ‌بندی متنوع. مناسب برای تمام فصول.',
    (SELECT id FROM categories WHERE slug = 'accessories' LIMIT 1),
    '["ابریشم", "شال", "لوکس"]'::jsonb,
    '/images/harir.jpeg',
    true,
    true,
    '[]'::jsonb
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'silk-scarf');

-- Add product details (prices) for featured products
INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT 
    p.id,
    2500000.00,
    3000000.00,
    15,
    17
FROM products p
WHERE p.slug = 'women-classic-suit'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT 
    p.id,
    1800000.00,
    2200000.00,
    20,
    18
FROM products p
WHERE p.slug = 'women-leather-handbag'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT 
    p.id,
    15000000.00,
    18000000.00,
    5,
    17
FROM products p
WHERE p.slug = 'gold-necklace-with-gem'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT 
    p.id,
    3500000.00,
    4200000.00,
    12,
    17
FROM products p
WHERE p.slug = 'men-formal-suit'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT 
    p.id,
    1200000.00,
    1500000.00,
    25,
    20
FROM products p
WHERE p.slug = 'men-leather-shoes'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT 
    p.id,
    800000.00,
    1000000.00,
    18,
    20
FROM products p
WHERE p.slug = 'handmade-silver-bracelet'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT 
    p.id,
    450000.00,
    550000.00,
    30,
    18
FROM products p
WHERE p.slug = 'silk-scarf'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

-- migrate:down
-- Remove seed data
DELETE FROM product_details WHERE product_id IN (
    SELECT id FROM products WHERE slug IN (
        'women-classic-suit',
        'women-leather-handbag',
        'gold-necklace-with-gem',
        'men-formal-suit',
        'men-leather-shoes',
        'handmade-silver-bracelet',
        'silk-scarf'
    )
);

DELETE FROM products WHERE slug IN (
    'women-classic-suit',
    'women-leather-handbag',
    'gold-necklace-with-gem',
    'men-formal-suit',
    'men-leather-shoes',
    'handmade-silver-bracelet',
    'silk-scarf'
);

DELETE FROM categories WHERE slug IN (
    'women-clothing',
    'bags-shoes',
    'jewelry',
    'men-clothing',
    'accessories'
);

