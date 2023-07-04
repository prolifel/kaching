create table if not exists users (
   user_id bigint primary key,
   email varchar unique not null,
   name varchar not null
);