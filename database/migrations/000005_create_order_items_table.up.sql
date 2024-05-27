CREATE TABLE order_items(
    id SERIAL PRIMARY KEY NOT NULL,
    order_id INT NOT NULL,
    item_id INT NOT NULL,
    current_price FLOAT NOT NULL,
    Foreign KEY (order_id) REFERENCES orders (id),
    Foreign KEY (item_id) REFERENCES items (id)
);