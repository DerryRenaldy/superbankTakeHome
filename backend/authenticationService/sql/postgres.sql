-- Create schema (optional, if using namespacing)
CREATE SCHEMA IF NOT EXISTS auth;

-- Users table
CREATE TABLE IF NOT EXISTS auth.users (
    user_id BIGSERIAL PRIMARY KEY,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash CHAR(60) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Sessions table
CREATE TABLE IF NOT EXISTS auth.sessions (
    refresh_token TEXT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    is_revoked BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL,
    CONSTRAINT session_user_FK FOREIGN KEY (user_id) REFERENCES auth.users(user_id) ON DELETE CASCADE
);

-- Roles table
CREATE TABLE IF NOT EXISTS auth.roles (
    role_id SERIAL PRIMARY KEY,  -- Auto-increment
    role_name VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- User Roles table (many-to-many)
CREATE TABLE IF NOT EXISTS auth.user_roles (
    user_id BIGINT NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, role_id),
    CONSTRAINT user_roles_user_FK FOREIGN KEY (user_id) REFERENCES auth.users(user_id) ON DELETE CASCADE,
    CONSTRAINT user_roles_role_FK FOREIGN KEY (role_id) REFERENCES auth.roles(role_id) ON DELETE CASCADE
);

-- Indexes for performance
CREATE INDEX idx_users_created_at ON auth.users (created_at);
CREATE INDEX idx_sessions_refresh_token ON auth.sessions (refresh_token);
CREATE INDEX idx_sessions_expires_at ON auth.sessions (expires_at);
CREATE INDEX idx_sessions_user_id ON auth.sessions (user_id);
CREATE INDEX idx_sessions_cleanup ON auth.sessions (is_revoked, expires_at);
CREATE INDEX idx_user_roles_role_id ON auth.user_roles (role_id);

-- Insert default roles (PostgreSQL way)
INSERT INTO auth.roles (role_name) VALUES
    ('admin'),
    ('user')
ON CONFLICT (role_name) DO NOTHING;  -- Prevent duplicate entries