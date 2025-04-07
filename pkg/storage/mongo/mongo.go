package mongo

// поддержка базы данных под управлением СУБД MongoDB

import (
	"GoNews/pkg/storage"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Хранилище данных.
type Store struct {
	*mongo.Client
}

const (
	databaseName   = "admin"
	collectionName = "Posts"
)

// Конструктор объекта хранилища.
func New(connString string) (*Store, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		return nil, fmt.Errorf("failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Проверка соединения
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	s := Store{client}

	return &s, nil
}

// Posts возвращает все посты из базы данных.
func (s *Store) Posts() ([]storage.Post, error) {

	collection := s.Database(databaseName).Collection(collectionName)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var posts []storage.Post
	for cur.Next(context.Background()) {
		var l storage.Post
		err := cur.Decode(&l)
		if err != nil {
			return nil, err
		}
		posts = append(posts, l)
	}
	return posts, cur.Err()
}

// AddPost добавляет новый пост в базу данных.
func (s *Store) AddPost(post storage.Post) error {
	collection := s.Database(databaseName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePost обновляет существующий пост в базе данных.
func (s *Store) UpdatePost(post storage.Post) error {
	collection := s.Database(databaseName).Collection(collectionName)
	_, err := collection.UpdateByID(context.Background(), post.ID, post)
	if err != nil {
		return err
	}
	return nil
}

// DeletePost удаляет пост из базы данных.
func (s *Store) DeletePost(post storage.Post) error {
	collection := s.Database(databaseName).Collection(collectionName)
	_, err := collection.DeleteOne(context.Background(), post.ID)
	if err != nil {
		return err
	}
	return nil
}
