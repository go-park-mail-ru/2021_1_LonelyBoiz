DROP TABLE sessions;
DROP TABLE secretPermission;
DROP TABLE secretphotos;
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
    datePreference varchar(10),
    photos varchar(50)[] DEFAULT array[]::varchar[],
    isActive BOOLEAN NOT NULL,
    isDeleted BOOLEAN NOT NULL,
    scrolls INT DEFAULT 20
);


CREATE TABLE photos(
    photoId SERIAL PRIMARY KEY,
    photoUuid UUID UNIQUE NOT NULL ,
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

CREATE TABLE secretPhotos(
    photoId SERIAL PRIMARY KEY,
    photos varchar(50)[] DEFAULT array[]::varchar[],
    userId INT,
    FOREIGN KEY (userId) REFERENCES users (id)
);

CREATE TABLE secretPermission(
    ownerId INT,
    FOREIGN KEY (ownerId) REFERENCES users (id),
    getterId INT,
    FOREIGN KEY (getterId) REFERENCES users (id)
);

select photos, email from users;

select * from chats;

select * from secretphotos ;

insert into chats(userid1, userid2) values(2,3)

update users set photos='{"e7816294-3529-42cd-a612-4f323d6d972a"}' where id=2;

insert into secretphotos values (4, '{}', 1) 