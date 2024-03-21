-- +goose Up
CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    date TIMESTAMP
);

-- +goose Down
DROP TABLE events;