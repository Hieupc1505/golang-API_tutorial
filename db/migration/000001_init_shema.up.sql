CREATE TYPE "task_status" AS ENUM (
  'delete',
  'doing',
  'done'
);

CREATE TABLE "todo_list" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "image" varchar,
  "description" text,
  "status" task_status NOT NULL DEFAULT 'doing',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()), 
  "user" varchar NOT NULL
);