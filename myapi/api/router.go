package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yucatty/Go-learning/myapi/api/middlewares"
	"github.com/yucatty/Go-learning/myapi/controllers"
	"github.com/yucatty/Go-learning/myapi/services"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	r := mux.NewRouter()
	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ListArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ShowArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.NiceArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.CommentHandler).Methods(http.MethodPost)

	r.Use(middlewares.LoggingMiddleware)
	return r
}
