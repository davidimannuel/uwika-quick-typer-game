-- Remove seed data (optional - usually we don't rollback seed data)
DELETE FROM phrases WHERE id::text LIKE '20000000-%';
DELETE FROM stages WHERE id::text LIKE '10000000-%';
DELETE FROM themes WHERE id::text LIKE '30000000-%';
DELETE FROM users WHERE id = '00000000-0000-0000-0000-000000000001'::uuid;

