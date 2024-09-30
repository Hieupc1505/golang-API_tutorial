ALTER TABLE IF EXISTS "todo_list" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

DROP TABLE IF EXISTS "users";
DROP TYPE IF EXISTS "user_roles";