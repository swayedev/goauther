-- Drop oauth_refresh_tokens table
DROP TABLE IF EXISTS oauth_refresh_tokens;
-- Drop oauth_access_tokens table
DROP TABLE IF EXISTS oauth_access_tokens;
-- Drop oauth_auth_codes table
DROP TABLE IF EXISTS oauth_auth_codes;
-- Drop oauth_personal_access_clients table
DROP TABLE IF EXISTS oauth_personal_access_clients;
-- Drop oauth_scopes table
DROP TABLE IF EXISTS oauth_scopes;
-- Drop oauth_clients table last because other tables reference it
DROP TABLE IF EXISTS oauth_clients;
