CREATE TABLE users (
    id SERIAL PRIMARY KEY, 
    email varchar(50) NOT NULL UNIQUE, 
    name varchar(50) NOT NULL, 
    passwordHash varchar(50) NOT NULL,
    birthday INT NOT NULL,
    description varchar(250),
    city varchar(50),
    sex varchar(10) NOT NULL,
    isActive BOOLEAN NOT NULL,
    isDeleted BOOLEAN NOT NULL
);

INSERT INTO users (
    email, 
    name, 
    passwordHash, 
    birthday, 
    description, 
    city,
    sex,
    isActive,
    isDeleted
    ) VALUES ('mail1');

DELETE FROM users WHERE id=1;

ALTER TABLE users ADD COLUMN email VARCHAR (50);

ALTER TABLE users DROP COLUMN name;

SELECT * FROM users;

SELECT current_catalog;

CREATE TABLE sessions 
(
    userId INT, 
    FOREIGN KEY (userId) REFERENCES users (id), 
    token varchar(40) NOT NULL UNIQUE, 
    expiration INT
)

INSERT INTO sessions (userId, token, expiration) VALUES (4,'123',1616974081);

DROP TABLE person;

SELECT userId FROM sessions WHERE token='23';