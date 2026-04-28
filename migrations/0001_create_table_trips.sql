-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS trips (
    id UUID PRIMARY KEY,
    status VARCHAR(50) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS trips;
-- +goose StatementEnd
