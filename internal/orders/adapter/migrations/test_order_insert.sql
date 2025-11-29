DO $$
DECLARE
    v_order_id BIGINT;
    v_user_id BIGINT;
    v_category_id BIGINT;
    v_product1_id BIGINT;
    v_product2_id BIGINT;
BEGIN
    -- Step 1: Get or create test user
    SELECT id INTO v_user_id 
    FROM users 
    WHERE phone = '09123456789' AND deleted_at IS NULL;
    
    IF v_user_id IS NULL THEN
        -- Create test user
        INSERT INTO users (
            avatar_identifier,
            first_name,
            last_name,
            email,
            phone,
            password,
            created_at,
            updated_at
        ) VALUES (
            'test-avatar-001',
            'علی',
            'احمدی',
            'ali.ahmadi@example.com',
            '09123456789',
            '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', -- password: "password123"
            CURRENT_TIMESTAMP,
            CURRENT_TIMESTAMP
        ) RETURNING id INTO v_user_id;
        
        RAISE NOTICE '✅ Test user created with ID: %', v_user_id;
    ELSE
        RAISE NOTICE 'ℹ️  Using existing user with ID: %', v_user_id;
    END IF;

    -- Step 2: Get or create test category
    BEGIN
        SELECT id INTO v_category_id 
        FROM categories 
        WHERE slug = 'test-category' AND deleted_at IS NULL
        LIMIT 1;
        
        IF v_category_id IS NULL THEN
            -- Try to find any category (even soft deleted) to reuse
            SELECT id INTO v_category_id 
            FROM categories 
            WHERE slug = 'test-category'
            LIMIT 1;
            
            IF v_category_id IS NULL THEN
                -- Create new category
                INSERT INTO categories (
                    name,
                    slug,
                    description,
                    created_at,
                    updated_at
                ) VALUES (
                    'دسته‌بندی تست',
                    'test-category',
                    'دسته‌بندی تستی برای محصولات',
                    CURRENT_TIMESTAMP,
                    CURRENT_TIMESTAMP
                ) RETURNING id INTO v_category_id;
                
                RAISE NOTICE '✅ Test category created with ID: %', v_category_id;
            ELSE
                -- Restore soft deleted category
                UPDATE categories 
                SET deleted_at = NULL, updated_at = CURRENT_TIMESTAMP
                WHERE id = v_category_id;
                
                RAISE NOTICE '✅ Test category restored with ID: %', v_category_id;
            END IF;
        ELSE
            RAISE NOTICE 'ℹ️  Using existing category with ID: %', v_category_id;
        END IF;
    EXCEPTION
        WHEN unique_violation THEN
            -- If duplicate key error, just select the existing one
            SELECT id INTO v_category_id 
            FROM categories 
            WHERE slug = 'test-category'
            LIMIT 1;
            RAISE NOTICE '⚠️  Category already exists, using ID: %', v_category_id;
    END;

    -- Verify category_id is set, if not use any available category
    IF v_category_id IS NULL THEN
        -- Fallback: use any available category
        SELECT id INTO v_category_id 
        FROM categories 
        WHERE deleted_at IS NULL
        LIMIT 1;
        
        IF v_category_id IS NULL THEN
            RAISE EXCEPTION 'No category found. Please create at least one category first.';
        ELSE
            RAISE NOTICE '⚠️  Using fallback category with ID: %', v_category_id;
        END IF;
    END IF;

    RAISE NOTICE '📦 Using category ID: %', v_category_id;

    -- Step 3: Get or create test products
    SELECT id INTO v_product1_id 
    FROM products 
    WHERE slug = 't-shirt-classic-male' AND deleted_at IS NULL
    LIMIT 1;
    
    IF v_product1_id IS NULL THEN
        -- Try to find any product (even soft deleted)
        SELECT id INTO v_product1_id 
        FROM products 
        WHERE slug = 't-shirt-classic-male'
        LIMIT 1;
        
        IF v_product1_id IS NOT NULL THEN
            -- Restore soft deleted product
            UPDATE products 
            SET deleted_at = NULL, updated_at = CURRENT_TIMESTAMP
            WHERE id = v_product1_id;
            
            RAISE NOTICE '✅ Test product 1 restored with ID: %', v_product1_id;
        END IF;
    END IF;
    
    IF v_product1_id IS NULL THEN
        -- Verify category_id before inserting
        IF v_category_id IS NULL THEN
            RAISE EXCEPTION 'Cannot create product: category_id is NULL';
        END IF;
        
        BEGIN
            INSERT INTO products (
                title,
                slug,
                brand,
                category_id,
                price,
                discount,
                stock,
                thumbnail,
                description,
                created_at,
                updated_at
            ) VALUES (
                'تی‌شرت مردانه کلاسیک',
                't-shirt-classic-male',
                'شیک‌پوشان',
                v_category_id,
                500000,
                10,
                100,
                'https://example.com/images/product1.jpg',
                'تی‌شرت مردانه کلاسیک با کیفیت عالی',
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP
            ) RETURNING id INTO v_product1_id;
        EXCEPTION
            WHEN unique_violation THEN
                SELECT id INTO v_product1_id 
                FROM products 
                WHERE slug = 't-shirt-classic-male'
                LIMIT 1;
                UPDATE products 
                SET deleted_at = NULL, updated_at = CURRENT_TIMESTAMP
                WHERE id = v_product1_id;
        END;
        
        RAISE NOTICE '✅ Test product 1 created/restored with ID: %', v_product1_id;
    ELSE
        RAISE NOTICE 'ℹ️  Using existing product 1 with ID: %', v_product1_id;
    END IF;

    SELECT id INTO v_product2_id 
    FROM products 
    WHERE slug = 'jeans-slim-fit' AND deleted_at IS NULL
    LIMIT 1;
    
    IF v_product2_id IS NULL THEN
        -- Try to find any product (even soft deleted)
        SELECT id INTO v_product2_id 
        FROM products 
        WHERE slug = 'jeans-slim-fit'
        LIMIT 1;
        
        IF v_product2_id IS NOT NULL THEN
            -- Restore soft deleted product
            UPDATE products 
            SET deleted_at = NULL, updated_at = CURRENT_TIMESTAMP
            WHERE id = v_product2_id;
            
            RAISE NOTICE '✅ Test product 2 restored with ID: %', v_product2_id;
        END IF;
    END IF;
    
    IF v_product2_id IS NULL THEN
        -- Verify category_id before inserting
        IF v_category_id IS NULL THEN
            RAISE EXCEPTION 'Cannot create product: category_id is NULL';
        END IF;
        
        BEGIN
            INSERT INTO products (
                title,
                slug,
                brand,
                category_id,
                price,
                discount,
                stock,
                thumbnail,
                description,
                created_at,
                updated_at
            ) VALUES (
                'شلوار جین اسلیم فیت',
                'jeans-slim-fit',
                'شیک‌پوشان',
                v_category_id,
                800000,
                12,
                50,
                'https://example.com/images/product2.jpg',
                'شلوار جین اسلیم فیت با برش مدرن',
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP
            ) RETURNING id INTO v_product2_id;
        EXCEPTION
            WHEN unique_violation THEN
                SELECT id INTO v_product2_id 
                FROM products 
                WHERE slug = 'jeans-slim-fit'
                LIMIT 1;
                UPDATE products 
                SET deleted_at = NULL, updated_at = CURRENT_TIMESTAMP
                WHERE id = v_product2_id;
        END;
        
        RAISE NOTICE '✅ Test product 2 created/restored with ID: %', v_product2_id;
    ELSE
        RAISE NOTICE 'ℹ️  Using existing product 2 with ID: %', v_product2_id;
    END IF;

    -- Step 4: Insert Order
    INSERT INTO orders (
        order_number,
        user_id,
        status,
        total_amount,
        discount_amount,
        shipping_cost,
        final_amount,
        payment_method,
        payment_status,
        tracking_number,
        created_at,
        updated_at
    ) VALUES (
        'ORD-' || TO_CHAR(CURRENT_DATE, 'YYYYMMDD') || '-' || LPAD(FLOOR(RANDOM() * 1000000)::TEXT, 6, '0'),
        v_user_id,
        'processing',
        1500000,  -- 1,500,000 تومان
        150000,   -- 150,000 تومان تخفیف
        50000,    -- 50,000 تومان ارسال
        1400000,  -- 1,400,000 تومان نهایی
        'online',
        'paid',
        'TRACK' || LPAD(FLOOR(RANDOM() * 1000000000)::TEXT, 9, '0'),
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ) RETURNING id INTO v_order_id;

    RAISE NOTICE '✅ Order created with ID: %', v_order_id;

    -- Step 5: Insert Order Items
    INSERT INTO order_items (
        order_id, product_id, product_name, product_slug, product_image,
        quantity, price, discount, color, size,
        created_at, updated_at
    ) VALUES 
    (v_order_id, v_product1_id, 'تی‌شرت مردانه کلاسیک', 't-shirt-classic-male', 
     'https://example.com/images/product1.jpg', 2, 500000, 50000, 'آبی', 'L',
     CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (v_order_id, v_product2_id, 'شلوار جین اسلیم فیت', 'jeans-slim-fit',
     'https://example.com/images/product2.jpg', 1, 800000, 100000, 'مشکی', '32',
     CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

    -- Step 6: Insert Order Address
    INSERT INTO order_addresses (
        order_id, full_name, phone, address, city, province, postal_code,
        created_at, updated_at
    ) VALUES (
        v_order_id, 'علی احمدی', '09123456789',
        'خیابان ولیعصر، پلاک 123، واحد 4', 'تهران', 'تهران', '1234567890',
        CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
    );

    RAISE NOTICE '🎉 Test order created successfully! Order ID: %', v_order_id;
END $$;