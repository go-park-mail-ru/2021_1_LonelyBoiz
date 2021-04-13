DROP TABLE sessions;
DROP TABLE messages;
DROP TABLE photos;
DROP TABLE chats;
DROP TABLE feed;
DROP TABLE users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email varchar(50) UNIQUE,
    name varchar(50),
    passwordHash bytea,
    birthday INT,
    instagram varchar(50),
    description varchar(250),
    city varchar(50),
    sex varchar(10),
    instagram varchar(70),
    datePreference varchar(10),
    isActive BOOLEAN NOT NULL,
    isDeleted BOOLEAN NOT NULL
);

CREATE TABLE photos(
    photoId SERIAL PRIMARY KEY,
    userId INT,
    FOREIGN KEY (userId) REFERENCES users (id),
    value BYTEA
);

CREATE TABLE chats (
    Id SERIAL PRIMARY KEY,
    userId1 INT,
    FOREIGN KEY (userId1) REFERENCES users (id),
    userId2 INT,
    FOREIGN KEY (userId2) REFERENCES users (id)
);

CREATE TABLE messages (
    messageId SERIAL PRIMARY KEY,
    chatId INT,
    FOREIGN KEY (chatId) REFERENCES chats (id),
    authorId INT,
    FOREIGN KEY (authorId) REFERENCES users (id),
    text varchar(200) DEFAULT 'empty',
    time INT,
    reaction int DEFAULT -1,
    messageOrder INT
);

CREATE TABLE feed (
    userId1 INT,
    FOREIGN KEY (userId1) REFERENCES users (id),
    userId2 INT,
    FOREIGN KEY (userId2) REFERENCES users (id),
    rating varchar(10) DEFAULT 'empty'
);

CREATE TABLE sessions (
    userId INT,
    FOREIGN KEY (userId) REFERENCES users (id),
    token varchar(40) NOT NULL UNIQUE,
    expiration INT
);