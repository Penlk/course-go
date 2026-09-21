-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS screenings_starts_at_idx ON screenings (starts_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS screenings_starts_at_idx;
-- +goose StatementEnd
