CREATE TABLE IF NOT EXISTS oauth_auth_codes (
    id VARCHAR(100) PRIMARY KEY,
    user_id BIGINT,
    client_id UUID NOT NULL,
    scopes TEXT NULL,
    revoked BOOLEAN NOT NULL,
    expires_at TIMESTAMP NULL
);

CREATE INDEX ON oauth_auth_codes (user_id);

CREATE TABLE IF NOT EXISTS oauth_access_tokens (
    id VARCHAR(100) PRIMARY KEY,
    user_id BIGINT NULL,
    client_id UUID NOT NULL,
    name VARCHAR(255) NULL,
    scopes TEXT NULL,
    revoked BOOLEAN NOT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    expires_at TIMESTAMP NULL
);

CREATE INDEX ON oauth_access_tokens (user_id);

CREATE TABLE oauth_refresh_tokens (
    id VARCHAR(100) PRIMARY KEY,
    access_token_id VARCHAR(100),
    revoked BOOLEAN NOT NULL,
    expires_at TIMESTAMP NULL
);

CREATE INDEX ON oauth_refresh_tokens (access_token_id);


CREATE TABLE oauth_clients (
    id UUID PRIMARY KEY,
    user_id BIGINT NULL,
    name VARCHAR(255) NOT NULL,
    secret VARCHAR(100) NULL,
    provider VARCHAR(255) NULL,
    redirect TEXT NOT NULL,
    personal_access_client BOOLEAN NOT NULL,
    password_client BOOLEAN NOT NULL,
    revoked BOOLEAN NOT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);

CREATE INDEX ON oauth_clients (user_id);

CREATE TABLE oauth_personal_access_clients (
    id BIGSERIAL PRIMARY KEY,
    client_id BIGINT NOT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);
