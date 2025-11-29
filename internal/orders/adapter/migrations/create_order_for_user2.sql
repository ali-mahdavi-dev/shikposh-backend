-- Create order for user_id = 2 (the user from your JWT token)
DO $$
DECLARE
    v_order_id BIGINT;
    v_user_id BIGINT := 2;  -- user_id از token شما
    v_category_id BIGINT;
    v_product1_id BIGINT;
    v_product2_id BIGINT;
BEGIN
    -- Get first available category
    SELECT id INTO v_category_id 
    FROM categories 
    WHERE deleted_at IS NULL 
    LIMIT 1;
    
    IF v_category_id IS NULL THEN
        RAISE EXCEPTION 'No category found. Please create at least one category first.';
    END IF;

    -- Get first available products
    SELECT id INTO v_product1_id 
    FROM products 
    WHERE deleted_at IS NULL 
    LIMIT 1 OFFSET 0;
    
    SELECT id INTO v_product2_id 
    FROM products 
    WHERE deleted_at IS NULL 
    LIMIT 1 OFFSET 1;
    
    IF v_product1_id IS NULL OR v_product2_id IS NULL THEN
        RAISE EXCEPTION 'Not enough products found. Need at least 2 products.';
    END IF;

    RAISE NOTICE 'Using category_id: %, product1_id: %, product2_id: %', v_category_id, v_product1_id, v_product2_id;

    -- Create order for user_id = 2
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

    RAISE NOTICE '✅ Order created with ID: % for user_id: %', v_order_id, v_user_id;

    -- Insert order items
    INSERT INTO order_items (
        order_id,
        product_id,
        product_name,
        product_slug,
        product_image,
        quantity,
        price,
        discount,
        color,
        size,
        created_at,
        updated_at
    ) VALUES 
    (
        v_order_id,
        v_product1_id,
        (SELECT title FROM products WHERE id = v_product1_id),
        (SELECT slug FROM products WHERE id = v_product1_id),
        (SELECT thumbnail FROM products WHERE id = v_product1_id),
        2,
        500000,
        50000,
        'آبی',
        'L',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        v_order_id,
        v_product2_id,
        (SELECT title FROM products WHERE id = v_product2_id),
        (SELECT slug FROM products WHERE id = v_product2_id),
        (SELECT thumbnail FROM products WHERE id = v_product2_id),
        1,
        800000,
        100000,
        'مشکی',
        '32',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    );

    RAISE NOTICE '✅ Order items created';

    -- Insert order address
    INSERT INTO order_addresses (
        order_id,
        full_name,
        phone,
        address,
        city,
        province,
        postal_code,
        created_at,
        updated_at
    ) VALUES (
        v_order_id,
        'علی احمدی',
        '09123456789',
        'خیابان ولیعصر، پلاک 123، واحد 4',
        'تهران',
        'تهران',
        '1234567890',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    );

    RAISE NOTICE '✅ Order address created';
    RAISE NOTICE '🎉 Order created successfully! Order ID: %, User ID: %', v_order_id, v_user_id;
END $$;

