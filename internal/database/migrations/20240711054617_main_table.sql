-- +goose Up
CREATE TABLE pairs (
    id SERIAL PRIMARY KEY,
    original text,
    short text,
    created_at timestamp
);

-- +goose Down
DROP TABLE pairs;
