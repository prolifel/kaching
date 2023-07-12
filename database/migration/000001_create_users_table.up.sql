create table if not exists users (
   user_id serial primary key,
   email varchar unique not null,
   name varchar not null
);