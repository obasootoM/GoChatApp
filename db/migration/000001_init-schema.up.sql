CREATE TABLE "post" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "content" varchar NOT NULL,
  "created_at" timestamptz DEFAULT 'now()'
);