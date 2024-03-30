-- Creates the "user" table on the schema named "root"
CREATE TABLE IF NOT EXISTS root.users (
    id SERIAL PRIMARY KEY, 
    name VARCHAR(255) NOT NULL, 
    address TEXT NOT NULL,
    email VARCHAR(255) NOT NULL
);