-- rambler up

CREATE TABLE users (
id SERIAL PRIMARY KEY,
name VARCHAR(50) UNIQUE NOT NULL,
created_on TIMESTAMP NOT NULL,
password VARCHAR(255),
top_score INTEGER
);

-- rambler down

DROP TABLE users;
