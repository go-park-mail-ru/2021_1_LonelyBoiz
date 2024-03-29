DROP TABLE IF EXISTS sessions CASCADE;
DROP TABLE IF EXISTS secretPermission CASCADE;
DROP TABLE IF EXISTS secretphotos CASCADE;
DROP TABLE IF EXISTS messages CASCADE;
DROP TABLE IF EXISTS photos CASCADE;
DROP TABLE IF EXISTS chats CASCADE;
DROP TABLE IF EXISTS feed CASCADE;
DROP TABLE IF EXISTS users CASCADE;
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
    datePreference varchar(10),
    photos varchar(50) [] DEFAULT array []::varchar [],
    isActive BOOLEAN NOT NULL,
    isDeleted BOOLEAN NOT NULL,
    scrolls INT DEFAULT 20,
    height INT DEFAULT -1,
    partnerHeightTop INT DEFAULT -1,
    partnerHeightBot INT DEFAULT -1,
    weight INT DEFAULT -1,
    partnerWeightTop INT DEFAULT -1,
    partnerWeightBot INT DEFAULT -1,
    partnerAgeTop INT DEFAULT -1,
    partnerAgeBot INT DEFAULT -1,
    interests int [] DEFAULT array []::int []
);
CREATE TABLE photos(
    photoId SERIAL PRIMARY KEY,
    photoUuid UUID UNIQUE NOT NULL,
    userId INT,
    FOREIGN KEY (userId) REFERENCES users (id) ON DELETE CASCADE
);
CREATE TABLE chats (
    Id SERIAL PRIMARY KEY,
    userId1 INT,
    FOREIGN KEY (userId1) REFERENCES users (id) ON DELETE CASCADE,
    userId2 INT,
    FOREIGN KEY (userId2) REFERENCES users (id) ON DELETE CASCADE,
    isOpened BOOLEAN DEFAULT FALSE
);
CREATE TABLE messages (
    messageId SERIAL PRIMARY KEY,
    chatId INT,
    FOREIGN KEY (chatId) REFERENCES chats (id) ON DELETE CASCADE,
    authorId INT,
    FOREIGN KEY (authorId) REFERENCES users (id) ON DELETE CASCADE,
    text varchar(200) DEFAULT 'empty',
    time INT,
    reaction int DEFAULT -1,
    messageOrder INT
);
CREATE TABLE feed (
    userId1 INT,
    FOREIGN KEY (userId1) REFERENCES users (id) ON DELETE CASCADE,
    userId2 INT,
    FOREIGN KEY (userId2) REFERENCES users (id) ON DELETE CASCADE,
    rating varchar(10) DEFAULT 'empty'
);
CREATE TABLE sessions (
    userId INT,
    FOREIGN KEY (userId) REFERENCES users (id) ON DELETE CASCADE,
    token varchar(40) NOT NULL UNIQUE,
    expiration INT
);
CREATE TABLE secretPhotos(
    photoId SERIAL PRIMARY KEY,
    photos varchar(50) [] DEFAULT array []::varchar [],
    userId INT,
    FOREIGN KEY (userId) REFERENCES users (id) ON DELETE CASCADE
);
CREATE TABLE secretPermission(
    ownerId INT,
    FOREIGN KEY (ownerId) REFERENCES users (id) ON DELETE CASCADE,
    getterId INT,
    FOREIGN KEY (getterId) REFERENCES users (id) ON DELETE CASCADE
);