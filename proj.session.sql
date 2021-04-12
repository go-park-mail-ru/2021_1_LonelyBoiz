CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email varchar(51) UNIQUE,
    name varchar(52),
    passwordHash bytea,
    birthday INT,
    description varchar(250),
    city varchar(53),
    photos varchar(500),
    sex varchar(10),
    datePreference varchar(10),
    isActive BOOLEAN NOT NULL,
    isDeleted BOOLEAN NOT NULL
);

CREATE TABLE messages (
    messageId SERIAL PRIMARY KEY,
    chatId INT,
    FOREIGN KEY (chatId) REFERENCES chats (id),
    authorId INT,
    FOREIGN KEY (authorId) REFERENCES users (id),
    text varchar(200) NOT NULL,
    time INT,
    reaction int DEFAULT -1,
    messageOrder INT
);


CREATE TABLE photos(
    photoId SERIAL PRIMARY KEY,
    userId INT,
    FOREIGN KEY (userId) REFERENCES users (id)
);

CREATE TABLE chats (
    Id SERIAL PRIMARY KEY,
    userId1 INT,
    FOREIGN KEY (userId1) REFERENCES users (id),
    userId2 INT,
    FOREIGN KEY (userId2) REFERENCES users (id)
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