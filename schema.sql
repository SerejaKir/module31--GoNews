DROP TABLE IF EXISTS posts, authors;

CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,              -- Уникальный идентификатор статьи
    author_id INTEGER REFERENCES authors(id) NOT NULL,   -- Имя автора статьи
    title TEXT  NOT NULL,               -- Заголовок статьи
    content TEXT NOT NULL,              -- Текст статьи
    created_at BIGINT NOT NULL          -- Время создания статьи
);

INSERT INTO authors (id, name) VALUES (0, 'Дмитрий');
INSERT INTO posts (id, author_id, title, content, created_at) VALUES (0, 0, 'Статья', 'Содержание статьи', CURRENT_TIMESTAMP);