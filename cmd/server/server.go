package main

import (
	"GoNews/pkg/api"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/memdb"
	"GoNews/pkg/storage/mongo"
	"GoNews/pkg/storage/postgres"
	"log"
	"net/http"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

const (
	host           = "172.16.87.110"
	portPostges    = 5432
	portMongo      = "27017"
	userDB         = "sergey"
	password       = "password"
	dbnamePostges  = "postgres"
	dbnameMongo    = "admin"
	collectionName = "Posts"
)

func main() {
	// Создаём объект сервера.
	var srv server

	// Создаём объекты баз данных.
	//
	// БД в памяти.
	db := memdb.New()

	// реляционная БД postgresql
	db2, err := postgres.New("postgres://" + userDB + ":" + password + "@" + host + "/" + dbnamePostges)
	if err != nil {
		log.Fatal(err)
	}
	// Документная БД MongoDB.
	db3, err := mongo.New("mongodb://" + userDB + ":" + password + "@" + host + ":" + portMongo)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = db2, db3

	// Инициализируем хранилище сервера конкретной БД: db2, db3
	srv.db = db

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
	// Предаём серверу маршрутизатор запросов,
	// поэтому сервер будет все запросы отправлять на маршрутизатор.
	// Маршрутизатор будет выбирать нужный обработчик.
	http.ListenAndServe(":8080", srv.api.Router())
	/*
		var postsAdd = []storage.Post{
			{
				ID:         1,
				Title:      "Effective Go whith PostgreSQL",
				Content:    "Go is a new language. PostgreSQL is a free and open-source relational database management system (RDBMS) emphasizing extensibility and SQL compliance. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand.",
				AuthorID:   1,
				AuthorName: "Mitch Dresdner",
			},
			{
				ID:         2,
				Title:      "The Go Memory Model",
				Content:    "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
				AuthorID:   2,
				AuthorName: "go.dev",
			},
		}

		for _, p := range postsAdd {
			err := srv.db.AddPost(p)
			if err != nil {
				fmt.Println(err)
			}
		}
	*/
}
