CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  username VARCHAR(255),
  password_hash VARCHAR(255),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  UNIQUE(username)
);