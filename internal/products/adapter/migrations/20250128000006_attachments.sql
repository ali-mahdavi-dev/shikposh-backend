-- migrate:up
CREATE TABLE attachments (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    attachable_type VARCHAR(100) NOT NULL,
    attachable_id TEXT NOT NULL,
    file_type VARCHAR(50) NOT NULL DEFAULT 'image',
    file_name VARCHAR(255) NOT NULL,
    file_path VARCHAR(500) NOT NULL,
    file_size BIGINT DEFAULT 0,
    mime_type VARCHAR(100),
    "order" INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT uk_attachments_unique UNIQUE (attachable_type, attachable_id, "order")
);

CREATE INDEX idx_attachments_deleted_at ON attachments(deleted_at);
CREATE INDEX idx_attachments_attachable ON attachments(attachable_type, attachable_id);
CREATE INDEX idx_attachments_file_type ON attachments(file_type);

-- migrate:down
DROP INDEX IF EXISTS idx_attachments_file_type;
DROP INDEX IF EXISTS idx_attachments_attachable;
DROP INDEX IF EXISTS idx_attachments_deleted_at;
DROP TABLE IF EXISTS attachments;

