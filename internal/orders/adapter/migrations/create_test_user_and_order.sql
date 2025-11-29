-- ============================================
-- Complete Script: Create User + Order
-- ============================================
-- This script will:
-- 1. Create a test user if it doesn't exist
-- 2. Create a test order with items and address
-- ============================================

DO $$
DECLARE
    v_order_id BIGINT;
    v_user_id BIGINT;
    v_product1_id BIGINT := 1;  -- Change to your actual product IDs
    v_product2_id BIGINT := 2;  -- Change to your actual product IDs
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
        RAISE NOTICE 'ℹ️  Test user already exists with ID: %', v_user_id;
    END IF;

    -- Step 2: Create order
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

    -- Step 3: Insert order items
    INSERT INTO order_items (
        order_id, product_id, product_name, product_slug, product_image,
        quantity, price, discount, color, size,
        created_at, updated_at
    ) VALUES 
    (
        v_order_id, v_product1_id, 'تی‌شرت مردانه کلاسیک', 't-shirt-classic-male',
        'https://example.com/images/product1.jpg', 2, 500000, 50000, 'آبی', 'L',
        CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
    ),
    (
        v_order_id, v_product2_id, 'شلوار جین اسلیم فیت', 'jeans-slim-fit',
        'https://example.com/images/product2.jpg', 1, 800000, 100000, 'مشکی', '32',
        CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
    );

    RAISE NOTICE '✅ Order items created';

    -- Step 4: Insert order address
    INSERT INTO order_addresses (
        order_id, full_name, phone, address, city, province, postal_code,
        created_at, updated_at
    ) VALUES (
        v_order_id, 'علی احمدی', '09123456789',
        'خیابان ولیعصر، پلاک 123، واحد 4', 'تهران', 'تهران', '1234567890',
        CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
    );

    RAISE NOTICE '✅ Order address created';
    RAISE NOTICE '🎉 Test order created successfully! Order ID: %', v_order_id;
END $$;

