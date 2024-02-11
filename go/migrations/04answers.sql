-- +migrate Up

CREATE TABLE IF NOT EXISTS codeforge.answers(
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    question_id BIGINT NOT NULL,
    content TEXT DEFAULT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT DEFAULT NULL,
    deleted_at BIGINT DEFAULT NULL,
    CONSTRAINT fk_answers_users FOREIGN KEY (user_id) REFERENCES codeforge.users(id) ON DELETE CASCADE,
    CONSTRAINT fk_answers_questions FOREIGN KEY (question_id) REFERENCES codeforge.questions(id) ON DELETE CASCADE
);

-- +migrate Down

DROP TABLE IF EXISTS codeforge.answers;
