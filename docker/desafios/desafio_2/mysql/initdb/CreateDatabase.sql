USE nodedb;

CREATE TABLE user (
    id integer not null auto_increment primary key,
    name varchar(200)
);

INSERT INTO user (name) VALUES ('João Fulano');