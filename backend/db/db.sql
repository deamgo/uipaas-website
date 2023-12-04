USE `uipaas-home`;

CREATE TABLE IF NOT EXISTS user (
    userid VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255)
);


CREATE TABLE IF NOT EXISTS company (
    id int PRIMARY KEY AUTO_INCREMENT ,
    company_name varchar(255) NOT NULL,
    company_size varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    business_email  varchar(255) NOT NULL,
    requirement_description varchar(255) NOT NULL,
    date datetime(3)
);

insert into user values (1,'admin','','admin');