CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR UNIQUE,
    parent_id INT REFERENCES users(id) NULL
);

INSERT INTO users (name, parent_id)
VALUES 
    ('Zaki', 2),
    ('Ilham', NULL),
    ('Irwan', 2),
    ('Arka', 3);


SELECT u.id, u.name, p.name AS parent_name
FROM users AS u
LEFT JOIN users AS p ON p.id = u.parent_id;