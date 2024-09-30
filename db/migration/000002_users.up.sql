
CREATE TYPE "user_roles" AS ENUM (
  'user',
  'editor'
);

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "change_password_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "role" user_roles NOT NULL DEFAULT 'user'
);

ALTER TABLE "todo_list" ADD FOREIGN KEY ("user") REFERENCES "users" ("username");
