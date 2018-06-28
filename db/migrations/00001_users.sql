-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE swivelserver.users(
  id serial PRIMARY KEY,
  user_id VARCHAR (50) UNIQUE NOT NULL,
  tags text[]
);
-- +goose Down
DROP TABLE swivelserver.users;
-- SQL in this section is executed when the migration is rolled back.
