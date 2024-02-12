-- +migrate Up

CREATE TABLE IF NOT EXISTS codeforge.blogs (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    blog_image TEXT,
    created_at BIGINT NOT NULL,
    updated_at BIGINT DEFAULT NULL,
    deleted_at BIGINT DEFAULT NULL,
    CONSTRAINT fk_blogs FOREIGN KEY (user_id) REFERENCES codeforge.users (id) ON DELETE CASCADE
);

-- +migrate Down

DROP TABLE IF EXISTS codeforge.blogs;
