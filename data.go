package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var Db *sql.DB

func init() {
	dataSource := os.Getenv("WEB_APP_DB_URI")
	var err error
	Db, err = sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}
}

func retrieve(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from post where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) create() (err error) {
	stmt, err := Db.Prepare("insert into post (content, author) values ($1, $2) returning id")
	if err != nil {
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}(stmt)
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) update() (err error) {
	_, err = Db.Exec("update post set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from post where id = $1", post.Id)
	return
}
