CREATE TABLE player (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) NOT NULL
);
CREATE TABLE game_item (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL
);
CREATE TABLE game_item_metadata (
    type VARCHAR(255) NOT NULL,
    game_item_id INTEGER REFERENCES game_item(id) ON DELETE CASCADE,
    data JSONB NOT NULL,
    PRIMARY KEY (game_item_id, type)
);
CREATE TABLE inventory_item (
    player_id INTEGER REFERENCES player(id) NOT NULL,
    item_id INTEGER REFERENCES game_item(id) NOT NULL,
    quantity INTEGER NOT NULL,
    PRIMARY KEY (player_id, item_id)
);