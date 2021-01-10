CREATE TABLE refresh_tokens (
  id SERIAL PRIMARY KEY,
  token_hash VARCHAR(255),
  user_id integer REFERENCES users (id),
  user_agent TEXT,
  expires_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT NOW(),
  UNIQUE(token_hash)
);