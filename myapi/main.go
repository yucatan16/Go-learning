package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/yucatty/Go-learning/myapi/controllers"
	"github.com/yucatty/Go-learning/myapi/services"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	log.Println(dbConn)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}
	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppController(ser)
	r := mux.NewRouter()
	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ListArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ShowArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.NiceArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.CommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
