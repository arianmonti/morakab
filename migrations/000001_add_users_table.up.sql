CREATE TABLE IF NOT EXISTS users (
   id serial PRIMARY KEY,
   username VARCHAR (128) UNIQUE NOT NULL,
   email VARCHAR (256) UNIQUE NOT NULL,
   password_hash VARCHAR (256) NOT NULL
);