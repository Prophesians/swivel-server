-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE swivelserver.search_history(
  id serial PRIMARY KEY,
  tags text[] NOT NULL,
  summary VARCHAR(200),
  url VARCHAR(100)
);
-- +goose Down
DROP TABLE swivelserver.search_history;
-- SQL in this section is executed when the migration is rolled back.
