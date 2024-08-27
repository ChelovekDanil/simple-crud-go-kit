CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY, 
    first_name VARCHAR(50), 
    last_name VARCHAR(50)
);

INSERT INTO users(first_name, last_name) 
VALUES('danil', 'antonov');
