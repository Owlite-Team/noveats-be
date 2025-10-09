CREATE TABLE IF NOT EXISTS Users
(
    id         BIGSERIAL PRIMARY KEY,
    email      VARCHAR(255) UNIQUE,
    name       VARCHAR(255) NOT NULL,
    password   VARCHAR(255),
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_users_created_at ON Users (created_at);

-- Add comment
COMMENT ON TABLE users IS 'Users table for authentication and user management';
