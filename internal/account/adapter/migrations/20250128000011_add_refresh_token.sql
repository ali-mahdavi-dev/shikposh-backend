-- migrate:up
ALTER TABLE tokens ADD COLUMN refresh_token TEXT;

-- migrate:down
ALTER TABLE tokens DROP COLUMN refresh_token;

