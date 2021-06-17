set client_encoding = 'UTF8';

create table users (
  id integer not null,
  name varchar not null,
  email varchar not null,
  password varchar not null
);

create table todos (
  id integer not null,
  title varchar not null,
  context varchar not null,
  user_id integer not null
);
