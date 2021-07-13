package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Id int
	Content string
	Author string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Yukako"}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)
}