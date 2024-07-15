-- +goose Up
CREATE TABLE pairs (
    id SERIAL PRIMARY KEY,
    original text,
    short text,
    created_at timestamp
--     PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE pairs;
