CREATE TABLE IF NOT EXISTS flag (id SERIAL PRIMARY KEY, contents VARCHAR(256));

INSERT INTO
  flag (contents)
VALUES
  (
    'If you see this, create a ticket in Discord please :('
  )
ON CONFLICT (contents) DO NOTHING;