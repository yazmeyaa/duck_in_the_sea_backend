CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    telegram_id INTEGER UNIQUE NOT NULL,
    created_at TIME DEFAULT now()
);