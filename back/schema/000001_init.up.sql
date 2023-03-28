CREATE TABLE users (
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    email varchar(255) not null unique,
    tel varchar(255),
    subscribe varchar(255) not null,
    id serial not null unique
);
CREATE TABLE datasets (
    title varchar(255) not null,
    description varchar(255) not null,
    id serial not null unique,
);
