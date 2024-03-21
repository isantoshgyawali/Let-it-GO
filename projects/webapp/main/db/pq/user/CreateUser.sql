CREATE TABLE IF NOT EXISTS root.users (
    id SERIAL PRIMARY KEY, 
    name VARCHAR(255) NOT NULL, 
    address TEXT NOT NULL,
    email VARCHAR(255) NOT NULL
);

-- INSERT INTO root.users (name,address,email)
-- VALUES ('RAM', 'BHAKTAPUR', 'ram@example.com'),
--        ('SITA', 'RUPANDEHI', 'sita@example.com');