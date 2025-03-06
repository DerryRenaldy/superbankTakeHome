CREATE DATABASE `auth` /*!40100 DEFAULT CHARACTER SET utf8 */ /*!80016 DEFAULT ENCRYPTION='N' */;

CREATE TABLE IF NOT EXISTS auth.users (
    user_id BIGINT NOT NULL AUTO_INCREMENT,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash CHAR(60) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT users_PK PRIMARY KEY (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE IF NOT EXISTS auth.sessions (
    session_id CHAR(36) NOT NULL,
    user_id BIGINT NOT NULL,
    refresh_token VARCHAR(512) NOT NULL,
    is_revoked BOOL NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at DATETIME NOT NULL,

    CONSTRAINT session_PK PRIMARY KEY (session_id),
    CONSTRAINT session_user_FK FOREIGN KEY (user_id) REFERENCES auth.users(user_id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS auth.roles (
    role_id INT NOT NULL AUTO_INCREMENT,
    role_name VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT roles_PK PRIMARY KEY (role_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS auth.user_roles (
    user_id BIGINT NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT user_roles_PK PRIMARY KEY (user_id, role_id),
    CONSTRAINT user_roles_user_FK FOREIGN KEY (user_id) REFERENCES auth.users(user_id) ON DELETE CASCADE,
    CONSTRAINT user_roles_role_FK FOREIGN KEY (role_id) REFERENCES auth.roles(role_id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE INDEX idx_users_created_at ON auth.users (created_at);
CREATE INDEX idx_sessions_refresh_token ON auth.sessions (refresh_token);
CREATE INDEX idx_sessions_expires_at ON auth.sessions (expires_at);
CREATE INDEX idx_sessions_user_id ON auth.sessions (user_id);
CREATE INDEX idx_sessions_cleanup ON auth.sessions (is_revoked, expires_at);
CREATE INDEX idx_user_roles_role_id ON auth.user_roles (role_id);

INSERT INTO auth.roles (role_name) VALUES 
('admin'),
('user'),
ON DUPLICATE KEY UPDATE role_name = VALUES(role_name);