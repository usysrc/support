-- +goose Up
-- +goose StatementBegin
CREATE TABLE tickets (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    status TEXT NOT NULL,
    priority TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE INDEX idx_tickets_status ON tickets (status);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tickets;
-- +goose StatementEnd
