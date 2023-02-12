package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yucatty/Go-learning/myapi/controllers"
)

func NewRouter(con *controllers.MyAppController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ListArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ShowArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.NiceArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.CommentHandler).Methods(http.MethodPost)

	return r
}
