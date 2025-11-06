-- migrate:up
CREATE TABLE tokens (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    token TEXT NOT NULL,
    user_id BIGINT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Optional: Index for soft delete queries
CREATE INDEX idx_tokens_deleted_at ON tokens(deleted_at);

-- migrate:down
DROP INDEX IF EXISTS idx_tokens_deleted_at;
DROP TABLE IF EXISTS tokens;
