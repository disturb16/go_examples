CREATE TABLE users (
  id serial constraint users_id PRIMARY KEY,
  name TEXT,
  email TEXT,
  password TEXT
);