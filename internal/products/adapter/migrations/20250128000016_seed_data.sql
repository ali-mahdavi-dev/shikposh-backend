-- migrate:up
-- Seed data for categories
INSERT INTO categories (name, slug, description, image) VALUES
('لباس زنانه', 'women-clothing', 'لباس‌های زنانه شامل انواع پوشاک روزمره و رسمی', 'https://picsum.photos/seed/category-women/800/800'),
('کیف و کفش', 'bags-shoes', 'کیف و کفش زنانه و مردانه', 'https://picsum.photos/seed/category-bags/800/800'),
('زیورآلات', 'jewelry', 'انواع زیورآلات طلا و نقره', 'https://picsum.photos/seed/category-jewelry/800/800'),
('پوشاک مردانه', 'men-clothing', 'لباس‌های مردانه شامل کت و شلوار و پوشاک روزمره', 'https://picsum.photos/seed/category-men/800/800'),
('اکسسوری', 'accessories', 'انواع اکسسوری و لوازم جانبی', 'https://picsum.photos/seed/category-accessories/800/800')
ON CONFLICT (slug) DO NOTHING;

-- Seed data for 30 products
-- لباس زنانه (8 products)
INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'کت و شلوار زنانه کلاسیک',
    'women-classic-suit-1',
    'شیک‌پوش',
    4.8,
    120,
    'کت و شلوار زنانه کلاسیک با طراحی مدرن و کیفیت عالی. مناسب برای مجالس و محیط‌های کاری.',
    (SELECT id FROM categories WHERE slug = 'women-clothing' LIMIT 1),
    'https://picsum.photos/seed/fashion-women-suit/800/800',
    false,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'women-classic-suit-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'لباس مجلسی زنانه ابریشمی',
    'women-silk-dress-1',
    'الیت',
    4.6,
    95,
    'لباس مجلسی ابریشمی با طراحی لوکس و رنگ‌بندی زیبا. مناسب برای مهمانی‌های رسمی.',
    (SELECT id FROM categories WHERE slug = 'women-clothing' LIMIT 1),
    'https://picsum.photos/seed/fashion-women-suit/800/800',
    true,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'women-silk-dress-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'شلوارک زنانه کتان',
    'women-cotton-shorts-1',
    'کاج',
    4.2,
    78,
    'شلوارک کتان راحت و خنک برای فصل تابستان. مناسب برای استایل کژوال.',
    (SELECT id FROM categories WHERE slug = 'women-clothing' LIMIT 1),
    'https://picsum.photos/seed/fashion-women-dress/800/800',
    false,
    false
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'women-cotton-shorts-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'بلوز زنانه یقه‌دار',
    'women-collared-blouse-1',
    'شیک‌پوش',
    4.5,
    112,
    'بلوز یقه‌دار با پارچه مرغوب و برش عالی. مناسب برای محیط کار و مجالس.',
    (SELECT id FROM categories WHERE slug = 'women-clothing' LIMIT 1),
    'https://picsum.photos/seed/fashion-women-shorts/800/800',
    true,
    false
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'women-collared-blouse-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'دامن زنانه مدادی',
    'women-pencil-skirt-1',
    'الیت',
    4.7,
    89,
    'دامن مدادی کلاسیک با پارچه باکیفیت. مناسب برای استایل رسمی و کاری.',
    (SELECT id FROM categories WHERE slug = 'women-clothing' LIMIT 1),
    'https://picsum.photos/seed/fashion-women-blouse/800/800',
    false,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'women-pencil-skirt-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'کت زنانه پشمی',
    'women-wool-coat-1',
    'کاج',
    4.4,
    134,
    'کت پشمی گرم و شیک برای فصل زمستان. با طراحی مدرن و رنگ‌بندی متنوع.',
    (SELECT id FROM categories WHERE slug = 'women-clothing' LIMIT 1),
    'https://picsum.photos/seed/fashion-women-skirt/800/800',
    false,
    false
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'women-wool-coat-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'تی‌شرت زنانه ساده',
    'women-basic-tshirt-1',
    'شیک‌پوش',
    4.1,
    156,
    'تی‌شرت ساده و راحت با پارچه نخی. مناسب برای استفاده روزمره.',
    (SELECT id FROM categories WHERE slug = 'women-clothing' LIMIT 1),
    'https://picsum.photos/seed/fashion-women-coat/800/800',
    false,
    false
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'women-basic-tshirt-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'لباس مجلسی زنانه شب',
    'women-evening-dress-1',
    'الیت',
    4.9,
    67,
    'لباس مجلسی شب با طراحی لوکس و پارچه ابریشمی. مناسب برای مهمانی‌های خاص.',
    (SELECT id FROM categories WHERE slug = 'women-clothing' LIMIT 1),
    'https://picsum.photos/seed/fashion-women-tshirt/800/800',
    true,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'women-evening-dress-1');

-- کیف و کفش (7 products)
INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'کیف دستی چرمی زنانه',
    'women-leather-handbag-1',
    'شیک‌پوش',
    4.3,
    95,
    'کیف دستی چرمی با کیفیت بالا و طراحی شیک. مناسب برای استفاده روزمره و مجالس.',
    (SELECT id FROM categories WHERE slug = 'bags-shoes' LIMIT 1),
    'https://picsum.photos/seed/fashion-women-evening/800/800',
    false,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'women-evening-dress-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'کفش پاشنه‌بلند زنانه',
    'women-heels-1',
    'الیت',
    4.6,
    123,
    'کفش پاشنه‌بلند با طراحی مدرن و راحت. مناسب برای مجالس و مهمانی‌ها.',
    (SELECT id FROM categories WHERE slug = 'bags-shoes' LIMIT 1),
    'https://picsum.photos/seed/fashion-women-evening/800/800',
    true,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'women-heels-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'کفش چرمی مردانه',
    'men-leather-shoes-1',
    'شیک‌پوش',
    4.9,
    180,
    'کفش چرمی مردانه با کیفیت عالی و طراحی کلاسیک. مناسب برای مجالس و محیط کار.',
    (SELECT id FROM categories WHERE slug = 'bags-shoes' LIMIT 1),
    'https://picsum.photos/seed/fashion-shoes-heels/800/800',
    true,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'men-leather-shoes-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'کیف کمری چرمی',
    'leather-belt-bag-1',
    'کاج',
    4.0,
    64,
    'کیف کمری چرمی با طراحی مدرن و کاربردی. مناسب برای سفر و استفاده روزمره.',
    (SELECT id FROM categories WHERE slug = 'bags-shoes' LIMIT 1),
    'https://picsum.photos/seed/fashion-shoes-leather/800/800',
    false,
    false
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'leather-belt-bag-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'کفش ورزشی مردانه',
    'men-sneakers-1',
    'الیت',
    4.7,
    201,
    'کفش ورزشی با کیفیت بالا و طراحی مدرن. مناسب برای پیاده‌روی و ورزش.',
    (SELECT id FROM categories WHERE slug = 'bags-shoes' LIMIT 1),
    'https://picsum.photos/seed/fashion-bag-belt/800/800',
    false,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'men-sneakers-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'کیف شانه‌ای زنانه',
    'women-shoulder-bag-1',
    'شیک‌پوش',
    4.4,
    88,
    'کیف شانه‌ای با طراحی شیک و کاربردی. مناسب برای استفاده روزمره.',
    (SELECT id FROM categories WHERE slug = 'bags-shoes' LIMIT 1),
    'https://picsum.photos/seed/fashion-shoes-sneakers/800/800',
    true,
    false
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'women-shoulder-bag-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'کفش کتانی زنانه',
    'women-sneakers-1',
    'کاج',
    4.5,
    145,
    'کفش کتانی راحت و شیک با طراحی مدرن. مناسب برای پیاده‌روی و استفاده روزمره.',
    (SELECT id FROM categories WHERE slug = 'bags-shoes' LIMIT 1),
    'https://picsum.photos/seed/fashion-bag-shoulder/800/800',
    false,
    false
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'women-sneakers-1');

-- زیورآلات (6 products)
INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'گردنبند طلا با نگین',
    'gold-necklace-with-gem-1',
    'شیک‌پوش',
    4.7,
    150,
    'گردنبند طلای 18 عیار با نگین الماس. طراحی مدرن و لوکس.',
    (SELECT id FROM categories WHERE slug = 'jewelry' LIMIT 1),
    'https://picsum.photos/seed/fashion-shoes-women/800/800',
    true,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'women-sneakers-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'دستبند نقره دست‌ساز',
    'handmade-silver-bracelet-1',
    'الیت',
    4.4,
    75,
    'دستبند نقره دست‌ساز با طراحی سنتی و مدرن. مناسب برای هدیه.',
    (SELECT id FROM categories WHERE slug = 'jewelry' LIMIT 1),
    'https://picsum.photos/seed/fashion-shoes-women/800/800',
    false,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'handmade-silver-bracelet-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'گوشواره طلا کلاسیک',
    'gold-earrings-classic-1',
    'شیک‌پوش',
    4.6,
    98,
    'گوشواره طلای 18 عیار با طراحی کلاسیک و زیبا. مناسب برای مجالس.',
    (SELECT id FROM categories WHERE slug = 'jewelry' LIMIT 1),
    'https://picsum.photos/seed/jewelry-bracelet/800/800',
    true,
    false
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'gold-earrings-classic-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'انگشتر نقره با نگین',
    'silver-ring-with-stone-1',
    'کاج',
    4.3,
    112,
    'انگشتر نقره با نگین طبیعی. طراحی مدرن و زیبا.',
    (SELECT id FROM categories WHERE slug = 'jewelry' LIMIT 1),
    'https://picsum.photos/seed/jewelry-earrings/800/800',
    false,
    false
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'silver-ring-with-stone-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'زنجیر طلا ساده',
    'simple-gold-chain-1',
    'الیت',
    4.5,
    134,
    'زنجیر طلای 18 عیار با طراحی ساده و کلاسیک. مناسب برای استفاده روزمره.',
    (SELECT id FROM categories WHERE slug = 'jewelry' LIMIT 1),
    'https://picsum.photos/seed/jewelry-ring/800/800',
    false,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'simple-gold-chain-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'ست زیورآلات طلا',
    'gold-jewelry-set-1',
    'شیک‌پوش',
    4.8,
    67,
    'ست کامل زیورآلات طلا شامل گردنبند، گوشواره و دستبند. طراحی هماهنگ و لوکس.',
    (SELECT id FROM categories WHERE slug = 'jewelry' LIMIT 1),
    'https://picsum.photos/seed/jewelry-chain/800/800',
    true,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'gold-jewelry-set-1');

-- پوشاک مردانه (5 products)
INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'کت و شلوار مردانه رسمی',
    'men-formal-suit-1',
    'شیک‌پوش',
    4.6,
    200,
    'کت و شلوار مردانه رسمی با برش عالی و پارچه مرغوب. مناسب برای مجالس و محیط کار.',
    (SELECT id FROM categories WHERE slug = 'men-clothing' LIMIT 1),
    'https://picsum.photos/seed/jewelry-set/800/800',
    true,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'gold-jewelry-set-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'پیراهن مردانه رسمی',
    'men-formal-shirt-1',
    'الیت',
    4.4,
    178,
    'پیراهن رسمی با پارچه مرغوب و برش عالی. مناسب برای محیط کار و مجالس.',
    (SELECT id FROM categories WHERE slug = 'men-clothing' LIMIT 1),
    'https://picsum.photos/seed/jewelry-set/800/800',
    false,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'men-formal-shirt-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'شلوار جین مردانه',
    'men-jeans-1',
    'کاج',
    4.5,
    223,
    'شلوار جین با کیفیت بالا و طراحی مدرن. مناسب برای استایل کژوال.',
    (SELECT id FROM categories WHERE slug = 'men-clothing' LIMIT 1),
    'https://picsum.photos/seed/fashion-men-shirt/800/800',
    false,
    false
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'men-jeans-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'تی‌شرت مردانه ساده',
    'men-basic-tshirt-1',
    'شیک‌پوش',
    4.2,
    189,
    'تی‌شرت ساده و راحت با پارچه نخی. مناسب برای استفاده روزمره.',
    (SELECT id FROM categories WHERE slug = 'men-clothing' LIMIT 1),
    'https://picsum.photos/seed/fashion-men-jeans/800/800',
    false,
    false
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'men-basic-tshirt-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'کت مردانه پشمی',
    'men-wool-coat-1',
    'الیت',
    4.7,
    156,
    'کت پشمی گرم و شیک برای فصل زمستان. با طراحی کلاسیک و رنگ‌بندی متنوع.',
    (SELECT id FROM categories WHERE slug = 'men-clothing' LIMIT 1),
    'https://picsum.photos/seed/fashion-men-tshirt/800/800',
    true,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'men-wool-coat-1');

-- اکسسوری (4 products)
INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'شال حریر ابریشمی',
    'silk-scarf-1',
    'شیک‌پوش',
    4.3,
    60,
    'شال حریر ابریشمی با طرح‌های زیبا و رنگ‌بندی متنوع. مناسب برای تمام فصول.',
    (SELECT id FROM categories WHERE slug = 'accessories' LIMIT 1),
    'https://picsum.photos/seed/fashion-men-coat/800/800',
    true,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'men-wool-coat-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'کمربند چرمی مردانه',
    'men-leather-belt-1',
    'کاج',
    4.4,
    92,
    'کمربند چرمی با کیفیت بالا و طراحی کلاسیک. مناسب برای استفاده روزمره.',
    (SELECT id FROM categories WHERE slug = 'accessories' LIMIT 1),
    'https://picsum.photos/seed/fashion-men-coat/800/800',
    false,
    false
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'men-leather-belt-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'عینک آفتابی',
    'sunglasses-1',
    'الیت',
    4.6,
    145,
    'عینک آفتابی با طراحی مدرن و محافظت کامل از چشم. مناسب برای تمام فصول.',
    (SELECT id FROM categories WHERE slug = 'accessories' LIMIT 1),
    'https://picsum.photos/seed/fashion-belt/800/800',
    true,
    true
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'sunglasses-1');

INSERT INTO products (name, slug, brand, rating, review_count, description, category_id, image, is_new, is_featured)
SELECT 
    'کلاه بافتنی',
    'knitted-hat-1',
    'شیک‌پوش',
    4.1,
    78,
    'کلاه بافتنی گرم و راحت برای فصل زمستان. با طراحی ساده و زیبا.',
    (SELECT id FROM categories WHERE slug = 'accessories' LIMIT 1),
    'https://picsum.photos/seed/fashion-sunglasses/800/800',
    false,
    false
WHERE NOT EXISTS (SELECT 1 FROM products WHERE slug = 'knitted-hat-1');
-- Add tags and sizes for all products
DO $$
DECLARE
    product_id_var BIGINT;
    tag_name TEXT;
    size_name TEXT;
    tag_slug TEXT;
    size_slug TEXT;
    tag_id_var BIGINT;
    size_id_var BIGINT;
BEGIN
    -- women-classic-suit-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-classic-suit-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['کلاسیک', 'رسمی', 'کت و شلوار'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            -- Check by name first (unique constraint is on name)
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                -- If not found by name, try by slug
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['S', 'M', 'L', 'XL'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- women-silk-dress-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-silk-dress-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['مجلسی', 'ابریشم', 'لوکس'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['S', 'M', 'L'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- women-cotton-shorts-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-cotton-shorts-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['کتان', 'راحت', 'تابستانی'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['S', 'M', 'L', 'XL'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- women-collared-blouse-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-collared-blouse-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['یقه‌دار', 'رسمی', 'کار'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['S', 'M', 'L'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- women-pencil-skirt-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-pencil-skirt-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['مدادی', 'کلاسیک', 'رسمی'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['S', 'M', 'L'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- women-wool-coat-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-wool-coat-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['پشمی', 'زمستانی', 'گرم'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['S', 'M', 'L', 'XL'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- women-basic-tshirt-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-basic-tshirt-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['ساده', 'راحت', 'روزمره'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['S', 'M', 'L', 'XL'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- women-evening-dress-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-evening-dress-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['مجلسی', 'شب', 'لوکس'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['S', 'M', 'L'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- women-leather-handbag-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-leather-handbag-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['چرم', 'کیف دستی', 'کلاسیک'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- women-heels-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-heels-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['پاشنه‌بلند', 'مجلسی', 'راحت'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['37', '38', '39', '40'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- men-leather-shoes-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'men-leather-shoes-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['چرم', 'رسمی', 'کلاسیک'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['40', '41', '42', '43', '44', '45'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- leather-belt-bag-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'leather-belt-bag-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['چرم', 'کمری', 'کاربردی'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- men-sneakers-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'men-sneakers-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['ورزشی', 'راحت', 'مدرن'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['40', '41', '42', '43', '44', '45'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- women-shoulder-bag-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-shoulder-bag-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['شانه‌ای', 'کاربردی', 'روزمره'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- women-sneakers-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-sneakers-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['کتانی', 'راحت', 'مدرن'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['37', '38', '39', '40'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- gold-necklace-with-gem-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'gold-necklace-with-gem-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['طلا', 'الماس', 'لوکس'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- handmade-silver-bracelet-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'handmade-silver-bracelet-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['نقره', 'دست‌ساز', 'سنتی'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- gold-earrings-classic-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'gold-earrings-classic-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['طلا', 'کلاسیک', 'مجلسی'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- silver-ring-with-stone-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'silver-ring-with-stone-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['نقره', 'نگین', 'مدرن'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- simple-gold-chain-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'simple-gold-chain-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['طلا', 'ساده', 'کلاسیک'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- gold-jewelry-set-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'gold-jewelry-set-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['طلا', 'ست', 'لوکس'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- men-formal-suit-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'men-formal-suit-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['رسمی', 'کت و شلوار', 'کلاسیک'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['44', '46', '48', '50', '52'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- men-formal-shirt-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'men-formal-shirt-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['رسمی', 'پیراهن', 'کار'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['L', 'XL', 'XXL'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- men-jeans-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'men-jeans-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['جین', 'کژوال', 'مدرن'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['30', '32', '34', '36', '38'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- men-basic-tshirt-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'men-basic-tshirt-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['ساده', 'راحت', 'روزمره'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['L', 'XL', 'XXL'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- men-wool-coat-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'men-wool-coat-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['پشمی', 'زمستانی', 'گرم'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['L', 'XL', 'XXL'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- silk-scarf-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'silk-scarf-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['ابریشم', 'شال', 'لوکس'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- men-leather-belt-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'men-leather-belt-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['چرم', 'کمربند', 'کلاسیک'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['100', '105', '110', '115'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- sunglasses-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'sunglasses-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['عینک', 'آفتابی', 'مدرن'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

    -- knitted-hat-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'knitted-hat-1';
    IF product_id_var IS NOT NULL THEN
        FOREACH tag_name IN ARRAY ARRAY['بافتنی', 'زمستانی', 'گرم'] LOOP
            tag_slug := lower(regexp_replace(tag_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL;
            IF tag_id_var IS NULL THEN
                SELECT id INTO tag_id_var FROM tags WHERE slug = tag_slug AND deleted_at IS NULL;
                IF tag_id_var IS NULL THEN
                    -- Use subquery approach to avoid conflict with partial unique index
                    INSERT INTO tags (name, slug)
                    SELECT tag_name, tag_slug
                    WHERE NOT EXISTS (SELECT 1 FROM tags WHERE name = tag_name AND deleted_at IS NULL)
                    RETURNING id INTO tag_id_var;
                    IF tag_id_var IS NULL THEN SELECT id INTO tag_id_var FROM tags WHERE name = tag_name AND deleted_at IS NULL; END IF;
                END IF;
            END IF;
            IF tag_id_var IS NOT NULL THEN
                INSERT INTO product_tags (product_id, tag_id) VALUES (product_id_var, tag_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
        FOREACH size_name IN ARRAY ARRAY['یک سایز'] LOOP
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) VALUES (size_name, size_slug) ON CONFLICT (slug) DO NOTHING RETURNING id INTO size_id_var;
                IF size_id_var IS NULL THEN SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL; END IF;
            END IF;
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id) VALUES (product_id_var, size_id_var) ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;
    END IF;

END $$;

-- Add product details (prices) for all products
-- لباس زنانه
INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 2500000.00, 3000000.00, 15, 17
FROM products p WHERE p.slug = 'women-classic-suit-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 3200000.00, 3800000.00, 12, 16
FROM products p WHERE p.slug = 'women-silk-dress-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 450000.00, 550000.00, 25, 18
FROM products p WHERE p.slug = 'women-cotton-shorts-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 850000.00, 1000000.00, 20, 15
FROM products p WHERE p.slug = 'women-collared-blouse-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 1200000.00, 1400000.00, 18, 14
FROM products p WHERE p.slug = 'women-pencil-skirt-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 2800000.00, 3300000.00, 10, 15
FROM products p WHERE p.slug = 'women-wool-coat-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 350000.00, 400000.00, 30, 13
FROM products p WHERE p.slug = 'women-basic-tshirt-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 4500000.00, 5500000.00, 8, 18
FROM products p WHERE p.slug = 'women-evening-dress-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

-- کیف و کفش
INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 1800000.00, 2200000.00, 20, 18
FROM products p WHERE p.slug = 'women-leather-handbag-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 1500000.00, 1800000.00, 15, 17
FROM products p WHERE p.slug = 'women-heels-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 1200000.00, 1500000.00, 25, 20
FROM products p WHERE p.slug = 'men-leather-shoes-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 650000.00, 800000.00, 18, 19
FROM products p WHERE p.slug = 'leather-belt-bag-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 2200000.00, 2600000.00, 22, 15
FROM products p WHERE p.slug = 'men-sneakers-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 950000.00, 1150000.00, 16, 17
FROM products p WHERE p.slug = 'women-shoulder-bag-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 1800000.00, 2200000.00, 20, 18
FROM products p WHERE p.slug = 'women-sneakers-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

-- زیورآلات
INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 15000000.00, 18000000.00, 5, 17
FROM products p WHERE p.slug = 'gold-necklace-with-gem-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 800000.00, 1000000.00, 18, 20
FROM products p WHERE p.slug = 'handmade-silver-bracelet-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 3200000.00, 3800000.00, 12, 16
FROM products p WHERE p.slug = 'gold-earrings-classic-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 550000.00, 650000.00, 25, 15
FROM products p WHERE p.slug = 'silver-ring-with-stone-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 4500000.00, 5300000.00, 10, 15
FROM products p WHERE p.slug = 'simple-gold-chain-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 25000000.00, 30000000.00, 3, 17
FROM products p WHERE p.slug = 'gold-jewelry-set-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

-- پوشاک مردانه
INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 3500000.00, 4200000.00, 12, 17
FROM products p WHERE p.slug = 'men-formal-suit-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 950000.00, 1150000.00, 20, 17
FROM products p WHERE p.slug = 'men-formal-shirt-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 1200000.00, 1450000.00, 28, 17
FROM products p WHERE p.slug = 'men-jeans-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 450000.00, 550000.00, 35, 18
FROM products p WHERE p.slug = 'men-basic-tshirt-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 3200000.00, 3800000.00, 14, 16
FROM products p WHERE p.slug = 'men-wool-coat-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

-- اکسسوری
INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 450000.00, 550000.00, 30, 18
FROM products p WHERE p.slug = 'silk-scarf-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 750000.00, 900000.00, 22, 17
FROM products p WHERE p.slug = 'men-leather-belt-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 1800000.00, 2200000.00, 18, 18
FROM products p WHERE p.slug = 'sunglasses-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

INSERT INTO product_details (product_id, price, original_price, stock, discount)
SELECT p.id, 280000.00, 350000.00, 25, 20
FROM products p WHERE p.slug = 'knitted-hat-1'
AND NOT EXISTS (SELECT 1 FROM product_details WHERE product_id = p.id);

-- Add product_details, variants, and images for all products
DO $$
DECLARE
    product_id_var BIGINT;
    detail_id_var BIGINT;
BEGIN
    -- Product: women-classic-suit-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-classic-suit-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 2500000, 2875000, 10, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'S', 2500000, 2875000, 10, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-classic-suit-1-black-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-classic-suit-1-black-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-classic-suit-1-black-S-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'M', 2500000, 2875000, 10, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-classic-suit-1-black-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-classic-suit-1-black-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-classic-suit-1-black-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'L', 2500000, 2875000, 10, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-classic-suit-1-black-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-classic-suit-1-black-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-classic-suit-1-black-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'XL', 2500000, 2875000, 10, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-classic-suit-1-black-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-classic-suit-1-black-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-classic-suit-1-black-XL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: آبی دریایی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 2500000, 2875000, 10, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: آبی دریایی / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'S', 2500000, 2875000, 10, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-classic-suit-1-navy-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-classic-suit-1-navy-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-classic-suit-1-navy-S-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی دریایی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'M', 2500000, 2875000, 10, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-classic-suit-1-navy-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-classic-suit-1-navy-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-classic-suit-1-navy-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی دریایی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'L', 2500000, 2875000, 10, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-classic-suit-1-navy-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-classic-suit-1-navy-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-classic-suit-1-navy-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی دریایی / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'XL', 2500000, 2875000, 10, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-classic-suit-1-navy-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-classic-suit-1-navy-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-classic-suit-1-navy-XL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: خاکستری
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 2500000, 2875000, 10, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: خاکستری / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'S', 2500000, 2875000, 10, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-classic-suit-1-gray-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-classic-suit-1-gray-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-classic-suit-1-gray-S-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'M', 2500000, 2875000, 10, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-classic-suit-1-gray-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-classic-suit-1-gray-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-classic-suit-1-gray-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'L', 2500000, 2875000, 10, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-classic-suit-1-gray-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-classic-suit-1-gray-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-classic-suit-1-gray-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'XL', 2500000, 2875000, 10, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-classic-suit-1-gray-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-classic-suit-1-gray-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-classic-suit-1-gray-XL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: women-silk-dress-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-silk-dress-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: قرمز
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'red', 'قرمز', 1800000, NULL, 15, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: قرمز / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'red', 'قرمز', 'S', 1800000, NULL, 15, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'red' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-silk-dress-1-red-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-silk-dress-1-red-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-silk-dress-1-red-S-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-silk-dress-1-red-S-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: قرمز / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'red', 'قرمز', 'M', 1800000, NULL, 15, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'red' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-silk-dress-1-red-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-silk-dress-1-red-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-silk-dress-1-red-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-silk-dress-1-red-M-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: قرمز / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'red', 'قرمز', 'L', 1800000, NULL, 15, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'red' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-silk-dress-1-red-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-silk-dress-1-red-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-silk-dress-1-red-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-silk-dress-1-red-L-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: آبی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'blue', 'آبی', 1800000, NULL, 15, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: آبی / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'blue', 'آبی', 'S', 1800000, NULL, 15, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'blue' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-silk-dress-1-blue-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-silk-dress-1-blue-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-silk-dress-1-blue-S-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-silk-dress-1-blue-S-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'blue', 'آبی', 'M', 1800000, NULL, 15, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'blue' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-silk-dress-1-blue-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-silk-dress-1-blue-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-silk-dress-1-blue-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-silk-dress-1-blue-M-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'blue', 'آبی', 'L', 1800000, NULL, 15, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'blue' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-silk-dress-1-blue-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-silk-dress-1-blue-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-silk-dress-1-blue-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-silk-dress-1-blue-L-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: صورتی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', 1800000, NULL, 15, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: صورتی / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', 'S', 1800000, NULL, 15, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'pink' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-silk-dress-1-pink-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-silk-dress-1-pink-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-silk-dress-1-pink-S-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-silk-dress-1-pink-S-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: صورتی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', 'M', 1800000, NULL, 15, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'pink' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-silk-dress-1-pink-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-silk-dress-1-pink-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-silk-dress-1-pink-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-silk-dress-1-pink-M-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: صورتی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', 'L', 1800000, NULL, 15, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'pink' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-silk-dress-1-pink-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-silk-dress-1-pink-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-silk-dress-1-pink-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-silk-dress-1-pink-L-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: women-cotton-shorts-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-cotton-shorts-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: سفید
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 450000, 540000, 20, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: سفید / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'S', 450000, 540000, 20, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-white-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-white-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'M', 450000, 540000, 20, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-white-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-white-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'L', 450000, 540000, 20, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-white-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-white-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'XL', 450000, 540000, 20, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-white-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-white-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: بژ
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'beige', 'بژ', 450000, 540000, 20, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: بژ / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'beige', 'بژ', 'S', 450000, 540000, 20, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'beige' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-beige-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-beige-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: بژ / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'beige', 'بژ', 'M', 450000, 540000, 20, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'beige' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-beige-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-beige-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: بژ / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'beige', 'بژ', 'L', 450000, 540000, 20, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'beige' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-beige-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-beige-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: بژ / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'beige', 'بژ', 'XL', 450000, 540000, 20, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'beige' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-beige-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-beige-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: خاکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'khaki', 'خاکی', 450000, 540000, 20, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: خاکی / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'khaki', 'خاکی', 'S', 450000, 540000, 20, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'khaki' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-khaki-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-khaki-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'khaki', 'خاکی', 'M', 450000, 540000, 20, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'khaki' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-khaki-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-khaki-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'khaki', 'خاکی', 'L', 450000, 540000, 20, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'khaki' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-khaki-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-khaki-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکی / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'khaki', 'خاکی', 'XL', 450000, 540000, 20, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'khaki' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-khaki-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-cotton-shorts-1-khaki-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: women-collared-blouse-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-collared-blouse-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: سفید
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 650000, 715000, 12, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: سفید / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'S', 650000, 715000, 12, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-white-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-white-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-white-S-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'M', 650000, 715000, 12, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-white-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-white-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-white-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'L', 650000, 715000, 12, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-white-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-white-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-white-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: کرم
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'cream', 'کرم', 650000, 715000, 12, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: کرم / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'cream', 'کرم', 'S', 650000, 715000, 12, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'cream' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-cream-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-cream-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-cream-S-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: کرم / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'cream', 'کرم', 'M', 650000, 715000, 12, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'cream' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-cream-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-cream-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-cream-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: کرم / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'cream', 'کرم', 'L', 650000, 715000, 12, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'cream' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-cream-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-cream-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-cream-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: آبی روشن
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'light-blue', 'آبی روشن', 650000, 715000, 12, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: آبی روشن / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'light-blue', 'آبی روشن', 'S', 650000, 715000, 12, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'light-blue' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-light-blue-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-light-blue-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-light-blue-S-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی روشن / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'light-blue', 'آبی روشن', 'M', 650000, 715000, 12, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'light-blue' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-light-blue-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-light-blue-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-light-blue-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی روشن / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'light-blue', 'آبی روشن', 'L', 650000, 715000, 12, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'light-blue' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-light-blue-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-light-blue-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-collared-blouse-1-light-blue-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: women-pencil-skirt-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-pencil-skirt-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 850000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'S', 850000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-black-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-black-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-black-S-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'M', 850000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-black-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-black-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-black-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'L', 850000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-black-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-black-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-black-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: آبی دریایی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 850000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: آبی دریایی / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'S', 850000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-navy-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-navy-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-navy-S-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی دریایی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'M', 850000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-navy-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-navy-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-navy-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی دریایی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'L', 850000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-navy-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-navy-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-navy-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: قهوه‌ای
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', 850000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: قهوه‌ای / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', 'S', 850000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'brown' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-brown-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-brown-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-brown-S-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: قهوه‌ای / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', 'M', 850000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'brown' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-brown-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-brown-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-brown-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: قهوه‌ای / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', 'L', 850000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'brown' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-brown-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-brown-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-pencil-skirt-1-brown-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: women-wool-coat-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-wool-coat-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 3200000, 4000000, 5, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'M', 3200000, 4000000, 5, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-wool-coat-1-black-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-wool-coat-1-black-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-wool-coat-1-black-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-wool-coat-1-black-M-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'L', 3200000, 4000000, 5, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-wool-coat-1-black-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-wool-coat-1-black-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-wool-coat-1-black-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-wool-coat-1-black-L-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'XL', 3200000, 4000000, 5, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-wool-coat-1-black-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-wool-coat-1-black-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-wool-coat-1-black-XL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-wool-coat-1-black-XL-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: شتری
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'camel', 'شتری', 3200000, 4000000, 5, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: شتری / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'camel', 'شتری', 'M', 3200000, 4000000, 5, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'camel' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-wool-coat-1-camel-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-wool-coat-1-camel-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-wool-coat-1-camel-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-wool-coat-1-camel-M-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: شتری / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'camel', 'شتری', 'L', 3200000, 4000000, 5, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'camel' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-wool-coat-1-camel-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-wool-coat-1-camel-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-wool-coat-1-camel-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-wool-coat-1-camel-L-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: شتری / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'camel', 'شتری', 'XL', 3200000, 4000000, 5, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'camel' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-wool-coat-1-camel-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-wool-coat-1-camel-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-wool-coat-1-camel-XL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-wool-coat-1-camel-XL-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: خاکستری
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 3200000, 4000000, 5, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: خاکستری / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'M', 3200000, 4000000, 5, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-wool-coat-1-gray-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-wool-coat-1-gray-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-wool-coat-1-gray-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-wool-coat-1-gray-M-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'L', 3200000, 4000000, 5, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-wool-coat-1-gray-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-wool-coat-1-gray-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-wool-coat-1-gray-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-wool-coat-1-gray-L-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'XL', 3200000, 4000000, 5, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-wool-coat-1-gray-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-wool-coat-1-gray-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-wool-coat-1-gray-XL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-wool-coat-1-gray-XL-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: women-basic-tshirt-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-basic-tshirt-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: سفید
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: سفید / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'S', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-white-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-white-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'M', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-white-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-white-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'L', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-white-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-white-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'XL', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-white-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-white-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'S', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-black-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-black-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'M', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-black-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-black-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'L', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-black-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-black-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'XL', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-black-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-black-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: خاکستری
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: خاکستری / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'S', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-gray-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-gray-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'M', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-gray-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-gray-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'L', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-gray-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-gray-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'XL', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-gray-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-gray-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: صورتی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: صورتی / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', 'S', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'pink' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-pink-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-pink-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: صورتی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', 'M', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'pink' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-pink-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-pink-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: صورتی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', 'L', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'pink' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-pink-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-pink-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: صورتی / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', 'XL', 280000, NULL, 25, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'pink' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-pink-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-basic-tshirt-1-pink-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: women-evening-dress-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-evening-dress-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 4500000, 5850000, 3, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'S', 4500000, 5850000, 3, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-evening-dress-1-black-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-evening-dress-1-black-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-evening-dress-1-black-S-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-evening-dress-1-black-S-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-5.jpg', 'https://picsum.photos/seed/women-evening-dress-1-black-S-4/800/800', 4)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'M', 4500000, 5850000, 3, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-evening-dress-1-black-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-evening-dress-1-black-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-evening-dress-1-black-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-evening-dress-1-black-M-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-5.jpg', 'https://picsum.photos/seed/women-evening-dress-1-black-M-4/800/800', 4)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'L', 4500000, 5850000, 3, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-evening-dress-1-black-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-evening-dress-1-black-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-evening-dress-1-black-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-evening-dress-1-black-L-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-5.jpg', 'https://picsum.photos/seed/women-evening-dress-1-black-L-4/800/800', 4)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: آبی دریایی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 4500000, 5850000, 3, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: آبی دریایی / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'S', 4500000, 5850000, 3, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-evening-dress-1-navy-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-evening-dress-1-navy-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-evening-dress-1-navy-S-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-evening-dress-1-navy-S-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-5.jpg', 'https://picsum.photos/seed/women-evening-dress-1-navy-S-4/800/800', 4)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی دریایی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'M', 4500000, 5850000, 3, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-evening-dress-1-navy-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-evening-dress-1-navy-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-evening-dress-1-navy-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-evening-dress-1-navy-M-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-5.jpg', 'https://picsum.photos/seed/women-evening-dress-1-navy-M-4/800/800', 4)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی دریایی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'L', 4500000, 5850000, 3, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-evening-dress-1-navy-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-evening-dress-1-navy-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-evening-dress-1-navy-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-evening-dress-1-navy-L-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-5.jpg', 'https://picsum.photos/seed/women-evening-dress-1-navy-L-4/800/800', 4)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: زرشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'burgundy', 'زرشکی', 4500000, 5850000, 3, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: زرشکی / S
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'burgundy', 'زرشکی', 'S', 4500000, 5850000, 3, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'burgundy' AND product_details.size_key = 'S' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-evening-dress-1-burgundy-S-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-evening-dress-1-burgundy-S-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-evening-dress-1-burgundy-S-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-evening-dress-1-burgundy-S-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-5.jpg', 'https://picsum.photos/seed/women-evening-dress-1-burgundy-S-4/800/800', 4)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: زرشکی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'burgundy', 'زرشکی', 'M', 4500000, 5850000, 3, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'burgundy' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-evening-dress-1-burgundy-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-evening-dress-1-burgundy-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-evening-dress-1-burgundy-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-evening-dress-1-burgundy-M-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-5.jpg', 'https://picsum.photos/seed/women-evening-dress-1-burgundy-M-4/800/800', 4)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: زرشکی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'burgundy', 'زرشکی', 'L', 4500000, 5850000, 3, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'burgundy' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-evening-dress-1-burgundy-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-evening-dress-1-burgundy-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-evening-dress-1-burgundy-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-evening-dress-1-burgundy-L-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-5.jpg', 'https://picsum.photos/seed/women-evening-dress-1-burgundy-L-4/800/800', 4)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: women-leather-handbag-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-leather-handbag-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: قهوه‌ای
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', 2800000, 3219999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: قهوه‌ای / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', 'one-size', 2800000, 3219999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'brown' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-leather-handbag-1-brown-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-leather-handbag-1-brown-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-leather-handbag-1-brown-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-leather-handbag-1-brown-one-size-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 2800000, 3219999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'one-size', 2800000, 3219999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-leather-handbag-1-black-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-leather-handbag-1-black-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-leather-handbag-1-black-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-leather-handbag-1-black-one-size-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: برنزه
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'tan', 'برنزه', 2800000, 3219999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: برنزه / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'tan', 'برنزه', 'one-size', 2800000, 3219999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'tan' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-leather-handbag-1-tan-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-leather-handbag-1-tan-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-leather-handbag-1-tan-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-leather-handbag-1-tan-one-size-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: women-heels-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-heels-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / 36
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '36', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '36' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-heels-1-black-36-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-heels-1-black-36-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-heels-1-black-36-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 37
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '37', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '37' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-heels-1-black-37-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-heels-1-black-37-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-heels-1-black-37-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 38
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '38', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '38' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-heels-1-black-38-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-heels-1-black-38-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-heels-1-black-38-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 39
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '39', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '39' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-heels-1-black-39-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-heels-1-black-39-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-heels-1-black-39-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 40
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '40', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '40' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-heels-1-black-40-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-heels-1-black-40-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-heels-1-black-40-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: نود
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'nude', 'نود', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: نود / 36
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'nude', 'نود', '36', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'nude' AND product_details.size_key = '36' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-heels-1-nude-36-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-heels-1-nude-36-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-heels-1-nude-36-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: نود / 37
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'nude', 'نود', '37', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'nude' AND product_details.size_key = '37' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-heels-1-nude-37-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-heels-1-nude-37-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-heels-1-nude-37-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: نود / 38
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'nude', 'نود', '38', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'nude' AND product_details.size_key = '38' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-heels-1-nude-38-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-heels-1-nude-38-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-heels-1-nude-38-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: نود / 39
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'nude', 'نود', '39', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'nude' AND product_details.size_key = '39' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-heels-1-nude-39-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-heels-1-nude-39-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-heels-1-nude-39-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: نود / 40
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'nude', 'نود', '40', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'nude' AND product_details.size_key = '40' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-heels-1-nude-40-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-heels-1-nude-40-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-heels-1-nude-40-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: قرمز
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'red', 'قرمز', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: قرمز / 36
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'red', 'قرمز', '36', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'red' AND product_details.size_key = '36' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-heels-1-red-36-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-heels-1-red-36-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-heels-1-red-36-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: قرمز / 37
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'red', 'قرمز', '37', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'red' AND product_details.size_key = '37' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-heels-1-red-37-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-heels-1-red-37-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-heels-1-red-37-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: قرمز / 38
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'red', 'قرمز', '38', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'red' AND product_details.size_key = '38' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-heels-1-red-38-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-heels-1-red-38-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-heels-1-red-38-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: قرمز / 39
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'red', 'قرمز', '39', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'red' AND product_details.size_key = '39' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-heels-1-red-39-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-heels-1-red-39-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-heels-1-red-39-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: قرمز / 40
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'red', 'قرمز', '40', 1200000, 1440000, 6, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'red' AND product_details.size_key = '40' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-heels-1-red-40-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-heels-1-red-40-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-heels-1-red-40-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: men-leather-shoes-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'men-leather-shoes-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: قهوه‌ای
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', 3500000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: قهوه‌ای / 40
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', '40', 3500000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'brown' AND product_details.size_key = '40' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-brown-40-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-brown-40-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-brown-40-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: قهوه‌ای / 41
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', '41', 3500000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'brown' AND product_details.size_key = '41' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-brown-41-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-brown-41-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-brown-41-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: قهوه‌ای / 42
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', '42', 3500000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'brown' AND product_details.size_key = '42' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-brown-42-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-brown-42-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-brown-42-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: قهوه‌ای / 43
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', '43', 3500000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'brown' AND product_details.size_key = '43' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-brown-43-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-brown-43-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-brown-43-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: قهوه‌ای / 44
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', '44', 3500000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'brown' AND product_details.size_key = '44' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-brown-44-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-brown-44-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-brown-44-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 3500000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / 40
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '40', 3500000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '40' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-black-40-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-black-40-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-black-40-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 41
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '41', 3500000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '41' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-black-41-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-black-41-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-black-41-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 42
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '42', 3500000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '42' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-black-42-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-black-42-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-black-42-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 43
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '43', 3500000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '43' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-black-43-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-black-43-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-black-43-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 44
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '44', 3500000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '44' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-black-44-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-black-44-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-leather-shoes-1-black-44-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: leather-belt-bag-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'leather-belt-bag-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: قهوه‌ای
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', 950000, 1045000, 12, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: قهوه‌ای / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', 'one-size', 950000, 1045000, 12, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'brown' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/leather-belt-bag-1-brown-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/leather-belt-bag-1-brown-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/leather-belt-bag-1-brown-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 950000, 1045000, 12, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'one-size', 950000, 1045000, 12, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/leather-belt-bag-1-black-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/leather-belt-bag-1-black-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/leather-belt-bag-1-black-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: men-sneakers-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'men-sneakers-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: سفید
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: سفید / 40
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', '40', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = '40' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-40-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-40-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-40-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / 41
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', '41', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = '41' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-41-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-41-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-41-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / 42
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', '42', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = '42' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-42-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-42-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-42-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / 43
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', '43', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = '43' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-43-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-43-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-43-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / 44
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', '44', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = '44' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-44-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-44-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-44-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / 45
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', '45', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = '45' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-45-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-45-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-white-45-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / 40
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '40', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '40' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-40-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-40-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-40-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 41
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '41', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '41' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-41-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-41-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-41-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 42
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '42', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '42' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-42-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-42-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-42-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 43
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '43', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '43' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-43-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-43-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-43-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 44
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '44', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '44' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-44-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-44-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-44-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 45
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '45', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '45' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-45-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-45-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-black-45-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: خاکستری
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: خاکستری / 40
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', '40', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = '40' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-40-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-40-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-40-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / 41
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', '41', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = '41' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-41-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-41-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-41-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / 42
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', '42', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = '42' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-42-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-42-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-42-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / 43
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', '43', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = '43' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-43-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-43-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-43-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / 44
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', '44', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = '44' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-44-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-44-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-44-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / 45
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', '45', 1800000, 2250000, 10, 25)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = '45' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-45-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-45-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-sneakers-1-gray-45-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: women-shoulder-bag-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-shoulder-bag-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 1500000, NULL, 9, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'one-size', 1500000, NULL, 9, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-shoulder-bag-1-black-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-shoulder-bag-1-black-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-shoulder-bag-1-black-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-shoulder-bag-1-black-one-size-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: بژ
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'beige', 'بژ', 1500000, NULL, 9, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: بژ / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'beige', 'بژ', 'one-size', 1500000, NULL, 9, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'beige' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-shoulder-bag-1-beige-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-shoulder-bag-1-beige-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-shoulder-bag-1-beige-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-shoulder-bag-1-beige-one-size-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: قرمز
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'red', 'قرمز', 1500000, NULL, 9, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: قرمز / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'red', 'قرمز', 'one-size', 1500000, NULL, 9, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'red' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-shoulder-bag-1-red-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-shoulder-bag-1-red-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-shoulder-bag-1-red-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/women-shoulder-bag-1-red-one-size-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: women-sneakers-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'women-sneakers-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: سفید
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: سفید / 36
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', '36', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = '36' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-sneakers-1-white-36-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-sneakers-1-white-36-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-sneakers-1-white-36-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / 37
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', '37', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = '37' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-sneakers-1-white-37-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-sneakers-1-white-37-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-sneakers-1-white-37-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / 38
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', '38', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = '38' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-sneakers-1-white-38-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-sneakers-1-white-38-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-sneakers-1-white-38-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / 39
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', '39', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = '39' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-sneakers-1-white-39-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-sneakers-1-white-39-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-sneakers-1-white-39-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / 40
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', '40', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = '40' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-sneakers-1-white-40-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-sneakers-1-white-40-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-sneakers-1-white-40-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: صورتی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: صورتی / 36
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', '36', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'pink' AND product_details.size_key = '36' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-sneakers-1-pink-36-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-sneakers-1-pink-36-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-sneakers-1-pink-36-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: صورتی / 37
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', '37', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'pink' AND product_details.size_key = '37' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-sneakers-1-pink-37-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-sneakers-1-pink-37-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-sneakers-1-pink-37-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: صورتی / 38
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', '38', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'pink' AND product_details.size_key = '38' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-sneakers-1-pink-38-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-sneakers-1-pink-38-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-sneakers-1-pink-38-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: صورتی / 39
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', '39', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'pink' AND product_details.size_key = '39' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-sneakers-1-pink-39-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-sneakers-1-pink-39-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-sneakers-1-pink-39-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: صورتی / 40
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', '40', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'pink' AND product_details.size_key = '40' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-sneakers-1-pink-40-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-sneakers-1-pink-40-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-sneakers-1-pink-40-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / 36
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '36', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '36' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-sneakers-1-black-36-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-sneakers-1-black-36-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-sneakers-1-black-36-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 37
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '37', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '37' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-sneakers-1-black-37-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-sneakers-1-black-37-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-sneakers-1-black-37-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 38
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '38', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '38' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-sneakers-1-black-38-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-sneakers-1-black-38-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-sneakers-1-black-38-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 39
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '39', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '39' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-sneakers-1-black-39-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-sneakers-1-black-39-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-sneakers-1-black-39-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 40
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '40', 1600000, 1839999, 7, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '40' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/women-sneakers-1-black-40-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/women-sneakers-1-black-40-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/women-sneakers-1-black-40-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: gold-necklace-with-gem-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'gold-necklace-with-gem-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: طلایی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'gold', 'طلایی', 8500000, NULL, 4, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: طلایی / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gold', 'طلایی', 'one-size', 8500000, NULL, 4, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gold' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/gold-necklace-with-gem-1-gold-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/gold-necklace-with-gem-1-gold-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/gold-necklace-with-gem-1-gold-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/gold-necklace-with-gem-1-gold-one-size-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: طلای صورتی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'rose-gold', 'طلای صورتی', 8500000, NULL, 4, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: طلای صورتی / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'rose-gold', 'طلای صورتی', 'one-size', 8500000, NULL, 4, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'rose-gold' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/gold-necklace-with-gem-1-rose-gold-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/gold-necklace-with-gem-1-rose-gold-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/gold-necklace-with-gem-1-rose-gold-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/gold-necklace-with-gem-1-rose-gold-one-size-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: handmade-silver-bracelet-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'handmade-silver-bracelet-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: نقره‌ای
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'silver', 'نقره‌ای', 3200000, 3520000, 6, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: نقره‌ای / small
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'silver', 'نقره‌ای', 'small', 3200000, 3520000, 6, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'silver' AND product_details.size_key = 'small' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/handmade-silver-bracelet-1-silver-small-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/handmade-silver-bracelet-1-silver-small-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/handmade-silver-bracelet-1-silver-small-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: نقره‌ای / medium
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'silver', 'نقره‌ای', 'medium', 3200000, 3520000, 6, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'silver' AND product_details.size_key = 'medium' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/handmade-silver-bracelet-1-silver-medium-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/handmade-silver-bracelet-1-silver-medium-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/handmade-silver-bracelet-1-silver-medium-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: نقره‌ای / large
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'silver', 'نقره‌ای', 'large', 3200000, 3520000, 6, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'silver' AND product_details.size_key = 'large' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/handmade-silver-bracelet-1-silver-large-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/handmade-silver-bracelet-1-silver-large-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/handmade-silver-bracelet-1-silver-large-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: gold-earrings-classic-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'gold-earrings-classic-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: طلایی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'gold', 'طلایی', 2800000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: طلایی / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gold', 'طلایی', 'one-size', 2800000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gold' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/gold-earrings-classic-1-gold-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/gold-earrings-classic-1-gold-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/gold-earrings-classic-1-gold-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: نقره‌ای
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'silver', 'نقره‌ای', 2800000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: نقره‌ای / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'silver', 'نقره‌ای', 'one-size', 2800000, NULL, 8, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'silver' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/gold-earrings-classic-1-silver-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/gold-earrings-classic-1-silver-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/gold-earrings-classic-1-silver-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: silver-ring-with-stone-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'silver-ring-with-stone-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: نقره‌ای
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'silver', 'نقره‌ای', 1800000, 2160000, 5, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: نقره‌ای / 16
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'silver', 'نقره‌ای', '16', 1800000, 2160000, 5, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'silver' AND product_details.size_key = '16' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/silver-ring-with-stone-1-silver-16-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/silver-ring-with-stone-1-silver-16-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/silver-ring-with-stone-1-silver-16-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/silver-ring-with-stone-1-silver-16-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: نقره‌ای / 17
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'silver', 'نقره‌ای', '17', 1800000, 2160000, 5, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'silver' AND product_details.size_key = '17' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/silver-ring-with-stone-1-silver-17-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/silver-ring-with-stone-1-silver-17-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/silver-ring-with-stone-1-silver-17-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/silver-ring-with-stone-1-silver-17-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: نقره‌ای / 18
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'silver', 'نقره‌ای', '18', 1800000, 2160000, 5, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'silver' AND product_details.size_key = '18' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/silver-ring-with-stone-1-silver-18-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/silver-ring-with-stone-1-silver-18-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/silver-ring-with-stone-1-silver-18-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/silver-ring-with-stone-1-silver-18-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: نقره‌ای / 19
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'silver', 'نقره‌ای', '19', 1800000, 2160000, 5, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'silver' AND product_details.size_key = '19' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/silver-ring-with-stone-1-silver-19-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/silver-ring-with-stone-1-silver-19-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/silver-ring-with-stone-1-silver-19-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/silver-ring-with-stone-1-silver-19-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: simple-gold-chain-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'simple-gold-chain-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: طلایی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'gold', 'طلایی', 4200000, NULL, 7, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: طلایی / 45cm
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gold', 'طلایی', '45cm', 4200000, NULL, 7, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gold' AND product_details.size_key = '45cm' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/simple-gold-chain-1-gold-45cm-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/simple-gold-chain-1-gold-45cm-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/simple-gold-chain-1-gold-45cm-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: طلایی / 50cm
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gold', 'طلایی', '50cm', 4200000, NULL, 7, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gold' AND product_details.size_key = '50cm' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/simple-gold-chain-1-gold-50cm-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/simple-gold-chain-1-gold-50cm-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/simple-gold-chain-1-gold-50cm-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: طلایی / 55cm
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gold', 'طلایی', '55cm', 4200000, NULL, 7, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gold' AND product_details.size_key = '55cm' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/simple-gold-chain-1-gold-55cm-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/simple-gold-chain-1-gold-55cm-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/simple-gold-chain-1-gold-55cm-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: طلایی / 60cm
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gold', 'طلایی', '60cm', 4200000, NULL, 7, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gold' AND product_details.size_key = '60cm' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/simple-gold-chain-1-gold-60cm-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/simple-gold-chain-1-gold-60cm-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/simple-gold-chain-1-gold-60cm-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: gold-jewelry-set-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'gold-jewelry-set-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: طلایی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'gold', 'طلایی', 12000000, 16200000, 2, 35)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: طلایی / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gold', 'طلایی', 'one-size', 12000000, 16200000, 2, 35)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gold' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/gold-jewelry-set-1-gold-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/gold-jewelry-set-1-gold-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/gold-jewelry-set-1-gold-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/gold-jewelry-set-1-gold-one-size-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-5.jpg', 'https://picsum.photos/seed/gold-jewelry-set-1-gold-one-size-4/800/800', 4)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: طلای صورتی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'rose-gold', 'طلای صورتی', 12000000, 16200000, 2, 35)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: طلای صورتی / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'rose-gold', 'طلای صورتی', 'one-size', 12000000, 16200000, 2, 35)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'rose-gold' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/gold-jewelry-set-1-rose-gold-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/gold-jewelry-set-1-rose-gold-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/gold-jewelry-set-1-rose-gold-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/gold-jewelry-set-1-rose-gold-one-size-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-5.jpg', 'https://picsum.photos/seed/gold-jewelry-set-1-rose-gold-one-size-4/800/800', 4)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: men-formal-suit-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'men-formal-suit-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: آبی دریایی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 4500000, NULL, 6, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: آبی دریایی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'M', 4500000, NULL, 6, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-suit-1-navy-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-suit-1-navy-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-suit-1-navy-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-formal-suit-1-navy-M-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی دریایی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'L', 4500000, NULL, 6, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-suit-1-navy-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-suit-1-navy-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-suit-1-navy-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-formal-suit-1-navy-L-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی دریایی / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'XL', 4500000, NULL, 6, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-suit-1-navy-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-suit-1-navy-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-suit-1-navy-XL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-formal-suit-1-navy-XL-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی دریایی / XXL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'XXL', 4500000, NULL, 6, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'XXL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-suit-1-navy-XXL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-suit-1-navy-XXL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-suit-1-navy-XXL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-formal-suit-1-navy-XXL-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 4500000, NULL, 6, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'M', 4500000, NULL, 6, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-suit-1-black-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-suit-1-black-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-suit-1-black-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-formal-suit-1-black-M-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'L', 4500000, NULL, 6, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-suit-1-black-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-suit-1-black-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-suit-1-black-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-formal-suit-1-black-L-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'XL', 4500000, NULL, 6, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-suit-1-black-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-suit-1-black-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-suit-1-black-XL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-formal-suit-1-black-XL-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / XXL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'XXL', 4500000, NULL, 6, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'XXL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-suit-1-black-XXL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-suit-1-black-XXL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-suit-1-black-XXL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-formal-suit-1-black-XXL-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: زغالی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'charcoal', 'زغالی', 4500000, NULL, 6, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: زغالی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'charcoal', 'زغالی', 'M', 4500000, NULL, 6, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'charcoal' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-suit-1-charcoal-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-suit-1-charcoal-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-suit-1-charcoal-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-formal-suit-1-charcoal-M-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: زغالی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'charcoal', 'زغالی', 'L', 4500000, NULL, 6, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'charcoal' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-suit-1-charcoal-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-suit-1-charcoal-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-suit-1-charcoal-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-formal-suit-1-charcoal-L-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: زغالی / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'charcoal', 'زغالی', 'XL', 4500000, NULL, 6, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'charcoal' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-suit-1-charcoal-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-suit-1-charcoal-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-suit-1-charcoal-XL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-formal-suit-1-charcoal-XL-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: زغالی / XXL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'charcoal', 'زغالی', 'XXL', 4500000, NULL, 6, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'charcoal' AND product_details.size_key = 'XXL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-suit-1-charcoal-XXL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-suit-1-charcoal-XXL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-suit-1-charcoal-XXL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-formal-suit-1-charcoal-XXL-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: men-formal-shirt-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'men-formal-shirt-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: سفید
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 1200000, 1380000, 15, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: سفید / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'M', 1200000, 1380000, 15, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-white-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-white-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-white-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'L', 1200000, 1380000, 15, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-white-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-white-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-white-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'XL', 1200000, 1380000, 15, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-white-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-white-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-white-XL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: آبی روشن
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'light-blue', 'آبی روشن', 1200000, 1380000, 15, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: آبی روشن / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'light-blue', 'آبی روشن', 'M', 1200000, 1380000, 15, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'light-blue' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-light-blue-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-light-blue-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-light-blue-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی روشن / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'light-blue', 'آبی روشن', 'L', 1200000, 1380000, 15, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'light-blue' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-light-blue-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-light-blue-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-light-blue-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی روشن / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'light-blue', 'آبی روشن', 'XL', 1200000, 1380000, 15, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'light-blue' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-light-blue-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-light-blue-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-light-blue-XL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: صورتی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', 1200000, 1380000, 15, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: صورتی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', 'M', 1200000, 1380000, 15, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'pink' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-pink-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-pink-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-pink-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: صورتی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', 'L', 1200000, 1380000, 15, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'pink' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-pink-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-pink-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-pink-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: صورتی / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', 'XL', 1200000, 1380000, 15, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'pink' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-pink-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-pink-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-formal-shirt-1-pink-XL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: men-jeans-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'men-jeans-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: آبی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'blue', 'آبی', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: آبی / 30
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'blue', 'آبی', '30', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'blue' AND product_details.size_key = '30' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-jeans-1-blue-30-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-jeans-1-blue-30-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-jeans-1-blue-30-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی / 32
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'blue', 'آبی', '32', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'blue' AND product_details.size_key = '32' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-jeans-1-blue-32-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-jeans-1-blue-32-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-jeans-1-blue-32-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی / 34
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'blue', 'آبی', '34', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'blue' AND product_details.size_key = '34' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-jeans-1-blue-34-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-jeans-1-blue-34-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-jeans-1-blue-34-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی / 36
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'blue', 'آبی', '36', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'blue' AND product_details.size_key = '36' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-jeans-1-blue-36-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-jeans-1-blue-36-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-jeans-1-blue-36-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی / 38
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'blue', 'آبی', '38', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'blue' AND product_details.size_key = '38' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-jeans-1-blue-38-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-jeans-1-blue-38-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-jeans-1-blue-38-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / 30
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '30', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '30' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-jeans-1-black-30-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-jeans-1-black-30-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-jeans-1-black-30-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 32
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '32', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '32' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-jeans-1-black-32-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-jeans-1-black-32-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-jeans-1-black-32-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 34
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '34', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '34' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-jeans-1-black-34-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-jeans-1-black-34-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-jeans-1-black-34-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 36
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '36', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '36' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-jeans-1-black-36-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-jeans-1-black-36-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-jeans-1-black-36-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 38
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '38', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '38' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-jeans-1-black-38-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-jeans-1-black-38-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-jeans-1-black-38-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: خاکستری
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: خاکستری / 30
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', '30', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = '30' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-jeans-1-gray-30-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-jeans-1-gray-30-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-jeans-1-gray-30-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / 32
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', '32', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = '32' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-jeans-1-gray-32-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-jeans-1-gray-32-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-jeans-1-gray-32-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / 34
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', '34', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = '34' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-jeans-1-gray-34-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-jeans-1-gray-34-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-jeans-1-gray-34-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / 36
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', '36', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = '36' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-jeans-1-gray-36-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-jeans-1-gray-36-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-jeans-1-gray-36-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / 38
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', '38', 1800000, 2160000, 12, 20)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = '38' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-jeans-1-gray-38-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-jeans-1-gray-38-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-jeans-1-gray-38-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: men-basic-tshirt-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'men-basic-tshirt-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: سفید
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: سفید / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'M', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-white-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-white-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'L', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-white-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-white-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'XL', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-white-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-white-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: سفید / XXL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'white', 'سفید', 'XXL', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'white' AND product_details.size_key = 'XXL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-white-XXL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-white-XXL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'M', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-black-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-black-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'L', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-black-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-black-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'XL', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-black-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-black-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / XXL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'XXL', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'XXL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-black-XXL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-black-XXL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: خاکستری
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: خاکستری / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'M', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-gray-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-gray-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'L', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-gray-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-gray-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'XL', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-gray-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-gray-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: خاکستری / XXL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'XXL', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'XXL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-gray-XXL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-gray-XXL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: آبی دریایی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: آبی دریایی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'M', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-navy-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-navy-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی دریایی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'L', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-navy-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-navy-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی دریایی / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'XL', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-navy-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-navy-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی دریایی / XXL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'XXL', 350000, NULL, 30, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'XXL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-navy-XXL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-basic-tshirt-1-navy-XXL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: men-wool-coat-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'men-wool-coat-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 3800000, 4940000, 4, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'M', 3800000, 4940000, 4, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-wool-coat-1-black-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-wool-coat-1-black-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-wool-coat-1-black-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-wool-coat-1-black-M-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'L', 3800000, 4940000, 4, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-wool-coat-1-black-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-wool-coat-1-black-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-wool-coat-1-black-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-wool-coat-1-black-L-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'XL', 3800000, 4940000, 4, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-wool-coat-1-black-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-wool-coat-1-black-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-wool-coat-1-black-XL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-wool-coat-1-black-XL-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: شتری
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'camel', 'شتری', 3800000, 4940000, 4, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: شتری / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'camel', 'شتری', 'M', 3800000, 4940000, 4, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'camel' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-wool-coat-1-camel-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-wool-coat-1-camel-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-wool-coat-1-camel-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-wool-coat-1-camel-M-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: شتری / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'camel', 'شتری', 'L', 3800000, 4940000, 4, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'camel' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-wool-coat-1-camel-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-wool-coat-1-camel-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-wool-coat-1-camel-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-wool-coat-1-camel-L-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: شتری / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'camel', 'شتری', 'XL', 3800000, 4940000, 4, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'camel' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-wool-coat-1-camel-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-wool-coat-1-camel-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-wool-coat-1-camel-XL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-wool-coat-1-camel-XL-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: آبی دریایی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 3800000, 4940000, 4, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: آبی دریایی / M
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'M', 3800000, 4940000, 4, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'M' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-wool-coat-1-navy-M-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-wool-coat-1-navy-M-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-wool-coat-1-navy-M-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-wool-coat-1-navy-M-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی دریایی / L
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'L', 3800000, 4940000, 4, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'L' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-wool-coat-1-navy-L-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-wool-coat-1-navy-L-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-wool-coat-1-navy-L-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-wool-coat-1-navy-L-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: آبی دریایی / XL
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'navy', 'آبی دریایی', 'XL', 3800000, 4940000, 4, 30)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'navy' AND product_details.size_key = 'XL' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-wool-coat-1-navy-XL-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-wool-coat-1-navy-XL-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/men-wool-coat-1-navy-XL-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-4.jpg', 'https://picsum.photos/seed/men-wool-coat-1-navy-XL-3/800/800', 3)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: silk-scarf-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'silk-scarf-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: چند رنگ
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'multicolor', 'چند رنگ', 650000, NULL, 18, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: چند رنگ / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'multicolor', 'چند رنگ', 'one-size', 650000, NULL, 18, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'multicolor' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/silk-scarf-1-multicolor-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/silk-scarf-1-multicolor-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/silk-scarf-1-multicolor-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: آبی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'blue', 'آبی', 650000, NULL, 18, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: آبی / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'blue', 'آبی', 'one-size', 650000, NULL, 18, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'blue' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/silk-scarf-1-blue-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/silk-scarf-1-blue-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/silk-scarf-1-blue-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: صورتی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', 650000, NULL, 18, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: صورتی / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'pink', 'صورتی', 'one-size', 650000, NULL, 18, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'pink' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/silk-scarf-1-pink-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/silk-scarf-1-pink-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/silk-scarf-1-pink-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: men-leather-belt-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'men-leather-belt-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: قهوه‌ای
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', 450000, 495000, 14, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: قهوه‌ای / 100cm
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', '100cm', 450000, 495000, 14, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'brown' AND product_details.size_key = '100cm' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-leather-belt-1-brown-100cm-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-leather-belt-1-brown-100cm-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: قهوه‌ای / 110cm
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', '110cm', 450000, 495000, 14, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'brown' AND product_details.size_key = '110cm' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-leather-belt-1-brown-110cm-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-leather-belt-1-brown-110cm-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: قهوه‌ای / 120cm
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', '120cm', 450000, 495000, 14, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'brown' AND product_details.size_key = '120cm' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-leather-belt-1-brown-120cm-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-leather-belt-1-brown-120cm-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 450000, 495000, 14, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / 100cm
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '100cm', 450000, 495000, 14, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '100cm' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-leather-belt-1-black-100cm-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-leather-belt-1-black-100cm-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 110cm
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '110cm', 450000, 495000, 14, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '110cm' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-leather-belt-1-black-110cm-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-leather-belt-1-black-110cm-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Variant: مشکی / 120cm
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', '120cm', 450000, 495000, 14, 10)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = '120cm' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/men-leather-belt-1-black-120cm-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/men-leather-belt-1-black-120cm-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: sunglasses-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'sunglasses-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 850000, NULL, 20, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'one-size', 850000, NULL, 20, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/sunglasses-1-black-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/sunglasses-1-black-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/sunglasses-1-black-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: قهوه‌ای
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', 850000, NULL, 20, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: قهوه‌ای / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'brown', 'قهوه‌ای', 'one-size', 850000, NULL, 20, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'brown' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/sunglasses-1-brown-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/sunglasses-1-brown-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/sunglasses-1-brown-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: خاکستری
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 850000, NULL, 20, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: خاکستری / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'one-size', 850000, NULL, 20, 0)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/sunglasses-1-gray-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/sunglasses-1-gray-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-3.jpg', 'https://picsum.photos/seed/sunglasses-1-gray-one-size-2/800/800', 2)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

    -- Product: knitted-hat-1
    SELECT id INTO product_id_var FROM products WHERE slug = 'knitted-hat-1' LIMIT 1;
    IF product_id_var IS NOT NULL THEN
        -- Color: خاکستری
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 280000, 322000, 25, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: خاکستری / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'gray', 'خاکستری', 'one-size', 280000, 322000, 25, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'gray' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/knitted-hat-1-gray-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/knitted-hat-1-gray-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: مشکی
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 280000, 322000, 25, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: مشکی / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'black', 'مشکی', 'one-size', 280000, 322000, 25, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'black' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/knitted-hat-1-black-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/knitted-hat-1-black-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
        -- Color: بژ
        INSERT INTO product_details (product_id, color_key, color_name, price, original_price, stock, discount)
        VALUES (product_id_var, 'beige', 'بژ', 280000, 322000, 25, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING;
        -- Variant: بژ / one-size
        INSERT INTO product_details (product_id, color_key, color_name, size_key, price, original_price, stock, discount)
        VALUES (product_id_var, 'beige', 'بژ', 'one-size', 280000, 322000, 25, 15)
        ON CONFLICT (product_id, COALESCE(color_key, ''), COALESCE(size_key, '')) DO NOTHING
        RETURNING id INTO detail_id_var;
        IF detail_id_var IS NULL THEN
            SELECT id INTO detail_id_var FROM product_details 
            WHERE product_details.product_id = product_id_var AND product_details.color_key = 'beige' AND product_details.size_key = 'one-size' 
            LIMIT 1;
        END IF;
        IF detail_id_var IS NOT NULL THEN
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-1.jpg', 'https://picsum.photos/seed/knitted-hat-1-beige-one-size-0/800/800', 0)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
            VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-2.jpg', 'https://picsum.photos/seed/knitted-hat-1-beige-one-size-1/800/800', 1)
            ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
        END IF;
    END IF;

END $$;


-- Update prices based on size and color, and ensure all variants have attachments
DO $$
DECLARE
    detail_rec RECORD;
    base_price DECIMAL;
    new_price DECIMAL;
    new_original_price DECIMAL;
    size_mult DECIMAL;
    color_mult DECIMAL;
    detail_id_var BIGINT;
    img_count INTEGER;
    img_seed TEXT;
    img_url TEXT;
BEGIN
    -- Update prices for all product_details
    FOR detail_rec IN 
        SELECT pd.id, pd.product_id, pd.color_key, pd.size_key, pd.price, pd.discount, p.slug as product_slug
        FROM product_details pd
        JOIN products p ON pd.product_id = p.id
        WHERE pd.deleted_at IS NULL
    LOOP
        -- Calculate size multiplier
        size_mult := CASE detail_rec.size_key
            WHEN 'S' THEN 1.0
            WHEN 'M' THEN 1.05
            WHEN 'L' THEN 1.1
            WHEN 'XL' THEN 1.15
            WHEN 'XXL' THEN 1.2
            WHEN '36' THEN 1.0
            WHEN '37' THEN 1.02
            WHEN '38' THEN 1.05
            WHEN '39' THEN 1.08
            WHEN '40' THEN 1.1
            WHEN '41' THEN 1.12
            WHEN '42' THEN 1.15
            WHEN '43' THEN 1.18
            WHEN '44' THEN 1.2
            WHEN '45' THEN 1.22
            WHEN '30' THEN 1.0
            WHEN '32' THEN 1.05
            WHEN '34' THEN 1.1
            WHEN '36' THEN 1.15
            WHEN '38' THEN 1.2
            WHEN '16' THEN 1.0
            WHEN '17' THEN 1.05
            WHEN '18' THEN 1.1
            WHEN '19' THEN 1.15
            WHEN '45cm' THEN 1.0
            WHEN '50cm' THEN 1.1
            WHEN '55cm' THEN 1.2
            WHEN '60cm' THEN 1.3
            WHEN '100cm' THEN 1.0
            WHEN '110cm' THEN 1.1
            WHEN '120cm' THEN 1.2
            WHEN 'small' THEN 1.0
            WHEN 'medium' THEN 1.1
            WHEN 'large' THEN 1.2
            ELSE 1.0
        END;
        
        -- Calculate color multiplier
        color_mult := CASE detail_rec.color_key
            WHEN 'black' THEN 1.1
            WHEN 'navy' THEN 1.1
            WHEN 'charcoal' THEN 1.1
            WHEN 'brown' THEN 1.05
            WHEN 'tan' THEN 1.05
            WHEN 'gold' THEN 1.15
            WHEN 'rose-gold' THEN 1.15
            ELSE 1.0
        END;
        
        -- Get base price from the first variant of this product (smallest size, base color)
        SELECT price INTO base_price
        FROM product_details
        WHERE product_id = detail_rec.product_id
        AND deleted_at IS NULL
        ORDER BY 
            CASE size_key
                WHEN 'S' THEN 1
                WHEN 'M' THEN 2
                WHEN 'L' THEN 3
                WHEN 'XL' THEN 4
                WHEN 'XXL' THEN 5
                WHEN '36' THEN 1
                WHEN '37' THEN 2
                WHEN '38' THEN 3
                WHEN '39' THEN 4
                WHEN '40' THEN 5
                WHEN 'one-size' THEN 1
                ELSE 99
            END,
            CASE color_key
                WHEN 'white' THEN 1
                WHEN 'black' THEN 2
                ELSE 3
            END
        LIMIT 1;
        
        -- If no base price found, use current price
        IF base_price IS NULL THEN
            base_price := detail_rec.price;
        END IF;
        
        -- Apply multipliers
        new_price := base_price * size_mult * color_mult;
        
        -- Calculate original price if discount exists
        IF detail_rec.discount > 0 THEN
            new_original_price := new_price * (1 + detail_rec.discount::DECIMAL / 100);
        ELSE
            new_original_price := NULL;
        END IF;
        
        -- Update price
        UPDATE product_details
        SET price = new_price,
            original_price = new_original_price
        WHERE id = detail_rec.id;
        
        -- Ensure attachments exist for this variant
        detail_id_var := detail_rec.id;
        
        -- Count existing attachments
        SELECT COUNT(*) INTO img_count
        FROM attachments
        WHERE attachable_type = 'ProductDetail'
        AND attachable_id = detail_id_var::TEXT
        AND deleted_at IS NULL;
        
        -- If no attachments, add them
        IF img_count = 0 AND detail_rec.size_key IS NOT NULL THEN
            -- Determine number of images based on product type
            img_count := CASE 
                WHEN detail_rec.product_slug LIKE '%dress%' OR detail_rec.product_slug LIKE '%suit%' THEN 4
                WHEN detail_rec.product_slug LIKE '%jewelry%' OR detail_rec.product_slug LIKE '%set%' THEN 5
                WHEN detail_rec.product_slug LIKE '%bag%' OR detail_rec.product_slug LIKE '%coat%' THEN 4
                ELSE 3
            END;
            
            -- Insert images
            FOR i IN 0..img_count-1 LOOP
                img_seed := detail_rec.product_slug || '-' || COALESCE(detail_rec.color_key, 'default') || '-' || COALESCE(detail_rec.size_key, 'one-size') || '-' || i;
                img_url := 'https://picsum.photos/seed/' || img_seed || '/800/800';
                
                INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
                VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-' || (i+1) || '.jpg', img_url, i)
                ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            END LOOP;
        END IF;
    END LOOP;
END $$;


-- Update prices based on size and color, and ensure all variants have attachments
DO $$
DECLARE
    detail_rec RECORD;
    base_price DECIMAL;
    new_price DECIMAL;
    new_original_price DECIMAL;
    size_mult DECIMAL;
    color_mult DECIMAL;
    detail_id_var BIGINT;
    img_count INTEGER;
    img_seed TEXT;
    img_url TEXT;
    i INTEGER;
BEGIN
    -- Update prices for all product_details
    FOR detail_rec IN 
        SELECT pd.id, pd.product_id, pd.color_key, pd.size_key, pd.price, pd.discount, p.slug as product_slug
        FROM product_details pd
        JOIN products p ON pd.product_id = p.id
        WHERE pd.deleted_at IS NULL
        AND pd.size_key IS NOT NULL
    LOOP
        -- Calculate size multiplier
        size_mult := CASE detail_rec.size_key
            WHEN 'S' THEN 1.0
            WHEN 'M' THEN 1.05
            WHEN 'L' THEN 1.1
            WHEN 'XL' THEN 1.15
            WHEN 'XXL' THEN 1.2
            WHEN '36' THEN 1.0
            WHEN '37' THEN 1.02
            WHEN '38' THEN 1.05
            WHEN '39' THEN 1.08
            WHEN '40' THEN 1.1
            WHEN '41' THEN 1.12
            WHEN '42' THEN 1.15
            WHEN '43' THEN 1.18
            WHEN '44' THEN 1.2
            WHEN '45' THEN 1.22
            WHEN '30' THEN 1.0
            WHEN '32' THEN 1.05
            WHEN '34' THEN 1.1
            WHEN '36' THEN 1.15
            WHEN '38' THEN 1.2
            WHEN '16' THEN 1.0
            WHEN '17' THEN 1.05
            WHEN '18' THEN 1.1
            WHEN '19' THEN 1.15
            WHEN '45cm' THEN 1.0
            WHEN '50cm' THEN 1.1
            WHEN '55cm' THEN 1.2
            WHEN '60cm' THEN 1.3
            WHEN '100cm' THEN 1.0
            WHEN '110cm' THEN 1.1
            WHEN '120cm' THEN 1.2
            WHEN 'small' THEN 1.0
            WHEN 'medium' THEN 1.1
            WHEN 'large' THEN 1.2
            ELSE 1.0
        END;
        
        -- Calculate color multiplier
        color_mult := CASE detail_rec.color_key
            WHEN 'black' THEN 1.1
            WHEN 'navy' THEN 1.1
            WHEN 'charcoal' THEN 1.1
            WHEN 'brown' THEN 1.05
            WHEN 'tan' THEN 1.05
            WHEN 'gold' THEN 1.15
            WHEN 'rose-gold' THEN 1.15
            ELSE 1.0
        END;
        
        -- Get base price from smallest size of this product
        SELECT MIN(price) INTO base_price
        FROM product_details
        WHERE product_id = detail_rec.product_id
        AND deleted_at IS NULL
        AND size_key IS NOT NULL;
        
        -- If no base price found, use current price
        IF base_price IS NULL THEN
            base_price := detail_rec.price;
        END IF;
        
        -- Apply multipliers to base price
        new_price := base_price * size_mult * color_mult;
        
        -- Calculate original price if discount exists
        IF detail_rec.discount > 0 THEN
            new_original_price := new_price * (1 + detail_rec.discount::DECIMAL / 100);
        ELSE
            new_original_price := NULL;
        END IF;
        
        -- Update price
        UPDATE product_details
        SET price = new_price,
            original_price = new_original_price
        WHERE id = detail_rec.id;
        
        -- Ensure attachments exist for this variant
        detail_id_var := detail_rec.id;
        
        -- Count existing attachments
        SELECT COUNT(*) INTO img_count
        FROM attachments
        WHERE attachable_type = 'ProductDetail'
        AND attachable_id = detail_id_var::TEXT
        AND deleted_at IS NULL;
        
        -- If no attachments, add them
        IF img_count = 0 THEN
            -- Determine number of images based on product type
            img_count := CASE 
                WHEN detail_rec.product_slug LIKE '%dress%' OR detail_rec.product_slug LIKE '%suit%' THEN 4
                WHEN detail_rec.product_slug LIKE '%jewelry%' OR detail_rec.product_slug LIKE '%set%' THEN 5
                WHEN detail_rec.product_slug LIKE '%bag%' OR detail_rec.product_slug LIKE '%coat%' THEN 4
                ELSE 3
            END;
            
            -- Insert images
            FOR i IN 0..img_count-1 LOOP
                img_seed := detail_rec.product_slug || '-' || COALESCE(detail_rec.color_key, 'default') || '-' || COALESCE(detail_rec.size_key, 'one-size') || '-' || i;
                img_url := 'https://picsum.photos/seed/' || img_seed || '/800/800';
                
                INSERT INTO attachments (attachable_type, attachable_id, file_type, file_name, file_path, "order")
                VALUES ('ProductDetail', detail_id_var::TEXT, 'image', 'image-' || (i+1) || '.jpg', img_url, i)
                ON CONFLICT (attachable_type, attachable_id, "order") DO NOTHING;
            END LOOP;
        END IF;
    END LOOP;
END $$;


-- migrate:down
-- Remove seed data
DELETE FROM product_details WHERE product_id IN (
    SELECT id FROM products WHERE slug IN (
        'women-classic-suit-1', 'women-silk-dress-1', 'women-cotton-shorts-1', 'women-collared-blouse-1',
        'women-pencil-skirt-1', 'women-wool-coat-1', 'women-basic-tshirt-1', 'women-evening-dress-1',
        'women-leather-handbag-1', 'women-heels-1', 'men-leather-shoes-1', 'leather-belt-bag-1',
        'men-sneakers-1', 'women-shoulder-bag-1', 'women-sneakers-1',
        'gold-necklace-with-gem-1', 'handmade-silver-bracelet-1', 'gold-earrings-classic-1',
        'silver-ring-with-stone-1', 'simple-gold-chain-1', 'gold-jewelry-set-1',
        'men-formal-suit-1', 'men-formal-shirt-1', 'men-jeans-1', 'men-basic-tshirt-1', 'men-wool-coat-1',
        'silk-scarf-1', 'men-leather-belt-1', 'sunglasses-1', 'knitted-hat-1'
    )
);

DELETE FROM products WHERE slug IN (
    'women-classic-suit-1', 'women-silk-dress-1', 'women-cotton-shorts-1', 'women-collared-blouse-1',
    'women-pencil-skirt-1', 'women-wool-coat-1', 'women-basic-tshirt-1', 'women-evening-dress-1',
    'women-leather-handbag-1', 'women-heels-1', 'men-leather-shoes-1', 'leather-belt-bag-1',
    'men-sneakers-1', 'women-shoulder-bag-1', 'women-sneakers-1',
    'gold-necklace-with-gem-1', 'handmade-silver-bracelet-1', 'gold-earrings-classic-1',
    'silver-ring-with-stone-1', 'simple-gold-chain-1', 'gold-jewelry-set-1',
    'men-formal-suit-1', 'men-formal-shirt-1', 'men-jeans-1', 'men-basic-tshirt-1', 'men-wool-coat-1',
    'silk-scarf-1', 'men-leather-belt-1', 'sunglasses-1', 'knitted-hat-1'
);

DELETE FROM categories WHERE slug IN (
    'women-clothing', 'bags-shoes', 'jewelry', 'men-clothing', 'accessories'
);
