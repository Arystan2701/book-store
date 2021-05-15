create table users
(
    id            serializable       not null unique,
    user_name     varchar(255) not null unique,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null
);

create table moderators
(
    id            serializable       not null unique,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null,
    role          varchar(255) not null
);