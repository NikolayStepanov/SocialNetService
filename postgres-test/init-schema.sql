CREATE TABLE IF NOT EXISTS users
(
    id serial primary key,
    name varchar(1024) not null,
    age integer not null,
    friends integer ARRAY
);
INSERT INTO users (name, age, friends) VALUES ('Vassily Petrov', 22, ARRAY[2,5]);
INSERT INTO users (name, age, friends) VALUES ('Mark Anderson', 30, ARRAY[1,3,4]);
INSERT INTO users (name, age, friends) VALUES ('David Freeman', 43, ARRAY[2,4,5]);
INSERT INTO users (name, age, friends) VALUES ('Frank Fields', 24, ARRAY[2,3,5]);
INSERT INTO users (name, age, friends) VALUES ('Richard Hardy', 29, ARRAY[1,3,4]);
INSERT INTO users (name, age, friends) VALUES ('Vernon Pearson', 19, ARRAY[7,8]);
INSERT INTO users (name, age, friends) VALUES ('Milton Cook', 43, ARRAY[6]);
INSERT INTO users (name, age, friends) VALUES ('Kathleen Smith', 32, ARRAY[6]);
INSERT INTO users (name, age) VALUES ('Jennifer Cruz', 20);
INSERT INTO users (name, age) VALUES ('Anna Adams', 32);