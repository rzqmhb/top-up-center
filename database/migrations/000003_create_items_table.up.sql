CREATE TABLE items(
    id SERIAL PRIMARY KEY NOT NULL,
    game_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    price FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (game_id) REFERENCES games (id)
);