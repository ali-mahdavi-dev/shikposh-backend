-- migrate:up
CREATE TABLE products (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    rating DECIMAL(3,2) DEFAULT 0.00,
    review_count INTEGER DEFAULT 0,
    description TEXT,
    category_id BIGINT NOT NULL,
    tags JSONB DEFAULT '[]'::jsonb,
    image VARCHAR(500),
    is_new BOOLEAN DEFAULT false,
    is_featured BOOLEAN DEFAULT false,
    sizes JSONB DEFAULT '[]'::jsonb,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_products_category FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE RESTRICT
);

CREATE INDEX idx_products_deleted_at ON products(deleted_at);
CREATE INDEX idx_products_category_id ON products(category_id);
CREATE INDEX idx_products_is_featured ON products(is_featured);
CREATE INDEX idx_products_rating ON products(rating);
CREATE INDEX idx_products_price ON products(price);
CREATE INDEX idx_products_tags ON products USING GIN(tags);

-- migrate:down
DROP INDEX IF EXISTS idx_products_tags;
DROP INDEX IF EXISTS idx_products_price;
DROP INDEX IF EXISTS idx_products_rating;
DROP INDEX IF EXISTS idx_products_is_featured;
DROP INDEX IF EXISTS idx_products_category_id;
DROP INDEX IF EXISTS idx_products_deleted_at;
DROP TABLE IF EXISTS products;

