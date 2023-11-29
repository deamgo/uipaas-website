USE `uipass-waitlist-page`;

CREATE TABLE IF NOT EXISTS user (
    userid VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255)
);

<<<<<<< HEAD
CREATE TABLE IF NOT EXISTS company (
=======
CREATE TABLE IF NOT EXISTS company_info (
>>>>>>> main
    id uint primary key,
    company_name varchar(255),
    company_size varchar(255),
    name varchar(255),
    phone varchar(255),
    requirement_description varchar(255),
    date datetime(3)
);

insert into user values (1,'admin','','admin');