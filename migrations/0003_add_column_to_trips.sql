-- +goose Up
-- +goose StatementBegin
ALTER TABLE trips add column version int NOT NULL DEFAULT 1;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin


ALTER TABLE trips DROP COLUMN IF EXISTS version;

-- +goose StatementEnd