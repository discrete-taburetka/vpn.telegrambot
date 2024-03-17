CREATE USER tg_bot WITH PASSWORD 'tgbotpwd12345';
CREATE DATABASE tg_bot;
GRANT ALL PRIVILEGES ON DATABASE "tg_bot" to tg_bot;
\connect tg_bot
CREATE TABLE tg_users (
    tg_user_name    varchar(1000) CONSTRAINT firstkey PRIMARY KEY,
    user_name       varchar(1000) NOT NULL,
    user_password   varchar(1000) NOT NULL,
    user_email      varchar(1000) NOT NULL
);
ALTER TABLE tg_users OWNER TO tg_bot;
