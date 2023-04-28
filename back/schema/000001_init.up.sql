CREATE TABLE users (
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    email varchar(255) not null unique,
    tel varchar(255),
    subscribe varchar(255) not null,
    time_of_sub int not null,
    id serial not null unique
);
CREATE TABLE datasets (
    title varchar(255) not null,
    description varchar(2000) not null,
    id serial not null unique
);
CREATE TABLE users_datasets (
    id serial not null unique,
    id_user int references users(id) on delete cascade  not null,
    id_dataset int references datasets(id) on delete cascade  not null
);
CREATE TABLE datasetItem (
    id serial not null unique,
    datainfo varchar(2000),
    solution varchar(2000)
);
CREATE TABLE datasets_items (
    id serial not null unique,
    id_dataset int references datasets(id) on delete cascade  not null,
    id_item int references datasetItem(id) on delete cascade  not null
);