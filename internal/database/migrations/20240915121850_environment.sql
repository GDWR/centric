-- +goose Up
-- +goose StatementBegin
CREATE TABLE environment (
  id TEXT PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  icon VARCHAR(255) NOT NULL,
  uri VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE environment;
-- +goose StatementEnd
