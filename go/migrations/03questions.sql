-- +migrate Up

CREATE TABLE IF NOT EXISTS codeforge.questions (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    question_language TEXT NOT NULL,
    picture TEXT DEFAULT NULL,
    created_at BIGINT,
    updated_at BIGINT,
    deleted_at BIGINT,
    CONSTRAINT fk_questions_users FOREIGN KEY (user_id) REFERENCES codeforge.users (id) ON DELETE CASCADE
);

-- +migrate Down

DROP TABLE IF EXISTS codeforge.questions;
