CREATE TABLE sessions(
    id SERIAL PRIMARY KEY NOT NULL,
    user_name VARCHAR(100) NOT NULL UNIQUE,
    token VARCHAR(255) NOT NULL,
    expiry TIMESTAMP NOT NULL,
    Foreign Key (user_name) REFERENCES users (name)
);