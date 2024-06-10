CREATE TABLE orders(
    id SERIAL PRIMARY KEY NOT NULL,
    item_id INT NOT NULL,
    user_id INT NOT NULL,
    item_current_price FLOAT NOT NULL,
    in_game_user_id VARCHAR(100) NOT NULL UNIQUE,
    status VARCHAR(100) NOT NULL,
    date DATE NOT NULL,
    Foreign KEY (item_id) REFERENCES items (id),
    Foreign KEY (user_id) REFERENCES users (id)
);