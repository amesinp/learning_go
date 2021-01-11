ALTER TABLE refresh_tokens
ADD COLUMN is_used BOOLEAN DEFAULT false;