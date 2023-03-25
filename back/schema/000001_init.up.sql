CREATE TABLE users (
    username varchar(255) not null,
    password_hash varchar(255) not null,
    email varchar(255) not null,
    tel varchar(255),
    subscribe varchar(255) DEFAULT "free",
    id serial not null unique
);