package postgres

// поддержка базы данных под управлением СУБД PostgreSQL

import (
	"GoNews/pkg/storage"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool" // command to download the pgx package and its dependencies: "go get github.com/jackc/pgx/v5/pgxpool" "go mod tidy"
)

// Хранилище данных.
type Store struct {
	*pgxpool.Pool
}

// Конструктор объекта хранилища.
func New(connString string) (*Store, error) {
	var ctx context.Context = context.Background()
	db_sql, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	if err = db_sql.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to reach database: %v", err)
	}
	s := Store{db_sql}

	return &s, nil
}

// Posts возвращает все посты из базы данных.
func (s *Store) Posts() ([]storage.Post, error) {
	rows, err := s.Query(context.Background(), `
		SELECT
			p.id AS post_id,  
			p.title AS post_title, 
			p.content AS post_content, 
			p.author_id AS post_author_id,
			a.name AS author_name,
			p.created_at AS post_created_at
		FROM posts
			authors a
		JOIN 
			posts p ON a.id = p.author_id
		ORDER BY p.id;
	`,
	)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t storage.Post
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&t.AuthorID,
			&t.AuthorName,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		posts = append(posts, t)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return posts, rows.Err()
}

// AddPost добавляет новый пост в базу данных.
func (s *Store) AddPost(post storage.Post) error {
	var id int
	err := s.QueryRow(context.Background(), `
		INSERT INTO posts (title, content, created_at)
		VALUES ($1, $2, CURRENT_TIMESTAMP) RETURNING id;
		`,
		post.Title,
		post.Content,
	).Scan(&id)
	return err
}

// UpdatePost обновляет существующий пост в базе данных.
func (s *Store) UpdatePost(post storage.Post) error {
	_, err := s.Query(context.Background(), `
		UPDATE posts
		SET (title=$1, content=$2)
		WHERE $3;
		`,
		post.Title,
		post.Content,
		post.ID,
	)
	return err
}

// DeletePost удаляет пост из базы данных.
func (s *Store) DeletePost(post storage.Post) error {
	_, err := s.Query(context.Background(), `
		DELETE FROM posts
		WHERE $1;
		`,
		post,
	)
	return err
}
