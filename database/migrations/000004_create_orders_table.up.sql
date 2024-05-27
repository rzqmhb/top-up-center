CREATE TABLE orders(
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INT NOT NULL,
    status VARCHAR(100) NOT NULL,
    date DATE NOT NULL,
    Foreign KEY (user_id) REFERENCES users (id)
);