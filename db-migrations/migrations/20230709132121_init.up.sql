CREATE TABLE IF NOT EXISTS todos
(
    id serial PRIMARY KEY,
    text VARCHAR,
    status BOOLEAN
);