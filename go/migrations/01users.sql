-- +migrate Up

CREATE SCHEMA IF NOT EXISTS codeforge;


CREATE TABLE IF NOT EXISTS codeforge.users (
    id SERIAL PRIMARY KEY,
    user_name TEXT,
    user_bio TEXT DEFAULT NULL,
    user_profile TEXT DEFAULT NULL,
    user_email TEXT UNIQUE,
    user_password VARCHAR(255),
    gender TEXT,
    created_at BIGINT,
    updated_at BIGINT,
    deleted_at BIGINT
);

-- +migrate Down

DROP TABLE IF EXISTS codeforge.users;
