-- migrate:up
-- Create sizes table
CREATE TABLE sizes (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_sizes_deleted_at ON sizes(deleted_at);
CREATE INDEX idx_sizes_slug ON sizes(slug);
CREATE UNIQUE INDEX idx_sizes_name_unique ON sizes(name) WHERE deleted_at IS NULL;

-- Create product_sizes junction table for many-to-many relationship
CREATE TABLE product_sizes (
    product_id BIGINT NOT NULL,
    size_id BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (product_id, size_id),
    CONSTRAINT fk_product_sizes_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    CONSTRAINT fk_product_sizes_size FOREIGN KEY (size_id) REFERENCES sizes(id) ON DELETE CASCADE
);

CREATE INDEX idx_product_sizes_product_id ON product_sizes(product_id);
CREATE INDEX idx_product_sizes_size_id ON product_sizes(size_id);

-- Migrate existing sizes from products.sizes JSONB column to the new tables
-- This will extract unique sizes from the JSONB array and create size records
DO $$
DECLARE
    product_record RECORD;
    size_name TEXT;
    size_slug TEXT;
    size_id_var BIGINT;
    size_array TEXT[];
BEGIN
    -- Loop through all products that have sizes
    FOR product_record IN 
        SELECT id, sizes FROM products WHERE sizes IS NOT NULL AND sizes::text != '[]'::text
    LOOP
        -- Parse the JSONB array
        SELECT ARRAY(SELECT jsonb_array_elements_text(product_record.sizes)) INTO size_array;
        
        -- Process each size
        FOREACH size_name IN ARRAY size_array
        LOOP
            -- Generate slug from size name (simple version - you might want to use a proper slug function)
            size_slug := lower(regexp_replace(size_name, '[^a-zA-Z0-9]+', '-', 'g'));
            size_slug := trim(both '-' from size_slug);
            
            -- Check if size already exists, if not create it
            SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
            
            IF size_id_var IS NULL THEN
                INSERT INTO sizes (name, slug) 
                VALUES (size_name, size_slug)
                ON CONFLICT (slug) DO NOTHING
                RETURNING id INTO size_id_var;
                
                -- If still null (conflict), get the existing size
                IF size_id_var IS NULL THEN
                    SELECT id INTO size_id_var FROM sizes WHERE slug = size_slug AND deleted_at IS NULL;
                END IF;
            END IF;
            
            -- Link product to size if not already linked
            IF size_id_var IS NOT NULL THEN
                INSERT INTO product_sizes (product_id, size_id)
                VALUES (product_record.id, size_id_var)
                ON CONFLICT (product_id, size_id) DO NOTHING;
            END IF;
        END LOOP;
    END LOOP;
END $$;

-- migrate:down
DROP INDEX IF EXISTS idx_product_sizes_size_id;
DROP INDEX IF EXISTS idx_product_sizes_product_id;
DROP TABLE IF EXISTS product_sizes;
DROP INDEX IF EXISTS idx_sizes_name_unique;
DROP INDEX IF EXISTS idx_sizes_slug;
DROP INDEX IF EXISTS idx_sizes_deleted_at;
DROP TABLE IF EXISTS sizes;

