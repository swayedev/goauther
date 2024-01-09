CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.Updated_At = CURRENT_TIMESTAMP;
  RETURN NEW; 
END;
$$ language 'plpgsql';

CREATE TABLE IF NOT EXISTS oauth_clients (
    Id UUID PRIMARY KEY,
    User_Id UUID NULL,
    Name VARCHAR(255) NOT NULL,
    Secret VARCHAR(100) NULL,
    Provider VARCHAR(255) NULL,
    Redirect_Uris TEXT NOT NULL,
    Personal_Access_Client BOOLEAN NOT NULL,
    Password_Client BOOLEAN NOT NULL,
    Revoked BOOLEAN NOT NULL,
    Created_At TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    Updated_At TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_user_id ON oauth_clients (User_Id);
CREATE TRIGGER update_oauth_clients_modtime BEFORE UPDATE ON oauth_clients FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TABLE oauth_personal_access_clients (
    Id BIGSERIAL PRIMARY KEY,
    Client_Id UUID NOT NULL,
    Created_At TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    Updated_At TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (Client_Id) REFERENCES oauth_clients(Id) ON DELETE CASCADE
);
CREATE TRIGGER update_oauth_personal_access_clients_modtime BEFORE UPDATE ON oauth_personal_access_clients FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TABLE IF NOT EXISTS oauth_scopes (
    Id BIGSERIAL PRIMARY KEY,
    Scope VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS oauth_auth_codes (
    Id VARCHAR(100) PRIMARY KEY,
    User_Id UUID NULL,
    Client_Id UUID NOT NULL,
    Scopes TEXT NULL,
    Revoked BOOLEAN NOT NULL,
    Expires_At TIMESTAMP NULL,
    FOREIGN KEY (Client_Id) REFERENCES oauth_clients(Id) ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS idx_user_id ON oauth_auth_codes (User_Id);


CREATE TABLE IF NOT EXISTS oauth_access_tokens (
    Id VARCHAR(100) PRIMARY KEY,
    User_Id UUID NULL,
    Client_Id UUID NOT NULL,
    Name VARCHAR(255) NULL,
    Scopes TEXT NULL,
    Revoked BOOLEAN NOT NULL,
    Created_At TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    Updated_At TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    Expires_At TIMESTAMP NULL,
    FOREIGN KEY (Client_Id) REFERENCES oauth_clients(Id) ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS idx_user_id ON oauth_access_tokens (User_Id);
CREATE TRIGGER update_oauth_access_tokens_modtime BEFORE UPDATE ON oauth_access_tokens FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TABLE IF NOT EXISTS oauth_refresh_tokens (
    Id VARCHAR(100) PRIMARY KEY,
    Access_Token_Id VARCHAR(100),
    Revoked BOOLEAN NOT NULL,
    Expires_At TIMESTAMP NULL
);
CREATE INDEX IF NOT EXISTS idx_access_token_id ON oauth_refresh_tokens (Access_Token_Id);

