CREATE TABLE IF NOT EXISTS flag (id SERIAL PRIMARY KEY, contents VARCHAR(50));

INSERT INTO
  flag (contents)
VALUES
  (
    'HZ2025{l@t3r@l_m00000v3m3nt_r0cks_u_b3c0m1ng_a_hecker_Oik=}'
  )
ON CONFLICT (contents) DO NOTHING;