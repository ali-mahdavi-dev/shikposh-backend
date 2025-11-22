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
