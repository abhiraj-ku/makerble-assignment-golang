-- Remove foreign key constraint first
ALTER TABLE patients DROP CONSTRAINT IF EXISTS patients_updated_by_fkey;

-- Remove columns added by this migration
ALTER TABLE patients DROP COLUMN IF EXISTS updated_by;
ALTER TABLE patients DROP COLUMN IF EXISTS updated_at;
ALTER TABLE patients DROP COLUMN IF EXISTS created_at;
ALTER TABLE patients DROP COLUMN IF EXISTS contact;
ALTER TABLE patients DROP COLUMN IF EXISTS address;
ALTER TABLE patients DROP COLUMN IF EXISTS gender;
ALTER TABLE patients DROP COLUMN IF EXISTS age;
ALTER TABLE patients DROP COLUMN IF EXISTS last_name;
ALTER TABLE patients DROP COLUMN IF EXISTS first_name;

-- Remove the table structure (optional and safe here since the table would be empty without columns)
DROP TABLE IF EXISTS patients;

-- Remove the users table (only if no real data is in use)
ALTER TABLE users DROP COLUMN IF EXISTS created_at;
ALTER TABLE users DROP COLUMN IF EXISTS role;
ALTER TABLE users DROP COLUMN IF EXISTS password;
ALTER TABLE users DROP COLUMN IF EXISTS username;

-- Remove users table entirely (optional, based on use case)
DROP TABLE IF EXISTS users;
