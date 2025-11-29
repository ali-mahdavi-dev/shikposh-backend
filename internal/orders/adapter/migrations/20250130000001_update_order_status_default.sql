-- migrate:up
-- Update default value for order status from 'pending' to 'payment_confirmed'
ALTER TABLE orders ALTER COLUMN status SET DEFAULT 'payment_confirmed';

-- Update existing records with 'pending' status to 'payment_confirmed' (if any)
UPDATE orders SET status = 'payment_confirmed' WHERE status = 'pending';

-- migrate:down
-- Revert default value back to 'pending'
ALTER TABLE orders ALTER COLUMN status SET DEFAULT 'pending';

-- Revert existing records back to 'pending' (if needed)
UPDATE orders SET status = 'pending' WHERE status = 'payment_confirmed';

