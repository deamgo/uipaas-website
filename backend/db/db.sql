USE `uipass-waitlist-page`;

CREATE TABLE IF NOT EXISTS user (
    id VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255)
);


DROP TABLE user;