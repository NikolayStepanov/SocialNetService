CREATE TABLE IF NOT EXISTS users
(
    id serial not null,
    name varchar(1024) not null,
    age integer not null,
    friends integer ARRAY
);