CREATE TABLE users
(
    id            serial not null unique,
    username      varchar(255) not null unique,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE moderators
(
    id            serial not null unique,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null,
    role          varchar(255) not null
);