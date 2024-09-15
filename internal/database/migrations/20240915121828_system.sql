-- +goose Up
-- +goose StatementBegin
CREATE TABLE system (
  owner_id VARCHAR(255) PRIMARY KEY NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE system;
-- +goose StatementEnd