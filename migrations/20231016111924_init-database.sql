-- Create "user_infos" table
CREATE TABLE "user_infos" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "first_name" text NULL,
  "last_name" text NULL,
  "middle_nam" text NULL,
  "display_name" text NULL DEFAULT 'data123',
  PRIMARY KEY ("id")
);
-- Create index "idx_user_infos_deleted_at" to table: "user_infos"
CREATE INDEX "idx_user_infos_deleted_at" ON "user_infos" ("deleted_at");
-- Create "users" table
CREATE TABLE "users" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "email" text NULL,
  "is_active" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
