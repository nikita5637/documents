CREATE TABLE documents (
    id bigserial not null primary key,
    name varchar(100) not null,
    date varchar(100) not null,
    number bigint not null unique,
    sum varchar(64) not null
);