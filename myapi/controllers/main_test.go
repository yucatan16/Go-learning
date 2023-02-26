package controllers_test

import (
	"testing"

	"github.com/yucatty/Go-learning/myapi/controllers"
	"github.com/yucatty/Go-learning/myapi/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
