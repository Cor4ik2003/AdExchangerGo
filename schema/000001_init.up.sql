CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    email varchar(255) not null,
    password_hash varchar(255) not null
)

CREATE TABLE advertisements
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    category varchar(255) not null,
    
)